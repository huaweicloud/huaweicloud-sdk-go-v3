/*
 * ecs
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

//
type NovaLink struct {
	// 相应资源的链接。
	Href string `json:"href"`
	// 有三种取值。self：自助链接包含版本链接的资源。立即链接后使用这些链接。bookmark：书签链接提供了一个永久资源的永久链接，该链接适合于长期存储。alternate：备用链接可以包含资源的替换表示形式。例如，OpenStack计算映像可能在OpenStack映像服务中有一个替代表示。
	Rel NovaLinkRel `json:"rel"`
}

func (o NovaLink) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NovaLink", string(data)}, " ")
}

type NovaLinkRel struct {
	value string
}

type NovaLinkRelEnum struct {
	SELF        NovaLinkRel
	BOOKMARK    NovaLinkRel
	ALTERNATE   NovaLinkRel
	DESCRIBEDBY NovaLinkRel
}

func GetNovaLinkRelEnum() NovaLinkRelEnum {
	return NovaLinkRelEnum{
		SELF: NovaLinkRel{
			value: "self",
		},
		BOOKMARK: NovaLinkRel{
			value: "bookmark",
		},
		ALTERNATE: NovaLinkRel{
			value: "alternate",
		},
		DESCRIBEDBY: NovaLinkRel{
			value: "describedby",
		},
	}
}

func (c NovaLinkRel) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *NovaLinkRel) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
