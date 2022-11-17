package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListServiceItemsDetailsRequest struct {

	// 服务组id
	SetId string `json:"set_id"`

	// 查询字段
	KeyWord *string `json:"key_word,omitempty"`

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`
}

func (o ListServiceItemsDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceItemsDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListServiceItemsDetailsRequest", string(data)}, " ")
}
