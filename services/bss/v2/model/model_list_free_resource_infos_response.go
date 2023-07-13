package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFreeResourceInfosResponse Response Object
type ListFreeResourceInfosResponse struct {

	// 资源包信息列表，具体参见表2。
	FreeResourcePackages *[]FreeResourcePackageV3 `json:"free_resource_packages,omitempty"`

	// 总条数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFreeResourceInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFreeResourceInfosResponse struct{}"
	}

	return strings.Join([]string{"ListFreeResourceInfosResponse", string(data)}, " ")
}
