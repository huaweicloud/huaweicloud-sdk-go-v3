package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FeatureConfig 参数解释：特性配置信息。
type FeatureConfig struct {

	// 参数解释：配置的ID。
	Id string `json:"id"`

	// 参数解释：创建时间。
	CreatedAt string `json:"created_at"`

	// 参数解释：更新时间。
	UpdatedAt string `json:"updated_at"`

	// 参数解释：所属服务，固定ELB。
	Service string `json:"service"`

	// 参数解释：租户ID，含义同project_id。
	TenantId string `json:"tenant_id"`

	// 参数解释：特性名称。
	Feature string `json:"feature"`

	// 参数解释：特性配置启用开关，表示当前配置是否生效。  取值范围： - true：特性配置已生效。 - false: 特性配置未生效。
	Switch bool `json:"switch"`

	// 参数解释：特性配置值(value字段)的类型，如：INT，表示整型。
	Type string `json:"type"`

	// 参数解释：特性配置值。如开关类型的特性配置取值true/false，表示特性开启关闭；配额类型的特性配置取值整数，表示限制配额。
	Value string `json:"value"`

	// 参数解释：特性配置描述。
	Description string `json:"description"`

	// 参数解释：配置创建者。
	Caller string `json:"caller"`
}

func (o FeatureConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FeatureConfig struct{}"
	}

	return strings.Join([]string{"FeatureConfig", string(data)}, " ")
}
