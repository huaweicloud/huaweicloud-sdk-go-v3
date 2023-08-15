package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateNatTagRequest Request Object
type DeletePrivateNatTagRequest struct {

	// 标签key。
	Key string `json:"key"`

	// 私网NAT网关的ID。
	ResourceId string `json:"resource_id"`
}

func (o DeletePrivateNatTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateNatTagRequest struct{}"
	}

	return strings.Join([]string{"DeletePrivateNatTagRequest", string(data)}, " ")
}
