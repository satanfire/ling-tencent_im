package tencentim

import "bytes"

// API interface
type API interface {
	Name() string
	QueryString() string
	Body() (error, *bytes.Buffer)
}
