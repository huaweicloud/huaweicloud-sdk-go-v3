package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTransitSubnetTagRequest Request Object
type DeleteTransitSubnetTagRequest struct {

	// 标签key。
	Key string `json:"key"`

	// 中转子网的ID。
	ResourceId string `json:"resource_id"`
}

func (o DeleteTransitSubnetTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransitSubnetTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteTransitSubnetTagRequest", string(data)}, " ")
}
