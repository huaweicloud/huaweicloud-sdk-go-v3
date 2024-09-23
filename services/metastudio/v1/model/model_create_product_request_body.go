package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProductRequestBody 服务开通请求
type CreateProductRequestBody struct {

	// 商品名称
	Name string `json:"name"`

	// 商品描述
	Description *string `json:"description,omitempty"`

	// 标签。单个标签16字节，多个用逗号分隔，最多50个。
	Tags *[]string `json:"tags,omitempty"`

	Cover *ProductCoverInfo `json:"cover,omitempty"`

	// 文本列表
	TextList *[]ProductTextInfo `json:"text_list,omitempty"`

	// 素材资产列表
	AssetList *[]ProductMediaInfo `json:"asset_list,omitempty"`

	// **参数解释**： 自动激活商品
	AutoActive *bool `json:"auto_active,omitempty"`
}

func (o CreateProductRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductRequestBody struct{}"
	}

	return strings.Join([]string{"CreateProductRequestBody", string(data)}, " ")
}
