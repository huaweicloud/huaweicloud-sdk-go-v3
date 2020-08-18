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
type ListVpcRoutesRequest struct {
	Limit       *int32                   `json:"limit,omitempty"`
	Marker      *string                  `json:"marker,omitempty"`
	Id          *string                  `json:"id,omitempty"`
	Type        ListVpcRoutesRequestType `json:"type,omitempty"`
	VpcId       *string                  `json:"vpc_id,omitempty"`
	Destination *string                  `json:"destination,omitempty"`
	TenantId    *string                  `json:"tenant_id,omitempty"`
}

func (o ListVpcRoutesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListVpcRoutesRequest", string(data)}, " ")
}

type ListVpcRoutesRequestType struct {
	value string
}

type ListVpcRoutesRequestTypeEnum struct {
	PEERING ListVpcRoutesRequestType
}

func GetListVpcRoutesRequestTypeEnum() ListVpcRoutesRequestTypeEnum {
	return ListVpcRoutesRequestTypeEnum{
		PEERING: ListVpcRoutesRequestType{
			value: "peering",
		},
	}
}

func (c ListVpcRoutesRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListVpcRoutesRequestType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
