package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelSqlJobRequest Request Object
type CancelSqlJobRequest struct {

	// 作业ID
	JobId string `json:"job_id"`
}

func (o CancelSqlJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSqlJobRequest struct{}"
	}

	return strings.Join([]string{"CancelSqlJobRequest", string(data)}, " ")
}
