package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceAvailableNums 所有类型的资源可用数量
type ResourceAvailableNums struct {

	// flexus版资源数。
	Flexus *int32 `json:"flexus,omitempty"`

	// 基础版资源数。
	Basic *int32 `json:"basic,omitempty"`

	// 进阶版资源数。
	Middle *int32 `json:"middle,omitempty"`

	// 进阶测试版资源数。
	MiddleOnDemand *int32 `json:"middle_on_demand,omitempty"`

	// 高级版资源数。
	Advance *int32 `json:"advance,omitempty"`

	// 出门问问资源数。
	ThirdPartyCmww *int32 `json:"third_party_cmww,omitempty"`
}

func (o ResourceAvailableNums) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceAvailableNums struct{}"
	}

	return strings.Join([]string{"ResourceAvailableNums", string(data)}, " ")
}
