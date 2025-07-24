package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IRack struct {

	// 机柜ID，资源的唯一标识
	Id *string `json:"id,omitempty"`

	// 机柜名称
	Name *string `json:"name,omitempty"`

	// 机房名称
	Dc *string `json:"dc,omitempty"`

	// 局点信息
	Region *string `json:"region,omitempty"`

	// 机柜位置
	Location *string `json:"location,omitempty"`

	// 机柜尺寸
	Size *string `json:"size,omitempty"`

	// 机柜U位
	Unit *string `json:"unit,omitempty"`

	// 机柜额定功率
	Power *string `json:"power,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 订单状态
	OrderStatus *string `json:"order_status,omitempty"`

	// 是否云化纳管柜
	IsCloudBased *string `json:"is_cloud_based,omitempty"`

	// 操作状态
	OperationStatus *int32 `json:"operation_status,omitempty"`

	// 冻结效果
	FreezeEffect *int32 `json:"freeze_effect,omitempty"`

	// 冻结场景
	Scene *string `json:"scene,omitempty"`
}

func (o IRack) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IRack struct{}"
	}

	return strings.Join([]string{"IRack", string(data)}, " ")
}
