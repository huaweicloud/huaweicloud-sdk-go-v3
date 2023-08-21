package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowStatisticsEventRequest Request Object
type ShowStatisticsEventRequest struct {

	// 开始时间（13位时间戳），需要和end_time同时使用
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳），需要和start_time同时使用
	EndTime int64 `json:"end_time"`

	// 类型：  - attack_count——不同类型攻击事件次数   - flow_drop_count——访问与攻击次数   - ddos_attack_count——DDos攻击次数
	Type ShowStatisticsEventRequestType `json:"type"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowStatisticsEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticsEventRequest struct{}"
	}

	return strings.Join([]string{"ShowStatisticsEventRequest", string(data)}, " ")
}

type ShowStatisticsEventRequestType struct {
	value string
}

type ShowStatisticsEventRequestTypeEnum struct {
	ATTACK_COUNT      ShowStatisticsEventRequestType
	FLOW_DROP_COUNT   ShowStatisticsEventRequestType
	DDOS_ATTACK_COUNT ShowStatisticsEventRequestType
}

func GetShowStatisticsEventRequestTypeEnum() ShowStatisticsEventRequestTypeEnum {
	return ShowStatisticsEventRequestTypeEnum{
		ATTACK_COUNT: ShowStatisticsEventRequestType{
			value: "attack_count",
		},
		FLOW_DROP_COUNT: ShowStatisticsEventRequestType{
			value: "flow_drop_count",
		},
		DDOS_ATTACK_COUNT: ShowStatisticsEventRequestType{
			value: "ddos_attack_count",
		},
	}
}

func (c ShowStatisticsEventRequestType) Value() string {
	return c.value
}

func (c ShowStatisticsEventRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowStatisticsEventRequestType) UnmarshalJSON(b []byte) error {
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
