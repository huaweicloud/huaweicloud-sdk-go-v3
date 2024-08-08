package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGenJobResponse Response Object
type CreateGenJobResponse struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 限制的并发量
	LimitConcurrency *int32 `json:"limit_concurrency,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o CreateGenJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGenJobResponse struct{}"
	}

	return strings.Join([]string{"CreateGenJobResponse", string(data)}, " ")
}
