package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteGeipSegmentTagsResponse Response Object
type BatchDeleteGeipSegmentTagsResponse struct {
	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteGeipSegmentTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteGeipSegmentTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteGeipSegmentTagsResponse", string(data)}, " ")
}
