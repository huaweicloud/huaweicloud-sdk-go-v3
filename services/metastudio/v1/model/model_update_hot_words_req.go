package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHotWordsReq 修改热词记录请求。
type UpdateHotWordsReq struct {
	HotWordsType *HotWordsTypeEnum `json:"hot_words_type"`

	SisHotWords *UpdateSisHotWords `json:"sis_hot_words"`
}

func (o UpdateHotWordsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHotWordsReq struct{}"
	}

	return strings.Join([]string{"UpdateHotWordsReq", string(data)}, " ")
}
