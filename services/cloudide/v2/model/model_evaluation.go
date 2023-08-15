package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Evaluation struct {

	// 插件id
	ExtensionId string `json:"extension_id"`

	// 评论内容
	Text string `json:"text"`
}

func (o Evaluation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Evaluation struct{}"
	}

	return strings.Join([]string{"Evaluation", string(data)}, " ")
}
