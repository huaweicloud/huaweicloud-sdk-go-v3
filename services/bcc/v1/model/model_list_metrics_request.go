package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMetricsRequest Request Object
type ListMetricsRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 指标名称，取值范围：BCC.VAULT_UTIL,BCC.VAULT_USED,BCC.EVENT
	Name *ListMetricsRequestName `json:"name,omitempty"`

	// RegionID
	RegionId *string `json:"region_id,omitempty"`

	// 查询范围起始时间，例如：2025-03-20T09:31:45Z。
	StartTime *string `json:"start_time,omitempty"`

	// 查询范围结束时间，例如：2025-03-20T09:31:45Z。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListMetricsRequest", string(data)}, " ")
}

type ListMetricsRequestName struct {
	value string
}

type ListMetricsRequestNameEnum struct {
	BCC_VAULT_UTIL ListMetricsRequestName
	BCC_VAULT_USED ListMetricsRequestName
	BCC_EVENT      ListMetricsRequestName
}

func GetListMetricsRequestNameEnum() ListMetricsRequestNameEnum {
	return ListMetricsRequestNameEnum{
		BCC_VAULT_UTIL: ListMetricsRequestName{
			value: "BCC.VAULT_UTIL",
		},
		BCC_VAULT_USED: ListMetricsRequestName{
			value: "BCC.VAULT_USED",
		},
		BCC_EVENT: ListMetricsRequestName{
			value: "BCC.EVENT",
		},
	}
}

func (c ListMetricsRequestName) Value() string {
	return c.value
}

func (c ListMetricsRequestName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMetricsRequestName) UnmarshalJSON(b []byte) error {
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
