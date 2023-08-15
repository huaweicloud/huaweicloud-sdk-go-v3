package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateNatTagRequest Request Object
type CreatePrivateNatTagRequest struct {

	// 私网NAT网关的ID。
	ResourceId string `json:"resource_id"`

	Body *CreateResourceTagRequestBody `json:"body,omitempty"`
}

func (o CreatePrivateNatTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateNatTagRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivateNatTagRequest", string(data)}, " ")
}
