package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTemplatesRequest Request Object
type BatchDeleteTemplatesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *TemplateListRo `json:"body,omitempty"`
}

func (o BatchDeleteTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTemplatesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTemplatesRequest", string(data)}, " ")
}
