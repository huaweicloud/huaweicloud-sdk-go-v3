package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBlockchainsResponse struct {

	// 服务实例简要信息
	Blockchains *[]BlockchainInfo `json:"blockchains,omitempty" xml:"blockchains"`

	// 实例总数
	Count          *int64 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBlockchainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBlockchainsResponse struct{}"
	}

	return strings.Join([]string{"ListBlockchainsResponse", string(data)}, " ")
}
