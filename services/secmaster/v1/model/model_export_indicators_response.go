package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportIndicatorsResponse Response Object
type ExportIndicatorsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportIndicatorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportIndicatorsResponse struct{}"
	}

	return strings.Join([]string{"ExportIndicatorsResponse", string(data)}, " ")
}
