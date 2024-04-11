package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompoundMetricVo struct {

	// 编码。
	Id *int64 `json:"id,omitempty"`

	// 字段名。
	NameEn string `json:"name_en"`

	// 业务属性。
	NameCh string `json:"name_ch"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 颗粒度ID。
	DimensionGroup string `json:"dimension_group"`

	// 颗粒度名称。
	GroupName *string `json:"group_name,omitempty"`

	// 颗粒度编码。
	GroupCode *string `json:"group_code,omitempty"`

	// 指标信息。
	MetricIds []int64 `json:"metric_ids"`

	// 指标名称信息。
	MetricNames *[]string `json:"metric_names,omitempty"`

	// 引用函数ID。
	CalFnIds *[]int64 `json:"cal_fn_ids,omitempty"`

	// 计算表达式，形如${index_id} + ${compound#index_id}，其中index_id代表引用的衍生指标ID，compound#index_id代表引用的复合指标ID。
	CalExp string `json:"cal_exp"`

	// 主题域分组ID。
	L1Id *int64 `json:"l1_id,omitempty"`

	// 主题域ID，只读，创建和更新时无需填写。
	L2Id *string `json:"l2_id,omitempty"`

	// 业务对象ID。
	L3Id *int64 `json:"l3_id,omitempty"`

	// 字段类型。
	DataType *string `json:"data_type,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 创建时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	Monitor *MetricMonitorVo `json:"monitor,omitempty"`

	// 主题域分组中文名，只读，创建和更新时无需填写。
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名，只读，创建和更新时无需填写。
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名，只读，创建和更新时无需填写。
	L3 *string `json:"l3,omitempty"`

	// 汇总表ID。
	SummaryTableId *int64 `json:"summary_table_id,omitempty"`
}

func (o CompoundMetricVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompoundMetricVo struct{}"
	}

	return strings.Join([]string{"CompoundMetricVo", string(data)}, " ")
}
