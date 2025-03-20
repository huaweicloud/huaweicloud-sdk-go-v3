package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesignDataLayersResultData 接口返回的数据。
type ListDesignDataLayersResultData struct {

	// 数仓分层数组。
	Value *[]DataLayerVo `json:"value,omitempty"`
}

func (o ListDesignDataLayersResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesignDataLayersResultData struct{}"
	}

	return strings.Join([]string{"ListDesignDataLayersResultData", string(data)}, " ")
}
