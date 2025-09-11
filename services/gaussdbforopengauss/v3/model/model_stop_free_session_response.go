package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopFreeSessionResponse Response Object
type StopFreeSessionResponse struct {

	// **参数解释**： 是否成功。 **取值范围**: - true：成功。 - false：失败。
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o StopFreeSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopFreeSessionResponse struct{}"
	}

	return strings.Join([]string{"StopFreeSessionResponse", string(data)}, " ")
}
