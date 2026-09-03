package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksBySqlTemplateIdRequest Request Object
type ListTasksBySqlTemplateIdRequest struct {
	Body *ListTasksBySqlTemplateIdRequestBody `json:"body,omitempty"`
}

func (o ListTasksBySqlTemplateIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksBySqlTemplateIdRequest struct{}"
	}

	return strings.Join([]string{"ListTasksBySqlTemplateIdRequest", string(data)}, " ")
}
