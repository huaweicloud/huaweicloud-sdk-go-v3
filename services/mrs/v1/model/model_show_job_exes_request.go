package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJobExesRequest struct {

	// 作业ID。
	JobExeId string `json:"job_exe_id"`
}

func (o ShowJobExesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobExesRequest struct{}"
	}

	return strings.Join([]string{"ShowJobExesRequest", string(data)}, " ")
}
