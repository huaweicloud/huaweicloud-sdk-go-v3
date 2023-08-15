package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStudyJobResponse Response Object
type ListStudyJobResponse struct {

	// 作业总数
	Count *int32 `json:"count,omitempty"`

	// 作业列表
	Jobs           *[]StudyJobRsp `json:"jobs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListStudyJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStudyJobResponse struct{}"
	}

	return strings.Join([]string{"ListStudyJobResponse", string(data)}, " ")
}
