package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowClassroomDetailRequest struct {
	// 课堂ID

	ClassroomId string `json:"classroom_id"`
}

func (o ShowClassroomDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowClassroomDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowClassroomDetailRequest", string(data)}, " ")
}
