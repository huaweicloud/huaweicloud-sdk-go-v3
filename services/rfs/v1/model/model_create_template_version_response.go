package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplateVersionResponse Response Object
type CreateTemplateVersionResponse struct {

	// 模板模板版本ID
	VersionId      *string `json:"version_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTemplateVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateVersionResponse struct{}"
	}

	return strings.Join([]string{"CreateTemplateVersionResponse", string(data)}, " ")
}
