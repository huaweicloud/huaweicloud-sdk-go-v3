package model

import (
	"encoding/json"

	"strings"
)

// UpdateCaseRequestBody
type UpdateCaseRequestBody struct {
	// contents

	Contents *[]UpdateCaseRequestBodyContents `json:"contents,omitempty"`
	// for_loop_params

	ForLoopParams *[]interface{} `json:"for_loop_params,omitempty"`
}

func (o UpdateCaseRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCaseRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCaseRequestBody", string(data)}, " ")
}
