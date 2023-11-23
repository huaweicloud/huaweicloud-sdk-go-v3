package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFunctionAppResponse Response Object
type CreateFunctionAppResponse struct {

	// 应用id
	ApplicationId  *string `json:"application_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFunctionAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionAppResponse struct{}"
	}

	return strings.Join([]string{"CreateFunctionAppResponse", string(data)}, " ")
}
