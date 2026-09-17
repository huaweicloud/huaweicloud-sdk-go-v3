package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyConnectionRequestBody 修改实例连接请求体
type ModifyConnectionRequestBody struct {

	// 用户名
	Username string `json:"username"`

	// 密码
	Password string `json:"password"`

	// 是否保存密码
	IsSavePassword bool `json:"is_save_password"`

	// 节点ID列表
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`

	// 端口
	Port *int32 `json:"port,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// SQL记录开关
	SqlRecordFlag *bool `json:"sql_record_flag,omitempty"`
}

func (o ModifyConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyConnectionRequestBody", string(data)}, " ")
}
