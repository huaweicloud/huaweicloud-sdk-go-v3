package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportOtherResourceRequest Request Object
type ImportOtherResourceRequest struct {
	Body *ImportOtherResourceRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportOtherResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportOtherResourceRequest struct{}"
	}

	return strings.Join([]string{"ImportOtherResourceRequest", string(data)}, " ")
}
