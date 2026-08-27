package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProviderTemplatesRequest Request Object
type ListProviderTemplatesRequest struct {
}

func (o ListProviderTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProviderTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListProviderTemplatesRequest", string(data)}, " ")
}
