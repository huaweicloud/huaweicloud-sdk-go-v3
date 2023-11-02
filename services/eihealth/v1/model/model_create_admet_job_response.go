package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAdmetJobResponse Response Object
type CreateAdmetJobResponse struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 限制的并发量
	LimitConcurrency *int32 `json:"limit_concurrency,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o CreateAdmetJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAdmetJobResponse struct{}"
	}

	return strings.Join([]string{"CreateAdmetJobResponse", string(data)}, " ")
}
