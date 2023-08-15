package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSynthesisJobResponse Response Object
type CreateSynthesisJobResponse struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 限制的并发量
	LimitConcurrency *int32 `json:"limit_concurrency,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o CreateSynthesisJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSynthesisJobResponse struct{}"
	}

	return strings.Join([]string{"CreateSynthesisJobResponse", string(data)}, " ")
}
