package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportExportObsObjectsResponse Response Object
type ImportExportObsObjectsResponse struct {

	// OBS桶名
	BucketName *string `json:"bucket_name,omitempty"`

	// 列举桶内对象列表时，指定一个标识符，作为列举时的起始位置
	Marker *string `json:"marker,omitempty"`

	// 如果本次没有返回全部结果，响应请求中将包含此字段，用于标明本次请求列举到的最后一个对象
	NextMarker *string `json:"next_marker,omitempty"`

	// 表明本次请求是否返回了全部结果
	Truncated *bool `json:"truncated,omitempty"`

	// 列举对象的最大数目
	MaxKeys *int32 `json:"max_keys,omitempty"`

	// 列举桶内对象列表时，指定一个前缀
	Prefix *string `json:"prefix,omitempty"`

	// 分组信息
	CommonPrefixes *[]string `json:"common_prefixes,omitempty"`

	// 对象的元数据信息
	Contents       *[]ObsObjectInfo `json:"contents,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ImportExportObsObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportExportObsObjectsResponse struct{}"
	}

	return strings.Join([]string{"ImportExportObsObjectsResponse", string(data)}, " ")
}
