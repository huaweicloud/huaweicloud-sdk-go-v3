package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePasswordComplexityStatusResponse Response Object
type ChangePasswordComplexityStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangePasswordComplexityStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePasswordComplexityStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangePasswordComplexityStatusResponse", string(data)}, " ")
}
