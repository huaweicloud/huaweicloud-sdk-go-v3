package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// BizMetricVo 业务指标信息。
type BizMetricVo struct {

	// 编码，更新时必填，创建时为空，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 指标名称。
	Name string `json:"name"`

	// 指标编码，只读。
	Code *string `json:"code,omitempty"`

	// 指标别名。
	NameAlias *string `json:"name_alias,omitempty"`

	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 归属的流程架构的ID，填写String类型替代Long类型。
	BizCatalogId string `json:"biz_catalog_id"`

	// 归属的流程架构路径，只读。
	BizCatalogPath *string `json:"biz_catalog_path,omitempty"`

	// 创建人，只读。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人，只读。
	UpdateBy *string `json:"update_by,omitempty"`

	// 数据来源。
	DataOrigin *string `json:"data_origin,omitempty"`

	// 计量单位。
	Unit *string `json:"unit,omitempty"`

	// 统计周期(时间限定)。
	TimeFilters string `json:"time_filters"`

	// 统计维度。
	Dimensions *string `json:"dimensions,omitempty"`

	// 统计口径和修饰词。
	GeneralFilters *string `json:"general_filters,omitempty"`

	// 刷新频率。 枚举值：   - MINUTE: 每分钟   - HOUR: 每小时   - DAY: 每天   - WEEK: 每周   - MONTH: 每月   - YEAR: 每年   - REAL_TIME: 实时   - HALF_HOUR: 每半小时   - QUART: 每15分钟   - DOUBLE_WEEK: 每两周   - HALF_YEAR: 每半年   - HALF_DAY: 每半天
	IntervalType BizMetricVoIntervalType `json:"interval_type"`

	// 应用场景。
	ApplyScenario *string `json:"apply_scenario,omitempty"`

	// 关联技术指标，填写String类型替代Long类型。
	TechnicalMetric *string `json:"technical_metric,omitempty"`

	// 关联技术指标名称，只读。
	TechnicalMetricName *string `json:"technical_metric_name,omitempty"`

	TechnicalMetricType *BizTypeEnum `json:"technical_metric_type,omitempty"`

	// 度量对象。
	Measure *string `json:"measure,omitempty"`

	// 指标责任人。
	Owner string `json:"owner"`

	// 指标管理部门。
	OwnerDepartment string `json:"owner_department"`

	// 设置目的。
	Destination string `json:"destination"`

	// 资产同步后的guid，只读。
	Guid *string `json:"guid,omitempty"`

	// 指标定义。
	Definition string `json:"definition"`

	// 计算公式。
	Expression string `json:"expression"`

	// 备注。
	Remark *string `json:"remark,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 主题域分组中文名，只读，创建和更新时无需填写。
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名，只读，创建和更新时无需填写。
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名，只读，创建和更新时无需填写。
	L3 *string `json:"l3,omitempty"`

	BizMetric *SyncStatusEnum `json:"biz_metric,omitempty"`

	SummaryStatus *SyncStatusEnum `json:"summary_status,omitempty"`

	// 自定义项
	SelfDefinedFields *[]SelfDefinedFieldVo `json:"self_defined_fields,omitempty"`
}

func (o BizMetricVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BizMetricVo struct{}"
	}

	return strings.Join([]string{"BizMetricVo", string(data)}, " ")
}

type BizMetricVoIntervalType struct {
	value string
}

type BizMetricVoIntervalTypeEnum struct {
	MINUTE      BizMetricVoIntervalType
	HOUR        BizMetricVoIntervalType
	DAY         BizMetricVoIntervalType
	WEEK        BizMetricVoIntervalType
	MONTH       BizMetricVoIntervalType
	YEAR        BizMetricVoIntervalType
	REAL_TIME   BizMetricVoIntervalType
	HALF_HOUR   BizMetricVoIntervalType
	QUART       BizMetricVoIntervalType
	DOUBLE_WEEK BizMetricVoIntervalType
	HALF_YEAR   BizMetricVoIntervalType
	HALF_DAY    BizMetricVoIntervalType
}

func GetBizMetricVoIntervalTypeEnum() BizMetricVoIntervalTypeEnum {
	return BizMetricVoIntervalTypeEnum{
		MINUTE: BizMetricVoIntervalType{
			value: "MINUTE",
		},
		HOUR: BizMetricVoIntervalType{
			value: "HOUR",
		},
		DAY: BizMetricVoIntervalType{
			value: "DAY",
		},
		WEEK: BizMetricVoIntervalType{
			value: "WEEK",
		},
		MONTH: BizMetricVoIntervalType{
			value: "MONTH",
		},
		YEAR: BizMetricVoIntervalType{
			value: "YEAR",
		},
		REAL_TIME: BizMetricVoIntervalType{
			value: "REAL_TIME",
		},
		HALF_HOUR: BizMetricVoIntervalType{
			value: "HALF_HOUR",
		},
		QUART: BizMetricVoIntervalType{
			value: "QUART",
		},
		DOUBLE_WEEK: BizMetricVoIntervalType{
			value: "DOUBLE_WEEK",
		},
		HALF_YEAR: BizMetricVoIntervalType{
			value: "HALF_YEAR",
		},
		HALF_DAY: BizMetricVoIntervalType{
			value: "HALF_DAY",
		},
	}
}

func (c BizMetricVoIntervalType) Value() string {
	return c.value
}

func (c BizMetricVoIntervalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BizMetricVoIntervalType) UnmarshalJSON(b []byte) error {
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
