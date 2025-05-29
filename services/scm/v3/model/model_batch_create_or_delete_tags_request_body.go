package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateOrDeleteTagsRequestBody struct {

	// 标签操作标识。 - create：创建。 - delete：删除。
	Action string `json:"action"`

	// 标签列表，key和value键值对的集合。
	Tags *[]ScsResourceTag `json:"tags,omitempty"`
}

func (o BatchCreateOrDeleteTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteTagsRequestBody", string(data)}, " ")
}
