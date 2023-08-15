package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaurusProxyScaleRequest proxy实例规格变更请求体
type TaurusProxyScaleRequest struct {

	// 需要变更的新规格ID。
	FlavorRef string `json:"flavor_ref"`
}

func (o TaurusProxyScaleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaurusProxyScaleRequest struct{}"
	}

	return strings.Join([]string{"TaurusProxyScaleRequest", string(data)}, " ")
}
