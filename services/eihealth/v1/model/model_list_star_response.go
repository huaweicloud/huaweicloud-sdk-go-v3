package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStarResponse Response Object
type ListStarResponse struct {

	// 资产总数
	Count *int32 `json:"count,omitempty"`

	// 资产列表
	Assets         *[]GetAssetRsp `json:"assets,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListStarResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStarResponse struct{}"
	}

	return strings.Join([]string{"ListStarResponse", string(data)}, " ")
}
