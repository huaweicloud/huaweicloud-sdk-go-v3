package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAzDetailsResponse Response Object
type ShowAzDetailsResponse struct {

	// 对应CMDB的region数据中的'zoneCode'字段。
	Id *string `json:"id,omitempty"`

	// 对应CMDB的region数据中的'regionCode'字段。
	RegionId *string `json:"region_id,omitempty"`

	// 当前AZ的类型: - Edge: 边缘云 - Workspace：华为云
	Type *string `json:"type,omitempty"`

	// 英文名。
	DisplayName *string `json:"display_name,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	SoldOut *SoldOutInfo `json:"sold_out,omitempty"`

	// 参品Id集。
	ProductIds *[]string `json:"product_ids,omitempty"`

	// 计费模式，专属 / 共享。
	Mode *string `json:"mode,omitempty"`

	// az的别名(中文、数字、字母、下划线、中划线，最大128字节)。
	Alias *string `json:"alias,omitempty"`

	// EIP所属的group。
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 分类的Id:  - 0: default - 21: HomeZone - 41: IES
	Category       *int32 `json:"category,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAzDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAzDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowAzDetailsResponse", string(data)}, " ")
}
