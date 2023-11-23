package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlJobProgressRequest Request Object
type ShowSqlJobProgressRequest struct {

	// 作业ID
	JobId string `json:"job_id"`
}

func (o ShowSqlJobProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlJobProgressRequest", string(data)}, " ")
}
