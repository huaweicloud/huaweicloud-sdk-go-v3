package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateDefectInfoVo 关联缺陷信息
type AssociateDefectInfoVo struct {

	// 是否已关联
	Associate *bool `json:"associate,omitempty"`

	// 关联缺陷数
	AssociateCount *int32 `json:"associate_count,omitempty"`
}

func (o AssociateDefectInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateDefectInfoVo struct{}"
	}

	return strings.Join([]string{"AssociateDefectInfoVo", string(data)}, " ")
}
