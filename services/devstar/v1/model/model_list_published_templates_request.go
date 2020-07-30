/*
    * Devstar
    *
    * Devstar API
    *
*/

package model

// Request Object
type ListPublishedTemplatesRequest struct {
	XLanguage string `json:"X-Language,omitempty"`
	Keyword string `json:"keyword,omitempty"`
	Offset int32 `json:"offset,omitempty"`
	Limit int32 `json:"limit,omitempty"`
}
