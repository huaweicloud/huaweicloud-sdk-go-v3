package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddImageTagRequest Request Object
type AddImageTagRequest struct {

	// 镜像ID。
	ImageId string `json:"image_id"`

	Body *AddImageTagRequestBody `json:"body,omitempty"`
}

func (o AddImageTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageTagRequest struct{}"
	}

	return strings.Join([]string{"AddImageTagRequest", string(data)}, " ")
}
