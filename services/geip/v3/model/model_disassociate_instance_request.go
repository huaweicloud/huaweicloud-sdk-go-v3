package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateInstanceRequest Request Object
type DisassociateInstanceRequest struct {

	// 解绑实例后是否保留GCB，false表示解绑实例后会同时删除GCB
	IsReserveGcb bool `json:"is_reserve_gcb"`

	GlobalEipId string `json:"global_eip_id"`
}

func (o DisassociateInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateInstanceRequest struct{}"
	}

	return strings.Join([]string{"DisassociateInstanceRequest", string(data)}, " ")
}
