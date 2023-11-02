package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertDetailDataclassRef 数据类对象
type AlertDetailDataclassRef struct {

	// 数据类唯一标识，UUID格式，最大36个字符
	Id *string `json:"id,omitempty"`

	// 数据类名称
	Name *string `json:"name,omitempty"`
}

func (o AlertDetailDataclassRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertDetailDataclassRef struct{}"
	}

	return strings.Join([]string{"AlertDetailDataclassRef", string(data)}, " ")
}
