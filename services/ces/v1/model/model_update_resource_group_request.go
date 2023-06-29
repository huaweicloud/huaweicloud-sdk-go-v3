package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceGroupRequest Request Object
type UpdateResourceGroupRequest struct {

	// 资源分组ID。
	GroupId string `json:"group_id"`

	Body *UpdateResourceGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateResourceGroupRequest", string(data)}, " ")
}
