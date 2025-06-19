package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpBlacklistRequest Request Object
type ListIpBlacklistRequest struct {

	// 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId string `json:"fw_instance_id"`

	// 在分页查询的情况下，每页查询的记录条数，范围为1-1024
	Limit int32 `json:"limit"`

	// 数据查询的偏移量，在分页查询的时候使用，指定查询记录的起始位置，必须为数字，取值范围为大于等于0
	Offset int32 `json:"offset"`
}

func (o ListIpBlacklistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpBlacklistRequest struct{}"
	}

	return strings.Join([]string{"ListIpBlacklistRequest", string(data)}, " ")
}
