package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EipOperateProtectReq EIP操作防护请求体
type EipOperateProtectReq struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用查询防火墙实例接口获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。具体可参考APIExlorer和帮助中心FAQ。
	ObjectId string `json:"object_id"`

	// EIP状态，0表示防护中，1表示未防护
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
