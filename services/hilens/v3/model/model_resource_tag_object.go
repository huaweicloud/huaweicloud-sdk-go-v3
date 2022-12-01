package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTagObject struct {

	// 资源ID，不同资源（节点，部署，配置项，密钥）有不同的资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 标签键，最大长度36个字符。不能为空，只能包含大小写字母，数字，中划线“-”，下划线“_”
	Key *string `json:"key,omitempty"`

	// 标签值，每个值最大长度43个字符，删除时如果value有值按照key/value删除，如果value没值，则按照key删除。不能为空，只能包含大小写字母，数字，中划线“-”，下划线“_”
	Value *string `json:"value,omitempty"`
}

func (o ResourceTagObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagObject struct{}"
	}

	return strings.Join([]string{"ResourceTagObject", string(data)}, " ")
}
