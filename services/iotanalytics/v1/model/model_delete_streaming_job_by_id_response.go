package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStreamingJobByIdResponse Response Object
type DeleteStreamingJobByIdResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteStreamingJobByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStreamingJobByIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteStreamingJobByIdResponse", string(data)}, " ")
}
