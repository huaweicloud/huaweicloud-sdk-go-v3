package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserListTemplateResponse Response Object
type ExportUserListTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportUserListTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserListTemplateResponse struct{}"
	}

	return strings.Join([]string{"ExportUserListTemplateResponse", string(data)}, " ")
}
