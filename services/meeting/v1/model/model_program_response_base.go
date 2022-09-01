package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节目响应信息
type ProgramResponseBase struct {

	// 节目ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 更新者
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty" xml:"lastUpdatedBy"`

	// 更新时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime"`

	// 节目名称
	ProgramName *string `json:"programName,omitempty" xml:"programName"`

	// 节目的总素材大小（含单位）
	MaterialSizeStr *string `json:"materialSizeStr,omitempty" xml:"materialSizeStr"`

	// 节目的总播放时长，单位秒
	PlayTime *int32 `json:"playTime,omitempty" xml:"playTime"`
}

func (o ProgramResponseBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProgramResponseBase struct{}"
	}

	return strings.Join([]string{"ProgramResponseBase", string(data)}, " ")
}
