package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIndicatorRequestBody 删除指标请求参数
type DeleteIndicatorRequestBody struct {

	// 威胁情报ID列表
	BatchIds *[]string `json:"batch_ids,omitempty"`
}

func (o DeleteIndicatorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIndicatorRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteIndicatorRequestBody", string(data)}, " ")
}
