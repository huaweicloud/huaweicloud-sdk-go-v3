package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SynchronizeInstanceListNewResponse Response Object
type SynchronizeInstanceListNewResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o SynchronizeInstanceListNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SynchronizeInstanceListNewResponse struct{}"
	}

	return strings.Join([]string{"SynchronizeInstanceListNewResponse", string(data)}, " ")
}
