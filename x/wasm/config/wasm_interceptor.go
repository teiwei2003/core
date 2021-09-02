package config

type CodeCreateInterceptor func(codeHash []byte) error
type CodeLoadInterceptor func(codeHash []byte) error

// DefaultInterceptorNoop safe default for interceptors.
func DefaultInterceptorNoop([]byte) error {
	return nil
}