package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEndpointVpcsRequest struct {

	// 待查询的vpc的ID。
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o ListEndpointVpcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointVpcsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointVpcsRequest", string(data)}, " ")
}
