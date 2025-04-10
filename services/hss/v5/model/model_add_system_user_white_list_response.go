package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSystemUserWhiteListResponse Response Object
type AddSystemUserWhiteListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddSystemUserWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSystemUserWhiteListResponse struct{}"
	}

	return strings.Join([]string{"AddSystemUserWhiteListResponse", string(data)}, " ")
}
