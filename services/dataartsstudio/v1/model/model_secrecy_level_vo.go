package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecrecyLevelVo struct {

	// 密级ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 数据安全主键。
	Uuid *string `json:"uuid,omitempty"`

	// 密级名。
	Name string `json:"name"`

	// 密级等级。
	Slevel *int32 `json:"slevel,omitempty"`

	// 密级描述。
	Description *string `json:"description,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`
}

func (o SecrecyLevelVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecrecyLevelVo struct{}"
	}

	return strings.Join([]string{"SecrecyLevelVo", string(data)}, " ")
}
