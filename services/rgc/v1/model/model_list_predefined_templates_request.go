package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPredefinedTemplatesRequest Request Object
type ListPredefinedTemplatesRequest struct {
}

func (o ListPredefinedTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPredefinedTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListPredefinedTemplatesRequest", string(data)}, " ")
}
