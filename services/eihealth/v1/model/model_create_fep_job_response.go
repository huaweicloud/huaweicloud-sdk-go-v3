package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFepJobResponse Response Object
type CreateFepJobResponse struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 限制的并发量
	LimitConcurrency *int32 `json:"limit_concurrency,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o CreateFepJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFepJobResponse struct{}"
	}

	return strings.Join([]string{"CreateFepJobResponse", string(data)}, " ")
}
