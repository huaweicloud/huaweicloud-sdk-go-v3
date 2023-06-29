package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBackupResponse Response Object
type CreateBackupResponse struct {

	// 数据作业ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackupResponse struct{}"
	}

	return strings.Join([]string{"CreateBackupResponse", string(data)}, " ")
}
