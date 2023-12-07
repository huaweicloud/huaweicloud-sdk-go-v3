package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NameAndValueVo 用户自定义状态对应的用例数目
type NameAndValueVo struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 值
	Value *int32 `json:"value,omitempty"`
}

func (o NameAndValueVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NameAndValueVo struct{}"
	}

	return strings.Join([]string{"NameAndValueVo", string(data)}, " ")
}
