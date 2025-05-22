package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFactoryPendingItemsPackageResponse Response Object
type CreateFactoryPendingItemsPackageResponse struct {

	// 发布包名称
	PackageName *string `json:"package_name,omitempty"`

	// 发布包ID
	ReleasePackageId *string `json:"release_package_id,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o CreateFactoryPendingItemsPackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFactoryPendingItemsPackageResponse struct{}"
	}

	return strings.Join([]string{"CreateFactoryPendingItemsPackageResponse", string(data)}, " ")
}
