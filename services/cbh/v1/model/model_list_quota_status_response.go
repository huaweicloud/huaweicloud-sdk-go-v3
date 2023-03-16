package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListQuotaStatusResponse struct {

	// 支持IPv6弹性云服务器资源状态。 - sellout 售罄 - normal  正常商用 - abandon 下线（即不显示）
	StatusV6 *string `json:"status_v6,omitempty"`

	// 弹性云服务器资源状态。 - sellout 售罄 - normal  正常商用 - abandon 下线（即不显示）
	Status *string `json:"status,omitempty"`

	// 弹性配额信息，返回默认值null。
	EipQuota *int32 `json:"eip_quota,omitempty"`

	// 剩余可创建云堡垒机配额信息，返回默认值null。
	Quota          *int32 `json:"quota,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListQuotaStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaStatusResponse struct{}"
	}

	return strings.Join([]string{"ListQuotaStatusResponse", string(data)}, " ")
}
