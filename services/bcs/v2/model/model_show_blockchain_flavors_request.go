package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBlockchainFlavorsRequest Request Object
type ShowBlockchainFlavorsRequest struct {

	// 取值范围(0,1000]，默认值为1000。用于限制本次返回的结果数据条数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始位置，为非负整数。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowBlockchainFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockchainFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ShowBlockchainFlavorsRequest", string(data)}, " ")
}
