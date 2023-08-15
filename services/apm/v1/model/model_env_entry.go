package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnvEntry 环境信息。
type EnvEntry struct {

	// 环境id。
	EnvId *int64 `json:"env_id,omitempty"`

	// 环境名称。
	EnvName *string `json:"env_name,omitempty"`
}

func (o EnvEntry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvEntry struct{}"
	}

	return strings.Join([]string{"EnvEntry", string(data)}, " ")
}
