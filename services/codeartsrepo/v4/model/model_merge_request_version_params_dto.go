package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MergeRequestVersionParamsDto 合并请求版本详情。
type MergeRequestVersionParamsDto struct {

	// **参数解释：** 版本id(由合并请求id和sha值共同组成)。
	DiffId *int32 `json:"diff_id,omitempty"`

	// **参数解释：** 起始sha值。
	StartSha *string `json:"start_sha,omitempty"`

	// **参数解释：** 提交记录id。
	CommitId *string `json:"commit_id,omitempty"`
}

func (o MergeRequestVersionParamsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestVersionParamsDto struct{}"
	}

	return strings.Join([]string{"MergeRequestVersionParamsDto", string(data)}, " ")
}
