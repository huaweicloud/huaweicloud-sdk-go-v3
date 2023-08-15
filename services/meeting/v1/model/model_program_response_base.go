package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProgramResponseBase 节目信息。
type ProgramResponseBase struct {

	// 节目ID。
	Id *string `json:"id,omitempty"`

	// 更新者。
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`

	// 更新时间。
	UpdateTime *int64 `json:"updateTime,omitempty"`

	// 节目名称。
	ProgramName *string `json:"programName,omitempty"`

	// 节目的总素材大小（含单位）。
	MaterialSizeStr *string `json:"materialSizeStr,omitempty"`

	// 节目的总播放时长，单位秒。
	PlayTime *int32 `json:"playTime,omitempty"`
}

func (o ProgramResponseBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProgramResponseBase struct{}"
	}

	return strings.Join([]string{"ProgramResponseBase", string(data)}, " ")
}
