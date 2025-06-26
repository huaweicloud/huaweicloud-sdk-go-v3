package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceRepositoryResponse Response Object
type DeleteInstanceRepositoryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceRepositoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceRepositoryResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceRepositoryResponse", string(data)}, " ")
}
