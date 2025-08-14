package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportResourceResponse Response Object
type ExportResourceResponse struct {

	// 返回结果
	Value          *string `json:"value,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportResourceResponse struct{}"
	}

	return strings.Join([]string{"ExportResourceResponse", string(data)}, " ")
}
