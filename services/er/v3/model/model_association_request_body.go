package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 关联
type AssociationRequestBody struct {

	// 连接唯一标识
	AttachmentId *string `json:"attachment_id,omitempty"`

	RoutePolicy *ExportRoutePolicy `json:"route_policy,omitempty"`
}

func (o AssociationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociationRequestBody struct{}"
	}

	return strings.Join([]string{"AssociationRequestBody", string(data)}, " ")
}
