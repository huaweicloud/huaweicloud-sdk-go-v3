package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConflictSectionLineMetaDataDto struct {

	// 旧列位置
	OldPos *int32 `json:"old_pos,omitempty"`

	// 新列位置
	NewPos *int32 `json:"new_pos,omitempty"`
}

func (o ConflictSectionLineMetaDataDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConflictSectionLineMetaDataDto struct{}"
	}

	return strings.Join([]string{"ConflictSectionLineMetaDataDto", string(data)}, " ")
}
