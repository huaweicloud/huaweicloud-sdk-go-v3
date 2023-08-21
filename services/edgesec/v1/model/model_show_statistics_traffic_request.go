package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowStatisticsTrafficRequest Request Object
type ShowStatisticsTrafficRequest struct {

	// 开始时间（13位时间戳），需要和end_time同时使用
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳），需要和start_time同时使用
	EndTime int64 `json:"end_time"`

	// 类型：  - max_flow_bandwidth——DDos入流量带宽峰值   - max_drop_bandwidth——DDos入流量带宽峰值   - ddos_flow——DDos入流量   - flow_drop_traffic——入流量与清洗流量   - attack_traffic——不同类型攻击流量
	Type ShowStatisticsTrafficRequestType `json:"type"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowStatisticsTrafficRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticsTrafficRequest struct{}"
	}

	return strings.Join([]string{"ShowStatisticsTrafficRequest", string(data)}, " ")
}

type ShowStatisticsTrafficRequestType struct {
	value string
}

type ShowStatisticsTrafficRequestTypeEnum struct {
	MAX_FLOW_BANDWIDTH ShowStatisticsTrafficRequestType
	MAX_DROP_BANDWIDTH ShowStatisticsTrafficRequestType
	DDOS_FLOW          ShowStatisticsTrafficRequestType
	FLOW_DROP_TRAFFIC  ShowStatisticsTrafficRequestType
	ATTACK_TRAFFIC     ShowStatisticsTrafficRequestType
}

func GetShowStatisticsTrafficRequestTypeEnum() ShowStatisticsTrafficRequestTypeEnum {
	return ShowStatisticsTrafficRequestTypeEnum{
		MAX_FLOW_BANDWIDTH: ShowStatisticsTrafficRequestType{
			value: "max_flow_bandwidth",
		},
		MAX_DROP_BANDWIDTH: ShowStatisticsTrafficRequestType{
			value: "max_drop_bandwidth",
		},
		DDOS_FLOW: ShowStatisticsTrafficRequestType{
			value: "ddos_flow",
		},
		FLOW_DROP_TRAFFIC: ShowStatisticsTrafficRequestType{
			value: "flow_drop_traffic",
		},
		ATTACK_TRAFFIC: ShowStatisticsTrafficRequestType{
			value: "attack_traffic",
		},
	}
}

func (c ShowStatisticsTrafficRequestType) Value() string {
	return c.value
}

func (c ShowStatisticsTrafficRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowStatisticsTrafficRequestType) UnmarshalJSON(b []byte) error {
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
