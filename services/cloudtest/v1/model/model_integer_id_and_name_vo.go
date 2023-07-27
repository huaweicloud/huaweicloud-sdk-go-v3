package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IntegerIdAndNameVo 测试类型信息
type IntegerIdAndNameVo struct {

	// 数据库存储数字
	Id *int32 `json:"id,omitempty"`

	// 页面显示值
	Name *string `json:"name,omitempty"`
}

func (o IntegerIdAndNameVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntegerIdAndNameVo struct{}"
	}

	return strings.Join([]string{"IntegerIdAndNameVo", string(data)}, " ")
}
