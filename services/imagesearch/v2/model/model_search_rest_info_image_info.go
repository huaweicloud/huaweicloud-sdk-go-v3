package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchRestInfoImageInfo 搜索图像的相关信息，不同服务类型返回信息不同，具体可参见服务类型说明。
type SearchRestInfoImageInfo struct {

	// 用于搜索的主体目标框。
	Box *string `json:"box,omitempty"`

	// 用于搜索的主体类目序号。
	Category float32 `json:"category,omitempty"`

	// 用于搜索的主体类目名称。
	CategoryName *string `json:"category_name,omitempty"`

	// 搜索图像中的所有主体列表。
	Objects *[]AddDataRestInfoImageInfoObjects `json:"objects,omitempty"`
}

func (o SearchRestInfoImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchRestInfoImageInfo struct{}"
	}

	return strings.Join([]string{"SearchRestInfoImageInfo", string(data)}, " ")
}
