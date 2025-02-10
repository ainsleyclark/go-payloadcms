package payloadcms

import (
	"fmt"
	"net/url"
	"strings"
)

// ListParams represents additional query parameters for the find endpoint.
type ListParams struct {
	Sort  string        `json:"sort" url:"sort"`   // Sort the returned documents by a specific field.
	Where *QueryBuilder `json:"where" url:"where"` // Constrain returned documents with a where query.
	Limit int           `json:"limit" url:"limit"` // Limit the returned documents to a certain number.
	Page  int           `json:"page" url:"page"`   // Get a specific page of documents.
}

// Encode encodes ListParams into a URL query string.
func (p ListParams) Encode() string {
	str := ""
	if p.Sort != "" {
		str += fmt.Sprintf("&sort=%s", url.QueryEscape(p.Sort))
	}
	if p.Where != nil {
		str += fmt.Sprintf("&%s", p.Where.Build())
	}
	if p.Limit > 0 {
		str += fmt.Sprintf("&limit=%d", p.Limit)
	}
	if p.Page > 0 {
		str += fmt.Sprintf("&page=%d", p.Page)
	}
	if str == "" {
		return ""
	}
	return "?" + strings.TrimPrefix(str, "&")
}
