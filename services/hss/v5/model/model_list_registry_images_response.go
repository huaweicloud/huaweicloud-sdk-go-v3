package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRegistryImagesResponse Response Object
type ListRegistryImagesResponse struct {

	// **参数解释**: 满足查询条件的镜像记录总数量 **取值范围**: 0-2147483547；单位：个
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 查询到的仓库镜像详细信息列表 **取值范围**: 数组长度0-当前查询的limit值（1-200），数组元素为RegistryImagesInfo对象
	DataList       *[]RegistryImagesInfo `json:"data_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListRegistryImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegistryImagesResponse struct{}"
	}

	return strings.Join([]string{"ListRegistryImagesResponse", string(data)}, " ")
}
