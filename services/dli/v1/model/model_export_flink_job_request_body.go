package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFlinkJobRequestBody 导出作业的请求参数。
type ExportFlinkJobRequestBody struct {

	// 导出作业文件的OBS保存路径。
	ObsDir string `json:"obs_dir"`

	// 是否导出指定的作业。
	IsSelected bool `json:"is_selected"`

	// 当is_selected=true时，该参数是待导出作业的作业ID集合。当is_selected=true时，此参数必填。
	JobSelected *[]int64 `json:"job_selected,omitempty"`
}

func (o ExportFlinkJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFlinkJobRequestBody struct{}"
	}

	return strings.Join([]string{"ExportFlinkJobRequestBody", string(data)}, " ")
}
