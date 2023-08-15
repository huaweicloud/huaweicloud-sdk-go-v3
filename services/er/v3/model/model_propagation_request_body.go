package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PropagationRequestBody struct {

	// 连接唯一标识
	AttachmentId *string `json:"attachment_id,omitempty"`

	RoutePolicy *ImportRoutePolicy `json:"route_policy,omitempty"`
}

func (o PropagationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropagationRequestBody struct{}"
	}

	return strings.Join([]string{"PropagationRequestBody", string(data)}, " ")
}
