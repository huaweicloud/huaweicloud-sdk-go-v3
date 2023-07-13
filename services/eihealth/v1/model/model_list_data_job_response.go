package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataJobResponse Response Object
type ListDataJobResponse struct {

	// 总条数
	Count *int32 `json:"count,omitempty"`

	// 数据作业列表
	DataJobs       *[]DataJobRsp `json:"data_jobs,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDataJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataJobResponse struct{}"
	}

	return strings.Join([]string{"ListDataJobResponse", string(data)}, " ")
}
