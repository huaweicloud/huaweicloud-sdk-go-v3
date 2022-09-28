package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateScheduleReq struct {

	// 是否可调度
	Schedulable bool `json:"schedulable"`
}

func (o UpdateScheduleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduleReq struct{}"
	}

	return strings.Join([]string{"UpdateScheduleReq", string(data)}, " ")
}
