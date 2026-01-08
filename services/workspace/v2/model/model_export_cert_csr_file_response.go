package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCertCsrFileResponse Response Object
type ExportCertCsrFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportCertCsrFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertCsrFileResponse struct{}"
	}

	return strings.Join([]string{"ExportCertCsrFileResponse", string(data)}, " ")
}
