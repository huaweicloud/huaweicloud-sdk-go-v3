package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DebugSecurityDlfDataWareHousesResponse Response Object
type DebugSecurityDlfDataWareHousesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DebugSecurityDlfDataWareHousesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugSecurityDlfDataWareHousesResponse struct{}"
	}

	return strings.Join([]string{"DebugSecurityDlfDataWareHousesResponse", string(data)}, " ")
}
