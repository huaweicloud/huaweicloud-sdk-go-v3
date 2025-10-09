package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyInstancePortResponse Response Object
type ModifyInstancePortResponse struct {

	// **参数解释**: 修改端口的任务ID。 **取值范围**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyInstancePortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstancePortResponse struct{}"
	}

	return strings.Join([]string{"ModifyInstancePortResponse", string(data)}, " ")
}
