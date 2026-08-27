package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProviderTemplatesResponse Response Object
type ListProviderTemplatesResponse struct {

	// 供应商模板列表。
	Templates      *[]ProviderTemplateInfo `json:"templates,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListProviderTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProviderTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListProviderTemplatesResponse", string(data)}, " ")
}
