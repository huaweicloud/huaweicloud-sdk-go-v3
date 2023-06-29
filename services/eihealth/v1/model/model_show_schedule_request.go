package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScheduleRequest Request Object
type ShowScheduleRequest struct {

	// 计算资源id
	Id string `json:"id"`
}

func (o ShowScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScheduleRequest struct{}"
	}

	return strings.Join([]string{"ShowScheduleRequest", string(data)}, " ")
}
