package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询结果返回体。
type ListPortInfosResponseModel struct {
	PageInfo *PageOffSet `json:"page_info,omitempty"`

	// 请求成功返回的数据。
	Data *[]ListPortInfosResponseModelData `json:"data,omitempty"`
}

func (o ListPortInfosResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortInfosResponseModel struct{}"
	}

	return strings.Join([]string{"ListPortInfosResponseModel", string(data)}, " ")
}
