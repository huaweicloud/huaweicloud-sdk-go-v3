package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationModelResponse Response Object
type ListApplicationModelResponse struct {
	Data           *ApplicationModelQueryResponseData `json:"data,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListApplicationModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationModelResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationModelResponse", string(data)}, " ")
}
