package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHtapDataStoreResponse Response Object
type ListHtapDataStoreResponse struct {

	// 数据库信息列表。
	Datastores     *[]SrDataStoresDatastores `json:"datastores,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListHtapDataStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHtapDataStoreResponse struct{}"
	}

	return strings.Join([]string{"ListHtapDataStoreResponse", string(data)}, " ")
}
