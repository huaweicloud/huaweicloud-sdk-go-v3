package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchStartInstanceResponse struct {

	// 任务列表对象。
	Jobs           *[]JobResult `json:"jobs,omitempty" xml:"jobs"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchStartInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartInstanceResponse struct{}"
	}

	return strings.Join([]string{"BatchStartInstanceResponse", string(data)}, " ")
}
