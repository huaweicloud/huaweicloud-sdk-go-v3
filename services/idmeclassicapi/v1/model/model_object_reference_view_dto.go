package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObjectReferenceViewDto struct {

	// 类名。
	Clazz *string `json:"clazz,omitempty"`

	// 唯一标识。
	Id string `json:"id"`
}

func (o ObjectReferenceViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectReferenceViewDto struct{}"
	}

	return strings.Join([]string{"ObjectReferenceViewDto", string(data)}, " ")
}
