package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCheckFeatureRuleResponse Response Object
type ListCheckFeatureRuleResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 检测特性规则列表
	DataList       *[]FeatureRuleInfo `json:"data_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListCheckFeatureRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCheckFeatureRuleResponse struct{}"
	}

	return strings.Join([]string{"ListCheckFeatureRuleResponse", string(data)}, " ")
}
