package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowResourcesListResponseBody struct {
	// 资源列表对象。

	Resources []ShowResourcesDetailResponseBody `json:"resources"`
}

func (o ShowResourcesListResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourcesListResponseBody struct{}"
	}

	return strings.Join([]string{"ShowResourcesListResponseBody", string(data)}, " ")
}
