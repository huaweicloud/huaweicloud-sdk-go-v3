package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type NovaShowKeypairRequest struct {

	// 密钥名称信息。
	KeypairName string `json:"keypair_name" xml:"keypair_name"`

	// 微版本头
	OpenStackAPIVersion *string `json:"OpenStack-API-Version,omitempty" xml:"OpenStack-API-Version"`
}

func (o NovaShowKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaShowKeypairRequest struct{}"
	}

	return strings.Join([]string{"NovaShowKeypairRequest", string(data)}, " ")
}
