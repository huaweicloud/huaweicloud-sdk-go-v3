package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// swagger文档导入结果  暂不支持
type Swagger struct {

	// swagger文档编号
	Id *string `json:"id,omitempty"`

	// 导入结果说明
	Result *string `json:"result,omitempty"`
}

func (o Swagger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Swagger struct{}"
	}

	return strings.Join([]string{"Swagger", string(data)}, " ")
}
