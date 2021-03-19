package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DisassociateDomainV2Request struct {
	InstanceId string `json:"instance_id"`

	DomainId string `json:"domain_id"`

	GroupId string `json:"group_id"`
}

func (o DisassociateDomainV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociateDomainV2Request struct{}"
	}

	return strings.Join([]string{"DisassociateDomainV2Request", string(data)}, " ")
}
