package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobConfigResponse struct {

	// 作业保存条数
	JobRetainNumber *int32 `json:"job_retain_number,omitempty"`

	// 作业保存时长，单位天
	JobRetainTime  *int32 `json:"job_retain_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowJobConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowJobConfigResponse", string(data)}, " ")
}
