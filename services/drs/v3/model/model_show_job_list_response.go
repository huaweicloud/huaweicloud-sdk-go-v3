package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobListResponse Response Object
type ShowJobListResponse struct {

	// 任务总数
	TotalRecord *int32 `json:"total_record,omitempty"`

	// 任务信息列表
	Jobs           *[]JobInfo `json:"jobs,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowJobListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobListResponse struct{}"
	}

	return strings.Join([]string{"ShowJobListResponse", string(data)}, " ")
}
