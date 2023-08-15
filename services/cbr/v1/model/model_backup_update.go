package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupUpdate
type BackupUpdate struct {

	// 备份名称
	Name *string `json:"name,omitempty"`
}

func (o BackupUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupUpdate struct{}"
	}

	return strings.Join([]string{"BackupUpdate", string(data)}, " ")
}
