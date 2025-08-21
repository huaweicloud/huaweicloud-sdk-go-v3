package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConflictSectionDto struct {

	// 是否冲突
	Conflict *bool `json:"conflict,omitempty"`

	// 冲突行列表
	Lines *[]ConflictSectionLineDto `json:"lines,omitempty"`

	// 编号
	Id *string `json:"id,omitempty"`
}

func (o ConflictSectionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConflictSectionDto struct{}"
	}

	return strings.Join([]string{"ConflictSectionDto", string(data)}, " ")
}
