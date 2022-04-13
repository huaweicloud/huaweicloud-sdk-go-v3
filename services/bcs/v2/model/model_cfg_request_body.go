package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CfgRequestBody struct {
	// 链代码名称，以小写字母开头，支持小写字母和数字，长度6-25位

	ChaincodeName string `json:"chaincode_name"`
	// SDK配置文件存放路径

	CertPath string `json:"cert_path"`
	// 通道名称

	ChannelName string `json:"channel_name"`
	// key：组织名，value：该组织下需要下载的peer节点信息，peer节点请按照0,1,2的顺序升序填写

	PeerOrgs map[string][]string `json:"peer_orgs"`
	// key：联盟成员名称，value：该联盟成员peer组织名称hash值数组

	UnionInfo map[string][]string `json:"union_info,omitempty"`
	// 是否是多通道请求，如此处设成true则必须传入channel_chaincode，chaincode_name和channel_name设为空即可

	IsMultiChan *bool `json:"is_multi_chan,omitempty"`
	// key：通道名称，value：该通道对应的链代码数组

	ChannelChaincode map[string][]string `json:"channel_chaincode,omitempty"`
}

func (o CfgRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CfgRequestBody struct{}"
	}

	return strings.Join([]string{"CfgRequestBody", string(data)}, " ")
}
