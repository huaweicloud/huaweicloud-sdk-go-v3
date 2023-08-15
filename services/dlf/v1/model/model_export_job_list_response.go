package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportJobListResponse Response Object
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
