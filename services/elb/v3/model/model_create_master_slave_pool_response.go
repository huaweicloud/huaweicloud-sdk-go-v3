package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMasterSlavePoolResponse Response Object
type CreateMasterSlavePoolResponse struct {

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId *string `json:"request_id,omitempty"`

	Pool           *MasterSlavePool `json:"pool,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateMasterSlavePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMasterSlavePoolResponse struct{}"
	}

	return strings.Join([]string{"CreateMasterSlavePoolResponse", string(data)}, " ")
}
