package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListObsBucketsResponseBody struct {

	// **参数解释**： 响应头。
	ResponseHeaders *interface{} `json:"responseHeaders,omitempty"`

	// **参数解释**： 原始响应头。
	OriginalHeaders *interface{} `json:"originalHeaders,omitempty"`

	// **参数解释**： 状态码。 **取值范围**： 不涉及。
	StatusCode *int32 `json:"statusCode,omitempty"`

	// **参数解释**： 桶名。 **取值范围**： 不涉及。
	BucketName *string `json:"bucketName,omitempty"`

	// **参数解释**： 桶拥有者信息。
	Owner *interface{} `json:"owner,omitempty"`

	// **参数解释**： 桶的创建时间。 **取值范围**： 长度为24的字符串。
	CreationDate *int64 `json:"creationDate,omitempty"`

	// **参数解释**： 桶所在的区域。 **取值范围**： 不涉及。
	Location *string `json:"location,omitempty"`

	// **参数解释**： 集群类型。 **取值范围**： 不涉及。
	Clustertype *string `json:"clustertype,omitempty"`

	// **参数解释**： 存储类型。 **取值范围**： 不涉及。
	StorageClass *string `json:"storageClass,omitempty"`

	// **参数解释**： 元数据。
	Metadata *interface{} `json:"metadata,omitempty"`

	// **参数解释**： 桶ACL策略。 **取值范围**： 不涉及。
	Acl *string `json:"acl,omitempty"`

	// **参数解释**： 桶的存储类型。 **取值范围**： - STANDARD：标准存储。 - WARM：低频访问存储。 - COLD：归档存储。 - DEEP_ARCHIVE：深度归档存储（受限公测）。
	BucketStorageClass *string `json:"bucketStorageClass,omitempty"`

	// **参数解释**： 桶类型。 **取值范围**： 不涉及。
	BucketType *string `json:"bucketType,omitempty"`

	// **参数解释**： 请求ID。 **取值范围**： 不涉及。
	RequestId *string `json:"requestId,omitempty"`
}

func (o ListObsBucketsResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsBucketsResponseBody struct{}"
	}

	return strings.Join([]string{"ListObsBucketsResponseBody", string(data)}, " ")
}
