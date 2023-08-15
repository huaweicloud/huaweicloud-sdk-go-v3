package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageRequest Request Object
type DeleteImageRequest struct {

	// 镜像ID
	ImageId string `json:"image_id"`
}

func (o DeleteImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageRequest struct{}"
	}

	return strings.Join([]string{"DeleteImageRequest", string(data)}, " ")
}
