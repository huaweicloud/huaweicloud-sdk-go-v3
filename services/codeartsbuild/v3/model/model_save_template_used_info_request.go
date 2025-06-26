package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveTemplateUsedInfoRequest Request Object
type SaveTemplateUsedInfoRequest struct {
	Body *SaveTemplateUsedInfoRequestBody `json:"body,omitempty"`
}

func (o SaveTemplateUsedInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTemplateUsedInfoRequest struct{}"
	}

	return strings.Join([]string{"SaveTemplateUsedInfoRequest", string(data)}, " ")
}
