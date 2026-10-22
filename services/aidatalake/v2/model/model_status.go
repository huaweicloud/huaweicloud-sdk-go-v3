package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Status **参数解释**：资源状态。 **取值范围**：   - PREPARING：创建中。   - RELEASING：释放中。   - RUNNING：运行中。   - SCALEIN：缩容中。   - SCALEOUT：扩容中。   - FREEZING：冻结中。   - DELETED：已删除。   - DELETING：删除中。   - FAILED：失败。   - ORDER_PAYING：订单支付中。
type Status struct {
}

func (o Status) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Status struct{}"
	}

	return strings.Join([]string{"Status", string(data)}, " ")
}
