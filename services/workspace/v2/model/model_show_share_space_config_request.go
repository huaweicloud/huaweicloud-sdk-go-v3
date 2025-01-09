package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShareSpaceConfigRequest Request Object
type ShowShareSpaceConfigRequest struct {

	// 查询协同桌面默认用户配置, share-space-user-conf
	Name string `json:"name"`
}

func (o ShowShareSpaceConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShareSpaceConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowShareSpaceConfigRequest", string(data)}, " ")
}
