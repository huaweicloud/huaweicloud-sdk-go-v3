package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性IP白名单。
type CreateClusterElbWhiteList struct {
	// 是否开启访问控制。

	EnableWhiteList bool `json:"enableWhiteList"`
	// 访问控制白名单。

	WhiteList *string `json:"whiteList,omitempty"`
}

func (o CreateClusterElbWhiteList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterElbWhiteList struct{}"
	}

	return strings.Join([]string{"CreateClusterElbWhiteList", string(data)}, " ")
}
