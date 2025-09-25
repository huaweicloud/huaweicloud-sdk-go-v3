package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageSynchronizeTaskResponse Response Object
type CreateImageSynchronizeTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateImageSynchronizeTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageSynchronizeTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateImageSynchronizeTaskResponse", string(data)}, " ")
}
