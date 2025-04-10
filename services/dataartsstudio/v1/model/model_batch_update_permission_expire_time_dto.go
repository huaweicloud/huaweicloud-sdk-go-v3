package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdatePermissionExpireTimeDto struct {

	// 数据密级id列表，数据密级id可以通过查询接口获取。
	Ids *[]string `json:"ids,omitempty"`

	// 到期时间
	ExpireTime *int64 `json:"expire_time,omitempty"`
}

func (o BatchUpdatePermissionExpireTimeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePermissionExpireTimeDto struct{}"
	}

	return strings.Join([]string{"BatchUpdatePermissionExpireTimeDto", string(data)}, " ")
}
