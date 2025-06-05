package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteComputingClusterResponse Response Object
type DeleteComputingClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteComputingClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComputingClusterResponse struct{}"
	}

	return strings.Join([]string{"DeleteComputingClusterResponse", string(data)}, " ")
}
