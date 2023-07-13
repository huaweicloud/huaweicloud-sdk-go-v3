package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmFileUploadResponse Response Object
type ConfirmFileUploadResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ConfirmFileUploadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmFileUploadResponse struct{}"
	}

	return strings.Join([]string{"ConfirmFileUploadResponse", string(data)}, " ")
}
