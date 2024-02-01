package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGcbResourceTagRequest Request Object
type DeleteGcbResourceTagRequest struct {

	// 资源唯一标识符。
	ResourceId string `json:"resource_id"`

	// 删除的tag的key。
	TagKey string `json:"tag_key"`
}

func (o DeleteGcbResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGcbResourceTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteGcbResourceTagRequest", string(data)}, " ")
}
