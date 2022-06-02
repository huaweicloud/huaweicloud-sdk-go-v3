package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListWorkitemStatusRecordsV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 偏移量 从0开始
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量 最小1,最大100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListWorkitemStatusRecordsV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkitemStatusRecordsV4Request struct{}"
	}

	return strings.Join([]string{"ListWorkitemStatusRecordsV4Request", string(data)}, " ")
}
