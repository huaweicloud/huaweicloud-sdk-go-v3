package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DashboardDto struct {

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 数据类型：0=用例成功率；1=用例时长
	DataType *string `json:"data_type,omitempty"`

	// 唯一ID，主键
	Id *string `json:"id,omitempty"`

	// 看板标题
	Name *string `json:"name,omitempty"`

	// 服务ID
	ServiceId *string `json:"service_id,omitempty"`

	// 任务ID列表
	TaskIds *[]string `json:"task_ids,omitempty"`

	// 任务类型，仅支持持续拨测和冒烟测试
	TaskType *string `json:"task_type,omitempty"`

	// 修改时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 修改者
	UpdateUser *string `json:"update_user,omitempty"`

	// 看板类型：0=折线图；1=散点图；2=饼图
	ViewType *string `json:"view_type,omitempty"`
}

func (o DashboardDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DashboardDto struct{}"
	}

	return strings.Join([]string{"DashboardDto", string(data)}, " ")
}
