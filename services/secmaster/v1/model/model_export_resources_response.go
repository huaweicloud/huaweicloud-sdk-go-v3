package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportResourcesResponse Response Object
type ExportResourcesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportResourcesResponse struct{}"
	}

	return strings.Join([]string{"ExportResourcesResponse", string(data)}, " ")
}
