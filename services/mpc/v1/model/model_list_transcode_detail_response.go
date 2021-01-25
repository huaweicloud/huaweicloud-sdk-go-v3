package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListTranscodeDetailResponse struct {
	// 转码详情任务组
	TaskArray      *[]TaskDetailInfo `json:"task_array,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListTranscodeDetailResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTranscodeDetailResponse struct{}"
	}

	return strings.Join([]string{"ListTranscodeDetailResponse", string(data)}, " ")
}
