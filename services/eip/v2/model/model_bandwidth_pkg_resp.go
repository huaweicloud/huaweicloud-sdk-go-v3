package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Response Object
type BandwidthPkgResp struct {

	// - 功能说明：加油包ID - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）
	ResourceId string `json:"resourceId"`

	// - 功能说明：加油包名称
	ResourceName string `json:"resourceName"`

	// - 功能说明：资源创建时间，UTC时间格式：2016-03-28T00:00:00Z
	ProcessedTime string `json:"processedTime"`

	// - 功能说明：加油包绑定的原带宽ID
	BandwidthId string `json:"bandwidthId"`

	// - 功能说明：加油包的大小，即在原始带宽之上提升的带宽大小 - 取值范围：>1M，pkgSize+原始带宽大小 < 云服务带宽接口限制的带宽上限
	PkgSize int32 `json:"pkgSize"`

	// - 功能说明：租户id
	TenantId string `json:"tenantId"`

	// - 功能说明：加油包订单相关信息格式：非空时值为''orderId:productId''
	BillingInfo string `json:"billingInfo"`

	// - 功能说明：加油包起始生效时间，UTC时间格式：2016-03-28T00:00:00Z - 取值范围：startTime >= 服务处理请求的时间
	StartTime string `json:"startTime"`

	// - 功能说明：加油包结束时间UTC时间格式：2016-03-28T00:00:00Z - 取值范围：startTime < endTime
	EndTime string `json:"endTime"`

	// - 功能说明：加油包资源状态，仅管理员权限可以变更状态 - 取值范围：''pending'', ''active'', ''completed'', ''error''
	Status string `json:"status"`
}

func (o BandwidthPkgResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthPkgResp struct{}"
	}

	return strings.Join([]string{"BandwidthPkgResp", string(data)}, " ")
}
