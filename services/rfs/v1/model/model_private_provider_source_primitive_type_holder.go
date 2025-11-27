package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateProviderSourcePrimitiveTypeHolder struct {

	// 用户使用私有provider，在Terraform模板中定义required_providers信息时，需要指明的source参数。  该参数按照“huawei.com/private-provider/{provider_name}”的形式拼接。关于provider_name和provider_source字段在模板中的使用细节，详见创建私有Provider的API描述。
	ProviderSource *string `json:"provider_source,omitempty"`
}

func (o PrivateProviderSourcePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateProviderSourcePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateProviderSourcePrimitiveTypeHolder", string(data)}, " ")
}
