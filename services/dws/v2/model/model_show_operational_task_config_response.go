package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOperationalTaskConfigResponse Response Object
type ShowOperationalTaskConfigResponse struct {
	Configuration  *OperationalTaskConfiguration `json:"configuration,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowOperationalTaskConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOperationalTaskConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowOperationalTaskConfigResponse", string(data)}, " ")
}
