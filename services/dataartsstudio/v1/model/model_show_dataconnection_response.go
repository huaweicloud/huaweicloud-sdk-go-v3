package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataconnectionResponse Response Object
type ShowDataconnectionResponse struct {

	// 数据连接名称
	DwName *string `json:"dw_name,omitempty"`

	// 数据连接类型
	DwType *string `json:"dw_type,omitempty"`

	// 连接动态变化配置项，每种连接略有区别，建议在界面进行调试
	DwConfig *interface{} `json:"dw_config,omitempty"`

	// 代理id
	AgentId *string `json:"agent_id,omitempty"`

	// 代理名称
	AgentName *string `json:"agent_name,omitempty"`

	// 0：开发模式 1：生产模式。默认为0
	EnvType *int32 `json:"env_type,omitempty"`

	// 数据连接限定名称
	QualifiedName *string `json:"qualified_name,omitempty"`

	// 数据连接id
	DwId *string `json:"dw_id,omitempty"`

	// 数据连接创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 数据连接创建时间，时间戳
	CreateTime float32 `json:"create_time,omitempty"`

	// 数据连接类别
	DwCatagory *string `json:"dw_catagory,omitempty"`

	// 连接描述信息
	Description *string `json:"description,omitempty"`

	// 0：创建 1：更新。默认为0
	UpdateType     *int32 `json:"update_type,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDataconnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataconnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowDataconnectionResponse", string(data)}, " ")
}
