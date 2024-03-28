package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatasourceConnectionRequest Request Object
type CreateDatasourceConnectionRequest struct {
	Body *CreateDatasourceConnectionRequestBody `json:"body,omitempty"`
}

func (o CreateDatasourceConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatasourceConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateDatasourceConnectionRequest", string(data)}, " ")
}
