package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateInstance struct {

	// **参数说明**：创建的实例类型。实例类型说明参见[[产品规格说明](https://support.huaweicloud.com/productdesc-iothub/iot_04_0014.html)](tag:hws)[[产品规格说明](https://support.huaweicloud.com/intl/zh-cn/productdesc-iothub/iot_04_0014.html)](tag:hws_hk)。 **取值范围**： - standard：标准版实例 - enterprise：企业版实例
	InstanceType string `json:"instance_type"`

	Flavor *Flavor `json:"flavor"`

	// **参数说明**：实例名称 **取值范围**：由中文字符，英文字母、数字及“_”、“-”组成，且长度为[1-64]个字符。
	Name string `json:"name"`

	ChargeInfo *ChargeInfo `json:"charge_info"`

	// **参数说明**：设备接入实例的描述信息。 **取值范围**：由中文，字母，数字，句号，逗号，下划线（“_”），中划线（“-”），空格组成，且长度为[1-256]个字符。
	Description *string `json:"description,omitempty"`

	// **参数说明**：企业项目Id。此字段填写明确的企业项目Id或者0(表示默认企业项目Id)时支持企业项目特性。可以企业项目管理服务中获取。 **取值范围**：长度不超过36，由小写字母[a-f]、数字、连接符（-）的组成。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数说明**：设备接入实例的标签信息。
	Tags *[]Tag `json:"tags,omitempty"`

	AdditionalParams *AdditionalParams `json:"additional_params,omitempty"`
}

func (o CreateInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstance struct{}"
	}

	return strings.Join([]string{"CreateInstance", string(data)}, " ")
}
