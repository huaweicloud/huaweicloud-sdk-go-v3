package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePipelineByTemplateRequest struct {
	// 语言类型 中文:zh-cn 英文:en-us，默认en-us

	XLanguage *string `json:"X-Language,omitempty"`

	Body *TemplateCddl `json:"body,omitempty"`
}

func (o CreatePipelineByTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePipelineByTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreatePipelineByTemplateRequest", string(data)}, " ")
}
