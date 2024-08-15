package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecutePublicScriptResponse Response Object
type ExecutePublicScriptResponse struct {

	// 执行公共脚本返回体:execute_uuid
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecutePublicScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutePublicScriptResponse struct{}"
	}

	return strings.Join([]string{"ExecutePublicScriptResponse", string(data)}, " ")
}
