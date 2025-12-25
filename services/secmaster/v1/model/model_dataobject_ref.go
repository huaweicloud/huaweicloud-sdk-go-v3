package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataobjectRef 数据对象关键字段
type DataobjectRef struct {

	// 唯一标识ID
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`
}

func (o DataobjectRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataobjectRef struct{}"
	}

	return strings.Join([]string{"DataobjectRef", string(data)}, " ")
}
