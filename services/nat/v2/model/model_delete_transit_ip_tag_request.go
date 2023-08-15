package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTransitIpTagRequest Request Object
type DeleteTransitIpTagRequest struct {

	// 标签key。
	Key string `json:"key"`

	// 中转IP的ID。
	ResourceId string `json:"resource_id"`
}

func (o DeleteTransitIpTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransitIpTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteTransitIpTagRequest", string(data)}, " ")
}
