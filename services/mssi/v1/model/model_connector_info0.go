package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectorInfo0 Connector元数据
type ConnectorInfo0 struct {
	AuthContent *AuthConfigA `json:"authContent,omitempty"`

	// 认证id
	AuthId *string `json:"authId,omitempty"`

	// 判断方式
	AuthType *string `json:"auth_type,omitempty"`

	// 内置连接器查询
	Category *string `json:"category,omitempty"`

	// 创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 连接器
	DefinitionRef *string `json:"definitionRef,omitempty"`

	// 连接器描述
	Description *string `json:"description,omitempty"`

	// logo base64编码
	Icon *string `json:"icon,omitempty"`

	// 连接器ID
	Id *string `json:"id,omitempty"`

	// 连接器名称
	Name *string `json:"name,omitempty"`

	// 是否需要验证
	NeedAuth *bool `json:"needAuth,omitempty"`

	// 视图数据
	Operations *[]interface{} `json:"operations,omitempty"`

	// 供应商
	Provider *string `json:"provider,omitempty"`

	// swagger文档，大文本
	Swagger *string `json:"swagger,omitempty"`

	// 操作json
	Triggers *[]interface{} `json:"triggers,omitempty"`

	// 连接器类型
	Type *string `json:"type,omitempty"`

	// 修改时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`
}

func (o ConnectorInfo0) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectorInfo0 struct{}"
	}

	return strings.Join([]string{"ConnectorInfo0", string(data)}, " ")
}
