package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type NovaListKeypairsRequest struct {

	// 查询返回秘钥数量限制。  在微版本2.35后支持
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 从marker指定的keypair的名称的下一条数据开始查询。  在微版本2.35后支持。
	Marker *string `json:"marker,omitempty" xml:"marker"`

	// 微版本头
	OpenStackAPIVersion *string `json:"OpenStack-API-Version,omitempty" xml:"OpenStack-API-Version"`
}

func (o NovaListKeypairsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListKeypairsRequest struct{}"
	}

	return strings.Join([]string{"NovaListKeypairsRequest", string(data)}, " ")
}
