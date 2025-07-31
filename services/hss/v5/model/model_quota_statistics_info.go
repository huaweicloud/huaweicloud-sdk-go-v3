package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QuotaStatisticsInfo 防护配额统计信息
type QuotaStatisticsInfo struct {

	// version
	Version *QuotaStatisticsInfoVersion `json:"version,omitempty"`

	// 空闲总数
	IdleNum *int32 `json:"idle_num,omitempty"`

	// 使用中总数
	UsedNum *int32 `json:"used_num,omitempty"`

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`
}

func (o QuotaStatisticsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaStatisticsInfo struct{}"
	}

	return strings.Join([]string{"QuotaStatisticsInfo", string(data)}, " ")
}

type QuotaStatisticsInfoVersion struct {
	value string
}

type QuotaStatisticsInfoVersionEnum struct {
	HSS_VERSION_BASIC                QuotaStatisticsInfoVersion
	HSS_VERSION_ENTERPRISE           QuotaStatisticsInfoVersion
	HSS_VERSION_PREMIUM              QuotaStatisticsInfoVersion
	HSS_VERSION_WTP                  QuotaStatisticsInfoVersion
	HSS_VERSION_CONTAINER_ENTERPRISE QuotaStatisticsInfoVersion
}

func GetQuotaStatisticsInfoVersionEnum() QuotaStatisticsInfoVersionEnum {
	return QuotaStatisticsInfoVersionEnum{
		HSS_VERSION_BASIC: QuotaStatisticsInfoVersion{
			value: "hss.version.basic",
		},
		HSS_VERSION_ENTERPRISE: QuotaStatisticsInfoVersion{
			value: "hss.version.enterprise",
		},
		HSS_VERSION_PREMIUM: QuotaStatisticsInfoVersion{
			value: "hss.version.premium",
		},
		HSS_VERSION_WTP: QuotaStatisticsInfoVersion{
			value: "hss.version.wtp",
		},
		HSS_VERSION_CONTAINER_ENTERPRISE: QuotaStatisticsInfoVersion{
			value: "hss.version.container.enterprise",
		},
	}
}

func (c QuotaStatisticsInfoVersion) Value() string {
	return c.value
}

func (c QuotaStatisticsInfoVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QuotaStatisticsInfoVersion) UnmarshalJSON(b []byte) error {
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
