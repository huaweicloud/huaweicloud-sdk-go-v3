package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTaskParameterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ListTaskParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskParameterResponse struct{}"
	}

	return strings.Join([]string{"ListTaskParameterResponse", string(data)}, " ")
}
