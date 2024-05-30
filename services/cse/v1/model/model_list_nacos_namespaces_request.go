package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNacosNamespacesRequest Request Object
type ListNacosNamespacesRequest struct {

	// 微服务引擎的实例ID
	XEngineId string `json:"x-engine-id"`

	// 企业项目ID
	XEnterpriseProjectID string `json:"X-Enterprise-Project-ID"`

	// 分页参数，偏移量，从0开始
	Offset int32 `json:"offset"`

	// 分页参数，分页大小，0表示不分页
	Limit int32 `json:"limit"`
}

func (o ListNacosNamespacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNacosNamespacesRequest struct{}"
	}

	return strings.Join([]string{"ListNacosNamespacesRequest", string(data)}, " ")
}
