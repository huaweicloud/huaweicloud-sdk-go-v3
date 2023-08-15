package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagVo struct {

	// 编码
	Id *int64 `json:"id,omitempty"`

	// 标签名
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`
}

func (o TagVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagVo struct{}"
	}

	return strings.Join([]string{"TagVo", string(data)}, " ")
}
