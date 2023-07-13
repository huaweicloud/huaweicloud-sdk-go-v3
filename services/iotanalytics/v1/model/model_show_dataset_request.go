package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatasetRequest Request Object
type ShowDatasetRequest struct {

	// 作业ID。
	JobId string `json:"job_id"`

	// 作业运行ID。
	RunId string `json:"run_id"`

	// 当前偏移量，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的最大作业个数，范围: [1, 100]。默认值：10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowDatasetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatasetRequest struct{}"
	}

	return strings.Join([]string{"ShowDatasetRequest", string(data)}, " ")
}
