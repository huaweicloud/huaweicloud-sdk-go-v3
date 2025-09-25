package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRegistryImagesResponse Response Object
type ListRegistryImagesResponse struct {

	// **参数解释**: 总数 **取值范围**: 0-2147483547
	TotalNum *int32 `json:"total_num,omitempty"`

	// 仓库镜像列表
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
