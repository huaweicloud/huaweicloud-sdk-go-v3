package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskTableReferenceDetailResponse Response Object
type ListTaskTableReferenceDetailResponse struct {

	// 引用的作业数量。
	Total *int32 `json:"total,omitempty"`

	// 引用作业的详情。
	TaskTableDetailList *[]TaskTableReferenceDetailResponse `json:"task_table_detail_list,omitempty"`
	HttpStatusCode      int                                 `json:"-"`
}

func (o ListTaskTableReferenceDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskTableReferenceDetailResponse struct{}"
	}

	return strings.Join([]string{"ListTaskTableReferenceDetailResponse", string(data)}, " ")
}
