package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTestCaseAndScriptResponse Response Object
type UpdateTestCaseAndScriptResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTestCaseAndScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestCaseAndScriptResponse struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseAndScriptResponse", string(data)}, " ")
}
