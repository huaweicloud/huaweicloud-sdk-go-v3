package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type VersionLinks struct {
	// 链接类型。self：自助链接包含了版本链接的资源。bookmark：书签链接提供了一个永久资源的永久链接。alternate：备用链接包含了资源的替换表示形式。

	Rel string `json:"rel"`
	// 资源链接地址。

	Href string `json:"href"`
}

func (o VersionLinks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionLinks struct{}"
	}

	return strings.Join([]string{"VersionLinks", string(data)}, " ")
}
