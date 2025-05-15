package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NonRequiredName 实例名称。
type NonRequiredName struct {

	// 实例名称。
	Name *string `json:"name,omitempty"`
}

func (o NonRequiredName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NonRequiredName struct{}"
	}

	return strings.Join([]string{"NonRequiredName", string(data)}, " ")
}
