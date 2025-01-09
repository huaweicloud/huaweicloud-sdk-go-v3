package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternetResponse Response Object
type ListInternetResponse struct {

	// 上网信息
	InternetInfos  *[]InternetInfo `json:"internet_infos,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListInternetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternetResponse struct{}"
	}

	return strings.Join([]string{"ListInternetResponse", string(data)}, " ")
}
