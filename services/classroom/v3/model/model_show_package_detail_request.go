package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPackageDetailRequest Request Object
type ShowPackageDetailRequest struct {

	// 需查询的习题库id
	PackageId string `json:"package_id"`
}

func (o ShowPackageDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPackageDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowPackageDetailRequest", string(data)}, " ")
}
