package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag 资源标签
type Tag struct {

	// 参数解释:  键 示例: key 约束限制:  长度大于等于1且小于等于128 取值范围: 无 默认取值: 无
	Key string `json:"key"`

	// 参数解释:  值 示例: value 约束限制:  长度大于等于1且小于等于255 取值范围: 无 默认取值: 无
	Value string `json:"value"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
