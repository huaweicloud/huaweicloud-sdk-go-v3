package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExtensionDetailResponse Response Object
type ShowExtensionDetailResponse struct {

	// 插件列表查询结果集合
	Results        *[]ExtensionQueryResult `json:"results,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowExtensionDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExtensionDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowExtensionDetailResponse", string(data)}, " ")
}
