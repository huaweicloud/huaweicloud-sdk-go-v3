package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUsersTemplateResponse Response Object
type ExportUsersTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportUsersTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUsersTemplateResponse struct{}"
	}

	return strings.Join([]string{"ExportUsersTemplateResponse", string(data)}, " ")
}
