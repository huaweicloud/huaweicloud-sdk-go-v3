package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CleanupModelResponse Response Object
type CleanupModelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CleanupModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CleanupModelResponse struct{}"
	}

	return strings.Join([]string{"CleanupModelResponse", string(data)}, " ")
}
