package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateClusterTagsBody struct {

	// 集群标签的key值。可输入的字符串长度为1~36个字符。只能包含数字、字母、中划线\"-\"和下划线\"_\"。
	Key string `json:"key"`

	// 集群标签的value值。可输入的字符串长度为0~43个字符。只能包含数字、字母、中划线\"-\"和下划线\"_\"。
	Value string `json:"value"`
}

func (o CreateClusterTagsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterTagsBody struct{}"
	}

	return strings.Join([]string{"CreateClusterTagsBody", string(data)}, " ")
}
