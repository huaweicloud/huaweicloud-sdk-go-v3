package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogtankRequest Request Object
type ListLogtankRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`
}

func (o ListLogtankRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogtankRequest struct{}"
	}

	return strings.Join([]string{"ListLogtankRequest", string(data)}, " ")
}
