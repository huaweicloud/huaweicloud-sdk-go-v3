package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteShareFilesRequest struct {
	Body *DeleteShareFilesRequestBody `json:"body,omitempty"`
}

func (o DeleteShareFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShareFilesRequest struct{}"
	}

	return strings.Join([]string{"DeleteShareFilesRequest", string(data)}, " ")
}
