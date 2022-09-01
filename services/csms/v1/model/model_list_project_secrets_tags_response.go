package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProjectSecretsTagsResponse struct {

	// 标签列表，key和value键值对的集合。  - key：表示标签键，一个凭据下最多包含10个key，key不能为空，不能重复，同一个key中value不能重复。key最大长度为36个字符。  - value：表示标签值。每个值最大长度43个字符，value之间为“与”的关系。
	Tags           *[]Tag `json:"tags,omitempty" xml:"tags"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProjectSecretsTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectSecretsTagsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectSecretsTagsResponse", string(data)}, " ")
}
