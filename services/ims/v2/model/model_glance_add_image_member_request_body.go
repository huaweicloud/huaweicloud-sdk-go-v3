package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 添加镜像成员请求参数
type GlanceAddImageMemberRequestBody struct {

	// 镜像成员。取值为目标用户的项目ID。
	Member string `json:"member"`
}

func (o GlanceAddImageMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceAddImageMemberRequestBody struct{}"
	}

	return strings.Join([]string{"GlanceAddImageMemberRequestBody", string(data)}, " ")
}
