package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIterationHistoriesRequest Request Object
type ListIterationHistoriesRequest struct {

	// 迭代id
	IterationId int32 `json:"iteration_id"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量，最大为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListIterationHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIterationHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListIterationHistoriesRequest", string(data)}, " ")
}
