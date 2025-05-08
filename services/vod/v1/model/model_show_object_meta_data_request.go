package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowObjectMetaDataRequest Request Object
type ShowObjectMetaDataRequest struct {

	// obs桶名称
	Bucket string `json:"bucket"`

	// obs对象路径
	Object *string `json:"object,omitempty"`

	// 列举桶内对象列表时，指定一个标识符，作为列举时的起始位置，从该标识符以后按对象名的字典顺序返回对象列表
	Marker *string `json:"marker,omitempty"`

	// 列举对象的最大数目，返回的对象列表将是按照字典顺序的最多前limit个对象
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowObjectMetaDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowObjectMetaDataRequest struct{}"
	}

	return strings.Join([]string{"ShowObjectMetaDataRequest", string(data)}, " ")
}
