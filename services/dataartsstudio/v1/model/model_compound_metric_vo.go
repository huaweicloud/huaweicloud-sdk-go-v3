package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type CompoundMetricVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 字段名。
	NameEn string `json:"name_en"`

	// 业务属性。
	NameCh string `json:"name_ch"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 颗粒度ID。
	DimensionGroup string `json:"dimension_group"`

	// 颗粒度名称，只读。
	GroupName *string `json:"group_name,omitempty"`

	// 颗粒度编码，只读。
	GroupCode *string `json:"group_code,omitempty"`

	// 复合指标类型。 枚举值：   - EXPRESSION: 表达式   - PERIODICITY_VALUED_COMPARISON: 环比   - INTERVAL_VALUED_COMPARISON: 同比
	CompoundType *CompoundMetricVoCompoundType `json:"compound_type,omitempty"`

	// 比较类型。 枚举值：   - YEAR_TO_YEAR: 年同比   - MONTH_TO_MONTH: 月同比   - WEEK_TO_WEEK: 周同比
	ComparisonType *CompoundMetricVoComparisonType `json:"comparison_type,omitempty"`

	// 指标信息，填写String类型替代Long类型。
	MetricIds []string `json:"metric_ids"`

	// 指标名称信息。
	MetricNames *[]string `json:"metric_names,omitempty"`

	// 复合指标信息，填写String类型替代Long类型。
	CompoundMetricIds *[]string `json:"compound_metric_ids,omitempty"`

	// 复合指标名称信息
	CompoundMetricNames *[]string `json:"compound_metric_names,omitempty"`

	// 引用函数ID，填写String类型替代Long类型。
	CalFnIds *[]string `json:"cal_fn_ids,omitempty"`

	// 计算表达式，形如${index_id} + ${compound#index_id}，其中index_id代表引用的衍生指标ID，compound#index_id代表引用的复合指标ID。
	CalExp string `json:"cal_exp"`

	// 主题域分组ID，只读，填写String类型替代Long类型。
	L1Id *string `json:"l1_id,omitempty"`

	// 主题域ID，只读，创建和更新时无需填写。
	L2Id *string `json:"l2_id,omitempty"`

	// 业务对象ID，填写String类型替代Long类型。
	L3Id *string `json:"l3_id,omitempty"`

	// 字段类型。
	DataType *string `json:"data_type,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
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

	// 汇总表ID，只读，填写String类型替代Long类型。
	SummaryTableId *string `json:"summary_table_id,omitempty"`
}

func (o CompoundMetricVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompoundMetricVo struct{}"
	}

	return strings.Join([]string{"CompoundMetricVo", string(data)}, " ")
}

type CompoundMetricVoCompoundType struct {
	value string
}

type CompoundMetricVoCompoundTypeEnum struct {
	EXPRESSION                    CompoundMetricVoCompoundType
	PERIODICITY_VALUED_COMPARISON CompoundMetricVoCompoundType
	INTERVAL_VALUED_COMPARISON    CompoundMetricVoCompoundType
}

func GetCompoundMetricVoCompoundTypeEnum() CompoundMetricVoCompoundTypeEnum {
	return CompoundMetricVoCompoundTypeEnum{
		EXPRESSION: CompoundMetricVoCompoundType{
			value: "EXPRESSION",
		},
		PERIODICITY_VALUED_COMPARISON: CompoundMetricVoCompoundType{
			value: "PERIODICITY_VALUED_COMPARISON",
		},
		INTERVAL_VALUED_COMPARISON: CompoundMetricVoCompoundType{
			value: "INTERVAL_VALUED_COMPARISON",
		},
	}
}

func (c CompoundMetricVoCompoundType) Value() string {
	return c.value
}

func (c CompoundMetricVoCompoundType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CompoundMetricVoCompoundType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CompoundMetricVoComparisonType struct {
	value string
}

type CompoundMetricVoComparisonTypeEnum struct {
	YEAR_TO_YEAR   CompoundMetricVoComparisonType
	MONTH_TO_MONTH CompoundMetricVoComparisonType
	WEEK_TO_WEEK   CompoundMetricVoComparisonType
}

func GetCompoundMetricVoComparisonTypeEnum() CompoundMetricVoComparisonTypeEnum {
	return CompoundMetricVoComparisonTypeEnum{
		YEAR_TO_YEAR: CompoundMetricVoComparisonType{
			value: "YEAR_TO_YEAR",
		},
		MONTH_TO_MONTH: CompoundMetricVoComparisonType{
			value: "MONTH_TO_MONTH",
		},
		WEEK_TO_WEEK: CompoundMetricVoComparisonType{
			value: "WEEK_TO_WEEK",
		},
	}
}

func (c CompoundMetricVoComparisonType) Value() string {
	return c.value
}

func (c CompoundMetricVoComparisonType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CompoundMetricVoComparisonType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
