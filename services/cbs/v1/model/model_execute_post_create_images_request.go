package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecutePostCreateImagesRequest Request Object
type ExecutePostCreateImagesRequest struct {
	Body *PostImagesReq `json:"body,omitempty" type:"multipart"`
}

func (o ExecutePostCreateImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutePostCreateImagesRequest struct{}"
	}

	return strings.Join([]string{"ExecutePostCreateImagesRequest", string(data)}, " ")
}
