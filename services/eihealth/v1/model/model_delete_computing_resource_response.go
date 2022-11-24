package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteComputingResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteComputingResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComputingResourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteComputingResourceResponse", string(data)}, " ")
}
