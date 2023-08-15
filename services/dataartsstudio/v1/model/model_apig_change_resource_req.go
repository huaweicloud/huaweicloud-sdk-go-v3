package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApigChangeResourceReq 规格变更信息请求体
type ApigChangeResourceReq struct {

	// 规格变更类型：10：升配；30：降配；40：续费；60：扩容；70：切换操作系统
	ChangeMode int32 `json:"change_mode"`

	// 资源id
	ResourceId string `json:"resource_id"`

	// 资源规格编码
	ResourceSpecCode string `json:"resource_spec_code"`

	// 产品id
	ProductId *string `json:"product_id,omitempty"`

	// 促销信息
	PromotionInfo *string `json:"promotion_info,omitempty"`
}

func (o ApigChangeResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigChangeResourceReq struct{}"
	}

	return strings.Join([]string{"ApigChangeResourceReq", string(data)}, " ")
}
