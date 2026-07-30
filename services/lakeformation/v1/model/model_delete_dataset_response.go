package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatasetResponse Response Object
type DeleteDatasetResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDatasetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatasetResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatasetResponse", string(data)}, " ")
}
