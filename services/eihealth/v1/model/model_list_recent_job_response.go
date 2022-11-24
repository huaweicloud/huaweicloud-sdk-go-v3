package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRecentJobResponse struct {

	// 总个数
	Count *int64 `json:"count,omitempty"`

	// 作业列表
	Jobs           *[]JobRsp `json:"jobs,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRecentJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecentJobResponse struct{}"
	}

	return strings.Join([]string{"ListRecentJobResponse", string(data)}, " ")
}
