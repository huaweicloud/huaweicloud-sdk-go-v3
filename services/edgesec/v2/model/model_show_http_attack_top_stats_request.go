package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHttpAttackTopStatsRequest Request Object
type ShowHttpAttackTopStatsRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id，默认为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 指标类型。例如：req_num（请求次数）、bw（带宽）。目前只支持req_num
	StatType ShowHttpAttackTopStatsRequestStatType `json:"stat_type"`

	// 分组类型。响应值按分组类型进行统计，类型可为：host（请求的服务器域名）、sip（请求的客户端IP）、url（请求URL）、rule（自定义的策略类型描述）、user-agent（用户代理）、method（请求方法）、country（国家维度统计）。
	GroupBy string `json:"group_by"`

	// 限制Top数量(不超过100，默认为10)
	Limit *int32 `json:"limit,omitempty"`

	// 时间枚举（LATEST（最近30mins）、TODAY（今天）、CUSTOMIZE（自定义，昨天到近30天任意整数天内））
	TimeType ShowHttpAttackTopStatsRequestTimeType `json:"time_type"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ShowHttpAttackTopStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpAttackTopStatsRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpAttackTopStatsRequest", string(data)}, " ")
}

type ShowHttpAttackTopStatsRequestStatType struct {
	value string
}

type ShowHttpAttackTopStatsRequestStatTypeEnum struct {
	REQ_NUM ShowHttpAttackTopStatsRequestStatType
	BW      ShowHttpAttackTopStatsRequestStatType
}

func GetShowHttpAttackTopStatsRequestStatTypeEnum() ShowHttpAttackTopStatsRequestStatTypeEnum {
	return ShowHttpAttackTopStatsRequestStatTypeEnum{
		REQ_NUM: ShowHttpAttackTopStatsRequestStatType{
			value: "req_num",
		},
		BW: ShowHttpAttackTopStatsRequestStatType{
			value: "bw",
		},
	}
}

func (c ShowHttpAttackTopStatsRequestStatType) Value() string {
	return c.value
}

func (c ShowHttpAttackTopStatsRequestStatType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHttpAttackTopStatsRequestStatType) UnmarshalJSON(b []byte) error {
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

type ShowHttpAttackTopStatsRequestTimeType struct {
	value string
}

type ShowHttpAttackTopStatsRequestTimeTypeEnum struct {
	LATEST    ShowHttpAttackTopStatsRequestTimeType
	TODAY     ShowHttpAttackTopStatsRequestTimeType
	CUSTOMIZE ShowHttpAttackTopStatsRequestTimeType
}

func GetShowHttpAttackTopStatsRequestTimeTypeEnum() ShowHttpAttackTopStatsRequestTimeTypeEnum {
	return ShowHttpAttackTopStatsRequestTimeTypeEnum{
		LATEST: ShowHttpAttackTopStatsRequestTimeType{
			value: "LATEST",
		},
		TODAY: ShowHttpAttackTopStatsRequestTimeType{
			value: "TODAY",
		},
		CUSTOMIZE: ShowHttpAttackTopStatsRequestTimeType{
			value: "CUSTOMIZE",
		},
	}
}

func (c ShowHttpAttackTopStatsRequestTimeType) Value() string {
	return c.value
}

func (c ShowHttpAttackTopStatsRequestTimeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHttpAttackTopStatsRequestTimeType) UnmarshalJSON(b []byte) error {
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
