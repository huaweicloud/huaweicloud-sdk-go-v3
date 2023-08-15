package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFunctionResponse Response Object
type ExportFunctionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportFunctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFunctionResponse struct{}"
	}

	return strings.Join([]string{"ExportFunctionResponse", string(data)}, " ")
}
