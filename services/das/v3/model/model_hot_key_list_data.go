package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HotKeyListData struct {

	// 总记录数
	Total *int64 `json:"total,omitempty"`

	// 热key列表
	KeyList *[]HotKeyInfo `json:"key_list,omitempty"`
}

func (o HotKeyListData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotKeyListData struct{}"
	}

	return strings.Join([]string{"HotKeyListData", string(data)}, " ")
}
