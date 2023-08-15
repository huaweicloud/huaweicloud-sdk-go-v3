package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportPointsResponse Response Object
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
