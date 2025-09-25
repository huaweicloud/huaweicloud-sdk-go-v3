package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReduceDnRequestBody struct {

	// **参数解释**: 缩容分片数量。 **约束限制**: 缩容分片数量需要大于0且集群至少保留一个分片。 **取值范围**: 大于0的正整数。 **默认取值**: 不涉及。
	ContractionNum int32 `json:"contraction_num"`
}

func (o ReduceDnRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReduceDnRequestBody struct{}"
	}

	return strings.Join([]string{"ReduceDnRequestBody", string(data)}, " ")
}
