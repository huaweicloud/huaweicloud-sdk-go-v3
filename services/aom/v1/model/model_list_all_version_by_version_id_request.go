package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllVersionByVersionIdRequest Request Object
type ListAllVersionByVersionIdRequest struct {
	Body *SearchScriptsRequestBody `json:"body,omitempty"`
}

func (o ListAllVersionByVersionIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllVersionByVersionIdRequest struct{}"
	}

	return strings.Join([]string{"ListAllVersionByVersionIdRequest", string(data)}, " ")
}
