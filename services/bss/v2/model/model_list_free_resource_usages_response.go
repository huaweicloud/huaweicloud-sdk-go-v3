package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFreeResourceUsagesResponse Response Object
type ListFreeResourceUsagesResponse struct {

	// 资源套餐内的资源项信息（资源项ID级的详情），具体参见表2。
	FreeResources  *[]FreeResourceDetail `json:"free_resources,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListFreeResourceUsagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFreeResourceUsagesResponse struct{}"
	}

	return strings.Join([]string{"ListFreeResourceUsagesResponse", string(data)}, " ")
}
