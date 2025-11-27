package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceGroupResponse Response Object
type UpdateInstanceGroupResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateInstanceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceGroupResponse", string(data)}, " ")
}
