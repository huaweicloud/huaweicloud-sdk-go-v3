package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPackagesResponse struct {

	// 习题库数量
	Total *int32 `json:"total,omitempty"`

	// 习题库列表
	Data           *[]PackageCard `json:"data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListPackagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPackagesResponse struct{}"
	}

	return strings.Join([]string{"ListPackagesResponse", string(data)}, " ")
}
