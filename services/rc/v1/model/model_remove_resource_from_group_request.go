package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveResourceFromGroupRequest Request Object
type RemoveResourceFromGroupRequest struct {

	// 资源组ID
	GroupId string `json:"group_id"`

	// 资源ID
	ResourceId string `json:"resource_id"`
}

func (o RemoveResourceFromGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveResourceFromGroupRequest struct{}"
	}

	return strings.Join([]string{"RemoveResourceFromGroupRequest", string(data)}, " ")
}
