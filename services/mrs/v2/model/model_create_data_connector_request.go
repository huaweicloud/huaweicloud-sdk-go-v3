package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataConnectorRequest Request Object
type CreateDataConnectorRequest struct {
	Body *DataConnectorReq `json:"body,omitempty"`
}

func (o CreateDataConnectorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataConnectorRequest struct{}"
	}

	return strings.Join([]string{"CreateDataConnectorRequest", string(data)}, " ")
}
