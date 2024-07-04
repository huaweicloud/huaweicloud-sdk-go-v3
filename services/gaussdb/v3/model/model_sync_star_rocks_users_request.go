package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncStarRocksUsersRequest Request Object
type SyncStarRocksUsersRequest struct {

	// StarRocks实例ID。
	InstanceId string `json:"instance_id"`

	// 内容类型。 取值：application/json。
	ContentType string `json:"Content-Type"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	Body *UserSyncReq `json:"body,omitempty"`
}

func (o SyncStarRocksUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncStarRocksUsersRequest struct{}"
	}

	return strings.Join([]string{"SyncStarRocksUsersRequest", string(data)}, " ")
}
