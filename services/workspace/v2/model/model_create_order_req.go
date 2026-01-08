package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrderReq 包周期订单请求体。
type CreateOrderReq struct {

	// 订单类型，createDesktops：创建桌面，addVolumes：添加磁盘，rebuildDesktops：重建系统盘，createDesktopPool 创建桌面池，expandDesktopPool 扩容桌面池，applyDesktopsInternet：开通桌面EIP上网，createExclusiveHosts：创建专享主机，subscribeUserSharer：订购用户协同资源，ApplySubnetBandwidth：开通云办公带宽。
	Type string `json:"type"`

	// 企业项目ID，默认\"0。\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 授权给Billing服务的委托URN。使用RAM共享密钥创建包周期云桌面或添加包周期磁盘时，需要传入该字段。
	AgencyUrn *string `json:"agency_urn,omitempty"`

	// 包周期资源。
	Resources []Resource `json:"resources"`

	ExtendParam *OrderExtendParam `json:"extend_param,omitempty"`
}

func (o CreateOrderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrderReq struct{}"
	}

	return strings.Join([]string{"CreateOrderReq", string(data)}, " ")
}
