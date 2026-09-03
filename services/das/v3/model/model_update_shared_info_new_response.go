package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSharedInfoNewResponse Response Object
type UpdateSharedInfoNewResponse struct {

	// 操作结果
	Status         *bool `json:"status,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateSharedInfoNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSharedInfoNewResponse struct{}"
	}

	return strings.Join([]string{"UpdateSharedInfoNewResponse", string(data)}, " ")
}
