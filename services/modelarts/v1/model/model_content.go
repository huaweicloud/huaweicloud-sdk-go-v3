package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Content 对象的元数据信息。
type Content struct {
	Owner *Owner `json:"owner,omitempty"`

	// **参数解释**：对象的base64编码的128位MD5摘要。ETag是对象内容的唯一标识，可以通过该值识别对象内容是否有变化。比如上传对象时ETag为A，下载对象时ETag为B，则说明对象内容发生了变化。实际的ETag是对象的哈希值。ETag只反映变化的内容，而不是其元数据。上传的对象或拷贝操作创建的对象，通过MD5加密后都有唯一的ETag。 **约束限制**：当对象是服务端加密的对象时，ETag值不是对象的MD5值，而是通过服务端加密计算出的唯一标识。 **取值范围**：长度为32的字符串。
	Etag *string `json:"etag,omitempty"`

	// **参数解释**：对象类型。 **约束限制**：非Normal对象时返回。 **取值范围**：不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：对象名。对象名是对象在存储桶中的唯一标识。对象名是对象在桶中的完整路径，路径中不包含桶名。 **取值范围**：长度大于0且不超过1024的字符串。
	Key *string `json:"key,omitempty"`

	// **参数解释**：对象最近一次被修改的时间（UTC时间）。 **约束限制**：日期格式为ISO8601的格式。 **取值范围**：不涉及。
	LastModified *string `json:"last_modified,omitempty"`

	// **参数解释**：对象的字节数。 **取值范围**：0~48.8TB，单位：字节。
	Size *int64 `json:"size,omitempty"`

	// **参数解释**：对象的存储类型。 **取值范围**：不涉及。
	StorageClass *string `json:"storage_class,omitempty"`
}

func (o Content) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Content struct{}"
	}

	return strings.Join([]string{"Content", string(data)}, " ")
}
