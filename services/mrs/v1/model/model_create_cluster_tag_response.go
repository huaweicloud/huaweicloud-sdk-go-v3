package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterTagResponse Response Object
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
