package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RunOption struct {

	// 并发数
	ThreadNum *int32 `json:"thread_num,omitempty"`

	BatchWindow *BatchWindow `json:"batch_window,omitempty"`
}

func (o RunOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunOption struct{}"
	}

	return strings.Join([]string{"RunOption", string(data)}, " ")
}
