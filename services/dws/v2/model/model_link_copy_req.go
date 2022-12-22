package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 快照复制请求
type LinkCopyReq struct {

	// 快照名称
	BackupName string `json:"backup_name"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o LinkCopyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinkCopyReq struct{}"
	}

	return strings.Join([]string{"LinkCopyReq", string(data)}, " ")
}
