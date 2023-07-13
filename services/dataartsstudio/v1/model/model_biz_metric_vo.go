package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type BizMetricVo struct {

	// 编码
	Id *int64 `json:"id,omitempty"`

	// 指标名称, 支持中英文, 数字, 下划线, 中划线, 中英文括号()（）/
	Name string `json:"name"`

	// 指标编码
	Code *string `json:"code,omitempty"`

	// 指标名称, 支持中英文, 数字, 下划线, 中划线, 中英文括号()（）/
	NameAlias *string `json:"name_alias,omitempty"`

	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 归属的流程架构的id
	BizCatalogId int64 `json:"biz_catalog_id"`

	// 归属的流程架构路径
	BizCatalogPath *string `json:"biz_catalog_path,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	// 数据来源
	DataOrigin *string `json:"data_origin,omitempty"`

	// 计量单位
	Unit *string `json:"unit,omitempty"`

	// 统计周期(时间限定)
	TimeFilters string `json:"time_filters"`

	// 统计维度
	Dimensions *string `json:"dimensions,omitempty"`

	// 统计口径/修饰词（通用限定）
	GeneralFilters *string `json:"general_filters,omitempty"`

	// 刷新频率
	IntervalType BizMetricVoIntervalType `json:"interval_type"`

	// 应用场景
	ApplyScenario *string `json:"apply_scenario,omitempty"`

	// 关联技术指标
	TechnicalMetric *int64 `json:"technical_metric,omitempty"`

	// 关联技术指标名称
	TechnicalMetricName *string `json:"technical_metric_name,omitempty"`

	TechnicalMetricType *BizTypeEnum `json:"technical_metric_type,omitempty"`

	// 度量对象
	Measure *string `json:"measure,omitempty"`

	// 负责人，指标解释人
	Owner string `json:"owner"`

	// 指标管理部门, 支持中英文, 数字, 下划线, 中划线, 中英文括号()（）/, 空格
	OwnerDepartment string `json:"owner_department"`

	// 设置目的
	Destination string `json:"destination"`

	// 资产同步后的guid
	Guid *string `json:"guid,omitempty"`

	// 指标定义
	Definition string `json:"definition"`

	// 计算公式
	Expression string `json:"expression"`

	// 备注
	Remark *string `json:"remark,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 主题域分组中文名
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名
	L3 *string `json:"l3,omitempty"`

	BizMetric *SyncStatusEnum `json:"biz_metric,omitempty"`

	SummaryStatus *SyncStatusEnum `json:"summary_status,omitempty"`
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
