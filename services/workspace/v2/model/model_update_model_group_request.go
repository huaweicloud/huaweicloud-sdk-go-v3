package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelGroupRequest Request Object
type UpdateModelGroupRequest struct {

	// 模型组id。
	GroupId string `json:"group_id"`

	Body *UpdateModelGroupReq `json:"body,omitempty"`
}

func (o UpdateModelGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateModelGroupRequest", string(data)}, " ")
}
