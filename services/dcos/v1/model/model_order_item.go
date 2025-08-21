package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderItem 客户工单-列表值
type OrderItem struct {

	// 服务单号
	Number *string `json:"number,omitempty"`

	// 标题
	Title string `json:"title"`

	// 工单类型:IDC运维 设备运维 设备检查 客户陪同
	Type *string `json:"type,omitempty"`

	// 具体操作类型:设备物理上下电
	SubType *string `json:"sub_type,omitempty"`

	// 工单类型编码
	ModelCode string `json:"model_code"`

	// 机房编码
	RoomCode *string `json:"room_code,omitempty"`

	// 机房编码
	RoomName *string `json:"room_name,omitempty"`

	// 机房编码
	DcName *string `json:"dc_name,omitempty"`

	// 当前阶段.DRAFT草稿、SUMMITED服务申请、IN_PROGRESS服务处理中、ACCEPTANCE服务验收、CLOSED服务关闭
	Stage string `json:"stage"`

	// 当前状态
	Status *string `json:"status,omitempty"`

	// 申请人
	Applicant *string `json:"applicant,omitempty"`

	// 申请时间/创建时间
	CreateTime *string `json:"create_time,omitempty"`
}

func (o OrderItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderItem struct{}"
	}

	return strings.Join([]string{"OrderItem", string(data)}, " ")
}
