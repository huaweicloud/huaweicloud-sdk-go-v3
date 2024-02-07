package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddGeipSegmentTagsResponse Response Object
type AddGeipSegmentTagsResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddGeipSegmentTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddGeipSegmentTagsResponse struct{}"
	}

	return strings.Join([]string{"AddGeipSegmentTagsResponse", string(data)}, " ")
}
