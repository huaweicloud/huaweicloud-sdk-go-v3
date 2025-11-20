package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Vport 虚拟IP端口信息
type Vport struct {

	// - 参数解释：虚拟IP地址的唯一标识。 - 约束限制：需保证该ID对应资源是当前租户可操作的。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Id string `json:"id"`
}

func (o Vport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vport struct{}"
	}

	return strings.Join([]string{"Vport", string(data)}, " ")
}
