package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObsBucketInfo struct {

	// 响应头
	ResponseHeaders *interface{} `json:"responseHeaders,omitempty"`

	// 请求头
	OriginalHeaders *interface{} `json:"originalHeaders,omitempty"`

	// 状态
	StatusCode *int32 `json:"statusCode,omitempty"`

	// 桶名称
	BucketName *string `json:"bucketName,omitempty"`

	Owner *ObsBucketInfoOwner `json:"owner,omitempty"`

	// 桶的创建时间
	CreationDate *string `json:"creationDate,omitempty"`

	// 桶的位置信息
	Location *string `json:"location,omitempty"`

	// 对象的存储类型
	StorageClass *interface{} `json:"storageClass,omitempty"`

	// 桶元数据
	Metadata *interface{} `json:"metadata,omitempty"`

	// 桶ACL
	Acl *interface{} `json:"acl,omitempty"`

	// 桶的存储类型
	BucketStorageClass *interface{} `json:"bucketStorageClass,omitempty"`

	// 桶类型
	BucketType *string `json:"bucketType,omitempty"`

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
}

func (o ObsBucketInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsBucketInfo struct{}"
	}

	return strings.Join([]string{"ObsBucketInfo", string(data)}, " ")
}
