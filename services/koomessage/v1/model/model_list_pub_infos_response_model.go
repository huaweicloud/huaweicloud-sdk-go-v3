package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPubInfosResponseModel 获取服务号详情返回体。
type ListPubInfosResponseModel struct {
	PageInfo *Page `json:"page_info,omitempty"`

	// 服务号详情列表。
	Data *[]PubDetail `json:"data,omitempty"`
}

func (o ListPubInfosResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPubInfosResponseModel struct{}"
	}

	return strings.Join([]string{"ListPubInfosResponseModel", string(data)}, " ")
}
