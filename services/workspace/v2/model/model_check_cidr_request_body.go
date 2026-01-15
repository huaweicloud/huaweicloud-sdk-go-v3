package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckCidrRequestBody 校验租户冲突网段请求体
type CheckCidrRequestBody struct {

	// 租户网段
	TenantCidrs *[]string `json:"tenant_cidrs,omitempty"`

	// 开通服务资源使用的可用分区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o CheckCidrRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCidrRequestBody struct{}"
	}

	return strings.Join([]string{"CheckCidrRequestBody", string(data)}, " ")
}
