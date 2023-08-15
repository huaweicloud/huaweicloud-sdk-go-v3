package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStopInstanceResponse Response Object
type BatchStopInstanceResponse struct {

	// 任务列表对象。
	Jobs           *[]JobResult `json:"jobs,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchStopInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopInstanceResponse struct{}"
	}

	return strings.Join([]string{"BatchStopInstanceResponse", string(data)}, " ")
}
