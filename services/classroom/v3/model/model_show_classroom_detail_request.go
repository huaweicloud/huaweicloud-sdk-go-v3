package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowClassroomDetailRequest struct {

	// 课堂ID
	ClassroomId string `json:"classroom_id" xml:"classroom_id"`
}

func (o ShowClassroomDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClassroomDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowClassroomDetailRequest", string(data)}, " ")
}
