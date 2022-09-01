package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationModify struct {

	// 应用名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 应用描述。
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o ApplicationModify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationModify struct{}"
	}

	return strings.Join([]string{"ApplicationModify", string(data)}, " ")
}
