package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopOrderRequestBody 创建桌面订单请求体。
type CreateDesktopOrderRequestBody struct {

	// 企业项目ID，默认\"0。\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 授权给Billing服务的委托URN。使用RAM共享密钥创建包周期云桌面或添加包周期磁盘时，需要传入该字段。
	AgencyUrn *string `json:"agency_urn,omitempty"`

	// 小时包资源。
	HourPackageResources []HourPackageResource `json:"hour_package_resources"`

	ExtendParam *OrderExtendParam `json:"extend_param,omitempty"`
}

func (o CreateDesktopOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopOrderRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDesktopOrderRequestBody", string(data)}, " ")
}
