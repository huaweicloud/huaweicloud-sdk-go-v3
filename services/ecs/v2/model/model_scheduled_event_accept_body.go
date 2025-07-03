package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScheduledEventAcceptBody struct {

	// 计划执行开始时间
	NotBefore *string `json:"not_before,omitempty"`
}

func (o ScheduledEventAcceptBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledEventAcceptBody struct{}"
	}

	return strings.Join([]string{"ScheduledEventAcceptBody", string(data)}, " ")
}
