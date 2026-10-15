package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataFilterRuleReq 数据过滤规则请求体
type DataFilterRuleReq struct {

	// 数据过滤条件列表。包含需要校验的数据过滤规则信息，当data_filter_type为common（简单条件过滤）时，此参数为必填项。列表中的每个元素代表一组过滤条件，包含过滤对象列表和源库过滤条件SQL。包含以下关键字段：filter_object_list（过滤对象列表）/source_filter_condition（源库过滤条件SQL）。适用条件：data_filter_type=common时必填。
	GeneralFilteringList *[]GeneralTransformationReq `json:"general_filtering_list,omitempty"`

	// 数据过滤高级设置列表。用于配置关联表场景下的高级过滤规则，包含数据库名、表名、列名、索引等详细过滤配置。当前Mysql-Mysql、Mysql-TaurusDB、TaurusDB-TaurusDB、TaurusDB-Mysql、Mysql-MariaDB、MariaDB-Mysql、MariaDB-TaurusDBL链路支持此功能。适用条件：data_filter_type=config时必填。
	AdvancedSettingList *[]DataFilteringAdvancedSetting `json:"advanced_setting_list,omitempty"`

	// 校验请求的来源场景。用于区分数据过滤规则校验是来自数据同步任务执行场景，还是来自数据对比任务执行场景。 取值：job（数据同步任务执行时的过滤规则校验）/compare（数据对比任务执行时的过滤规则校验）
	Source *DataFilterRuleReqSource `json:"source,omitempty"`

	// 数据过滤类型。用于指定数据过滤规则的模式，common表示简单条件过滤（单表过滤条件），config表示关联表过滤（多表关联条件过滤）。 取值：common（简单条件过滤）/config（关联表过滤）。 约束：当简单条件过滤是子查询时，需要先校验关联条件，关联条件校验成功之后，需要保存关联条件。
	DataFilterType *string `json:"data_filter_type,omitempty"`
}

func (o DataFilterRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataFilterRuleReq struct{}"
	}

	return strings.Join([]string{"DataFilterRuleReq", string(data)}, " ")
}

type DataFilterRuleReqSource struct {
	value string
}

type DataFilterRuleReqSourceEnum struct {
	JOB     DataFilterRuleReqSource
	COMPARE DataFilterRuleReqSource
}

func GetDataFilterRuleReqSourceEnum() DataFilterRuleReqSourceEnum {
	return DataFilterRuleReqSourceEnum{
		JOB: DataFilterRuleReqSource{
			value: "job",
		},
		COMPARE: DataFilterRuleReqSource{
			value: "compare",
		},
	}
}

func (c DataFilterRuleReqSource) Value() string {
	return c.value
}

func (c DataFilterRuleReqSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataFilterRuleReqSource) UnmarshalJSON(b []byte) error {
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
