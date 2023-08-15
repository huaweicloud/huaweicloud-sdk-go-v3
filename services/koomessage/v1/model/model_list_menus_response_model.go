package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMenusResponseModel 查询菜单返回体。
type ListMenusResponseModel struct {

	// 菜单信息。
	Data *[]MenusRsp `json:"data,omitempty"`

	PageInfo *PageOffSet `json:"page_info,omitempty"`
}

func (o ListMenusResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMenusResponseModel struct{}"
	}

	return strings.Join([]string{"ListMenusResponseModel", string(data)}, " ")
}
