package tencentim

import "bytes"

// API interface
type API interface {
	Name() string
	URI() string
	QueryString() string
	Body() (*bytes.Buffer, error)
}
