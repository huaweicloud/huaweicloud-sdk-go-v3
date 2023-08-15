package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LockViewResponse Response Object
type LockViewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o LockViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockViewResponse struct{}"
	}

	return strings.Join([]string{"LockViewResponse", string(data)}, " ")
}
