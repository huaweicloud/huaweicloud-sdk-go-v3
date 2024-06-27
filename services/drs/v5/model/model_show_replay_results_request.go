package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowReplayResultsRequest Request Object
type ShowReplayResultsRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowReplayResultsRequestXLanguage `json:"X-Language,omitempty"`

	// 结果类型。取值： - shard_statistics：回放概览基于时间维度统计信息。 - slow_sql：慢SQL详情。 - error_sql： 回放异常SQL详情。 - slow_sql_template：慢SQL统计信息。  - error_sql_template：异常SQL统计信息。 - replaying_sql：正在回放SQL详情。 - error_classification：回放异常SQL分类。
	Type ShowReplayResultsRequestType `json:"type"`

	// 查询数据的起始时间，在type为shard_statistics、slow_sql、error_sql时必填
	StartTime *string `json:"start_time,omitempty"`

	// 查询数据的结束时间，在type为shard_statistics、slow_sql、error_sql时必填
	EndTime *string `json:"end_time,omitempty"`

	// 分页查询数据表当前超始偏移量, 在type为slow_sql、error_sql、slow_sql_template、error_sql_template必填
	Offset *int64 `json:"offset,omitempty"`

	// 分页查询数据表当前页数据总量，在type为slow_sql、error_sql、slow_sql_template、error_sql_template必填
	Limit *int64 `json:"limit,omitempty"`

	// 返回结果按该关键字排序（slow_sql_template支持count，maxLatency、avgLatency关键字，error_sql_template支持count关键字）
	SortKey *string `json:"sort_key,omitempty"`

	// 排序规则，取值如下： - asc：升序 - desc：降序
	SortDir *ShowReplayResultsRequestSortDir `json:"sort_dir,omitempty"`

	// 回放数据库名称，用于在一致性回放策略场景，过滤目标库与源库镜像库回放结果。参数非必须，不提供则默认查询所有数据，其取值如下： - target：查询目标库回放结果 - target_mirror：查询源库镜像库回放结果
	TargetName *ShowReplayResultsRequestTargetName `json:"target_name,omitempty"`

	// 是否查询样例true/false，type=slow_sql/error_sql时生效，值为true时只查询一条样例数据。
	IsSample *bool `json:"is_sample,omitempty"`

	// 错误分类，type=error_sql/error_sql_template时生效，根据错误分类过滤数据。
	ErrorType *string `json:"error_type,omitempty"`

	// sql模板md5，type=slow_sql/error_sql时生效，根据模板过滤对应的异常SQL和慢SQL，该值为本接口type=slow_sql_template/error_sql_template时的返回字段。
	SqlTemplateMd5 *string `json:"sql_template_md5,omitempty"`
}

func (o ShowReplayResultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplayResultsRequest struct{}"
	}

	return strings.Join([]string{"ShowReplayResultsRequest", string(data)}, " ")
}

type ShowReplayResultsRequestXLanguage struct {
	value string
}

type ShowReplayResultsRequestXLanguageEnum struct {
	EN_US ShowReplayResultsRequestXLanguage
	ZH_CN ShowReplayResultsRequestXLanguage
}

func GetShowReplayResultsRequestXLanguageEnum() ShowReplayResultsRequestXLanguageEnum {
	return ShowReplayResultsRequestXLanguageEnum{
		EN_US: ShowReplayResultsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowReplayResultsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowReplayResultsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowReplayResultsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReplayResultsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ShowReplayResultsRequestType struct {
	value string
}

type ShowReplayResultsRequestTypeEnum struct {
	SHARD_STATISTICS     ShowReplayResultsRequestType
	SLOW_SQL             ShowReplayResultsRequestType
	ERROR_SQL            ShowReplayResultsRequestType
	SLOW_SQL_TEMPLATE    ShowReplayResultsRequestType
	ERROR_SQL_TEMPLATE   ShowReplayResultsRequestType
	REPLAYING_SQL        ShowReplayResultsRequestType
	ERROR_CLASSIFICATION ShowReplayResultsRequestType
}

func GetShowReplayResultsRequestTypeEnum() ShowReplayResultsRequestTypeEnum {
	return ShowReplayResultsRequestTypeEnum{
		SHARD_STATISTICS: ShowReplayResultsRequestType{
			value: "shard_statistics",
		},
		SLOW_SQL: ShowReplayResultsRequestType{
			value: "slow_sql",
		},
		ERROR_SQL: ShowReplayResultsRequestType{
			value: "error_sql",
		},
		SLOW_SQL_TEMPLATE: ShowReplayResultsRequestType{
			value: "slow_sql_template",
		},
		ERROR_SQL_TEMPLATE: ShowReplayResultsRequestType{
			value: "error_sql_template",
		},
		REPLAYING_SQL: ShowReplayResultsRequestType{
			value: "replaying_sql",
		},
		ERROR_CLASSIFICATION: ShowReplayResultsRequestType{
			value: "error_classification",
		},
	}
}

func (c ShowReplayResultsRequestType) Value() string {
	return c.value
}

func (c ShowReplayResultsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReplayResultsRequestType) UnmarshalJSON(b []byte) error {
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

type ShowReplayResultsRequestSortDir struct {
	value string
}

type ShowReplayResultsRequestSortDirEnum struct {
	ASC  ShowReplayResultsRequestSortDir
	DESC ShowReplayResultsRequestSortDir
}

func GetShowReplayResultsRequestSortDirEnum() ShowReplayResultsRequestSortDirEnum {
	return ShowReplayResultsRequestSortDirEnum{
		ASC: ShowReplayResultsRequestSortDir{
			value: "asc",
		},
		DESC: ShowReplayResultsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ShowReplayResultsRequestSortDir) Value() string {
	return c.value
}

func (c ShowReplayResultsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReplayResultsRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ShowReplayResultsRequestTargetName struct {
	value string
}

type ShowReplayResultsRequestTargetNameEnum struct {
	TARGET        ShowReplayResultsRequestTargetName
	TARGET_MIRROR ShowReplayResultsRequestTargetName
}

func GetShowReplayResultsRequestTargetNameEnum() ShowReplayResultsRequestTargetNameEnum {
	return ShowReplayResultsRequestTargetNameEnum{
		TARGET: ShowReplayResultsRequestTargetName{
			value: "target",
		},
		TARGET_MIRROR: ShowReplayResultsRequestTargetName{
			value: "target_mirror",
		},
	}
}

func (c ShowReplayResultsRequestTargetName) Value() string {
	return c.value
}

func (c ShowReplayResultsRequestTargetName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReplayResultsRequestTargetName) UnmarshalJSON(b []byte) error {
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
