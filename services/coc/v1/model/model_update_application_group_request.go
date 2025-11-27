package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationGroupRequest Request Object
type UpdateApplicationGroupRequest struct {

	// 分组ID。
	Id string `json:"id"`

	Body *GroupUpdateRequest `json:"body,omitempty"`
}

func (o UpdateApplicationGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateApplicationGroupRequest", string(data)}, " ")
}
