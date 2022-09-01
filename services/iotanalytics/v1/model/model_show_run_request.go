package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRunRequest struct {

	// 作业ID。
	JobId string `json:"job_id" xml:"job_id"`

	// 作业运行ID。
	RunId string `json:"run_id" xml:"run_id"`

	// 是否查询作业详情。true：查询；false：不查询。details属性为空。默认为false。
	WithDetails *bool `json:"with_details,omitempty" xml:"with_details"`
}

func (o ShowRunRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRunRequest struct{}"
	}

	return strings.Join([]string{"ShowRunRequest", string(data)}, " ")
}
