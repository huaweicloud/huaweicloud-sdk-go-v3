package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportDataobjectsResponse Response Object
type ExportDataobjectsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportDataobjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDataobjectsResponse struct{}"
	}

	return strings.Join([]string{"ExportDataobjectsResponse", string(data)}, " ")
}
