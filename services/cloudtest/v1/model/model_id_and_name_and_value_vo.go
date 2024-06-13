package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdAndNameAndValueVo 缺陷按照模块分布情况
type IdAndNameAndValueVo struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 值
	Value *int32 `json:"value,omitempty"`

	// id
	Id *string `json:"id,omitempty"`
}

func (o IdAndNameAndValueVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdAndNameAndValueVo struct{}"
	}

	return strings.Join([]string{"IdAndNameAndValueVo", string(data)}, " ")
}
