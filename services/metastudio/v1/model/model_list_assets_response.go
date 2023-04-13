package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAssetsResponse struct {

	// 资产总数。
	Count *int32 `json:"count,omitempty"`

	// 资产信息列表。
	Assets         *[]DigitalAssetInfo `json:"assets,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListAssetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetsResponse struct{}"
	}

	return strings.Join([]string{"ListAssetsResponse", string(data)}, " ")
}
