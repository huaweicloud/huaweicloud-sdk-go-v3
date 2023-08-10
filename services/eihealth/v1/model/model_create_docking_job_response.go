package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDockingJobResponse Response Object
type CreateDockingJobResponse struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 限制的并发量
	LimitConcurrency *int32 `json:"limit_concurrency,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o CreateDockingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDockingJobResponse struct{}"
	}

	return strings.Join([]string{"CreateDockingJobResponse", string(data)}, " ")
}
