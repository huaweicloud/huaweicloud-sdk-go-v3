package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandClusterComponentResponse Response Object
type ExpandClusterComponentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExpandClusterComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandClusterComponentResponse struct{}"
	}

	return strings.Join([]string{"ExpandClusterComponentResponse", string(data)}, " ")
}
