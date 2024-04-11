package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DerivativeIndexVo struct {

	// 编码。
	Id *int64 `json:"id,omitempty"`

	// 字段名。
	NameEn string `json:"name_en"`

	// 中文名。
	NameCh string `json:"name_ch"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 字段类型。
	DataType *string `json:"data_type,omitempty"`

	// 主题域分组ID。
	L1Id *int64 `json:"l1_id,omitempty"`

	// 主题域ID，只读，创建和更新时无需填写。
	L2Id *string `json:"l2_id,omitempty"`

	// 业务对象guid。
	L3Id int64 `json:"l3_id"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 原子指标ID。
	AtomicIndexId int64 `json:"atomic_index_id"`

	// 时间限定ID。
	TimeConditionId *int64 `json:"time_condition_id,omitempty"`

	// 时间限定关联字段ID。
	TimeFieldId *int64 `json:"time_field_id,omitempty"`

	// 时间限定关联字段名称。
	TimeFieldName *string `json:"time_field_name,omitempty"`

	// 通用限定信息。
	CommonConditions *[]CommonConditionVo `json:"common_conditions,omitempty"`

	// 维度组(颗粒度)。
	DimensionGroups *[]DerivativeIndexDimensionVo `json:"dimension_groups,omitempty"`

	Monitor *MetricMonitorVo `json:"monitor,omitempty"`

	AtomicIndex *AtomicIndexVo `json:"atomic_index,omitempty"`

	// 时间限定名称。
	TimeConditionName *string `json:"time_condition_name,omitempty"`

	// 创建时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 主题域分组中文名，只读，创建和更新时无需填写。
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名，只读，创建和更新时无需填写。
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名，只读，创建和更新时无需填写。
	L3 *string `json:"l3,omitempty"`

	// 汇总表ID。
	SummaryTableId *int64 `json:"summary_table_id,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`
}

func (o DerivativeIndexVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DerivativeIndexVo struct{}"
	}

	return strings.Join([]string{"DerivativeIndexVo", string(data)}, " ")
}
