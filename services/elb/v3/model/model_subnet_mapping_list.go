package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubnetMappingList 源子网信息和目标子网的对应关系
type SubnetMappingList struct {

	// **参数解释**：源监听器下后端服务器所在VPC的子网ID。  **约束限制**：不涉及  **取值范围**：标准的UUID格式，长度为36个字符。  **默认取值**：不涉及
	SubnetCidrId string `json:"subnet_cidr_id"`

	// **参数解释**：新监听器下后端服务器需要绑定的VPC子网ID。  **约束限制**：该VPC子网ID需要存在，且与源子网网段相同。  **取值范围**：标准的UUID格式，长度为36个字符。  **默认取值**：不涉及
	DstSubnetCidrId string `json:"dst_subnet_cidr_id"`
}

func (o SubnetMappingList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetMappingList struct{}"
	}

	return strings.Join([]string{"SubnetMappingList", string(data)}, " ")
}
