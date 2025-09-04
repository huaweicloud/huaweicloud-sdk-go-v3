package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityProfileRequest Request Object
type UpdateSecurityProfileRequest struct {

	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。您可以在IoTDA管理控制台界面，选择左侧导航栏“总览”页签查看当前实例的ID。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：安全态势感知配置ID，在创建安全态势感知配置时由物联网平台分配。 **取值范围**：长度不超过24，数字与字符串的组合。
	ProfileId string `json:"profile_id"`

	Body *UpdateSecurityProfileDto `json:"body,omitempty"`
}

func (o UpdateSecurityProfileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityProfileRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityProfileRequest", string(data)}, " ")
}
