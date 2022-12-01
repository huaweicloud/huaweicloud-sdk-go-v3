package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDataSourceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDataSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataSourceResponse struct{}"
	}

	return strings.Join([]string{"UpdateDataSourceResponse", string(data)}, " ")
}
