package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMissingIndexScriptResponse Response Object
type ShowMissingIndexScriptResponse struct {

	// 索引脚本
	Script         *string `json:"script,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMissingIndexScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMissingIndexScriptResponse struct{}"
	}

	return strings.Join([]string{"ShowMissingIndexScriptResponse", string(data)}, " ")
}
