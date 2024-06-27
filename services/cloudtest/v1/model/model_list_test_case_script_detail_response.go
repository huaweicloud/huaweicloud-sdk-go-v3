package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestCaseScriptDetailResponse Response Object
type ListTestCaseScriptDetailResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTestCaseScriptDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestCaseScriptDetailResponse struct{}"
	}

	return strings.Join([]string{"ListTestCaseScriptDetailResponse", string(data)}, " ")
}
