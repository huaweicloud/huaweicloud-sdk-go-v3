package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ThirdPartyAssociatedResultData 工作项关联外部链接查询结果数据集
type ThirdPartyAssociatedResultData struct {

	// 工作项关联外部链接查询结果数据集合
	Result *[]ThirdPartyAssociatedDto `json:"result,omitempty"`

	Page *PageVo `json:"page,omitempty"`
}

func (o ThirdPartyAssociatedResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThirdPartyAssociatedResultData struct{}"
	}

	return strings.Join([]string{"ThirdPartyAssociatedResultData", string(data)}, " ")
}
