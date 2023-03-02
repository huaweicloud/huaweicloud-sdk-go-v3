package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowOtaPackageRequest struct {

	// Sp用户Token。通过调用IoBPS服务获取SP用户Token
	SpAuthToken *string `json:"Sp-Auth-Token,omitempty"`

	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：升级包ID，用于唯一标识一个升级包。由物联网平台分配获得。 **取值范围**：长度不超过36，只允许字母、数字、连接符（-）的组合。
	PackageId string `json:"package_id"`
}

func (o ShowOtaPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOtaPackageRequest struct{}"
	}

	return strings.Join([]string{"ShowOtaPackageRequest", string(data)}, " ")
}
