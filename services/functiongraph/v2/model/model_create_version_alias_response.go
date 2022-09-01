package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateVersionAliasResponse struct {

	// 要获取的别名名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 别名对应的版本名称。
	Version *string `json:"version,omitempty" xml:"version"`

	// 别名描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 别名最后修改时间。
	LastModified *sdktime.SdkTime `json:"last_modified,omitempty" xml:"last_modified"`

	// 版本别名唯一标识。
	AliasUrn *string `json:"alias_urn,omitempty" xml:"alias_urn"`

	// 灰度版本信息
	AdditionalVersionWeights map[string]int64 `json:"additional_version_weights,omitempty" xml:"additional_version_weights"`
	HttpStatusCode           int              `json:"-"`
}

func (o CreateVersionAliasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVersionAliasResponse struct{}"
	}

	return strings.Join([]string{"CreateVersionAliasResponse", string(data)}, " ")
}
