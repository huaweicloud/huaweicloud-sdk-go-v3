package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceDeleteTagRequest Request Object
type GlanceDeleteTagRequest struct {

	// 镜像id
	ImageId string `json:"image_id"`

	// 新增的tag。字符串中不能包含“=”和“.”。
	Tag string `json:"tag"`
}

func (o GlanceDeleteTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceDeleteTagRequest struct{}"
	}

	return strings.Join([]string{"GlanceDeleteTagRequest", string(data)}, " ")
}
