package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnlockRepositoryResponse Response Object
type UnlockRepositoryResponse struct {

	// 锁定状态
	Locked         *string `json:"locked,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnlockRepositoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockRepositoryResponse struct{}"
	}

	return strings.Join([]string{"UnlockRepositoryResponse", string(data)}, " ")
}
