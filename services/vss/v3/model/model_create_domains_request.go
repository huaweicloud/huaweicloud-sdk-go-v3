package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDomainsRequest struct {
	Body *CreateDomainsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainsRequest struct{}"
	}

	return strings.Join([]string{"CreateDomainsRequest", string(data)}, " ")
}
