package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAssetsNewResponse struct {
	// 总数

	Count *int64 `json:"count,omitempty"`
	// 资产集，数量不超过limit

	Assets         *[]AssetResponse `json:"assets,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListAssetsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetsNewResponse struct{}"
	}

	return strings.Join([]string{"ListAssetsNewResponse", string(data)}, " ")
}
