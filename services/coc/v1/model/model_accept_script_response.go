package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptScriptResponse Response Object
type AcceptScriptResponse struct {

	// 审批结果
	Data           *bool `json:"data,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o AcceptScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptScriptResponse struct{}"
	}

	return strings.Join([]string{"AcceptScriptResponse", string(data)}, " ")
}
