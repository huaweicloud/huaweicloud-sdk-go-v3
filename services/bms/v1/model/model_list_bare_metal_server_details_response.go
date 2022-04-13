package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBareMetalServerDetailsResponse struct {
	Server         *ServerDetails `json:"server,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListBareMetalServerDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBareMetalServerDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListBareMetalServerDetailsResponse", string(data)}, " ")
}
