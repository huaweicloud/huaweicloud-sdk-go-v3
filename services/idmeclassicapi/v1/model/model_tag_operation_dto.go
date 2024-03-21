package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagOperationDto struct {

	// 模型对象ID。
	ObjectId string `json:"objectId"`

	// 标签ID。
	TagId string `json:"tagId"`
}

func (o TagOperationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagOperationDto struct{}"
	}

	return strings.Join([]string{"TagOperationDto", string(data)}, " ")
}
