package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssociateRegistriesResponse Response Object
type ListAssociateRegistriesResponse struct {

	// 镜像仓列表
	DataList *[]AssociateRegistriesResponseInfo `json:"data_list,omitempty"`

	// 同步任务关联的镜像仓总数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAssociateRegistriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociateRegistriesResponse struct{}"
	}

	return strings.Join([]string{"ListAssociateRegistriesResponse", string(data)}, " ")
}
