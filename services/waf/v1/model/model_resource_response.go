package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源信息
type ResourceResponse struct {

	// 资源id
	ResourceId *string `json:"resourceId,omitempty"`

	// 云服务产品对应的云服务类型
	CloudServiceType *string `json:"cloudServiceType,omitempty"`

	// 云服务产品的资源类型   - hws.resource.type.waf: 云模式waf   - hws.resource.type.waf.domain: 域名扩展包   - hws.resource.type.waf.bandwidth：带宽扩展包   - hws.resource.type.waf.rule：规则扩展包
	ResourceType *string `json:"resourceType,omitempty"`

	// 云服务产品的资源规格
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty"`

	// 资源状态   - 0：解冻/正常   - 1：冻结   - 2：删除
	Status *int32 `json:"status,omitempty"`

	// 资源到期时间
	ExpireTime *string `json:"expireTime,omitempty"`

	// 资源数量
	ResourceSize *int32 `json:"resourceSize,omitempty"`
}

func (o ResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceResponse struct{}"
	}

	return strings.Join([]string{"ResourceResponse", string(data)}, " ")
}
