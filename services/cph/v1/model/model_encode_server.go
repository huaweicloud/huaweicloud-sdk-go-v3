package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 编码容器信息。
type EncodeServer struct {

	// 编码服务的名称，不大于64个字节
	EncodeServerName *string `json:"encode_server_name,omitempty"`

	// 编码服务的唯一标识ID，不大于32个字节
	EncodeServerId *string `json:"encode_server_id,omitempty"`

	// 编码服务IP地址
	EncodeServerIp *string `json:"encode_server_ip,omitempty"`

	// 云手机服务器ID，不大于32个字节
	ServerId *string `json:"server_id,omitempty"`

	// 编码服务登录密钥名称
	KeypairName *string `json:"keypair_name,omitempty"`

	// 编码服务类型 - 0：服务器 - 1：容器
	Type *int32 `json:"type,omitempty"`

	// 编码服务状态 - 1：运行中 - 2：异常 - 3：重启中 - 0、4、5：创建中
	Status *int32 `json:"status,omitempty"`

	// 编码服务的访问信息
	AccessInfos *[]EncodeServerAccessInfo `json:"access_infos,omitempty"`
}

func (o EncodeServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncodeServer struct{}"
	}

	return strings.Join([]string{"EncodeServer", string(data)}, " ")
}
