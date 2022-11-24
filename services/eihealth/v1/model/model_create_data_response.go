package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataResponse struct{}"
	}

	return strings.Join([]string{"CreateDataResponse", string(data)}, " ")
}
