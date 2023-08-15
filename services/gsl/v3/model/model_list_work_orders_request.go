package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkOrdersRequest Request Object
type ListWorkOrdersRequest struct {

	// 业务受理ID
	MainSearchKey *string `json:"main_search_key,omitempty"`

	// 分页查询时每页显示的记录数，默认值为10，取值范围为10-500的整数
	Limit *int64 `json:"limit,omitempty"`

	// 分页查询时的页码数，默认值为1，取值范围为1-1000000的整数
	Offset *int64 `json:"offset,omitempty"`

	// SIM卡类型:3.实体卡
	SimType *int32 `json:"sim_type,omitempty"`

	// 业务受理类型：1.批量激活实体卡 2.批量转移实体卡 3.创建前向流量池 4.实体卡复机 5.实体卡停机 6.批量启用或复机 7.批量停机或停机 8.批量订购 9.批量退订 10.实体卡激活 11.申请断网 12.达量断网 13.机卡重绑 14.实名制信息清除 15.实体卡限速 16.批量补卡 17.批量机卡重绑 18.重启已废弃后向流量池 19.批量达量断网 20断网恢复 21取消达量断网 22批量取消达量断网 23批量拆机
	WorkOrderType *int32 `json:"work_order_type,omitempty"`

	// 业务受理状态:：1审核中、2已审核、3处理中、4已完成、5已取消、6失败、7 审核不通过
	Status *int32 `json:"status,omitempty"`
}

func (o ListWorkOrdersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkOrdersRequest struct{}"
	}

	return strings.Join([]string{"ListWorkOrdersRequest", string(data)}, " ")
}
