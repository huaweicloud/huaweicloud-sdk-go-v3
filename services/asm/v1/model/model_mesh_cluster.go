package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MeshCluster struct {

	// 集群ID，资源唯一标识，通过该ID查询需要添加的集群。
	ClusterID string `json:"clusterID"`

	// 集群所在的Region
	Region string `json:"region"`

	// 集群所属的projectID
	ProjectID string `json:"projectID"`

	// 集群提供方
	Provider MeshClusterProvider `json:"provider"`

	// 集群代理模式
	ProxyMode *MeshClusterProxyMode `json:"proxyMode,omitempty"`

	Injection *InjectionConfig `json:"injection,omitempty"`

	Installation *InstallationConfig `json:"installation"`
}

func (o MeshCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MeshCluster struct{}"
	}

	return strings.Join([]string{"MeshCluster", string(data)}, " ")
}

type MeshClusterProvider struct {
	value string
}

type MeshClusterProviderEnum struct {
	CCE MeshClusterProvider
}

func GetMeshClusterProviderEnum() MeshClusterProviderEnum {
	return MeshClusterProviderEnum{
		CCE: MeshClusterProvider{
			value: "CCE",
		},
	}
}

func (c MeshClusterProvider) Value() string {
	return c.value
}

func (c MeshClusterProvider) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MeshClusterProvider) UnmarshalJSON(b []byte) error {
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

type MeshClusterProxyMode struct {
	value string
}

type MeshClusterProxyModeEnum struct {
	SIDECAR MeshClusterProxyMode
	NODE    MeshClusterProxyMode
}

func GetMeshClusterProxyModeEnum() MeshClusterProxyModeEnum {
	return MeshClusterProxyModeEnum{
		SIDECAR: MeshClusterProxyMode{
			value: "sidecar",
		},
		NODE: MeshClusterProxyMode{
			value: "node",
		},
	}
}

func (c MeshClusterProxyMode) Value() string {
	return c.value
}

func (c MeshClusterProxyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MeshClusterProxyMode) UnmarshalJSON(b []byte) error {
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
