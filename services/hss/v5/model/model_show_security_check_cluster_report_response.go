package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSecurityCheckClusterReportResponse Response Object
type ShowSecurityCheckClusterReportResponse struct {

	// **参数解释**： 体检时间 **取值范围**： 不涉及
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**： 集群ID **取值范围**： 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 集群名称 **取值范围**： 不涉及
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 集群类型 **取值范围**： - CCE：CCE Standard集群 - Turbo：CCE Turbo集群
	ClusterCategory *ShowSecurityCheckClusterReportResponseClusterCategory `json:"cluster_category,omitempty"`

	// **参数解释**： 本地镜像漏洞风险数量 **取值范围**： 不涉及
	LocalImageVulNum *int32 `json:"local_image_vul_num,omitempty"`

	// **参数解释**： 基线风险数量 **取值范围**： 不涉及
	RiskConfigNum *int32 `json:"risk_config_num,omitempty"`

	// **参数解释**： 特权容器数量 **取值范围**： 不涉及
	PrivilegedContainerNum *int32 `json:"privileged_container_num,omitempty"`

	// **参数解释**： pod数量 **取值范围**： 不涉及
	PodNum *int32 `json:"pod_num,omitempty"`

	// **参数解释**： 节点数量 **取值范围**： 不涉及
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**： 容器数量 **取值范围**： 不涉及
	ContainerNum *int32 `json:"container_num,omitempty"`

	// **参数解释**： 端口数量 **取值范围**： 不涉及
	PortNum *int32 `json:"port_num,omitempty"`

	// **参数解释**： 进程数量 **取值范围**： 不涉及
	ProcessNum *int32 `json:"process_num,omitempty"`

	// **参数解释**： 软件数量 **取值范围**： 不涉及
	AppNum *int32 `json:"app_num,omitempty"`

	// **参数解释**： 本地镜像漏洞列表 **取值范围**： 不涉及
	LocalImageVulList *[]ClusterSecurityCheckLocalImageVulInfo `json:"local_image_vul_list,omitempty"`

	// **参数解释**： 基线检测列表 **取值范围**： 不涉及
	BaselineDetectionList *[]ClusterSecurityCheckBaseLineDetectionInfo `json:"baseline_detection_list,omitempty"`

	// **参数解释**： 特权容器列表 **取值范围**： 不涉及
	PrivilegedContainerList *[]ClusterSecurityCheckPrivilegedContainerInfo `json:"privileged_container_list,omitempty"`

	// **参数解释**： 端口列表 **取值范围**： 不涉及
	PortList *[]ClusterSecurityCheckPortInfo `json:"port_list,omitempty"`

	// **参数解释**： 软件列表 **取值范围**： 不涉及
	AppList *[]ClusterSecurityCheckAppInfo `json:"app_list,omitempty"`

	// **参数解释**： 进程列表 **取值范围**： 不涉及
	ProcessList    *[]ClusterSecurityCheckProcessInfo `json:"process_list,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ShowSecurityCheckClusterReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityCheckClusterReportResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityCheckClusterReportResponse", string(data)}, " ")
}

type ShowSecurityCheckClusterReportResponseClusterCategory struct {
	value string
}

type ShowSecurityCheckClusterReportResponseClusterCategoryEnum struct {
	CCE   ShowSecurityCheckClusterReportResponseClusterCategory
	TURBO ShowSecurityCheckClusterReportResponseClusterCategory
}

func GetShowSecurityCheckClusterReportResponseClusterCategoryEnum() ShowSecurityCheckClusterReportResponseClusterCategoryEnum {
	return ShowSecurityCheckClusterReportResponseClusterCategoryEnum{
		CCE: ShowSecurityCheckClusterReportResponseClusterCategory{
			value: "CCE",
		},
		TURBO: ShowSecurityCheckClusterReportResponseClusterCategory{
			value: "Turbo",
		},
	}
}

func (c ShowSecurityCheckClusterReportResponseClusterCategory) Value() string {
	return c.value
}

func (c ShowSecurityCheckClusterReportResponseClusterCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityCheckClusterReportResponseClusterCategory) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
