package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthenticationTemplateBody 鉴权模板详细内容。
type AuthenticationTemplateBody struct {

	// **参数说明**：鉴权模板参数，目前平台预制了mqtt协议连接参数中的clientId，username与设备证书中的属性，参数名定义如下: - iotda::mqtt::client_id ：mqtt连接参数三元组中的Client Id - iotda::mqtt::username ：mqtt连接参数三元组中的User Name - iotda::certificate::country (国家/地区,C ) - iotda::certificate::organization (组织,O) - iotda::certificate::organizational_unit (组织单位,OU) - iotda::certificate::distinguished_name_qualifier (可辨别名称限定符,dnQualifier) - iotda::certificate::state_name (省市,ST) - iotda::certificate::common_name (公用名,CN) - iotda::certificate::serial_number (序列号,serialNumber) - iotda::device::secret ：表示设备原始秘钥
	Parameters *interface{} `json:"parameters"`

	Resources *AuthenticationTemplateResource `json:"resources"`
}

func (o AuthenticationTemplateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthenticationTemplateBody struct{}"
	}

	return strings.Join([]string{"AuthenticationTemplateBody", string(data)}, " ")
}
