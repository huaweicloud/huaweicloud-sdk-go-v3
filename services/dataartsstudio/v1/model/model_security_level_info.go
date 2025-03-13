package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityLevelInfo 密级等级信息
type SecurityLevelInfo struct {

	// 密级名称
	Name *string `json:"name,omitempty"`

	// 密级等级
	Level *string `json:"level,omitempty"`
}

func (o SecurityLevelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityLevelInfo struct{}"
	}

	return strings.Join([]string{"SecurityLevelInfo", string(data)}, " ")
}
