package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryJobsResponse Response Object
type ListFactoryJobsResponse struct {

	// 作业数量
	Total *int32 `json:"total,omitempty"`

	// 作业列表
	Jobs           *[]JobResp `json:"jobs,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListFactoryJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryJobsResponse struct{}"
	}

	return strings.Join([]string{"ListFactoryJobsResponse", string(data)}, " ")
}
