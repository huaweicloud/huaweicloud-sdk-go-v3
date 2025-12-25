package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTableLogsRequestBody struct {

	// 检索查询条件, 语法介绍请参考帮助文档。
	Query string `json:"query"`

	// 毫秒时间戳
	From *int64 `json:"from,omitempty"`

	// 毫秒时间戳
	To *int64 `json:"to,omitempty"`

	// 每页数量
	Limit int64 `json:"limit"`

	// 页号
	Offset *int64 `json:"offset,omitempty"`

	// 脚本参数列表
	ScriptParams *[]SearchScriptParam `json:"script_params,omitempty"`
}

func (o ListTableLogsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableLogsRequestBody struct{}"
	}

	return strings.Join([]string{"ListTableLogsRequestBody", string(data)}, " ")
}
