package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourcePackageOrderReq 创建资源包订单请求体。
type CreateResourcePackageOrderReq struct {

	// 企业项目ID，上传则指定企业项目，不上传则表示所有企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源包。
	ResourcePackages []ResourcePackage `json:"resource_packages"`

	// 购买资源包数量。
	ResourceSize *int32 `json:"resource_size,omitempty"`

	ExtendParam *OrderExtendParam `json:"extend_param,omitempty"`
}

func (o CreateResourcePackageOrderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourcePackageOrderReq struct{}"
	}

	return strings.Join([]string{"CreateResourcePackageOrderReq", string(data)}, " ")
}
