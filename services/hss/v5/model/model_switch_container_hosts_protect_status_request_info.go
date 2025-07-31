package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SwitchContainerHostsProtectStatusRequestInfo struct {

	// 开通的版本，包含如下：   - hss.version.null：无   - hss.version.container.enterprise：容器版
	Version SwitchContainerHostsProtectStatusRequestInfoVersion `json:"version"`

	// 付费模式，包含如下：  - packet_cycle：包周期  - on_demand：按需
	ChargingMode string `json:"charging_mode"`

	// 资源实例ID
	ResourceId string `json:"resource_id"`

	// 服务器列表
	HostIdList []string `json:"host_id_list"`

	// 资源标签
	Tags *[]TagInfo `json:"tags,omitempty"`
}

func (o SwitchContainerHostsProtectStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchContainerHostsProtectStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"SwitchContainerHostsProtectStatusRequestInfo", string(data)}, " ")
}

type SwitchContainerHostsProtectStatusRequestInfoVersion struct {
	value string
}

type SwitchContainerHostsProtectStatusRequestInfoVersionEnum struct {
	HSS_VERSION_NULL                 SwitchContainerHostsProtectStatusRequestInfoVersion
	HSS_VERSION_CONTAINER_ENTERPRISE SwitchContainerHostsProtectStatusRequestInfoVersion
}

func GetSwitchContainerHostsProtectStatusRequestInfoVersionEnum() SwitchContainerHostsProtectStatusRequestInfoVersionEnum {
	return SwitchContainerHostsProtectStatusRequestInfoVersionEnum{
		HSS_VERSION_NULL: SwitchContainerHostsProtectStatusRequestInfoVersion{
			value: "hss.version.null",
		},
		HSS_VERSION_CONTAINER_ENTERPRISE: SwitchContainerHostsProtectStatusRequestInfoVersion{
			value: "hss.version.container.enterprise",
		},
	}
}

func (c SwitchContainerHostsProtectStatusRequestInfoVersion) Value() string {
	return c.value
}

func (c SwitchContainerHostsProtectStatusRequestInfoVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchContainerHostsProtectStatusRequestInfoVersion) UnmarshalJSON(b []byte) error {
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
