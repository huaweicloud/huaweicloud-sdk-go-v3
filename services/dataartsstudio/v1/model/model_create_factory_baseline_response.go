package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFactoryBaselineResponse Response Object
type CreateFactoryBaselineResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateFactoryBaselineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFactoryBaselineResponse struct{}"
	}

	return strings.Join([]string{"CreateFactoryBaselineResponse", string(data)}, " ")
}
