package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPageAssetListResultResponse Response Object
type ShowPageAssetListResultResponse struct {

	// 资产详情
	Assets *[]Asset `json:"assets,omitempty"`

	Pagination     *Pagination `json:"pagination,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowPageAssetListResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPageAssetListResultResponse struct{}"
	}

	return strings.Join([]string{"ShowPageAssetListResultResponse", string(data)}, " ")
}
