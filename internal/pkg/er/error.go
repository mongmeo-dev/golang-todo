package er

import (
	"encoding/json"
	"github.com/cockroachdb/errors"
	"os"
)

type Code int

var (
	CodeUnspecified Code = 0
	CodeBadRequest  Code = 1001
)

func (c Code) String() string {
	switch c {
	case CodeBadRequest:
		return "CodeBadRequest"
	default:
		return "CodeUnspecified"
	}
}

func (c Code) StatusCode() int {
	switch c {
	case CodeBadRequest:
		return 400
	default:
		return 500
	}
}

type businessError struct {
	Code        Code   `json:"code"`
	Message     string `json:"message"`
	Meaning     string `json:"code_meaning"`
	sourceError error
}

func New(message string) error {
	return new(errors.New(message))
}
func WrapError(sourceError error) error {
	return new(sourceError)
}

func new(err error) *businessError {
	if errors.Is(err, &businessError{}) {
		var targetErr *businessError
		if errors.As(err, &targetErr) {
			return targetErr
		}
	}
	return &businessError{
		Code:        CodeUnspecified,
		sourceError: err,
	}
}

func WithCode(err error, code Code) error {
	targetError := new(err)
	targetError.Code = code
	return targetError
}

func WithMessage(err error, message string) error {
	targetError := new(err)
	targetError.Message = message
	return targetError
}

func WithStackTrace(err error, depth int) error {
	targetError := new(err)
	targetError.sourceError = errors.WithStackDepth(targetError.sourceError, depth)
	return targetError
}

func HelperErrorResponse(err error) ([]byte, error) {
	targetError := new(err)
	targetError.Meaning = targetError.Code.String()
	env := os.Getenv("ENV")
	if env != "prod" {
		targetError.Meaning = ""
	}
	b, err := json.Marshal(targetError)
	if err != nil {
		err = errors.WithMessage(err, "golang-todo.internal.pkg.HelperErrorResponse: failed to marshal this message")
		err = new(err)
		return nil, err
	}
	return b, nil
}

func Is(err, target error) bool {
	sourceErr := new(err)
	targetErr := new(target)
	return errors.Is(sourceErr.sourceError, targetErr.sourceError)
}

func As(err, target error) bool {
	sourceErr := new(err)
	targetErr := new(target)
	return errors.As(sourceErr.sourceError, targetErr.sourceError)
}

func GetResponseStatus(err error) int {
	targetError := &businessError{}
	if errors.Is(err, targetError) {
		return targetError.Code.StatusCode()
	}
	return new(err).Code.StatusCode()
}

func (e *businessError) Error() string {
	return e.sourceError.Error()
}
