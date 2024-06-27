package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmRestoredDataResponse Response Object
type ConfirmRestoredDataResponse struct {

	// 任务id,可以通过接口查询任务信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ConfirmRestoredDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmRestoredDataResponse struct{}"
	}

	return strings.Join([]string{"ConfirmRestoredDataResponse", string(data)}, " ")
}
