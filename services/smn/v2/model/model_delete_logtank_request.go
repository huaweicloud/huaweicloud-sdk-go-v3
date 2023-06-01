package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteLogtankRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`

	// 云日志信息唯一的资源标识。可通过[查询云日志](ListLogtank.xml)获取该标识。
	LogtankId string `json:"logtank_id"`
}

func (o DeleteLogtankRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogtankRequest struct{}"
	}

	return strings.Join([]string{"DeleteLogtankRequest", string(data)}, " ")
}
