package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeviceCertificateRequest Request Object
type DeleteDeviceCertificateRequest struct {

	// **参数说明**：实例ID。物理多租下各实例的唯一标识，建议携带该参数，在使用专业版时必须携带该参数。您可以在IoTDA管理控制台界面，选择左侧导航栏“总览”页签查看当前实例的ID，具体获取方式请参考[[查看实例详情](https://support.huaweicloud.com/usermanual-iothub/iot_01_0121.html)](tag:hws) [[查看实例详情](https://support.huaweicloud.com/intl/zh-cn/usermanual-iothub/iot_01_0121.html)](tag:hws_hk)。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：设备证书ID，用于唯一标识一个设备证书。在注册设备证书时由物联网平台分配获得。
	CertificateId string `json:"certificate_id"`
}

func (o DeleteDeviceCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceCertificateRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeviceCertificateRequest", string(data)}, " ")
}
