package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValueResp **参数解释**： 告警阈值。具体阈值取值请参见附录中各服务监控指标中取值范围。 **取值范围**： 最小值为0，最大值为1.7976931348623157e+108。
type ValueResp struct {
}

func (o ValueResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueResp struct{}"
	}

	return strings.Join([]string{"ValueResp", string(data)}, " ")
}
