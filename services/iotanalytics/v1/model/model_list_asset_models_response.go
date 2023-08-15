package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetModelsResponse Response Object
type ListAssetModelsResponse struct {

	// 总数
	Count *int64 `json:"count,omitempty"`

	// 模型集，数量不超过limit
	AssetModels    *[]AssetModelResponse `json:"asset_models,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListAssetModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetModelsResponse struct{}"
	}

	return strings.Join([]string{"ListAssetModelsResponse", string(data)}, " ")
}
