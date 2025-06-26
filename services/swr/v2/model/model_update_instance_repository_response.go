package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceRepositoryResponse Response Object
type UpdateInstanceRepositoryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceRepositoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRepositoryResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRepositoryResponse", string(data)}, " ")
}
