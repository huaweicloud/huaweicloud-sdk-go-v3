package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteUpdateImageNameRequest Request Object
type ExecuteUpdateImageNameRequest struct {

	// 图片id
	ImageId string `json:"image_id"`

	Body *UpdateImageNameReq `json:"body,omitempty"`
}

func (o ExecuteUpdateImageNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteUpdateImageNameRequest struct{}"
	}

	return strings.Join([]string{"ExecuteUpdateImageNameRequest", string(data)}, " ")
}
