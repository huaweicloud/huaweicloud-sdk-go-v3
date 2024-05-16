package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessPointResponse Response Object
type ShowAccessPointResponse struct {
	Id *int32 `json:"id,omitempty"`

	// 创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`

	// 修改时间
	GmtModify *string `json:"gmt_modify,omitempty"`

	// 当前局点
	Region *string `json:"region,omitempty"`

	// 接入点地址
	AccessPoint *string `json:"accessPoint,omitempty"`

	// token
	Token *string `json:"token,omitempty"`

	// token隐藏字符
	HiddenToken *string `json:"hidden_token,omitempty"`

	// 应用ID
	SwBusinessId *int32 `json:"sw_business_id,omitempty"`

	// agent下载地址
	AgentDownloadUrl *string `json:"agent_download_url,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowAccessPointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessPointResponse struct{}"
	}

	return strings.Join([]string{"ShowAccessPointResponse", string(data)}, " ")
}
