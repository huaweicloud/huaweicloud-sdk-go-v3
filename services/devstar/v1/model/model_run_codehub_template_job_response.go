package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunCodehubTemplateJobResponse struct {

	// 任务id。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 文件列表。
	FileList       *[]FileTreeNode `json:"file_list,omitempty" xml:"file_list"`
	HttpStatusCode int             `json:"-"`
}

func (o RunCodehubTemplateJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCodehubTemplateJobResponse struct{}"
	}

	return strings.Join([]string{"RunCodehubTemplateJobResponse", string(data)}, " ")
}
