package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePocketMolDesignJobResponse Response Object
type CreatePocketMolDesignJobResponse struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 限制的并发量
	LimitConcurrency *int32 `json:"limit_concurrency,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o CreatePocketMolDesignJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePocketMolDesignJobResponse struct{}"
	}

	return strings.Join([]string{"CreatePocketMolDesignJobResponse", string(data)}, " ")
}
