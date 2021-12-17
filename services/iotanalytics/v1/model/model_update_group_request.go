package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateGroupRequest struct {
	// 存储组 ID

	GroupId string `json:"group_id"`

	Body *StorageGroup `json:"body,omitempty"`
}

func (o UpdateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupRequest", string(data)}, " ")
}
