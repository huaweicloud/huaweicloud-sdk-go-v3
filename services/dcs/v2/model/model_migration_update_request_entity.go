package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationUpdateRequestEntity 迁移更新请求体
type MigrationUpdateRequestEntity struct {

	// 模式
	ResumeMode *string `json:"resume_mode,omitempty"`
}

func (o MigrationUpdateRequestEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationUpdateRequestEntity struct{}"
	}

	return strings.Join([]string{"MigrationUpdateRequestEntity", string(data)}, " ")
}
