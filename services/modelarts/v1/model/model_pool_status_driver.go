package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolStatusDriver **参数解释**：资源池驱动状态信息。
type PoolStatusDriver struct {
	Gpu *PoolDriverStatus `json:"gpu,omitempty"`

	Npu *PoolDriverStatus `json:"npu,omitempty"`
}

func (o PoolStatusDriver) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolStatusDriver struct{}"
	}

	return strings.Join([]string{"PoolStatusDriver", string(data)}, " ")
}
