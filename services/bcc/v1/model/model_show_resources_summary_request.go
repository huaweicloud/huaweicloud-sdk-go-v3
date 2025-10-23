package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowResourcesSummaryRequest Request Object
type ShowResourcesSummaryRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 云服务类型,ecs,evs,sfsturbo,workspace,rds,gaussdb,cbr
	Provider *ShowResourcesSummaryRequestProvider `json:"provider,omitempty"`

	// 起始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowResourcesSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourcesSummaryRequest struct{}"
	}

	return strings.Join([]string{"ShowResourcesSummaryRequest", string(data)}, " ")
}

type ShowResourcesSummaryRequestProvider struct {
	value string
}

type ShowResourcesSummaryRequestProviderEnum struct {
	ECS       ShowResourcesSummaryRequestProvider
	EVS       ShowResourcesSummaryRequestProvider
	SFSTURBO  ShowResourcesSummaryRequestProvider
	WORKSPACE ShowResourcesSummaryRequestProvider
	RDS       ShowResourcesSummaryRequestProvider
	GAUSSDB   ShowResourcesSummaryRequestProvider
	CBR       ShowResourcesSummaryRequestProvider
}

func GetShowResourcesSummaryRequestProviderEnum() ShowResourcesSummaryRequestProviderEnum {
	return ShowResourcesSummaryRequestProviderEnum{
		ECS: ShowResourcesSummaryRequestProvider{
			value: "ecs",
		},
		EVS: ShowResourcesSummaryRequestProvider{
			value: "evs",
		},
		SFSTURBO: ShowResourcesSummaryRequestProvider{
			value: "sfsturbo",
		},
		WORKSPACE: ShowResourcesSummaryRequestProvider{
			value: "workspace",
		},
		RDS: ShowResourcesSummaryRequestProvider{
			value: "rds",
		},
		GAUSSDB: ShowResourcesSummaryRequestProvider{
			value: "gaussdb",
		},
		CBR: ShowResourcesSummaryRequestProvider{
			value: "cbr",
		},
	}
}

func (c ShowResourcesSummaryRequestProvider) Value() string {
	return c.value
}

func (c ShowResourcesSummaryRequestProvider) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourcesSummaryRequestProvider) UnmarshalJSON(b []byte) error {
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
