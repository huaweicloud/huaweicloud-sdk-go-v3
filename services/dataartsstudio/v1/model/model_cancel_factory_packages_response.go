package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelFactoryPackagesResponse Response Object
type CancelFactoryPackagesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelFactoryPackagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelFactoryPackagesResponse struct{}"
	}

	return strings.Join([]string{"CancelFactoryPackagesResponse", string(data)}, " ")
}
