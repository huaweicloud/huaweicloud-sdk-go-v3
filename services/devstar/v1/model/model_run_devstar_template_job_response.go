package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunDevstarTemplateJobResponse struct {

	// 任务id。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 文件列表。
	FileList       *[]FileTreeNode `json:"file_list,omitempty" xml:"file_list"`
	HttpStatusCode int             `json:"-"`
}

func (o RunDevstarTemplateJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDevstarTemplateJobResponse struct{}"
	}

	return strings.Join([]string{"RunDevstarTemplateJobResponse", string(data)}, " ")
}
