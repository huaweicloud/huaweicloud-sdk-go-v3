package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2AgencyResponse Response Object
type ShowV2AgencyResponse struct {

	// **参数解释**：委托名称。 **取值范围**：固定值ai_datalake_admin_trust。
	AgencyName *string `json:"agency_name,omitempty"`

	// **参数解释**：授权与否。如果用户的服务委托权限包含系统所需要的委托权限则返回true，否则返回false。 **取值范围**：   - true：已授权。   - false：未授权。
	Authorized *bool `json:"authorized,omitempty"`

	// **参数解释**：资源创建时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：委托策略列表。
	AgencyPolicies *[]V2AgencyPolicy `json:"agency_policies,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowV2AgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2AgencyResponse struct{}"
	}

	return strings.Join([]string{"ShowV2AgencyResponse", string(data)}, " ")
}
