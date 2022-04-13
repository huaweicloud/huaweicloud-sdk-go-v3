package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDomainsRequest struct {
	Body *CreateDomainsRequestBody `json:"body,omitempty"`
}

func (o CreateDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainsRequest struct{}"
	}

	return strings.Join([]string{"CreateDomainsRequest", string(data)}, " ")
}
