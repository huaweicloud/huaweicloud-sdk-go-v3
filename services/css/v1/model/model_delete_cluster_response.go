package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClusterResponse Response Object
type DeleteClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterResponse struct{}"
	}

	return strings.Join([]string{"DeleteClusterResponse", string(data)}, " ")
}
