package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ImportPointsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportPointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportPointsResponse struct{}"
	}

	return strings.Join([]string{"ImportPointsResponse", string(data)}, " ")
}
