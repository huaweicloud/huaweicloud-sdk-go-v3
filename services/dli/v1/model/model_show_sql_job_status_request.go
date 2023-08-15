package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlJobStatusRequest Request Object
type ShowSqlJobStatusRequest struct {

	// 作业ID。
	JobId string `json:"job_id"`
}

func (o ShowSqlJobStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlJobStatusRequest", string(data)}, " ")
}
