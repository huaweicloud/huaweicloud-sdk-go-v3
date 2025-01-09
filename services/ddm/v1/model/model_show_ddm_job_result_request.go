package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdmJobResultRequest Request Object
type ShowDdmJobResultRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o ShowDdmJobResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdmJobResultRequest struct{}"
	}

	return strings.Join([]string{"ShowDdmJobResultRequest", string(data)}, " ")
}
