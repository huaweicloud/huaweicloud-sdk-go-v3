package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageTemplateListRequest Request Object
type ShowImageTemplateListRequest struct {
}

func (o ShowImageTemplateListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageTemplateListRequest struct{}"
	}

	return strings.Join([]string{"ShowImageTemplateListRequest", string(data)}, " ")
}
