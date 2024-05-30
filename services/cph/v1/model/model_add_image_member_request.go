package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddImageMemberRequest Request Object
type AddImageMemberRequest struct {

	// 镜像id。
	ImageId string `json:"image_id"`

	Body *AddImageMemberRequestBody `json:"body,omitempty"`
}

func (o AddImageMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageMemberRequest struct{}"
	}

	return strings.Join([]string{"AddImageMemberRequest", string(data)}, " ")
}
