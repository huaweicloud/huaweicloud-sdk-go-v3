package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityDlfDataWareHousesResponse Response Object
type ListSecurityDlfDataWareHousesResponse struct {

	// 数据开发细粒度连接列表
	DwLists        *[]DataWareHouseDto `json:"dw_lists,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListSecurityDlfDataWareHousesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDlfDataWareHousesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityDlfDataWareHousesResponse", string(data)}, " ")
}
