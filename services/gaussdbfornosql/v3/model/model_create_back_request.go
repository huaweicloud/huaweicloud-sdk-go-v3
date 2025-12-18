package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBackRequest Request Object
type CreateBackRequest struct {

	// **参数解释：** 实例ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceId string `json:"instance_id"`

	Body *NoSqlCreateBackupRequestBody `json:"body,omitempty"`
}

func (o CreateBackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackRequest struct{}"
	}

	return strings.Join([]string{"CreateBackRequest", string(data)}, " ")
}
