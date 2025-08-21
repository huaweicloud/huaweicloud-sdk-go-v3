package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterChangeMainSubnet struct {

	// **参数解释**： 子网ID。 **约束限制**： 格式为标准的UUID格式：xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx **取值范围**： 格式为标准的UUID格式：xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx **默认取值**： 不涉及
	SubnetId string `json:"subnet_id"`
}

func (o ClusterChangeMainSubnet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterChangeMainSubnet struct{}"
	}

	return strings.Join([]string{"ClusterChangeMainSubnet", string(data)}, " ")
}
