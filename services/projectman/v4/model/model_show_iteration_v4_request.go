package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowIterationV4Request struct {
	// 迭代id

	IterationId int32 `json:"iteration_id"`
}

func (o ShowIterationV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIterationV4Request struct{}"
	}

	return strings.Join([]string{"ShowIterationV4Request", string(data)}, " ")
}
