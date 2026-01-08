package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCertCrlFileResponse Response Object
type ExportCertCrlFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportCertCrlFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertCrlFileResponse struct{}"
	}

	return strings.Join([]string{"ExportCertCrlFileResponse", string(data)}, " ")
}
