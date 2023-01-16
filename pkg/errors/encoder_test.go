package errors

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ErrorStruct(t *testing.T) {
	httpError := HttpError{
		Errors: make(map[string]any),
	}
	httpError.Errors["body"] = []string{"can not be empty"}

	jsonOut, err := json.Marshal(httpError)
	assert.NoError(t, err)
	log.Info(string(jsonOut))
}
