package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExportJobListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportJobListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportJobListResponse struct{}"
	}

	return strings.Join([]string{"ExportJobListResponse", string(data)}, " ")
}
