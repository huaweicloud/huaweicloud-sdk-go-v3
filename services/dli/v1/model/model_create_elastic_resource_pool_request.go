package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateElasticResourcePoolRequest Request Object
type CreateElasticResourcePoolRequest struct {
	Body *CreateElasticResourcePoolRequestBody `json:"body,omitempty"`
}

func (o CreateElasticResourcePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateElasticResourcePoolRequest struct{}"
	}

	return strings.Join([]string{"CreateElasticResourcePoolRequest", string(data)}, " ")
}
