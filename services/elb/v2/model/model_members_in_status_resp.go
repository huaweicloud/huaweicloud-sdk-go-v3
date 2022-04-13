package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 后端云服务器对象列表，用于状态树中
type MembersInStatusResp struct {
	// 后端云服务器ID

	Id string `json:"id"`
	// 后端云服务器的IP地址

	Address string `json:"address"`
	// 后端云服务器的端口号

	ProtocolPort int32 `json:"protocol_port"`
	// 后端云服务器的健康检状态；可以为：ONLINE：健康检查在线；OFFLINE：健康检查离线；DISABLED：后端云服务器无对应的弹性云服务器；NO_MONITOR：健康检查未开启

	OperatingStatus MembersInStatusRespOperatingStatus `json:"operating_status"`
	// 监听器的配置状态；该字段为预留字段，暂未启用。默认为ACTIVE。

	ProvisioningStatus string `json:"provisioning_status"`
}

func (o MembersInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MembersInStatusResp struct{}"
	}

	return strings.Join([]string{"MembersInStatusResp", string(data)}, " ")
}

type MembersInStatusRespOperatingStatus struct {
	value string
}

type MembersInStatusRespOperatingStatusEnum struct {
	ONLINE     MembersInStatusRespOperatingStatus
	OFFLINE    MembersInStatusRespOperatingStatus
	DISABLED   MembersInStatusRespOperatingStatus
	NO_MONITOR MembersInStatusRespOperatingStatus
}

func GetMembersInStatusRespOperatingStatusEnum() MembersInStatusRespOperatingStatusEnum {
	return MembersInStatusRespOperatingStatusEnum{
		ONLINE: MembersInStatusRespOperatingStatus{
			value: "ONLINE",
		},
		OFFLINE: MembersInStatusRespOperatingStatus{
			value: "OFFLINE",
		},
		DISABLED: MembersInStatusRespOperatingStatus{
			value: "DISABLED",
		},
		NO_MONITOR: MembersInStatusRespOperatingStatus{
			value: "NO_MONITOR",
		},
	}
}

func (c MembersInStatusRespOperatingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MembersInStatusRespOperatingStatus) UnmarshalJSON(b []byte) error {
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
