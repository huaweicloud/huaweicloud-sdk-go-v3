package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValueResp **参数解释**： 告警阈值。具体阈值取值请参见“[支持服务列表](ces_03_0059.xml)”。  **取值范围**： 最小值为-1.7976931348623157e+108，最大值为1.7976931348623157e+108。
type ValueResp struct {
}

func (o ValueResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueResp struct{}"
	}

	return strings.Join([]string{"ValueResp", string(data)}, " ")
}
