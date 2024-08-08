package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AdditionalParams 企业版附加参数信息。
type AdditionalParams struct {

	// **参数说明**：企业版实例的VPCID
	VpcId string `json:"vpc_id"`

	// **参数说明**：企业版实例的子网ID
	SubnetId string `json:"subnet_id"`

	// **参数说明**：企业版实例的安全组ID。请确保所选安全组已放通22端口（Linux SSH登录），3389端口（Windows远程登录）和ICMP协议（Ping）
	SecurityGroupId string `json:"security_group_id"`

	// **参数说明**：SMN的topic urn, 当企业版实例创建成功时，平台将通过该topic发送通知
	SmnTopicUrn *string `json:"smn_topic_urn,omitempty"`

	// **参数说明**：实例支持的加密算法 **取值范围**： - COMMON_ALGORITHM: 通用加密算法（支持RSA，SHA256等国际通用的密码算法） - SM_ALGORITHM: 支持SM系列商密算法（支持SM2，SM3，SM4等国密算法）
	CipheringAlgorithm *string `json:"ciphering_algorithm,omitempty"`

	PortInfo *Port `json:"port_info"`
}

func (o AdditionalParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdditionalParams struct{}"
	}

	return strings.Join([]string{"AdditionalParams", string(data)}, " ")
}
