package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBreifStructTemplateRequest struct {
}

func (o ListBreifStructTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBreifStructTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListBreifStructTemplateRequest", string(data)}, " ")
}
