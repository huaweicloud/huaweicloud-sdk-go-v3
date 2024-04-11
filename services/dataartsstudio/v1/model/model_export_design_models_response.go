package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportDesignModelsResponse Response Object
type ExportDesignModelsResponse struct {
	Data           *ExportDesignModelsResultData `json:"data,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ExportDesignModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDesignModelsResponse struct{}"
	}

	return strings.Join([]string{"ExportDesignModelsResponse", string(data)}, " ")
}
