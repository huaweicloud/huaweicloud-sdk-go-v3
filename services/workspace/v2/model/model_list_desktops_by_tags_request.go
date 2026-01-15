package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopsByTagsRequest Request Object
type ListDesktopsByTagsRequest struct {
	Body *QueryDesktopByTagReq `json:"body,omitempty"`
}

func (o ListDesktopsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopsByTagsRequest", string(data)}, " ")
}
