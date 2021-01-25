package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAnimatedGraphicsTaskResponse struct {
	// 任务总数
	Total *int32 `json:"total,omitempty"`
	// 任务列表
	Tasks          *[]AnimatedGraphicsTask `json:"tasks,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListAnimatedGraphicsTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAnimatedGraphicsTaskResponse struct{}"
	}

	return strings.Join([]string{"ListAnimatedGraphicsTaskResponse", string(data)}, " ")
}
