package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type NovaCreateKeypairRequest struct {

	// 微版本头
	OpenStackAPIVersion *string `json:"OpenStack-API-Version,omitempty" xml:"OpenStack-API-Version"`

	Body *NovaCreateKeypairRequestBody `json:"body,omitempty" xml:"body"`
}

func (o NovaCreateKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateKeypairRequest struct{}"
	}

	return strings.Join([]string{"NovaCreateKeypairRequest", string(data)}, " ")
}
