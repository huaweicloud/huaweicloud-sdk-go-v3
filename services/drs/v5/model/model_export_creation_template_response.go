package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCreationTemplateResponse Response Object
type ExportCreationTemplateResponse struct {

	// 异步ID。
	Id *string `json:"id,omitempty"`

	// 状态。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportCreationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCreationTemplateResponse struct{}"
	}

	return strings.Join([]string{"ExportCreationTemplateResponse", string(data)}, " ")
}
