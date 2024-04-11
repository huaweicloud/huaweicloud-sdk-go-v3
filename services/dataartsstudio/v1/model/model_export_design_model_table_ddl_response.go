package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportDesignModelTableDdlResponse Response Object
type ExportDesignModelTableDdlResponse struct {
	Data           *ExportDesignModelTableDdlResultData `json:"data,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ExportDesignModelTableDdlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDesignModelTableDdlResponse struct{}"
	}

	return strings.Join([]string{"ExportDesignModelTableDdlResponse", string(data)}, " ")
}
