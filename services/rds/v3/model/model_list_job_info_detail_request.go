package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListJobInfoDetailRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 开始时间，格式为UTC时间戳。
	StartTime string `json:"start_time" xml:"start_time"`

	// 结束时间，格式为UTC时间戳。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`
}

func (o ListJobInfoDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobInfoDetailRequest struct{}"
	}

	return strings.Join([]string{"ListJobInfoDetailRequest", string(data)}, " ")
}
