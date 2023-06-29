package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetsNewRequest Request Object
type ListAssetsNewRequest struct {

	// 每页记录数
	Limit *int32 `json:"limit,omitempty"`

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 查询过滤器 示例： {\"key\":\"xxx\"} {\"key1\":\"xxx\",\"key2\":\"xxx\"} {\"key\":{\"eq|like\":\"xxx\"}} {\"key\":{\"in\":[\"xxx\",\"xxx\"]}} {\"or\":{\"key1\":\"xxx\",\"key2\":{\"eq|like\":\"xxx\"},\"key3\":{\"in\":[\"xxx\",\"xxx\"]}}} {\"and\":{\"key1\":\"xxx\",\"key2\":{\"eq|like\":\"xxx\"},\"key3\":{\"in\":[\"xxx\",\"xxx\"]}}} 支持的key： asset_model_id，asset_id，parent，name，display_name，root，state，job_id 注意： job_id只在RELEASE态下生效，只支持contain过滤 {\"job_id\":{\"contain\":\"xxx\"}}
	Filter *string `json:"filter,omitempty"`

	// SKETCH：草稿态；RELEASE：发布态
	Type string `json:"type"`
}

func (o ListAssetsNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetsNewRequest struct{}"
	}

	return strings.Join([]string{"ListAssetsNewRequest", string(data)}, " ")
}
