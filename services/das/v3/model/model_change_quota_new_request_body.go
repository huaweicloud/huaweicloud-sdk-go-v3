package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeQuotaNewRequestBody 修改配额请求体
type ChangeQuotaNewRequestBody struct {

	// 修改配额的数量
	OpenNum *int32 `json:"open_num,omitempty"`
}

func (o ChangeQuotaNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeQuotaNewRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeQuotaNewRequestBody", string(data)}, " ")
}
