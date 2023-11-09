package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopByTagsRequest Request Object
type ListDesktopByTagsRequest struct {
	Body *QueryDesktopByTagReq `json:"body,omitempty"`
}

func (o ListDesktopByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopByTagsRequest", string(data)}, " ")
}
