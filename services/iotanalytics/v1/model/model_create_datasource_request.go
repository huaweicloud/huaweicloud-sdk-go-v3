package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDatasourceRequest struct {
	Body *CreateDatasourceReqDto `json:"body,omitempty"`
}

func (o CreateDatasourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatasourceRequest struct{}"
	}

	return strings.Join([]string{"CreateDatasourceRequest", string(data)}, " ")
}
