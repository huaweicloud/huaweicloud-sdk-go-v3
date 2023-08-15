package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortalInfosResponseModel 查询主页列表响应体。
type ListPortalInfosResponseModel struct {

	// 主页列表。
	Portals *[]PortalModel `json:"portals,omitempty"`

	PageInfo *PageOffSet `json:"page_info,omitempty"`
}

func (o ListPortalInfosResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortalInfosResponseModel struct{}"
	}

	return strings.Join([]string{"ListPortalInfosResponseModel", string(data)}, " ")
}
