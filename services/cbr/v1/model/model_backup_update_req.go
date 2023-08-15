package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupUpdateReq 更新请求参数体
type BackupUpdateReq struct {
	Backup *BackupUpdate `json:"backup,omitempty"`
}

func (o BackupUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupUpdateReq struct{}"
	}

	return strings.Join([]string{"BackupUpdateReq", string(data)}, " ")
}
