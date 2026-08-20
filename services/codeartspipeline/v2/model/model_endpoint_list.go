package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointList EndpointList具体信息
type EndpointList struct {

	// Endpoint的具体信息
	Endpoints *[]Endpoint `json:"endpoints,omitempty"`

	// 数量
	Total *int32 `json:"total,omitempty"`
}

func (o EndpointList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointList struct{}"
	}

	return strings.Join([]string{"EndpointList", string(data)}, " ")
}
