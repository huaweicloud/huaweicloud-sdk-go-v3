package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObjectList struct {

	// 对象文件名
	Object *string `json:"object,omitempty"`

	// 媒资ID
	AssetId *string `json:"asset_id,omitempty"`

	MetaData *ObjectMetaData `json:"meta_data,omitempty"`
}

func (o ObjectList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectList struct{}"
	}

	return strings.Join([]string{"ObjectList", string(data)}, " ")
}
