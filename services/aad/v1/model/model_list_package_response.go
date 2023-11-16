package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPackageResponse Response Object
type ListPackageResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 数据
	Items          *[]PackageResponse `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListPackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPackageResponse struct{}"
	}

	return strings.Join([]string{"ListPackageResponse", string(data)}, " ")
}
