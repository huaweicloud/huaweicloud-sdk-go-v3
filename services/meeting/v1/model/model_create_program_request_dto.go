package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProgramRequestDto 节目信息。
type CreateProgramRequestDto struct {

	// 节目名称。
	ProgramName string `json:"programName"`

	// 节目素材列表。
	ProgramItemList *[]ProgramItemRequestBase `json:"programItemList,omitempty"`
}

func (o CreateProgramRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProgramRequestDto struct{}"
	}

	return strings.Join([]string{"CreateProgramRequestDto", string(data)}, " ")
}
