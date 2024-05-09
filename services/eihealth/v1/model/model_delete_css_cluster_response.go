package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCssClusterResponse Response Object
type DeleteCssClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCssClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCssClusterResponse struct{}"
	}

	return strings.Join([]string{"DeleteCssClusterResponse", string(data)}, " ")
}
