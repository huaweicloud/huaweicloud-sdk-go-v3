package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDataclassTypeRequestBody 启用、禁用类型的请求体
type EnableDataclassTypeRequestBody struct {

	// 是否启用,启用：true，禁用：false
	Enable bool `json:"enable"`

	// 类型ids
	Ids []string `json:"ids"`
}

func (o EnableDataclassTypeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDataclassTypeRequestBody struct{}"
	}

	return strings.Join([]string{"EnableDataclassTypeRequestBody", string(data)}, " ")
}
