package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegistryInfo 镜像仓接入信息
type RegistryInfo struct {

	// **参数解释**: 镜像仓ID **取值范围**: 字符长度1-64位
	Id *string `json:"id,omitempty"`

	// **参数解释**: 仓库名称 **取值范围**: 字符长度1-128位
	RegistryName *string `json:"registry_name,omitempty"`

	// **参数解释**: 镜像仓类型 **取值范围**: - Harbor harbor仓库 - Jfrog jfrog仓库 - SwrPrivate swr私有 - SwrShared  swr共享 - SwrEnterprise  swr企业 - Other  其他仓库
	RegistryType *string `json:"registry_type,omitempty"`

	// **参数解释**： 镜像仓接口版本 **取值范围**：   - V1：V1版本   - V2：V2版本
	ApiVersion *string `json:"api_version,omitempty"`

	// **参数解释**： 协议类型 **取值范围**：   - http：http协议   - https：https协议
	Protocol *string `json:"protocol,omitempty"`

	// **参数解释**： 镜像仓地址 **取值范围**： 字符长度1-256位
	RegistryAddr *string `json:"registry_addr,omitempty"`

	// **参数解释**： 用于登录镜像仓的用户名。 **取值范围**： 字符长度1-128位
	RegistryUsername *string `json:"registry_username,omitempty"`

	// **参数解释**： 镜像仓项目,用来指定扫描组件要上传到的镜像仓目录。get_scan_image_channel为Other时返回此值。 **取值范围**： 字符长度1-128位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 承载集群id **取值范围**： 字符长度1-64位
	ConnectClusterId *string `json:"connect_cluster_id,omitempty"`

	// **参数解释**： 获取扫描组件的方式 **取值范围**： - Swr：访问swr获取扫描组件 - Other：手动上传扫描组件到承载集群。
	GetScanImageChannel *string `json:"get_scan_image_channel,omitempty"`

	// **参数解释**： 接入状态 **取值范围**： - success：接入成功 - fail：接入异常 - accessing：接入中
	Status *string `json:"status,omitempty"`

	// **参数解释**: 失败原因 **取值范围**: - CREATE_JOB_FAILED：集群接入状态异常，请检查集群接入状态。 - REQUEST_API_ERROR：网络不通，访问镜像仓失败。请检查承载集群是否能正常访问镜像仓 ，或前往三方镜像仓页面重新接入。 - SERVER_ERROR：系统内部错误，请稍后重试。
	FailReason *string `json:"fail_reason,omitempty"`

	// **参数解释**： 镜像数量 **取值范围**： 0-20000
	ImagesNum *int32 `json:"images_num,omitempty"`

	// **参数解释**： 最近同步时间，时间单位 毫秒（ms） **取值范围**： 0-9223372036854775807
	LatestSyncTime *int64 `json:"latest_sync_time,omitempty"`
}

func (o RegistryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegistryInfo struct{}"
	}

	return strings.Join([]string{"RegistryInfo", string(data)}, " ")
}
