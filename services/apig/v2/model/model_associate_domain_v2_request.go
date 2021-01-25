package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AssociateDomainV2Request struct {
	ProjectId  string     `json:"project_id"`
	InstanceId string     `json:"instance_id"`
	GroupId    string     `json:"group_id"`
	Body       *DomainReq `json:"body,omitempty"`
}

func (o AssociateDomainV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateDomainV2Request struct{}"
	}

	return strings.Join([]string{"AssociateDomainV2Request", string(data)}, " ")
}
