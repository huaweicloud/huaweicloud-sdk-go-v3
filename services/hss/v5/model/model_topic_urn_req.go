package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopicUrnReq **参数解释**: 主题urn **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
type TopicUrnReq struct {
}

func (o TopicUrnReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicUrnReq struct{}"
	}

	return strings.Join([]string{"TopicUrnReq", string(data)}, " ")
}
