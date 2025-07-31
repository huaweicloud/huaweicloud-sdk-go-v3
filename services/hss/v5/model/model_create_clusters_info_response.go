package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClustersInfoResponse Response Object
type CreateClustersInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateClustersInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClustersInfoResponse struct{}"
	}

	return strings.Join([]string{"CreateClustersInfoResponse", string(data)}, " ")
}
