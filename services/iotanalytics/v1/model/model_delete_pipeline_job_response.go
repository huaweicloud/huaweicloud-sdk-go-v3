package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeletePipelineJobResponse struct {
	Body           *string `json:"body,omitempty" xml:"body"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePipelineJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipelineJobResponse struct{}"
	}

	return strings.Join([]string{"DeletePipelineJobResponse", string(data)}, " ")
}
