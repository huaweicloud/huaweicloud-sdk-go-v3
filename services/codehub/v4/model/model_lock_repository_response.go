package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LockRepositoryResponse Response Object
type LockRepositoryResponse struct {

	// 锁定状态
	Locked         *string `json:"locked,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o LockRepositoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockRepositoryResponse struct{}"
	}

	return strings.Join([]string{"LockRepositoryResponse", string(data)}, " ")
}
