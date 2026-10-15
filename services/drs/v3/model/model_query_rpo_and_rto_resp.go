package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryRpoAndRtoResp RPO和RTO信息体
type QueryRpoAndRtoResp struct {

	// **参数解释**：  数据复制任务ID，用于唯一标识一个任务。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	JobId *string `json:"job_id,omitempty"`

	RpoInfo *RpoAndRtoInfo `json:"rpo_info,omitempty"`

	RtoInfo *RpoAndRtoInfo `json:"rto_info,omitempty"`

	// **参数解释**：  错误码，当查询任务RPO和RTO失败时返回的错误代码。  **约束限制**：  不涉及。  **取值范围**：  格式为DRS.XXXXXX。  **默认取值**：  不涉及
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  错误详细信息，当查询任务RPO和RTO失败时返回的具体错误描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o QueryRpoAndRtoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRpoAndRtoResp struct{}"
	}

	return strings.Join([]string{"QueryRpoAndRtoResp", string(data)}, " ")
}
