/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type ListVpcPeeringsRequest struct {
	Limit    int32                        `json:"limit,omitempty"`
	Marker   string                       `json:"marker,omitempty"`
	Id       string                       `json:"id,omitempty"`
	Name     string                       `json:"name,omitempty"`
	Status   ListVpcPeeringsRequestStatus `json:"status,omitempty"`
	TenantId string                       `json:"tenant_id,omitempty"`
	VpcId    string                       `json:"vpc_id,omitempty"`
}

func (o ListVpcPeeringsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListVpcPeeringsRequest", string(data)}, " ")
}

type ListVpcPeeringsRequestStatus struct {
	value string
}

type ListVpcPeeringsRequestStatusEnum struct {
	PENDING_ACCEPTANCE ListVpcPeeringsRequestStatus
	REJECTED           ListVpcPeeringsRequestStatus
	EXPIRED            ListVpcPeeringsRequestStatus
	DELETED            ListVpcPeeringsRequestStatus
	ACTIVE             ListVpcPeeringsRequestStatus
}

func GetListVpcPeeringsRequestStatusEnum() ListVpcPeeringsRequestStatusEnum {
	return ListVpcPeeringsRequestStatusEnum{
		PENDING_ACCEPTANCE: ListVpcPeeringsRequestStatus{
			value: "PENDING_ACCEPTANCE",
		},
		REJECTED: ListVpcPeeringsRequestStatus{
			value: "REJECTED",
		},
		EXPIRED: ListVpcPeeringsRequestStatus{
			value: "EXPIRED",
		},
		DELETED: ListVpcPeeringsRequestStatus{
			value: "DELETED",
		},
		ACTIVE: ListVpcPeeringsRequestStatus{
			value: "ACTIVE",
		},
	}
}

func (c ListVpcPeeringsRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListVpcPeeringsRequestStatus) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
