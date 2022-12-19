package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpgradeClusterResponse struct {

	// 升级任务ID，可通过调用获取集群升级任务详情API查询进展
	Uid            *string `json:"uid,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeClusterResponse struct{}"
	}

	return strings.Join([]string{"UpgradeClusterResponse", string(data)}, " ")
}
