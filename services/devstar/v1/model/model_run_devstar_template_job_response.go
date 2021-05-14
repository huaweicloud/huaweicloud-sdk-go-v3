package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RunDevstarTemplateJobResponse struct {
	// 任务id。

	JobId *string `json:"job_id,omitempty"`
	// 文件列表。

	FileList       *[]FileTreeNode `json:"file_list,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o RunDevstarTemplateJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunDevstarTemplateJobResponse struct{}"
	}

	return strings.Join([]string{"RunDevstarTemplateJobResponse", string(data)}, " ")
}
