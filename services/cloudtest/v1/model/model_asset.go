package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Asset struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	CreatorName *string `json:"creator_name,omitempty"`

	CreatorNum *string `json:"creator_num,omitempty"`

	UpdateName *string `json:"update_name,omitempty"`

	UpdateNum *string `json:"update_num,omitempty"`

	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o Asset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Asset struct{}"
	}

	return strings.Join([]string{"Asset", string(data)}, " ")
}
