package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompoundMetricVo struct {

	// 编码
	Id *int64 `json:"id,omitempty"`

	// 字段名
	NameEn string `json:"name_en"`

	// 业务属性
	NameCh string `json:"name_ch"`

	Description *string `json:"description,omitempty"`

	// 颗粒度id
	DimensionGroup string `json:"dimension_group"`

	// 颗粒度名称
	GroupName *string `json:"group_name,omitempty"`

	// 颗粒度编码
	GroupCode *string `json:"group_code,omitempty"`

	// 指标信息
	MetricIds []int64 `json:"metric_ids"`

	// 指标名称信息
	MetricNames *[]string `json:"metric_names,omitempty"`

	// 引用函数id
	CalFnIds *[]int64 `json:"cal_fn_ids,omitempty"`

	// 计算表达式, ${index_id} + ${index_id}
	CalExp string `json:"cal_exp"`

	// 主题域分组id
	L1Id *int64 `json:"l1_id,omitempty"`

	L2Id *string `json:"l2_id,omitempty"`

	// 业务对象id
	L3Id *int64 `json:"l3_id,omitempty"`

	// 字段类型
	DataType *string `json:"data_type,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	Monitor *MetricMonitorVo `json:"monitor,omitempty"`

	// 主题域分组中文名
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名
	L3 *string `json:"l3,omitempty"`

	// 汇总表id
	SummaryTableId *int64 `json:"summary_table_id,omitempty"`
}

func (o CompoundMetricVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompoundMetricVo struct{}"
	}

	return strings.Join([]string{"CompoundMetricVo", string(data)}, " ")
}
