package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppBaseResponse struct {

	// 创建的应用id
	Id *string `json:"id,omitempty"`

	// 创建应用名称
	Name *string `json:"name,omitempty"`

	// 应用所属区域
	Region *string `json:"region,omitempty"`

	// 部署任务列表信息
	ArrangeInfos *[]TaskBaseBody `json:"arrange_infos,omitempty"`
}

func (o AppBaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppBaseResponse struct{}"
	}

	return strings.Join([]string{"AppBaseResponse", string(data)}, " ")
}
