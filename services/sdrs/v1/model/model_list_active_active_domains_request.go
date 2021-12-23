package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListActiveActiveDomainsRequest struct {
}

func (o ListActiveActiveDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActiveActiveDomainsRequest struct{}"
	}

	return strings.Join([]string{"ListActiveActiveDomainsRequest", string(data)}, " ")
}
