package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkCapability 租户能力详情
type CentralNetworkCapability struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	Capability *CentralNetworkCapabilityEnum `json:"capability"`
}

func (o CentralNetworkCapability) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkCapability struct{}"
	}

	return strings.Join([]string{"CentralNetworkCapability", string(data)}, " ")
}
