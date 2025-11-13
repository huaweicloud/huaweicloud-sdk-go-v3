package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateResolverQueryLogConfigRequestBody struct {

	// 日志组ID。
	LtsGroupId string `json:"lts_group_id"`

	// 日志流ID。
	LtsTopicId string `json:"lts_topic_id"`

	Vpc *Vpc `json:"vpc"`
}

func (o CreateResolverQueryLogConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResolverQueryLogConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResolverQueryLogConfigRequestBody", string(data)}, " ")
}
