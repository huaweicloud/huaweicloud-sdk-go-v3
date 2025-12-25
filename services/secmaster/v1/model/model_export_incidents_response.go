package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportIncidentsResponse Response Object
type ExportIncidentsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportIncidentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportIncidentsResponse struct{}"
	}

	return strings.Join([]string{"ExportIncidentsResponse", string(data)}, " ")
}
