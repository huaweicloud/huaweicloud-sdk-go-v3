package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePackageNameResponse Response Object
type UpdatePackageNameResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePackageNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePackageNameResponse struct{}"
	}

	return strings.Join([]string{"UpdatePackageNameResponse", string(data)}, " ")
}
