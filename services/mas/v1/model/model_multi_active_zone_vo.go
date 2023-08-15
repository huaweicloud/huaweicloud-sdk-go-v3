package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiActiveZoneVo struct {
	AvailableZone *[]string `json:"available_zone,omitempty"`

	CreatedDate *sdktime.SdkTime `json:"created_date,omitempty"`

	Description *string `json:"description,omitempty"`

	Id *string `json:"id,omitempty"`

	IsMaster *bool `json:"is_master,omitempty"`

	Name *string `json:"name,omitempty"`

	NamespaceId *string `json:"namespace_id,omitempty"`

	Region *string `json:"region,omitempty"`

	RegionName *string `json:"region_name,omitempty"`

	Spec *BaseMultiActiveZoneSpec `json:"spec,omitempty"`

	Type *int32 `json:"type,omitempty"`

	UpdatedDate *sdktime.SdkTime `json:"updated_date,omitempty"`
}

func (o MultiActiveZoneVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiActiveZoneVo struct{}"
	}

	return strings.Join([]string{"MultiActiveZoneVo", string(data)}, " ")
}
