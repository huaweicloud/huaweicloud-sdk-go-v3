package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcePackagesResponse Response Object
type ListResourcePackagesResponse struct {

	// 资源包列表。
	DesktopResourcePackages *[]DesktopResourcePackage `json:"desktop_resource_packages,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourcePackagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcePackagesResponse struct{}"
	}

	return strings.Join([]string{"ListResourcePackagesResponse", string(data)}, " ")
}
