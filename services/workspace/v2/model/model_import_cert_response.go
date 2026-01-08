package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportCertResponse Response Object
type ImportCertResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ImportCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCertResponse struct{}"
	}

	return strings.Join([]string{"ImportCertResponse", string(data)}, " ")
}
