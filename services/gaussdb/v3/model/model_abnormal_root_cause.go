package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AbnormalRootCause struct {

	// **参数解释**： 异常类型。 **取值范围**： LOCK_WAIT：锁等待。
	AbnormalType *string `json:"abnormal_type,omitempty"`

	LockRootCause *LockRootCause `json:"lock_root_cause,omitempty"`
}

func (o AbnormalRootCause) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AbnormalRootCause struct{}"
	}

	return strings.Join([]string{"AbnormalRootCause", string(data)}, " ")
}
