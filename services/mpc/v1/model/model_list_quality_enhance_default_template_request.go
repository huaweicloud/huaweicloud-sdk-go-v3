package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListQualityEnhanceDefaultTemplateRequest struct {
}

func (o ListQualityEnhanceDefaultTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListQualityEnhanceDefaultTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListQualityEnhanceDefaultTemplateRequest", string(data)}, " ")
}
