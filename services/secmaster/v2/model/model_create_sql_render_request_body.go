package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSqlRenderRequestBody struct {

	// 查询语句
	Query *string `json:"query,omitempty"`

	// 脚本参数
	ScriptParams *[]ScriptParam `json:"script_params,omitempty"`

	// 查询语句
	TransformQuery *string `json:"transform_query,omitempty"`

	// Console会话ID
	SessionId *string `json:"session_id,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`

	// 交互参数列表
	InterActiveParams *[]InterActiveParams `json:"inter_active_params,omitempty"`

	// 起始
	From *int64 `json:"from,omitempty"`

	// 终止
	To *int64 `json:"to,omitempty"`
}

func (o CreateSqlRenderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlRenderRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSqlRenderRequestBody", string(data)}, " ")
}
