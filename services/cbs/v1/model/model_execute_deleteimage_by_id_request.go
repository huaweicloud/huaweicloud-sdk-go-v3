package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteDeleteimageByIdRequest struct {

	// 图片id
	ImageId string `json:"image_id"`
}

func (o ExecuteDeleteimageByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDeleteimageByIdRequest struct{}"
	}

	return strings.Join([]string{"ExecuteDeleteimageByIdRequest", string(data)}, " ")
}
