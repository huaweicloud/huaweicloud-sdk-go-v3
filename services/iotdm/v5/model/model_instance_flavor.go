package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceFlavor struct {

	// **参数说明**：创建的实例类型。实例类型说明参见[[产品规格说明](https://support.huaweicloud.com/productdesc-iothub/iot_04_0014.html)](tag:hws)[[产品规格说明](https://support.huaweicloud.com/intl/zh-cn/productdesc-iothub/iot_04_0014.html)](tag:hws_hk)。 **取值范围**： - standard：标准版实例 - enterprise：企业版实例
	InstanceType string `json:"instance_type"`

	// **参数说明**：设备接入实例的规格名称。详情请参见[[产品规格说明](https://support.huaweicloud.com/productdesc-iothub/iot_04_0014.html)](tag:hws)[[产品规格说明](https://support.huaweicloud.com/intl/zh-cn/productdesc-iothub/iot_04_0014.html)](tag:hws_hk)中的规格编码。
	Type string `json:"type"`

	// **参数说明**：实例规格的在售状态。 **取值范围**： - normal：在售 - sellout：已售罄
	Status string `json:"status"`

	// **参数说明**：实例规格支持的付费方式列表。 **取值范围**： - prePaid：包年/包月 - postPaid：按需计费
	ChargeModes []string `json:"charge_modes"`
}

func (o InstanceFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceFlavor struct{}"
	}

	return strings.Join([]string{"InstanceFlavor", string(data)}, " ")
}
