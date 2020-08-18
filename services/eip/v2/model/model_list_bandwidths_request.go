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
type ListBandwidthsRequest struct {
	Marker              *string                        `json:"marker,omitempty"`
	Limit               *int32                         `json:"limit,omitempty"`
	EnterpriseProjectId *string                        `json:"enterprise_project_id,omitempty"`
	ShareType           ListBandwidthsRequestShareType `json:"share_type,omitempty"`
}

func (o ListBandwidthsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBandwidthsRequest", string(data)}, " ")
}

type ListBandwidthsRequestShareType struct {
	value string
}

type ListBandwidthsRequestShareTypeEnum struct {
	WHOLE ListBandwidthsRequestShareType
	PER   ListBandwidthsRequestShareType
}

func GetListBandwidthsRequestShareTypeEnum() ListBandwidthsRequestShareTypeEnum {
	return ListBandwidthsRequestShareTypeEnum{
		WHOLE: ListBandwidthsRequestShareType{
			value: "WHOLE",
		},
		PER: ListBandwidthsRequestShareType{
			value: "PER",
		},
	}
}

func (c ListBandwidthsRequestShareType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListBandwidthsRequestShareType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
