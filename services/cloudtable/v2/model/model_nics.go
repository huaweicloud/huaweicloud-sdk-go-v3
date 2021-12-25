package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群所在的网络信息以及安全组信息。
type Nics struct {
	// CloudTable集群所在网络ID。

	NetId string `json:"net_id"`
	// CloudTable所在安全组对应的ID。

	SecurityGroupId string `json:"security_group_id"`
}

func (o Nics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nics struct{}"
	}

	return strings.Join([]string{"Nics", string(data)}, " ")
}
