package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetBackupPolicyResponse Response Object
type BatchSetBackupPolicyResponse struct {

	// **参数解释**: 设置成功的实例数量。 **取值范围**: 不涉及。
	SucceededNum *int32 `json:"succeeded_num,omitempty"`

	// **参数解释**: 设置失败的实例数量。 **取值范围**: 不涉及。
	FailedNum *int32 `json:"failed_num,omitempty"`

	// **参数解释**: 设置失败的实例记录。
	FailedInstances *[]BatchSetBackupPolicyFailedRecordResult `json:"failed_instances,omitempty"`
	HttpStatusCode  int                                       `json:"-"`
}

func (o BatchSetBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"BatchSetBackupPolicyResponse", string(data)}, " ")
}
