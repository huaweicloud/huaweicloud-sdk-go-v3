package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterSshStateResponse Response Object
type ListClusterSshStateResponse struct {

	// 查询指定集群节点授权状态 - -1：未开启集群节点授权 - 0：开启集群节点授权失败 - 1：正在开启集群节点授权 - 2：已开启集群节点授权 - 3：集群节点授权即将失效状态 - 4：集群节点授权已失效状态
	SshOpsStat     float32 `json:"sshOpsStat,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListClusterSshStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterSshStateResponse struct{}"
	}

	return strings.Join([]string{"ListClusterSshStateResponse", string(data)}, " ")
}
