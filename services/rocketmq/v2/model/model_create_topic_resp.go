package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTopicResp struct {

	// 主题名称。
	Id *string `json:"id,omitempty"`
}

func (o CreateTopicResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicResp struct{}"
	}

	return strings.Join([]string{"CreateTopicResp", string(data)}, " ")
}
