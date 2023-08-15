package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListProjectVersionsV4ResponseBodyIterations struct {

	// 迭代描述
	Description *string `json:"description,omitempty"`

	// 迭代结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 迭代id
	Id *int32 `json:"id,omitempty"`

	// 迭代标题
	Name *string `json:"name,omitempty"`

	// 迭代开始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 迭代状态
	Status *string `json:"status,omitempty"`

	// 迭代更新时间，长整型时间戳
	UpdatedTime *int64 `json:"updated_time,omitempty"`

	// 迭代是否已经删除，false, 未删除， true已经删除
	Deleted *bool `json:"deleted,omitempty"`
}

func (o ListProjectVersionsV4ResponseBodyIterations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectVersionsV4ResponseBodyIterations struct{}"
	}

	return strings.Join([]string{"ListProjectVersionsV4ResponseBodyIterations", string(data)}, " ")
}
