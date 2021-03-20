package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListBatchTaskFilesResponse struct {
	// 批量任务文件列表。

	Files          *[]BatchTaskFile `json:"files,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListBatchTaskFilesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBatchTaskFilesResponse struct{}"
	}

	return strings.Join([]string{"ListBatchTaskFilesResponse", string(data)}, " ")
}
