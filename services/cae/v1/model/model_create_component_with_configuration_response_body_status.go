package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateComponentWithConfigurationResponseBodyStatus struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o CreateComponentWithConfigurationResponseBodyStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentWithConfigurationResponseBodyStatus struct{}"
	}

	return strings.Join([]string{"CreateComponentWithConfigurationResponseBodyStatus", string(data)}, " ")
}
