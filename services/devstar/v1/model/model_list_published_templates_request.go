/*
 * Devstar
 *
 * Devstar API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type ListPublishedTemplatesRequest struct {
	XLanguage ListPublishedTemplatesRequestXLanguage `json:"X-Language,omitempty"`
	Keyword   *string                                `json:"keyword,omitempty"`
	Offset    *int32                                 `json:"offset,omitempty"`
	Limit     *int32                                 `json:"limit,omitempty"`
}

func (o ListPublishedTemplatesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPublishedTemplatesRequest", string(data)}, " ")
}

type ListPublishedTemplatesRequestXLanguage struct {
	value string
}

type ListPublishedTemplatesRequestXLanguageEnum struct {
	ZH_CN ListPublishedTemplatesRequestXLanguage
	EN_US ListPublishedTemplatesRequestXLanguage
}

func GetListPublishedTemplatesRequestXLanguageEnum() ListPublishedTemplatesRequestXLanguageEnum {
	return ListPublishedTemplatesRequestXLanguageEnum{
		ZH_CN: ListPublishedTemplatesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListPublishedTemplatesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListPublishedTemplatesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListPublishedTemplatesRequestXLanguage) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
