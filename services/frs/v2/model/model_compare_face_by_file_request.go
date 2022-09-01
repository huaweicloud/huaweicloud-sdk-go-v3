package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CompareFaceByFileRequest struct {
	Body *CompareFaceByFileRequestBody `json:"body,omitempty" xml:"body" type:"multipart"`
}

func (o CompareFaceByFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareFaceByFileRequest struct{}"
	}

	return strings.Join([]string{"CompareFaceByFileRequest", string(data)}, " ")
}
