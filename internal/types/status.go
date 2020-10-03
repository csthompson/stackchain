package types

const (
	StatusOk uint32 = 0
)

const (
	ErrorBadSignature uint32 = iota + 400
	ErrorUnauthorized
	ErrorInternal
)
