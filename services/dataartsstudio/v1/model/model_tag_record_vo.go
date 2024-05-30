package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagRecordVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 标签ID，填写String类型替代Long类型。
	TagId string `json:"tag_id"`

	// 标签名称。
	TagName *string `json:"tag_name,omitempty"`

	// 实体ID，填写String类型替代Long类型。
	BizId string `json:"biz_id"`

	BizType *BizTypeEnum `json:"biz_type"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o TagRecordVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagRecordVo struct{}"
	}

	return strings.Join([]string{"TagRecordVo", string(data)}, " ")
}
