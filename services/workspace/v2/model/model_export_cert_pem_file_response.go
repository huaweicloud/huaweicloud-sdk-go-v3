package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCertPemFileResponse Response Object
type ExportCertPemFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportCertPemFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertPemFileResponse struct{}"
	}

	return strings.Join([]string{"ExportCertPemFileResponse", string(data)}, " ")
}
