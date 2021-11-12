package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RetryInfo struct {
	// 任务ID

	JobId string `json:"job_id"`
	// 再编辑之后启动，必填为true。

	IsSyncReEdit *bool `json:"is_sync_re_edit,omitempty"`
}

func (o RetryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryInfo struct{}"
	}

	return strings.Join([]string{"RetryInfo", string(data)}, " ")
}
