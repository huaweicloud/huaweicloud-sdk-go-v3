package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchRebootInstanceResponse struct {

	// 任务列表对象。
	Jobs           *[]JobResult `json:"jobs,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchRebootInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootInstanceResponse struct{}"
	}

	return strings.Join([]string{"BatchRebootInstanceResponse", string(data)}, " ")
}
