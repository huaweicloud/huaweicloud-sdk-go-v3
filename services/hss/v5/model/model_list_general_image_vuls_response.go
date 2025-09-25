package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGeneralImageVulsResponse Response Object
type ListGeneralImageVulsResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 漏洞详情
	DataList       *[]GeneralImageVulsResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListGeneralImageVulsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeneralImageVulsResponse struct{}"
	}

	return strings.Join([]string{"ListGeneralImageVulsResponse", string(data)}, " ")
}
