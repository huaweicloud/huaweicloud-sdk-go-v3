package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteClusterTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteClusterTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteClusterTagResponse", string(data)}, " ")
}
