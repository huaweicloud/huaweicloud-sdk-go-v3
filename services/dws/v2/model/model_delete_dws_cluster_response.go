package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDwsClusterResponse Response Object
type DeleteDwsClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDwsClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDwsClusterResponse struct{}"
	}

	return strings.Join([]string{"DeleteDwsClusterResponse", string(data)}, " ")
}
