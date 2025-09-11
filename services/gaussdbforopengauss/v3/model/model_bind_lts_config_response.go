package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindLtsConfigResponse Response Object
type BindLtsConfigResponse struct {

	// **参数解释**: 关联LTS日志流任务ID。 **取值范围**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BindLtsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindLtsConfigResponse struct{}"
	}

	return strings.Join([]string{"BindLtsConfigResponse", string(data)}, " ")
}
