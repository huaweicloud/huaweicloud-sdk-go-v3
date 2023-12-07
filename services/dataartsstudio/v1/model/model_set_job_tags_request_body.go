package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetJobTagsRequestBody struct {

	// 标签列表, 若未空，则标识删除标签
	TagNames *[]string `json:"tag_names,omitempty"`
}

func (o SetJobTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetJobTagsRequestBody struct{}"
	}

	return strings.Join([]string{"SetJobTagsRequestBody", string(data)}, " ")
}
