package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Flavor **参数说明**：设备接入实例的规格参数。
type Flavor struct {

	// **参数说明**：待创建设备接入实例的规格名称。详情请参见[[产品规格说明](https://support.huaweicloud.com/productdesc-iothub/iot_04_0014.html)](tag:hws)[[产品规格说明](https://support.huaweicloud.com/intl/zh-cn/productdesc-iothub/iot_04_0014.html)](tag:hws_hk)中的规格编码。
	Type string `json:"type"`

	// **参数说明**：待创建设备接入标准版实例的单元数量。详情请参见[[产品规格说明](https://support.huaweicloud.com/productdesc-iothub/iot_04_0014.html)](tag:hws)[[产品规格说明](https://support.huaweicloud.com/intl/zh-cn/productdesc-iothub/iot_04_0014.html)](tag:hws_hk)。当instance_type是standard时，该参数必填。
	Size *int32 `json:"size,omitempty"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
