package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DisassociateDomainV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 域名的编号

	DomainId string `json:"domain_id"`
	// 分组的编号

	GroupId string `json:"group_id"`
}

func (o DisassociateDomainV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociateDomainV2Request struct{}"
	}

	return strings.Join([]string{"DisassociateDomainV2Request", string(data)}, " ")
}
