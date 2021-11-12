package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPoliciesV3Request struct {
	// 需要查询授权策略的通道名称。

	StreamName string `json:"stream_name"`
}

func (o ListPoliciesV3Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesV3Request struct{}"
	}

	return strings.Join([]string{"ListPoliciesV3Request", string(data)}, " ")
}
