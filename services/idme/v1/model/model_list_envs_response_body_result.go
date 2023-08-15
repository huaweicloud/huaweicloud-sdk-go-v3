package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListEnvsResponseBodyResult struct {

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 区域ID。
	RegionId *string `json:"region_id,omitempty"`

	// 运行服务的ID。
	EnvId *string `json:"env_id,omitempty"`

	// 运行服务的名称。
	EnvName *string `json:"env_name,omitempty"`

	// 运行服务的状态。
	EnvStatus *string `json:"env_status,omitempty"`

	// 运行服务与应用间的状态。
	EnvAppLinkStatus *string `json:"env_app_link_status,omitempty"`

	// 运行服务与应用间的状态信息。
	EnvAppLinkStatusMsg *string `json:"env_app_link_status_msg,omitempty"`

	// 访问方式。
	Endpoint *string `json:"endpoint,omitempty"`

	// 创建运行服务的jobId。
	JobId *string `json:"job_id,omitempty"`

	// 运行服务的配置信息。
	EnvConfInfo *string `json:"env_conf_info,omitempty"`

	// 部署的应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 部署的应用版本。
	AppVersion *string `json:"app_version,omitempty"`

	// 部署应用的英文名称。
	AppNameEn *string `json:"app_name_en,omitempty"`

	// 部署应用的中文名称。
	AppNameCn *string `json:"app_name_cn,omitempty"`

	// 应用是否可用。 - 0：被认为是false。 - 非0：被认为是true。
	Enabled *bool `json:"enabled,omitempty"`

	// 运行服务的过期时间。
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// 最后部署时间。
	LastDeployTime *int64 `json:"last_deploy_time,omitempty"`

	// 上次部署应用的IAM用户ID。
	DeployUserId *string `json:"deploy_user_id,omitempty"`

	// 计费模式。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 运行服务的创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 绑定主资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 是否支持部署。
	Deployable *bool `json:"deployable,omitempty"`

	// 是否支持卸载。
	Uninstallable *bool `json:"uninstallable,omitempty"`
}

func (o ListEnvsResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvsResponseBodyResult struct{}"
	}

	return strings.Join([]string{"ListEnvsResponseBodyResult", string(data)}, " ")
}
