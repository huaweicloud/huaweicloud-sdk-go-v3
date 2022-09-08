package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModResourceDto struct {

	// 资源标识
	Id string `json:"id"`

	// 资源类型，企业内ID和TYPE唯一标识一个资源项，若只传资源ID可能会修改多个资源的信息. - VMR        - 云会议室 - CONF_CALL  - 会议并发数 - HARD_1080P - 1080P硬终端 - HARD_720P  - 720P硬终端 - SOFT       - 软终端用户数 - ROOM       - 大屏软终端 - LIVE       - 直播推流 - RECORD     - 录播空间 - HARD_THIRD_PARTY - 第三方硬终端账号 - HUAWEI_VISION -智慧屏
	Type *string `json:"type,omitempty"`

	// 到期时间
	ExpireDate *int64 `json:"expireDate,omitempty"`

	// 资源是否被停用
	IsDisabled *bool `json:"isDisabled,omitempty"`
}

func (o ModResourceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModResourceDto struct{}"
	}

	return strings.Join([]string{"ModResourceDto", string(data)}, " ")
}
