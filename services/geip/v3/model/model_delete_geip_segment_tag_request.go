package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGeipSegmentTagRequest Request Object
type DeleteGeipSegmentTagRequest struct {

	// 全域弹性公网IP的id
	ResourceId string `json:"resource_id"`

	// 待删除标签的key
	TagKey string `json:"tag_key"`
}

func (o DeleteGeipSegmentTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGeipSegmentTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteGeipSegmentTagRequest", string(data)}, " ")
}
