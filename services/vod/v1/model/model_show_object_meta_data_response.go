package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowObjectMetaDataResponse Response Object
type ShowObjectMetaDataResponse struct {

	// 桶名
	Bucket *string `json:"bucket,omitempty"`

	// 下次列举对象请求的起始位置。
	NextMarker *string `json:"next_marker,omitempty"`

	// 媒体元数据列表
	ObjectList     *[]ObjectList `json:"object_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowObjectMetaDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowObjectMetaDataResponse struct{}"
	}

	return strings.Join([]string{"ShowObjectMetaDataResponse", string(data)}, " ")
}
