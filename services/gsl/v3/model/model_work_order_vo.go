package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkOrderVo struct {

	// 业务受理ID
	Id *int64 `json:"id,omitempty"`

	// SIM卡类型:3.实体卡
	SimType *int32 `json:"sim_type,omitempty"`

	// 业务受理类型：1.批量激活实体卡 2.批量转移实体卡 3.创建前向流量池 4.实体卡复机 5.实体卡停机 6.批量启用或复机 7.批量停机或停机 8.批量订购 9.批量退订 10.实体卡激活 11.申请断网 12.达量断网 13.机卡重绑 14.实名制信息清除 15.实体卡限速 16.批量补卡 17.批量机卡重绑 18.重启已废弃后向流量池 19.批量达量断网 20断网恢复 21取消达量断网 22批量取消达量断网 23批量拆机
	WorkOrderType *int32 `json:"work_order_type,omitempty"`

	// 请求详情
	ReqDetail *string `json:"req_detail,omitempty"`

	// 业务受理明细总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 业务受理明细成功数
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 业务受理明细失败数
	FailCount *int32 `json:"fail_count,omitempty"`

	// 业务受理明细处理中数
	ProcessCount *int32 `json:"process_count,omitempty"`

	// 业务受理状态：1审核中、2已审核、3处理中、4已完成、5已取消、6失败、7 审核不通过
	Status *int32 `json:"status,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 完成时间
	FinishTime *sdktime.SdkTime `json:"finish_time,omitempty"`

	// 失败原因
	FailReason *string `json:"fail_reason,omitempty"`

	// 响应内容
	Response *string `json:"response,omitempty"`

	// 业务受理单来源,1:运营人员生成,2:用户操作生成(console),3:自动化规则生成,4:后向流量池超阈值停用次月自动复机任务,5:单卡没流量停机定时任务,6:SIM卡到期自动停机定时任务,7:流量池停机定时任务,8:用户操作生成(api)
	WorkOrderSource *int32 `json:"work_order_source,omitempty"`

	// 业务受理单来源描述
	WorkOrderSourceDesc *string `json:"work_order_source_desc,omitempty"`
}

func (o WorkOrderVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkOrderVo struct{}"
	}

	return strings.Join([]string{"WorkOrderVo", string(data)}, " ")
}
