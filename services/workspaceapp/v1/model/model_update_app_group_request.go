package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppGroupRequest Request Object
type UpdateAppGroupRequest struct {

	// 应用组ID
	AppGroupId string `json:"app_group_id"`

	Body *UpdateAppGroupReq `json:"body,omitempty"`
}

func (o UpdateAppGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppGroupRequest", string(data)}, " ")
}
