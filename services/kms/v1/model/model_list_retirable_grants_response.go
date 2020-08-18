/*
 * kms
 *
 * KMS v1.0 API, open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type ListRetirableGrantsResponse struct {
	// grant列表，详情请参见grants字段数据结构说明。
	Grants []Grants `json:"grants,omitempty"`
	// 获取下一页所需要传递的marker值。 当“truncated”为“false”时，“next_marker”为空。
	NextMarker *string `json:"next_marker,omitempty"`
	// 是否还有下一页：  - “true”表示还有数据。  - “false”表示已经是最后一页。
	Truncated ListRetirableGrantsResponseTruncated `json:"truncated,omitempty"`
}

func (o ListRetirableGrantsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListRetirableGrantsResponse", string(data)}, " ")
}

type ListRetirableGrantsResponseTruncated struct {
	value string
}

type ListRetirableGrantsResponseTruncatedEnum struct {
	TRUE  ListRetirableGrantsResponseTruncated
	FALSE ListRetirableGrantsResponseTruncated
}

func GetListRetirableGrantsResponseTruncatedEnum() ListRetirableGrantsResponseTruncatedEnum {
	return ListRetirableGrantsResponseTruncatedEnum{
		TRUE: ListRetirableGrantsResponseTruncated{
			value: "true",
		},
		FALSE: ListRetirableGrantsResponseTruncated{
			value: "false",
		},
	}
}

func (c ListRetirableGrantsResponseTruncated) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListRetirableGrantsResponseTruncated) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
