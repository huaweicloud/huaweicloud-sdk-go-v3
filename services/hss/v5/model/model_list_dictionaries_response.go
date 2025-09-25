package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDictionariesResponse Response Object
type ListDictionariesResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 字典项数量
	DataList       *[]Dictionary `json:"data_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDictionariesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDictionariesResponse struct{}"
	}

	return strings.Join([]string{"ListDictionariesResponse", string(data)}, " ")
}
