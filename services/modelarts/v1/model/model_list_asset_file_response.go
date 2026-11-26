package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetFileResponse Response Object
type ListAssetFileResponse struct {

	// **参数解释**：列举的对象内容。
	Contents *[]Content `json:"contents,omitempty"`

	// **参数解释**：表明本次返回的结果列表是否被截断。“true”表示本次没有返回全部结果；“false”表示本次已经返回了全部结果。 **取值范围**：true,false。
	IsTruncated *bool `json:"is_truncated,omitempty"`

	// **参数解释**：列举桶内对象列表时，指定一个标识符，作为列举时的起始位置，从该标识符以后按对象名的字典顺序返回对象列表。 **取值范围**：长度大于0且不超过1024的字符串。
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：如果本次没有返回全部结果，响应请求中将包含此字段，用于标明本次请求列举到的最后一个对象。后续请求可以指定Marker等于该值来列举剩余的对象。 **取值范围**：长度大于0且不超过1024的字符串。
	NextMarker *string `json:"next_marker,omitempty"`

	// **参数解释**：列举对象的最大数目，返回的对象列表将是按照字典顺序的最多前max_keys个对象。 **取值范围**：1-1000。
	MaxKeys *int32 `json:"max_keys,omitempty"`

	// **参数解释**：本次请求的桶名。 **约束限制**：桶的名字需全局唯一，不能与已有的任何桶名称重复，包括其他用户创建的桶。   桶命名规则如下：   - 3～63个字符，数字或字母开头，支持小写字母、数字、“-”、“.”。   - 禁止使用IP地址。   - 禁止以“-”或“.”开头及结尾。   - 禁止两个“.”相邻（如：“my..bucket”）。   - 禁止“.”和“-”相邻（如：“my-.bucket”和“my.-bucket”）。   - 同一用户在同一个区域多次创建同名桶不会报错，创建的桶属性以第一次请求为准。 **取值范围**：不涉及。
	BucketName *string `json:"bucket_name,omitempty"`

	// **参数解释**：列举桶内对象列表时，指定一个前缀，限定返回的对象名必须带有prefix前缀。 **取值范围**：不涉及。
	Prefix *string `json:"prefix,omitempty"`

	// **参数解释**：列举桶内对象列表时，指定前缀列表，限定返回的对象名必须带有prefix前缀。
	CommonPrefixes *[]string `json:"common_prefixes,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAssetFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetFileResponse struct{}"
	}

	return strings.Join([]string{"ListAssetFileResponse", string(data)}, " ")
}
