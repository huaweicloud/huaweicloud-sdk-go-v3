package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEdgeDDoSDomainsRequest Request Object
type CreateEdgeDDoSDomainsRequest struct {
	Body *CreateEdgeDDoSDomainsRequestBody `json:"body,omitempty"`
}

func (o CreateEdgeDDoSDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeDDoSDomainsRequest struct{}"
	}

	return strings.Join([]string{"CreateEdgeDDoSDomainsRequest", string(data)}, " ")
}
