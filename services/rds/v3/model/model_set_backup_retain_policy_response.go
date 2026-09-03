package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetBackupRetainPolicyResponse Response Object
type SetBackupRetainPolicyResponse struct {

	// **参数解释**  设置备份保留策略同步接口返回成功响应结果  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetBackupRetainPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBackupRetainPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetBackupRetainPolicyResponse", string(data)}, " ")
}
