package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObsObjectsResponse Response Object
type ListObsObjectsResponse struct {

	// 桶名称
	BucketName *string `json:"bucket_name,omitempty"`

	// 当前页marker
	Marker *string `json:"marker,omitempty"`

	// 下一页marker
	NextMarker *string `json:"next_marker,omitempty"`

	// 文件夹列表
	CommonPrefixes *[]string `json:"common_prefixes,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListObsObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListObsObjectsResponse", string(data)}, " ")
}
