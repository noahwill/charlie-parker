package types

// BaseOutput contains fields found on every route output object
type BaseOutput struct {
	Error string `json:"error,omitempty"`
	Ok    bool   `json:"ok,omitempty"`
}
