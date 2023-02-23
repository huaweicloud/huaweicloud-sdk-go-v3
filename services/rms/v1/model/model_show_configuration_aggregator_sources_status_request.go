package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowConfigurationAggregatorSourcesStatusRequest struct {

	// 资源聚合器ID。
	AggregatorId string `json:"aggregator_id"`

	// 聚合帐号的状态。
	UpdateStatus *ShowConfigurationAggregatorSourcesStatusRequestUpdateStatus `json:"update_status,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o ShowConfigurationAggregatorSourcesStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationAggregatorSourcesStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationAggregatorSourcesStatusRequest", string(data)}, " ")
}

type ShowConfigurationAggregatorSourcesStatusRequestUpdateStatus struct {
	value string
}

type ShowConfigurationAggregatorSourcesStatusRequestUpdateStatusEnum struct {
	SUCCEEDED ShowConfigurationAggregatorSourcesStatusRequestUpdateStatus
	FAILED    ShowConfigurationAggregatorSourcesStatusRequestUpdateStatus
	OUTDATED  ShowConfigurationAggregatorSourcesStatusRequestUpdateStatus
}

func GetShowConfigurationAggregatorSourcesStatusRequestUpdateStatusEnum() ShowConfigurationAggregatorSourcesStatusRequestUpdateStatusEnum {
	return ShowConfigurationAggregatorSourcesStatusRequestUpdateStatusEnum{
		SUCCEEDED: ShowConfigurationAggregatorSourcesStatusRequestUpdateStatus{
			value: "SUCCEEDED",
		},
		FAILED: ShowConfigurationAggregatorSourcesStatusRequestUpdateStatus{
			value: "FAILED",
		},
		OUTDATED: ShowConfigurationAggregatorSourcesStatusRequestUpdateStatus{
			value: "OUTDATED",
		},
	}
}

func (c ShowConfigurationAggregatorSourcesStatusRequestUpdateStatus) Value() string {
	return c.value
}

func (c ShowConfigurationAggregatorSourcesStatusRequestUpdateStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowConfigurationAggregatorSourcesStatusRequestUpdateStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
