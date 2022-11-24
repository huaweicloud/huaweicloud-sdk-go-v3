package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EIP操作防护请求体
type EipOperateProtectReq struct {

	// 防护对象ID
	ObjectId string `json:"object_id"`

	// EIP状态
	Status int32 `json:"status"`

	// EIP信息列表
	IpInfos []EipOperateProtectReqIpInfos `json:"ip_infos"`
}

func (o EipOperateProtectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipOperateProtectReq struct{}"
	}

	return strings.Join([]string{"EipOperateProtectReq", string(data)}, " ")
}
