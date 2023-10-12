package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MeshMetadata struct {

	// 网格名称。 命名规则：以小写字母开头，由小写字母、数字、中划线(-)组成，长度范围4-64位，且不能以中划线(-)结尾。
	Name string `json:"name"`

	// 网格ID，资源唯一标识，创建成功后自动生成，填写无效
	Uid *string `json:"uid,omitempty"`

	// 网格注解，由key/value组成： ``` \"annotations\": {    \"key1\" : \"value1\",    \"key2\" : \"value2\" } ```
	Annotations map[string]string `json:"annotations,omitempty"`

	// 网格标签，由key/value组成：   ```  \"annotations\": {    \"key1\" : \"value1\",    \"key2\" : \"value2\" }  ```
	Labels map[string]string `json:"labels,omitempty"`

	// 网格创建时间
	CreateTimestamp *sdktime.SdkTime `json:"createTimestamp,omitempty"`
}

func (o MeshMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MeshMetadata struct{}"
	}

	return strings.Join([]string{"MeshMetadata", string(data)}, " ")
}
