package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTerminalsBindingDesktopsTemplateResponse Response Object
type ExportTerminalsBindingDesktopsTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportTerminalsBindingDesktopsTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTerminalsBindingDesktopsTemplateResponse struct{}"
	}

	return strings.Join([]string{"ExportTerminalsBindingDesktopsTemplateResponse", string(data)}, " ")
}
