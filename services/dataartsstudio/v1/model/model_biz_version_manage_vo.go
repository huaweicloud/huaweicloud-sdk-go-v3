package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BizVersionManageVo 业务版本管理。
type BizVersionManageVo struct {

	// ID信息。
	Id *int64 `json:"id,omitempty"`

	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	// 业务ID。
	BizId *int64 `json:"biz_id,omitempty"`

	// 业务对象信息。
	BizInfo *string `json:"biz_info,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 业务版本。
	BizVersion *int32 `json:"biz_version,omitempty"`

	// 创建时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o BizVersionManageVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BizVersionManageVo struct{}"
	}

	return strings.Join([]string{"BizVersionManageVo", string(data)}, " ")
}
