package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkCidr **参数解释**：路由网段范围，标准的IPv4格式。 **取值范围**：允许网段为10.0.0.0~10.255.0.0/8~22,172.16.0.0~172.31.0.0/12~22,192.168.0.0~192.168.0.0/16~22。
type NetworkCidr struct {
}

func (o NetworkCidr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkCidr struct{}"
	}

	return strings.Join([]string{"NetworkCidr", string(data)}, " ")
}
