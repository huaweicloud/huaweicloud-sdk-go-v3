package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAssetModelsResponse struct {

	// 总数
	Count *int64 `json:"count,omitempty" xml:"count"`

	// 模型集，数量不超过limit
	AssetModels    *[]AssetModelResponse `json:"asset_models,omitempty" xml:"asset_models"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListAssetModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetModelsResponse struct{}"
	}

	return strings.Join([]string{"ListAssetModelsResponse", string(data)}, " ")
}
