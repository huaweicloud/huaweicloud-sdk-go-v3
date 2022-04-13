package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeletePipelineJobResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePipelineJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipelineJobResponse struct{}"
	}

	return strings.Join([]string{"DeletePipelineJobResponse", string(data)}, " ")
}
