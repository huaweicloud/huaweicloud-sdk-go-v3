package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmBody struct {

	// SMN的topic urn
	TopicUrn string `json:"topic_urn"`
}

func (o AlarmBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmBody struct{}"
	}

	return strings.Join([]string{"AlarmBody", string(data)}, " ")
}
