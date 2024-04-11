package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportDesignResultResponse Response Object
type ExportDesignResultResponse struct {
	Data           *DsExportResultVoData `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ExportDesignResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDesignResultResponse struct{}"
	}

	return strings.Join([]string{"ExportDesignResultResponse", string(data)}, " ")
}
