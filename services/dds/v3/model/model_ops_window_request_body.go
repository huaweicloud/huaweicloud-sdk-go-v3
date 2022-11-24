package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpsWindowRequestBody struct {

	// 开始时间，格式必须为HH:MM且有效，当前时间指UTC时间。不能与结束时间相同，只能为整点。
	StartTime string `json:"start_time"`

	// 结束时间，格式必须为HH:MM且有效，当前时间指UTC时间。不能与开始时间相同，只能为整点。
	EndTime string `json:"end_time"`
}

func (o OpsWindowRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpsWindowRequestBody struct{}"
	}

	return strings.Join([]string{"OpsWindowRequestBody", string(data)}, " ")
}
