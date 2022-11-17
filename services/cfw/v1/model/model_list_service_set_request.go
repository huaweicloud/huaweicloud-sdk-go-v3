package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListServiceSetRequest struct {

	// 防护对象id
	ObjectId string `json:"object_id"`

	// 关键字
	KeyWord *string `json:"key_word,omitempty"`

	// 每页查询个数
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`
}

func (o ListServiceSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceSetRequest struct{}"
	}

	return strings.Join([]string{"ListServiceSetRequest", string(data)}, " ")
}
