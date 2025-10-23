package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListVaultTopsRequest Request Object
type ListVaultTopsRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 返回的类型：按容量排序返回或者按使用率排序返回，取值范围：vault_util,used_vault_size
	MetricName ListVaultTopsRequestMetricName `json:"metric_name"`

	// 返回的存储库数量，取值范围(0,100]，默认值为5，用于限制结果数据条数。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListVaultTopsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVaultTopsRequest struct{}"
	}

	return strings.Join([]string{"ListVaultTopsRequest", string(data)}, " ")
}

type ListVaultTopsRequestMetricName struct {
	value string
}

type ListVaultTopsRequestMetricNameEnum struct {
	VAULT_UTIL      ListVaultTopsRequestMetricName
	USED_VAULT_SIZE ListVaultTopsRequestMetricName
}

func GetListVaultTopsRequestMetricNameEnum() ListVaultTopsRequestMetricNameEnum {
	return ListVaultTopsRequestMetricNameEnum{
		VAULT_UTIL: ListVaultTopsRequestMetricName{
			value: "vault_util",
		},
		USED_VAULT_SIZE: ListVaultTopsRequestMetricName{
			value: "used_vault_size",
		},
	}
}

func (c ListVaultTopsRequestMetricName) Value() string {
	return c.value
}

func (c ListVaultTopsRequestMetricName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVaultTopsRequestMetricName) UnmarshalJSON(b []byte) error {
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
