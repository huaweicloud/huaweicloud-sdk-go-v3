package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEndpointServiceServerResourceResponse Response Object
type AddEndpointServiceServerResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddEndpointServiceServerResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEndpointServiceServerResourceResponse struct{}"
	}

	return strings.Join([]string{"AddEndpointServiceServerResourceResponse", string(data)}, " ")
}
