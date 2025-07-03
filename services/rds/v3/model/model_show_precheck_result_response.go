package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrecheckResultResponse Response Object
type ShowPrecheckResultResponse struct {

	// 检查是否通过，0代表通过，1代表未通过
	ResultCode *int32 `json:"result_code,omitempty"`

	// 检查状态
	Status *string `json:"status,omitempty"`

	// 检查结果更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 是否展示数据库检查结果
	Display *bool `json:"display,omitempty"`

	// 实例关联关系检查失败项
	InstanceStatusCheckList *[]string `json:"instance_status_check_list,omitempty"`

	DbUpgradePrecheck *DbUpgradePrecheck `json:"db_upgrade_precheck,omitempty"`

	// 检查结果下载链接
	DownloadLink *string `json:"download_link,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPrecheckResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrecheckResultResponse struct{}"
	}

	return strings.Join([]string{"ShowPrecheckResultResponse", string(data)}, " ")
}
