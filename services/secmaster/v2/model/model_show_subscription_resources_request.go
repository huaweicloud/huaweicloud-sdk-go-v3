package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSubscriptionResourcesRequest Request Object
type ShowSubscriptionResourcesRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**: SKU信息 - FLOW_DATA_BANDWIDTH 数据流量带宽 - CSS_CAPACITY CSS 存储容量 - PAIMON_CAPACITY Paimon 存储容量 - OBS_CAPACITY OBS 存储容量 - JOB_CAPACITY 作业处理容量 - AD_HOC_COUNT 即席查询次数  **约束限制** 不涉及 **取值范围**: - FLOW_DATA_BANDWIDTH - CSS_CAPACITY - PAIMON_CAPACITY - OBS_CAPACITY - JOB_CAPACITY - AD_HOC_COUNT  **默认值** 不涉及
	Sku ShowSubscriptionResourcesRequestSku `json:"sku"`
}

func (o ShowSubscriptionResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubscriptionResourcesRequest struct{}"
	}

	return strings.Join([]string{"ShowSubscriptionResourcesRequest", string(data)}, " ")
}

type ShowSubscriptionResourcesRequestSku struct {
	value string
}

type ShowSubscriptionResourcesRequestSkuEnum struct {
	FLOW_DATA_BANDWIDTH ShowSubscriptionResourcesRequestSku
	CSS_CAPACITY        ShowSubscriptionResourcesRequestSku
	PAIMON_CAPACITY     ShowSubscriptionResourcesRequestSku
	OBS_CAPACITY        ShowSubscriptionResourcesRequestSku
	JOB_CAPACITY        ShowSubscriptionResourcesRequestSku
	AD_HOC_COUNT        ShowSubscriptionResourcesRequestSku
}

func GetShowSubscriptionResourcesRequestSkuEnum() ShowSubscriptionResourcesRequestSkuEnum {
	return ShowSubscriptionResourcesRequestSkuEnum{
		FLOW_DATA_BANDWIDTH: ShowSubscriptionResourcesRequestSku{
			value: "FLOW_DATA_BANDWIDTH",
		},
		CSS_CAPACITY: ShowSubscriptionResourcesRequestSku{
			value: "CSS_CAPACITY",
		},
		PAIMON_CAPACITY: ShowSubscriptionResourcesRequestSku{
			value: "PAIMON_CAPACITY",
		},
		OBS_CAPACITY: ShowSubscriptionResourcesRequestSku{
			value: "OBS_CAPACITY",
		},
		JOB_CAPACITY: ShowSubscriptionResourcesRequestSku{
			value: "JOB_CAPACITY",
		},
		AD_HOC_COUNT: ShowSubscriptionResourcesRequestSku{
			value: "AD_HOC_COUNT",
		},
	}
}

func (c ShowSubscriptionResourcesRequestSku) Value() string {
	return c.value
}

func (c ShowSubscriptionResourcesRequestSku) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSubscriptionResourcesRequestSku) UnmarshalJSON(b []byte) error {
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
