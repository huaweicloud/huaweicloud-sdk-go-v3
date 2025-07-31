package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CicdConfigurationsResponseInfo CI/CD配置信息
type CicdConfigurationsResponseInfo struct {

	// **参数解释** CI/CD标识 **约束限制** 不涉及 **取值范围** 字符长度0-128位  **默认取值** 不涉及
	CicdId *string `json:"cicd_id,omitempty"`

	// **参数解释** CI/CD名称 **约束限制** 不涉及 **取值范围** 字符长度0-128位  **默认取值** 不涉及
	CicdName *string `json:"cicd_name,omitempty"`

	// **参数解释** 关联镜像扫描数量 **约束限制** 不涉及 **取值范围** 最小值0，最大值2147483647  **默认取值** 不涉及
	AssociatedImagesNum *int32 `json:"associated_images_num,omitempty"`

	// **参数解释** 关联配置扫描数量 **约束限制** 不涉及 **取值范围** 最小值0，最大值2147483647  **默认取值** 不涉及
	AssociatedConfigNum *int32 `json:"associated_config_num,omitempty"`
}

func (o CicdConfigurationsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CicdConfigurationsResponseInfo struct{}"
	}

	return strings.Join([]string{"CicdConfigurationsResponseInfo", string(data)}, " ")
}
