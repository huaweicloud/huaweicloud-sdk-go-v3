package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNextflowJobResponse Response Object
type ListNextflowJobResponse struct {

	// 作业列表
	Jobs *[]NextflowJobListDto `json:"jobs,omitempty"`

	// 作业总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNextflowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNextflowJobResponse struct{}"
	}

	return strings.Join([]string{"ListNextflowJobResponse", string(data)}, " ")
}
