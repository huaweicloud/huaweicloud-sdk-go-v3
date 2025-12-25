package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceRequestBody 删除资产请求body体
type DeleteResourceRequestBody struct {

	// 删除资产ID列表
	BatchIds []string `json:"batch_ids"`
}

func (o DeleteResourceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteResourceRequestBody", string(data)}, " ")
}
