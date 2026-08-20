package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceCrossVpcIpRespResults 修改broker跨VPC访问的结果。
type UpdateInstanceCrossVpcIpRespResults struct {

	// **参数解释**： advertised.listeners IP/域名。 **取值范围**： 不涉及。
	AdvertisedIp *string `json:"advertised_ip,omitempty"`

	// **参数解释**： 修改broker跨VPC访问的状态。 **取值范围**： - true：修改broker跨VPC访问成功。 - false：修改broker跨VPC访问失败。
	Success *bool `json:"success,omitempty"`

	// **参数解释**： listeners IP。 **取值范围**： 不涉及。
	Ip *string `json:"ip,omitempty"`
}

func (o UpdateInstanceCrossVpcIpRespResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceCrossVpcIpRespResults struct{}"
	}

	return strings.Join([]string{"UpdateInstanceCrossVpcIpRespResults", string(data)}, " ")
}
