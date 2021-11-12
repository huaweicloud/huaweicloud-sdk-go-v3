package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteScriptReq struct {
	// 脚本的执行参数

	Params *string `json:"params,omitempty"`
}

func (o ExecuteScriptReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScriptReq struct{}"
	}

	return strings.Join([]string{"ExecuteScriptReq", string(data)}, " ")
}
