package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceInternalEndpointResponse Response Object
type CreateInstanceInternalEndpointResponse struct {

	// 访问控制ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceInternalEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceInternalEndpointResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceInternalEndpointResponse", string(data)}, " ")
}
