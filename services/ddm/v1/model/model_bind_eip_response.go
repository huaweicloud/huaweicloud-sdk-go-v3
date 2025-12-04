package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindEipResponse Response Object
type BindEipResponse struct {

	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BindEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindEipResponse struct{}"
	}

	return strings.Join([]string{"BindEipResponse", string(data)}, " ")
}
