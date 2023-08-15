package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLiveDataQuotaV2Response Response Object
type ListLiveDataQuotaV2Response struct {

	// 数据源配额
	Datasource *string `json:"datasource,omitempty"`

	// 后端api配额
	Api *string `json:"api,omitempty"`

	// 脚本配额
	Scripts *string `json:"scripts,omitempty"`

	// 已使用的数据源数量
	DatasourceUsed *string `json:"datasource_used,omitempty"`

	// 已使用的后端api数量
	ApiUsed        *string `json:"api_used,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLiveDataQuotaV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLiveDataQuotaV2Response struct{}"
	}

	return strings.Join([]string{"ListLiveDataQuotaV2Response", string(data)}, " ")
}
