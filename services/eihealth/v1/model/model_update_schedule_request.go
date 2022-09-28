package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateScheduleRequest struct {

	// 计算资源id
	Id string `json:"id"`

	Body *UpdateScheduleReq `json:"body,omitempty"`
}

func (o UpdateScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduleRequest struct{}"
	}

	return strings.Join([]string{"UpdateScheduleRequest", string(data)}, " ")
}
