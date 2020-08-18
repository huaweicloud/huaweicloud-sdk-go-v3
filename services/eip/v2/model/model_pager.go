/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// marker分页结构
type Pager struct {
	// 页码url
	Href *string `json:"href,omitempty"`
	// next:下一页  previous:前一页
	Rel PagerRel `json:"rel,omitempty"`
}

func (o Pager) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Pager", string(data)}, " ")
}

type PagerRel struct {
	value string
}

type PagerRelEnum struct {
	NEXT     PagerRel
	PREVIOUS PagerRel
}

func GetPagerRelEnum() PagerRelEnum {
	return PagerRelEnum{
		NEXT: PagerRel{
			value: "next",
		},
		PREVIOUS: PagerRel{
			value: "previous",
		},
	}
}

func (c PagerRel) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *PagerRel) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
