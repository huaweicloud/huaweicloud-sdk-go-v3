package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadCertsResponse Response Object
type UploadCertsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UploadCertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadCertsResponse struct{}"
	}

	return strings.Join([]string{"UploadCertsResponse", string(data)}, " ")
}
