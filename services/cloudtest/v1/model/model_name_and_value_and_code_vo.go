package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NameAndValueAndCodeVo 用户自定义结果对应的用例数目
type NameAndValueAndCodeVo struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 值
	Value *int32 `json:"value,omitempty"`

	// 编码
	Code *string `json:"code,omitempty"`
}

func (o NameAndValueAndCodeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NameAndValueAndCodeVo struct{}"
	}

	return strings.Join([]string{"NameAndValueAndCodeVo", string(data)}, " ")
}
