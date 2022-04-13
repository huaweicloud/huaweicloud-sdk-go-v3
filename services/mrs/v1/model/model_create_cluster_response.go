package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateClusterResponse struct {
	//   操作结果。  true：操作成功  false：操作失败

	Result *bool `json:"result,omitempty"`
	// 系统提示信息，可为空。

	Msg *string `json:"msg,omitempty"`
	// 集群创建成功后系统返回的集群ID值。

	ClusterId      *string `json:"cluster_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterResponse", string(data)}, " ")
}
