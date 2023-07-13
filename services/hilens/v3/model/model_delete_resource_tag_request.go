package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceTagRequest Request Object
type DeleteResourceTagRequest struct {

	// 资源ID，不同资源（节点，部署，配置项，密钥）有不同的资源ID
	ResourceId string `json:"resource_id"`

	// 资源类型（节点，部署，配置项，密钥）
	ResourceType string `json:"resource_type"`

	// 标签键，最大长度36个字符。不能为空，只能包含大小写字母，数字，中划线“-”，下划线“_”
	Key string `json:"key"`
}

func (o DeleteResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceTagRequest", string(data)}, " ")
}
