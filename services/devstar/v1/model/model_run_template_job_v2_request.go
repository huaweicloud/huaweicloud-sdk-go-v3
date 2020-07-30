/*
    * Devstar
    *
    * Devstar API
    *
*/

package model

// Request Object
type RunTemplateJobV2Request struct {
	XLanguage string `json:"X-Language,omitempty"`
	Body *TemplateJobInfo `json:"body,omitempty"`
}
