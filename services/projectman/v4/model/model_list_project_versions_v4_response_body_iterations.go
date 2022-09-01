package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListProjectVersionsV4ResponseBodyIterations struct {

	// 迭代描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 迭代结束时间
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 迭代id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 迭代标题
	Name *string `json:"name,omitempty" xml:"name"`

	// 迭代开始时间
	BeginTime *string `json:"begin_time,omitempty" xml:"begin_time"`

	// 迭代状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 迭代更新时间，长整型时间戳
	UpdatedTime *int64 `json:"updated_time,omitempty" xml:"updated_time"`

	// 迭代是否已经删除，false, 未删除， true已经删除
	Deleted *bool `json:"deleted,omitempty" xml:"deleted"`
}

func (o ListProjectVersionsV4ResponseBodyIterations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectVersionsV4ResponseBodyIterations struct{}"
	}

	return strings.Join([]string{"ListProjectVersionsV4ResponseBodyIterations", string(data)}, " ")
}
