/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type ListPublicipsRequest struct {
	Marker              *string                       `json:"marker,omitempty"`
	Limit               *int32                        `json:"limit,omitempty"`
	IpVersion           ListPublicipsRequestIpVersion `json:"ip_version,omitempty"`
	EnterpriseProjectId *string                       `json:"enterprise_project_id,omitempty"`
	PortId              *string                       `json:"port_id,omitempty"`
	PublicIpAddress     *string                       `json:"public_ip_address,omitempty"`
	PrivateIpAddress    *string                       `json:"private_ip_address,omitempty"`
	Id                  *string                       `json:"id,omitempty"`
}

func (o ListPublicipsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPublicipsRequest", string(data)}, " ")
}

type ListPublicipsRequestIpVersion struct {
	value int32
}

type ListPublicipsRequestIpVersionEnum struct {
	_4 ListPublicipsRequestIpVersion
	_6 ListPublicipsRequestIpVersion
}

func GetListPublicipsRequestIpVersionEnum() ListPublicipsRequestIpVersionEnum {
	return ListPublicipsRequestIpVersionEnum{
		_4: ListPublicipsRequestIpVersion{
			value: 4,
		}, _6: ListPublicipsRequestIpVersion{
			value: 6,
		},
	}
}

func (c ListPublicipsRequestIpVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListPublicipsRequestIpVersion) UnmarshalJSON(b []byte) error {
	c.value = int32(strings.Trim(string(b[:]), "\""))
	return nil
}
