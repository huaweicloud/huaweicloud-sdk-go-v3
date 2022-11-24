package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteBackendInstanceV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBackendInstanceV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackendInstanceV2Response struct{}"
	}

	return strings.Join([]string{"DeleteBackendInstanceV2Response", string(data)}, " ")
}
