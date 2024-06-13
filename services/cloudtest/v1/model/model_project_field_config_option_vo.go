package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectFieldConfigOptionVo 实际的数据类型：单个对象，集合 或 NULL
type ProjectFieldConfigOptionVo struct {

	// 字段选项URI标识.新增不传，修改、删除使用必传
	Uri *string `json:"uri,omitempty"`

	// 可选项名称
	Name *string `json:"name,omitempty"`

	// 可选项code值
	Code *string `json:"code,omitempty"`

	Updator *NameAndIdVo `json:"updator,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 修改标识，0：不可修改 1：可修改，用于结果和状态的选项值
	Flag *int32 `json:"flag,omitempty"`

	// 顺序数值
	SortNumb *int32 `json:"sort_numb,omitempty"`

	Creator *NameAndIdVo `json:"creator,omitempty"`

	// 创建时间时间戳
	CreateTimeStamp *int64 `json:"create_time_stamp,omitempty"`

	// 更新时间时间戳
	UpdateTimeStamp *int64 `json:"update_time_stamp,omitempty"`
}

func (o ProjectFieldConfigOptionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectFieldConfigOptionVo struct{}"
	}

	return strings.Join([]string{"ProjectFieldConfigOptionVo", string(data)}, " ")
}
