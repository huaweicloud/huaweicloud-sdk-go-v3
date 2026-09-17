package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportBinlogResponse Response Object
type ExportBinlogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportBinlogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportBinlogResponse struct{}"
	}

	return strings.Join([]string{"ExportBinlogResponse", string(data)}, " ")
}
