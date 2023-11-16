package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPackageRequest Request Object
type ListPackageRequest struct {
}

func (o ListPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPackageRequest struct{}"
	}

	return strings.Join([]string{"ListPackageRequest", string(data)}, " ")
}
