package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type GlanceListImageMembersRequest struct {

	// 镜像id
	ImageId string `json:"image_id" xml:"image_id"`
}

func (o GlanceListImageMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageMembersRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImageMembersRequest", string(data)}, " ")
}
