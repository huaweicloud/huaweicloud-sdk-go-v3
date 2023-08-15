package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateDomainV2Request Request Object
type DisassociateDomainV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id"`

	// 域名的编号
	DomainId string `json:"domain_id"`
}

func (o DisassociateDomainV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateDomainV2Request struct{}"
	}

	return strings.Join([]string{"DisassociateDomainV2Request", string(data)}, " ")
}
