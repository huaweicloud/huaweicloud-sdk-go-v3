package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetTree struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	CreatorName *string `json:"creator_name,omitempty"`

	CreatorNum *string `json:"creator_num,omitempty"`

	UpdateName *string `json:"update_name,omitempty"`

	UpdateNum *string `json:"update_num,omitempty"`

	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	AssetId *string `json:"asset_id,omitempty"`

	ParentId *string `json:"parent_id,omitempty"`

	FactorCnt *int32 `json:"factor_cnt,omitempty"`
}

func (o AssetTree) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetTree struct{}"
	}

	return strings.Join([]string{"AssetTree", string(data)}, " ")
}
