package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHotWordsReq 创建热词请求。
type CreateHotWordsReq struct {
	HotWordsType *HotWordsTypeEnum `json:"hot_words_type"`

	// 应用ID。
	RobotId string `json:"robot_id"`

	SisHotWords *CreateSisHotWords `json:"sis_hot_words"`
}

func (o CreateHotWordsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHotWordsReq struct{}"
	}

	return strings.Join([]string{"CreateHotWordsReq", string(data)}, " ")
}
