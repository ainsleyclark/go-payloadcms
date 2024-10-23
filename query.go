package payloadcms

import (
	"fmt"
	"net/url"
	"strings"
)

type QueryBuilder struct {
	params url.Values
}

// Query creates a new instance of QueryBuilder
func Query() *QueryBuilder {
	return &QueryBuilder{
		params: url.Values{},
	}
}

// Equals adds an equals filter to the query
func (qb *QueryBuilder) Equals(field, value string) *QueryBuilder {
	qb.params.Add(fmt.Sprintf("where[%s][equals]", field), value)
	return qb
}

// NotEquals adds a not_equals filter to the query
func (qb *QueryBuilder) NotEquals(field, value string) *QueryBuilder {
	qb.params.Add(fmt.Sprintf("where[%s][not_equals]", field), value)
	return qb
}

// GreaterThan adds a greater_than filter to the query
func (qb *QueryBuilder) GreaterThan(field, value string) *QueryBuilder {
	qb.params.Add(fmt.Sprintf("where[%s][greater_than]", field), value)
	return qb
}

// LessThan adds a less_than filter to the query
func (qb *QueryBuilder) LessThan(field, value string) *QueryBuilder {
	qb.params.Add(fmt.Sprintf("where[%s][less_than]", field), value)
	return qb
}

// In adds an in filter to the query
func (qb *QueryBuilder) In(field string, values []string) *QueryBuilder {
	qb.params.Add(fmt.Sprintf("where[%s][in]", field), strings.Join(values, ","))
	return qb
}

// And adds an AND condition to the query
func (qb *QueryBuilder) And(subQuery *QueryBuilder) *QueryBuilder {
	for key, values := range subQuery.params {
		for _, value := range values {
			qb.params.Add("where[and][]"+fmt.Sprintf("[%s]", key), value)
		}
	}
	return qb
}

// Or adds an OR condition to the query
func (qb *QueryBuilder) Or(subQuery *QueryBuilder) *QueryBuilder {
	for key, values := range subQuery.params {
		for _, value := range values {
			qb.params.Add("where[or][]"+fmt.Sprintf("[%s]", key), value)
		}
	}
	return qb
}

// Exists adds an exists filter to the query
func (qb *QueryBuilder) Exists(field string, exists bool) *QueryBuilder {
	existsValue := "false"
	if exists {
		existsValue = "true"
	}
	qb.params.Add(fmt.Sprintf("where[%s][exists]", field), existsValue)
	return qb
}

// Build constructs the final query string
func (qb *QueryBuilder) Build() string {
	if len(qb.params) == 0 {
		return ""
	}
	return qb.params.Encode()
}
