package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgencyAccessResponse Response Object
type CreateAgencyAccessResponse struct {
	Body           *[]LtsAccessConfigInfoRespon200 `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o CreateAgencyAccessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyAccessResponse struct{}"
	}

	return strings.Join([]string{"CreateAgencyAccessResponse", string(data)}, " ")
}
