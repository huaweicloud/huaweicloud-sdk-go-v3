package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyRequestBody 删除应急策略请求body体
type DeletePolicyRequestBody struct {

	// 删除应急策略的ID列表
	BatchIds *[]string `json:"batch_ids,omitempty"`
}

func (o DeletePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"DeletePolicyRequestBody", string(data)}, " ")
}
