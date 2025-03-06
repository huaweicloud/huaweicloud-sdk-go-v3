package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateImageMemberRequest Request Object
type UpdateImageMemberRequest struct {

	// 镜像id。
	ImageId string `json:"image_id"`

	Body *UpdateImageMemberRequestBody `json:"body,omitempty"`
}

func (o UpdateImageMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageMemberRequest struct{}"
	}

	return strings.Join([]string{"UpdateImageMemberRequest", string(data)}, " ")
}
