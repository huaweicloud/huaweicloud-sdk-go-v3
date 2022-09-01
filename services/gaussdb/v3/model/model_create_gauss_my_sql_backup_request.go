package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateGaussMySqlBackupRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	Body *MysqlCreateBackupRequest `json:"body,omitempty" xml:"body"`
}

func (o CreateGaussMySqlBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlBackupRequest struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlBackupRequest", string(data)}, " ")
}
