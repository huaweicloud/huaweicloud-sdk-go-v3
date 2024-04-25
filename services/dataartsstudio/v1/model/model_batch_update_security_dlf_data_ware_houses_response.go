package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateSecurityDlfDataWareHousesResponse Response Object
type BatchUpdateSecurityDlfDataWareHousesResponse struct {

	// 细粒度认证更新结果
	Results        *[]FgacUpdateResult `json:"results,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o BatchUpdateSecurityDlfDataWareHousesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateSecurityDlfDataWareHousesResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateSecurityDlfDataWareHousesResponse", string(data)}, " ")
}
