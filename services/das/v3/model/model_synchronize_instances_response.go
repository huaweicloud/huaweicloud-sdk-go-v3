package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SynchronizeInstancesResponse Response Object
type SynchronizeInstancesResponse struct {

	// 任务创建是否成功。
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o SynchronizeInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SynchronizeInstancesResponse struct{}"
	}

	return strings.Join([]string{"SynchronizeInstancesResponse", string(data)}, " ")
}
