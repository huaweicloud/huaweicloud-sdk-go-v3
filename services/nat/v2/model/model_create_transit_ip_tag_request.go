package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransitIpTagRequest Request Object
type CreateTransitIpTagRequest struct {

	// 中转IP的ID。
	ResourceId string `json:"resource_id"`

	Body *CreateResourceTagRequestBody `json:"body,omitempty"`
}

func (o CreateTransitIpTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitIpTagRequest struct{}"
	}

	return strings.Join([]string{"CreateTransitIpTagRequest", string(data)}, " ")
}
