package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportIpTemplateResponse Response Object
type ExportIpTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportIpTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportIpTemplateResponse struct{}"
	}

	return strings.Join([]string{"ExportIpTemplateResponse", string(data)}, " ")
}
