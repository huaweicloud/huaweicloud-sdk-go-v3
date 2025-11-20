package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainConfigResponse Response Object
type UpdateDomainConfigResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateDomainConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainConfigResponse", string(data)}, " ")
}
