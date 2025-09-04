package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSyncHarmonySoftBusRequest Request Object
type CreateSyncHarmonySoftBusRequest struct {

	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。您可以在IoTDA管理控制台界面，选择左侧导航栏“总览”页签查看当前实例的ID
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：鸿蒙软总线ID，用于唯一标识一个鸿蒙软总线，在创建鸿蒙软总线时由物联网平台分配。 **取值范围**：长度不超过36，十六进制字符串和连接符（-）的组合。
	BusId string `json:"bus_id"`
}

func (o CreateSyncHarmonySoftBusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSyncHarmonySoftBusRequest struct{}"
	}

	return strings.Join([]string{"CreateSyncHarmonySoftBusRequest", string(data)}, " ")
}
