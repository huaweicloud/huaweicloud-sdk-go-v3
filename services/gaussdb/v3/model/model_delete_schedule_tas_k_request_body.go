package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteScheduleTasKRequestBody struct {

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o DeleteScheduleTasKRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduleTasKRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteScheduleTasKRequestBody", string(data)}, " ")
}
