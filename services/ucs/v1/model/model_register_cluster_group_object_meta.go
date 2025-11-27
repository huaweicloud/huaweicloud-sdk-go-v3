package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterClusterGroupObjectMeta 容器舰队元数据信息。
type RegisterClusterGroupObjectMeta struct {

	// 容器舰队名称
	Name string `json:"name"`
}

func (o RegisterClusterGroupObjectMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterClusterGroupObjectMeta struct{}"
	}

	return strings.Join([]string{"RegisterClusterGroupObjectMeta", string(data)}, " ")
}
