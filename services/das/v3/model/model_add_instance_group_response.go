package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddInstanceGroupResponse Response Object
type AddInstanceGroupResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 实例组ID
	GroupId        *int32 `json:"group_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o AddInstanceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddInstanceGroupResponse struct{}"
	}

	return strings.Join([]string{"AddInstanceGroupResponse", string(data)}, " ")
}
