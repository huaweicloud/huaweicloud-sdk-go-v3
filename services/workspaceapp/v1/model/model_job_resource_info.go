package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobResourceInfo 具体的资源信息
type JobResourceInfo struct {

	// 客户在云服务Console上可见的资源实例Id，全局唯一，且不可更改，最大64个字符。  注：“规格变更”场景下（包括升降配），是对某个资源实例的规格进行调整，  该资源实例其他信息（比如资源Id、资源名称）和运行的业务数据不变化。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称；创建、有最新资源名称场景，必填。
	ResourceName *string `json:"resource_name,omitempty"`

	// 云服务类型编码；新购、规格变更场景，必填。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 资源类型编码；新购、规格变更场景，必填。
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源规格编码；新购、规格变更场景，必填。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 规格类型，运营上需要呈现和使用的一些规格属性，多个使用K:V格式； 比如带宽的共享/独享(shareable:true/false)，数据盘的系统盘/数据盘类型(root:true/false) 当前针对共享带宽、共享盘使用，必填。
	SpecType map[string]interface{} `json:"spec_type,omitempty"`

	// 某些规格属性大小：比如带宽大小、数据盘大小
	SpecSize *float64 `json:"spec_size,omitempty"`

	// specSize的单位编码，比如GB、M，有specSize时，此字段必填。
	Measure *int32 `json:"measure,omitempty"`

	// 处理时间
	ProcessedTime *sdktime.SdkTime `json:"processed_time,omitempty"`

	// 该resourceId是否是主资源（仅开通场景使用，其他场景为空） * `1` - 是 * `0` - 否
	IsMainResource *int32 `json:"is_main_resource,omitempty"`

	// resourceId的主资源。  是挂载到/绑定到/依附到/包含于/关联到资源，比如IP的主资源‘云主机’、数据盘的主资源‘云主机’。  如果resourceId是依附在多个资源上，则有多个主资源，比如共享盘挂载到多个云主机上。  无关联主资源，则空，比如独立未挂载的数据盘。
	MainResources *[]RelativeResource `json:"main_resources,omitempty"`

	// expireTime：到期时间，域名注册服务使用。  UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ（2016-06-28T00:00:00Z）
	ExtendParams *string `json:"extend_params,omitempty"`

	// 仅针对ECS/BMS云服务的“切换操作系统”场景使用： 云主机切换操作系统的资源id会变化场景 填写变更前老的资源Id。资源Id未变化，无此字段
	OldResourceId *string `json:"old_resource_id,omitempty"`

	// 仅针对ECS/BMS云服务的“切换操作系统”场景使用：云主机切换操作系统的云服务类型编码会变化场景， 填写变更前老的云服务类型编码。云服务类型未变化，无此字段。
	OldCloudServiceType *string `json:"old_cloud_service_type,omitempty"`

	// 仅针对ECS/BMS云服务“切换操作系统”场景使用： 云主机切换操作系统的资源类型编码会变化场景， 填写变更前老的资源类型编码。资源类型未变化，无此字段
	OldResourceType *string `json:"old_resource_type,omitempty"`
}

func (o JobResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobResourceInfo struct{}"
	}

	return strings.Join([]string{"JobResourceInfo", string(data)}, " ")
}
