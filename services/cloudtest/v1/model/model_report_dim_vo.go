package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportDimVo 对比维度数据
type ReportDimVo struct {

	// id
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 值
	Value *int32 `json:"value,omitempty"`
}

func (o ReportDimVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportDimVo struct{}"
	}

	return strings.Join([]string{"ReportDimVo", string(data)}, " ")
}
