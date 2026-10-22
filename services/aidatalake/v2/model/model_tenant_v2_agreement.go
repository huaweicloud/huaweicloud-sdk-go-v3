package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TenantV2Agreement **参数解释**：租户协议。
type TenantV2Agreement struct {

	// **参数解释**：协议名称。 **取值范围**：长度为1~128的英文字母、数字、下划线和中划线的组合。
	AgreementName *string `json:"agreement_name,omitempty"`

	// **参数解释**：协议版本号。 **取值范围**：长度为1~128的英文字母、数字、下划线、中划线和小数点的组合。
	Version *string `json:"version,omitempty"`

	// **参数解释**：资源创建时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o TenantV2Agreement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantV2Agreement struct{}"
	}

	return strings.Join([]string{"TenantV2Agreement", string(data)}, " ")
}
