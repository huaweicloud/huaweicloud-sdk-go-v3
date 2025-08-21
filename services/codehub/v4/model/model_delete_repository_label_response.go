package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRepositoryLabelResponse Response Object
type DeleteRepositoryLabelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRepositoryLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRepositoryLabelResponse struct{}"
	}

	return strings.Join([]string{"DeleteRepositoryLabelResponse", string(data)}, " ")
}
