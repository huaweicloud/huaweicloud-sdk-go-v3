package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DebugSecurityDlfDataWareHousesRequest Request Object
type DebugSecurityDlfDataWareHousesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 数据连接id
	DwId string `json:"dw_id"`
}

func (o DebugSecurityDlfDataWareHousesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugSecurityDlfDataWareHousesRequest struct{}"
	}

	return strings.Join([]string{"DebugSecurityDlfDataWareHousesRequest", string(data)}, " ")
}
