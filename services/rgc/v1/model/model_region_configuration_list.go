package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RegionConfigurationList 区域纳管情况。
type RegionConfigurationList struct {

	// 区域名字。
	Region string `json:"region"`

	// 账号类型logging,security。 * ENABLED - 开启 * DISABLED - 关闭
	RegionConfigurationStatus RegionConfigurationListRegionConfigurationStatus `json:"region_configuration_status"`
}

func (o RegionConfigurationList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionConfigurationList struct{}"
	}

	return strings.Join([]string{"RegionConfigurationList", string(data)}, " ")
}

type RegionConfigurationListRegionConfigurationStatus struct {
	value string
}

type RegionConfigurationListRegionConfigurationStatusEnum struct {
	ENABLED  RegionConfigurationListRegionConfigurationStatus
	DISABLED RegionConfigurationListRegionConfigurationStatus
}

func GetRegionConfigurationListRegionConfigurationStatusEnum() RegionConfigurationListRegionConfigurationStatusEnum {
	return RegionConfigurationListRegionConfigurationStatusEnum{
		ENABLED: RegionConfigurationListRegionConfigurationStatus{
			value: "ENABLED",
		},
		DISABLED: RegionConfigurationListRegionConfigurationStatus{
			value: "DISABLED",
		},
	}
}

func (c RegionConfigurationListRegionConfigurationStatus) Value() string {
	return c.value
}

func (c RegionConfigurationListRegionConfigurationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegionConfigurationListRegionConfigurationStatus) UnmarshalJSON(b []byte) error {
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
