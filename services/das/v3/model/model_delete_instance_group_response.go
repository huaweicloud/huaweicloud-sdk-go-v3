package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceGroupResponse Response Object
type DeleteInstanceGroupResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteInstanceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceGroupResponse", string(data)}, " ")
}
