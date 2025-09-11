package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRegistryRequestBody struct {

	// 目标仓库名称，1-64字符组成，只能包含英文大小写、数字、汉字、中划线和下划线。
	Name string `json:"name"`

	// 目标仓库描述
	Description *string `json:"description,omitempty"`

	// 仓库类型，swr-pro(开源harbor仓库)、swr-pro-internal(SWR企业版仓库)、huawei-SWR(SWR基础版仓库)
	Type string `json:"type"`

	// 企业仓库实例所属的项目ID，当type为swr-pro-internal时必填
	ProjectId *string `json:"project_id,omitempty"`

	// 区域ID，当type为swr-pro-internal时必填
	RegionId *string `json:"region_id,omitempty"`

	// 企业仓库实例ID，当type为swr-pro-internal时必填
	InstanceId *string `json:"instance_id,omitempty"`

	// 目标仓库的地址
	Url string `json:"url"`

	Credential *Credential `json:"credential"`

	DnsConf *DnsConf `json:"dns_conf,omitempty"`

	// 是否验证远程证书
	Insecure bool `json:"insecure"`
}

func (o UpdateRegistryRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRegistryRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRegistryRequestBody", string(data)}, " ")
}
