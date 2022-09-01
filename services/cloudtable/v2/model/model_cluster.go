package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建集群参数对象实体。
type Cluster struct {

	// 是否开启IAM权限认证。 - false：不开启 - true：开启
	AuthMode *string `json:"auth_mode,omitempty" xml:"auth_mode"`

	// 是否开启Lemon(目前已关闭该参数，填false即可) - false：不开启 - true：开启
	EnableLemon *bool `json:"enable_lemon,omitempty" xml:"enable_lemon"`

	// 是否开启OpenTSDB。 - false：不开启 - true：开启
	EnableOpenTSDB *bool `json:"enable_openTSDB,omitempty" xml:"enable_openTSDB"`

	Instance *Instance `json:"instance" xml:"instance"`

	// CloudTable集群的名称。
	Name string `json:"name" xml:"name"`

	// 存储值的大小。  取值范围: 1-[10240-1024*1024*1024]
	StorageSize *int32 `json:"storage_size,omitempty" xml:"storage_size"`

	// 存储类型： - ULTRAHIGH：超高IO - COMMON：普通IO
	StorageType string `json:"storage_type" xml:"storage_type"`

	// 集群所在的（虚拟网络私有云）VPC。
	VpcId string `json:"vpc_id" xml:"vpc_id"`

	Datastore *Datastore `json:"datastore" xml:"datastore"`
}

func (o Cluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cluster struct{}"
	}

	return strings.Join([]string{"Cluster", string(data)}, " ")
}
