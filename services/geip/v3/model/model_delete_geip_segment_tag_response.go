package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGeipSegmentTagResponse Response Object
type DeleteGeipSegmentTagResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGeipSegmentTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGeipSegmentTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteGeipSegmentTagResponse", string(data)}, " ")
}
