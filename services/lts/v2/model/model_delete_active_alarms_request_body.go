package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteActiveAlarmsRequestBody struct {

	// 主题信息
	Events []Event `json:"events"`
}

func (o DeleteActiveAlarmsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteActiveAlarmsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteActiveAlarmsRequestBody", string(data)}, " ")
}
