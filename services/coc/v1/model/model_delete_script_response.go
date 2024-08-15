package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScriptResponse Response Object
type DeleteScriptResponse struct {

	// 被删除脚本的uuid
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScriptResponse struct{}"
	}

	return strings.Join([]string{"DeleteScriptResponse", string(data)}, " ")
}
