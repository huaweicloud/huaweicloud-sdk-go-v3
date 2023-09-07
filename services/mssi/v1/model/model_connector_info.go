package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectorInfo Connector元数据
type ConnectorInfo struct {

	// 执行动作数量
	ActionCount *int32 `json:"action_count,omitempty"`

	// 触发事件数量
	Actions *[]ActionBaseInfo `json:"actions,omitempty"`

	// 安全认证配置内容
	AuthContent *interface{} `json:"auth_content,omitempty"`

	// 认证配置ID
	AuthId *string `json:"auth_id,omitempty"`

	// 自定义连接器种类（连接器市场的tab分类）
	Category *string `json:"category,omitempty"`

	// 创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 自定义连接器描述
	Description *string `json:"description,omitempty"`

	// 是否收藏
	Favorite *bool `json:"favorite,omitempty"`

	// logo base64编码
	Icon *string `json:"icon,omitempty"`

	// 自定义连接器ID
	Id *string `json:"id,omitempty"`

	// 自定义连接器名称
	Name *string `json:"name,omitempty"`

	// 是否需要验证
	NeedAuth *bool `json:"need_auth,omitempty"`

	// 服务提供商
	ProviderName *string `json:"provider_name,omitempty"`

	// 发布版本
	ReleaseVersion *string `json:"release_version,omitempty"`

	// 权限
	RuntimePermissions *[]RuntimePermission `json:"runtime_permissions,omitempty"`

	// 状态(dev：草稿、released：已发布、onboard：已上架)
	Status *string `json:"status,omitempty"`

	// swagger文档（只包含基本信息+认证信息）
	Swagger *interface{} `json:"swagger,omitempty"`

	// 版本id
	SwaggerVersionId *string `json:"swagger_version_id,omitempty"`

	// 触发事件数量
	TriggerCount *int32 `json:"trigger_count,omitempty"`

	// 触发事件数量
	Triggers *[]TriggerBaseInfo `json:"triggers,omitempty"`

	// 自定义连接器类型
	Type *string `json:"type,omitempty"`

	// 修改时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`

	// 版本号
	Version *interface{} `json:"version,omitempty"`
}

func (o ConnectorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectorInfo struct{}"
	}

	return strings.Join([]string{"ConnectorInfo", string(data)}, " ")
}
