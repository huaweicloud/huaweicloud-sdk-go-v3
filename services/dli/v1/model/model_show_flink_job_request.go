package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkJobRequest Request Object
type ShowFlinkJobRequest struct {

	// 作业ID。
	JobId int64 `json:"job_id"`
}

func (o ShowFlinkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkJobRequest struct{}"
	}

	return strings.Join([]string{"ShowFlinkJobRequest", string(data)}, " ")
}
