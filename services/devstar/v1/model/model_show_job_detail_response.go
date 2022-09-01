package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobDetailResponse struct {

	// 任务的id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 任务的名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 任务的状态。
	JobStatus *interface{} `json:"job_status,omitempty" xml:"job_status"`

	// 任务结果信息。
	JobResult *string `json:"job_result,omitempty" xml:"job_result"`

	// 任务显示类型，页面显示使用字段
	ShowType       *string `json:"show_type,omitempty" xml:"show_type"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowJobDetailResponse", string(data)}, " ")
}
