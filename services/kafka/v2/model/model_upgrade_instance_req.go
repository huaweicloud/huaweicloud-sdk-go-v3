package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradeInstanceReq struct {

	// 是否作为定时任务执行。若非定时执行，is_schedule和execute_at字段可为空。若为定时执行，is_schedule为true，execute_at字段非空。
	IsSchedule *bool `json:"is_schedule,omitempty"`

	// 定时时间，格式为Unix时间戳，单位为毫秒
	ExecuteAt *int64 `json:"execute_at,omitempty"`
}

func (o UpgradeInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceReq struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceReq", string(data)}, " ")
}
