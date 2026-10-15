package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BigKeyListData struct {

	// 总记录数
	Total *int64 `json:"total,omitempty"`

	// 大key列表
	KeyList *[]BigKeyInfo `json:"key_list,omitempty"`
}

func (o BigKeyListData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BigKeyListData struct{}"
	}

	return strings.Join([]string{"BigKeyListData", string(data)}, " ")
}
