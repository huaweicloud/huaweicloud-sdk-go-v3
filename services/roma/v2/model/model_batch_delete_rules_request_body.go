package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteRulesRequestBody struct {

	// 规则ID列表，自动向下取整
	Resources []int32 `json:"resources" xml:"resources"`
}

func (o BatchDeleteRulesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRulesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteRulesRequestBody", string(data)}, " ")
}
