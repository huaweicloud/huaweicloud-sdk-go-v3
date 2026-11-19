package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopicUrnRes **参数解释**: 主题urn **取值范围**: 字符长度0-256位
type TopicUrnRes struct {
}

func (o TopicUrnRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicUrnRes struct{}"
	}

	return strings.Join([]string{"TopicUrnRes", string(data)}, " ")
}
