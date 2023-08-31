package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClusterRequestExtendParam 变更集群规格的扩展参数
type ResizeClusterRequestExtendParam struct {

	// 专属云CCE集群可指定控制节点的规格
	DecMasterFlavor *string `json:"decMasterFlavor,omitempty"`

	// 是否自动扣款 - “true”：自动扣款 - “false”：不自动扣款 > 包周期集群时生效，不填写此参数时默认不会自动扣款。
	IsAutoPay *string `json:"isAutoPay,omitempty"`
}

func (o ResizeClusterRequestExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterRequestExtendParam struct{}"
	}

	return strings.Join([]string{"ResizeClusterRequestExtendParam", string(data)}, " ")
}
