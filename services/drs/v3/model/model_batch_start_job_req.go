package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartJobReq 批量启动任务请求体。
type BatchStartJobReq struct {

	// **参数解释**：  批量启动任务请求列表，指定需要启动的一个或多个任务。详情请参见StartInfo对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Jobs []StartInfo `json:"jobs"`
}

func (o BatchStartJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartJobReq struct{}"
	}

	return strings.Join([]string{"BatchStartJobReq", string(data)}, " ")
}
