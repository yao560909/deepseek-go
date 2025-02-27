package apiquery

import "net/url"

type Queryer interface {
	URLQuery() url.Values
}
