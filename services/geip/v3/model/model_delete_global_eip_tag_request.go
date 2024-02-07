package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGlobalEipTagRequest Request Object
type DeleteGlobalEipTagRequest struct {

	// 全域弹性公网IP的id
	ResourceId string `json:"resource_id"`

	// 待删除标签的key
	TagKey string `json:"tag_key"`
}

func (o DeleteGlobalEipTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalEipTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteGlobalEipTagRequest", string(data)}, " ")
}
