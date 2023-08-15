package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStandardTemplateRequest Request Object
type DeleteStandardTemplateRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id
	Ids string `json:"ids"`
}

func (o DeleteStandardTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStandardTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteStandardTemplateRequest", string(data)}, " ")
}
