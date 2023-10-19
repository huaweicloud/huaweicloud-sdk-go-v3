package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DocumentTemplateVersion 文档模板版本。
type DocumentTemplateVersion struct {
	DocumentTemplateVersion *DocumentTemplateVersionEnum `json:"document_template_version"`
}

func (o DocumentTemplateVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentTemplateVersion struct{}"
	}

	return strings.Join([]string{"DocumentTemplateVersion", string(data)}, " ")
}
