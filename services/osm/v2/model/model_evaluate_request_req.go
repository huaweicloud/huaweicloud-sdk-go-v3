package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EvaluateRequestReq struct {

	// 评价结果
	Degree *int32 `json:"degree,omitempty"`
}

func (o EvaluateRequestReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EvaluateRequestReq struct{}"
	}

	return strings.Join([]string{"EvaluateRequestReq", string(data)}, " ")
}
