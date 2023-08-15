package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BizVersionManageVo 业务版本管理.
type BizVersionManageVo struct {

	// ID信息
	Id *int64 `json:"id,omitempty"`

	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	// 业务id
	BizId *int64 `json:"biz_id,omitempty"`

	// 业务对象信息
	BizInfo *string `json:"biz_info,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 业务版本
	BizVersion *int32 `json:"biz_version,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o BizVersionManageVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BizVersionManageVo struct{}"
	}

	return strings.Join([]string{"BizVersionManageVo", string(data)}, " ")
}
