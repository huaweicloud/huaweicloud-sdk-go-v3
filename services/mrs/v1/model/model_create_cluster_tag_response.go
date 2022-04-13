package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateClusterTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateClusterTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterTagResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterTagResponse", string(data)}, " ")
}
