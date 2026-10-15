package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchTestConnectionReq 批量测试连接任务请求体
type BatchTestConnectionReq struct {

	// **参数解释**：  批量测试连接请求列表，指定需要测试连接的任务及其数据库连接信息。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Jobs []TestEndPoint `json:"jobs"`
}

func (o BatchTestConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTestConnectionReq struct{}"
	}

	return strings.Join([]string{"BatchTestConnectionReq", string(data)}, " ")
}
