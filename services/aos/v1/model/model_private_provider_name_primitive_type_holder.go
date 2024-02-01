package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateProviderNamePrimitiveTypeHolder struct {

	// 私有provider（private-provider）的名称。此名字在domain_id+region下应唯一，可以使用小写英文、数字、中划线。仅支持以小写英文、数字开头结尾。  按照HCL最佳实践，该名称推荐为在模板中定义的provider的本地名称（local_name）。  创建私有Provider（CreatePrivateProvider）API 还会以 “huawei.com/private-provider”为固定前缀，并以“huawei.com/private-provider/{provider_name}”的形式返回provider_source字段。关于provider_name和provider_source字段在模板中的使用细节，详见创建私有Provider的API描述。
	ProviderName string `json:"provider_name"`
}

func (o PrivateProviderNamePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateProviderNamePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateProviderNamePrimitiveTypeHolder", string(data)}, " ")
}
