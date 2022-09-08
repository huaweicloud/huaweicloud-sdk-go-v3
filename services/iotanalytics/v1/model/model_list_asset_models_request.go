package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAssetModelsRequest struct {

	// 过滤条件 示例： {\"key\":\"xxx\"} {\"key1\":\"xxx\",\"key2\":\"xxx\"} {\"key\":{\"eq|like\":\"xxx\"}} {\"key\":{\"in\":[\"xxx\",\"xxx\"]}} {\"or\":{\"key1\":\"xxx\",\"key2\":{\"eq|like\":\"xxx\"},\"key3\":{\"in\":[\"xxx\",\"xxx\"]}}} {\"and\":{\"key1\":\"xxx\",\"key2\":{\"eq|like\":\"xxx\"},\"key3\":{\"in\":[\"xxx\",\"xxx\"]}}} 支持的key： asset_model_id，name，display_name，job_id 注意： job_id只支持contain过滤 {\"job_id\":{\"contain\":\"xxx\"}}
	Filter *string `json:"filter,omitempty"`

	// 每页记录数
	Limit *int32 `json:"limit,omitempty"`

	// 页码
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAssetModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetModelsRequest struct{}"
	}

	return strings.Join([]string{"ListAssetModelsRequest", string(data)}, " ")
}
