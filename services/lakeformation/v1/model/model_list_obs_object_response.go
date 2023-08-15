package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObsObjectResponse Response Object
type ListObsObjectResponse struct {

	// object名称列表
	ObjectNames *[]string `json:"object_names,omitempty"`

	// region编码
	Location *string `json:"location,omitempty"`

	// 搜索前缀
	Prefix *string `json:"prefix,omitempty"`

	// obs桶名
	BucketName *string `json:"bucket_name,omitempty"`

	// 下一个object起始位置
	NextMarker     *string `json:"next_marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListObsObjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsObjectResponse struct{}"
	}

	return strings.Join([]string{"ListObsObjectResponse", string(data)}, " ")
}
