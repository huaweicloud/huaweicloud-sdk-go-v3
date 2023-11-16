package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListUnblockQuotaStatisticsResponse Response Object
type ListUnblockQuotaStatisticsResponse struct {

	// 用户类型：common_user , native_protection_user
	Type *ListUnblockQuotaStatisticsResponseType `json:"type,omitempty"`

	// 解封总配额
	TotalUnblockingQuota *int32 `json:"total_unblocking_quota,omitempty"`

	// 剩余解封配额
	RemainingUnblockingQuota *int32 `json:"remaining_unblocking_quota,omitempty"`

	// 今日解封配额
	UnblockingQuotaToday *int32 `json:"unblocking_quota_today,omitempty"`

	// 今日剩余解封配额
	RemainingUnblockingQuotaToday *int32 `json:"remaining_unblocking_quota_today,omitempty"`
	HttpStatusCode                int    `json:"-"`
}

func (o ListUnblockQuotaStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUnblockQuotaStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListUnblockQuotaStatisticsResponse", string(data)}, " ")
}

type ListUnblockQuotaStatisticsResponseType struct {
	value string
}

type ListUnblockQuotaStatisticsResponseTypeEnum struct {
	COMMON_USER            ListUnblockQuotaStatisticsResponseType
	NATIVE_PROTECTION_USER ListUnblockQuotaStatisticsResponseType
}

func GetListUnblockQuotaStatisticsResponseTypeEnum() ListUnblockQuotaStatisticsResponseTypeEnum {
	return ListUnblockQuotaStatisticsResponseTypeEnum{
		COMMON_USER: ListUnblockQuotaStatisticsResponseType{
			value: "common_user",
		},
		NATIVE_PROTECTION_USER: ListUnblockQuotaStatisticsResponseType{
			value: "native_protection_user",
		},
	}
}

func (c ListUnblockQuotaStatisticsResponseType) Value() string {
	return c.value
}

func (c ListUnblockQuotaStatisticsResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListUnblockQuotaStatisticsResponseType) UnmarshalJSON(b []byte) error {
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
