package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProgramResponse struct {
	// 节目ID

	Id *string `json:"id,omitempty"`
	// 更新者

	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// 更新时间

	UpdateTime *int64 `json:"updateTime,omitempty"`
	// 节目名称

	ProgramName *string `json:"programName,omitempty"`
	// 节目的总素材大小（含单位）

	MaterialSizeStr *string `json:"materialSizeStr,omitempty"`
	// 节目的总播放时长，单位秒

	PlayTime *int32 `json:"playTime,omitempty"`
	// 节目素材列表

	ProgramItemList *[]ProgramItemResponseBase `json:"programItemList,omitempty"`
	HttpStatusCode  int                        `json:"-"`
}

func (o ShowProgramResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProgramResponse struct{}"
	}

	return strings.Join([]string{"ShowProgramResponse", string(data)}, " ")
}
