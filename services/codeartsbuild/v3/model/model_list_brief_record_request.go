package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBriefRecordRequest Request Object
type ListBriefRecordRequest struct {

	// 每页显示的条目数量，limit小于等于100
	Limit int32 `json:"limit"`

	// build_project_id，构建工程ID，36位UUID
	Body *[]string `json:"body,omitempty"`
}

func (o ListBriefRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBriefRecordRequest struct{}"
	}

	return strings.Join([]string{"ListBriefRecordRequest", string(data)}, " ")
}
