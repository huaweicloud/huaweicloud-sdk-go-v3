package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量数据加工规则请求体
type BatchDataTransformationReq struct {
	// 批量数据加工规则请求列表

	Jobs []CheckDataTransformationReq `json:"jobs"`
}

func (o BatchDataTransformationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDataTransformationReq struct{}"
	}

	return strings.Join([]string{"BatchDataTransformationReq", string(data)}, " ")
}
