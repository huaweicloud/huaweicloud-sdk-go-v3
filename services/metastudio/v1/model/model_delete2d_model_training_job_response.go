package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Delete2dModelTrainingJobResponse Response Object
type Delete2dModelTrainingJobResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o Delete2dModelTrainingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Delete2dModelTrainingJobResponse struct{}"
	}

	return strings.Join([]string{"Delete2dModelTrainingJobResponse", string(data)}, " ")
}
