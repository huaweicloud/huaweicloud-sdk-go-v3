package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugModelResponse Response Object
type CreateDrugModelResponse struct {

	// 模型id
	Id *string `json:"id,omitempty"`

	// 限制的并发量
	LimitConcurrency *int32 `json:"limit_concurrency,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o CreateDrugModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugModelResponse struct{}"
	}

	return strings.Join([]string{"CreateDrugModelResponse", string(data)}, " ")
}
