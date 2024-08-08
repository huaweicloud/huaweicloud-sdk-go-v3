package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AdditionalParamsResp 企业版附加参数信息。
type AdditionalParamsResp struct {

	// **参数说明**：企业版实例的VPCID
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数说明**：企业版实例的子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数说明**：企业版实例的安全组ID
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// **参数说明**：实例支持的加密算法 **取值范围**： - COMMON_ALGORITHM: 通用加密算法（支持RSA，SHA256等国际通用的密码算法） - SM_ALGORITHM: 支持SM系列商密算法（支持SM2，SM3，SM4等国密算法）
	CipheringAlgorithm *string `json:"ciphering_algorithm,omitempty"`

	ForwardingInfo *ForwardingInfo `json:"forwarding_info,omitempty"`
}

func (o AdditionalParamsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdditionalParamsResp struct{}"
	}

	return strings.Join([]string{"AdditionalParamsResp", string(data)}, " ")
}
