package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LeaveGroupResponse Response Object
type LeaveGroupResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o LeaveGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LeaveGroupResponse struct{}"
	}

	return strings.Join([]string{"LeaveGroupResponse", string(data)}, " ")
}
