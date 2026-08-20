package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointProxyParam 代理调用外部服务接口参数
type EndpointProxyParam struct {
	Authorization *EndpointAuthorizationBody `json:"authorization,omitempty"`

	//
	Data *interface{} `json:"data,omitempty"`

	// 数据源名称
	DatasourceName *string `json:"datasource_name,omitempty"`

	// 接入点uuid
	EndpointUuid *string `json:"endpoint_uuid,omitempty"`

	// 模块id
	ModuleId *string `json:"module_id,omitempty"`

	//
	Url *string `json:"url,omitempty"`

	//
	IsInner *bool `json:"is_inner,omitempty"`

	// 项目uuid
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 区域名
	RegionName *string `json:"region_name,omitempty"`
}

func (o EndpointProxyParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointProxyParam struct{}"
	}

	return strings.Join([]string{"EndpointProxyParam", string(data)}, " ")
}
