package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportJobResponse Response Object
type ExportJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportJobResponse struct{}"
	}

	return strings.Join([]string{"ExportJobResponse", string(data)}, " ")
}
