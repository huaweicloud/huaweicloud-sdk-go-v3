package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源配额详情
type QuotaDetail struct {

	// 资源类型 取值范围：  - edge_site：边缘小站  - compute_device：计算设备
	Type *string `json:"type,omitempty"`

	// 资源的总配额 约束：资源的默认配额数可以修改，默认配置：边缘小站（10）
	Quota *int32 `json:"quota,omitempty"`

	// 已创建的资源个数 取值范围：0~quota值
	Used *int32 `json:"used,omitempty"`
}

func (o QuotaDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetail struct{}"
	}

	return strings.Join([]string{"QuotaDetail", string(data)}, " ")
}
