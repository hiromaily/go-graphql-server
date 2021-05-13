package httpmethod

// HTTPMethod is HTTP method
type HTTPMethod string

// HTTP method
const (
	GET  HTTPMethod = "GET"
	POST HTTPMethod = "POST"
)

func (h HTTPMethod) String() string {
	return string(h)
}
