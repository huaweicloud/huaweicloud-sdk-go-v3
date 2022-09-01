package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchStopInstanceResponse struct {

	// 任务列表对象。
	Jobs           *[]JobResult `json:"jobs,omitempty" xml:"jobs"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchStopInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopInstanceResponse struct{}"
	}

	return strings.Join([]string{"BatchStopInstanceResponse", string(data)}, " ")
}
