package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRestoreInstanceResponse Response Object
type CreateRestoreInstanceResponse struct {
	Instance *CreateInstanceRespItem `json:"instance,omitempty"`

	// 恢复新实例的任务id。  仅恢复按需实例时会返回该参数。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRestoreInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRestoreInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateRestoreInstanceResponse", string(data)}, " ")
}
