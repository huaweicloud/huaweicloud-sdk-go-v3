package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeviceCertificateRequest Request Object
type ListDeviceCertificateRequest struct {

	// **参数说明**：实例ID。物理多租下各实例的唯一标识，建议携带该参数，在使用专业版时必须携带该参数。您可以在IoTDA管理控制台界面，选择左侧导航栏“总览”页签查看当前实例的ID，具体获取方式请参考[[查看实例详情](https://support.huaweicloud.com/usermanual-iothub/iot_01_0121.html)](tag:hws) [[查看实例详情](https://support.huaweicloud.com/intl/zh-cn/usermanual-iothub/iot_01_0121.html)](tag:hws_hk)。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：证书通用名称，支持前缀模糊查询。 **取值范围**：长度不超过32，只允许字母、数字、下划线（_）、连接符（-）以及（? '#()\\.,&%@!*）等字符的组合。
	CommonName *string `json:"common_name,omitempty"`

	// **参数说明**：证书指纹。 **取值范围**：长度为40的十六进制字符串或者长度为64的十六进制字符串。
	Fingerprint *string `json:"fingerprint,omitempty"`

	// **参数说明**：分页查询时每页显示的记录数。 **取值范围**：1-50的整数，默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数说明**：上一次分页查询结果中最后一条记录的ID，在上一次分页查询时由物联网平台返回获得。分页查询时物联网平台是按marker也就是记录ID降序查询的，越新的数据记录ID也会越大。若填写marker，则本次只查询记录ID小于marker的数据记录。若不填写，则从记录ID最大也就是最新的一条数据开始查询。如果需要依次查询所有数据，则每次查询时必须填写上一次查询响应中的marker值。 **取值范围**：长度为24的十六进制字符串，默认值为ffffffffffffffffffffffff。
	Marker *string `json:"marker,omitempty"`

	// **参数说明**：表示从marker后偏移offset条记录开始查询。默认为0，取值范围为0-500的整数。当offset为0时，表示从marker后第一条记录开始输出。限制offset最大值是出于API性能考虑，您可以搭配marker使用该参数实现翻页，例如每页50条记录，1-11页内都可以直接使用offset跳转到指定页，但到11页后，由于offset限制为500，您需要使用第11页返回的marker作为下次查询的marker，以实现翻页到12-22页。 **取值范围**：0-500的整数，默认为0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListDeviceCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeviceCertificateRequest struct{}"
	}

	return strings.Join([]string{"ListDeviceCertificateRequest", string(data)}, " ")
}
