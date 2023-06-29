package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProgramRequestDto 更新节目详情响应
type UpdateProgramRequestDto struct {

	// 节目名称。
	ProgramName string `json:"programName"`

	// 节目素材列表。
	ProgramItemList *[]ProgramItemRequestBase `json:"programItemList,omitempty"`
}

func (o UpdateProgramRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProgramRequestDto struct{}"
	}

	return strings.Join([]string{"UpdateProgramRequestDto", string(data)}, " ")
}
