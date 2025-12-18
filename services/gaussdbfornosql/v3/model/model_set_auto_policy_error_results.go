package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetAutoPolicyErrorResults struct {

	// **参数解释：** 实例ID。 **取值范围：** 不涉及
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释：** 错误信息。 **取值范围：** 不涉及
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o SetAutoPolicyErrorResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoPolicyErrorResults struct{}"
	}

	return strings.Join([]string{"SetAutoPolicyErrorResults", string(data)}, " ")
}
