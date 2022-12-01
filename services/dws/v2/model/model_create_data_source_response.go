package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDataSourceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDataSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataSourceResponse struct{}"
	}

	return strings.Join([]string{"CreateDataSourceResponse", string(data)}, " ")
}
