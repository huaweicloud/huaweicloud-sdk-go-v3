package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlJobDetailRequest Request Object
type ShowSqlJobDetailRequest struct {

	// 作业ID
	JobId string `json:"job_id"`
}

func (o ShowSqlJobDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlJobDetailRequest", string(data)}, " ")
}
