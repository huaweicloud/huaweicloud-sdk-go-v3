package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEdgeWafDomainsRequest Request Object
type CreateEdgeWafDomainsRequest struct {
	Body *CreateEdgeWafDomainsRequestBody `json:"body,omitempty"`
}

func (o CreateEdgeWafDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeWafDomainsRequest struct{}"
	}

	return strings.Join([]string{"CreateEdgeWafDomainsRequest", string(data)}, " ")
}
