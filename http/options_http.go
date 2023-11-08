package http

import (
	"crypto/tls"
	"time"
)

// BasicHTTPOptions defines only core wpkhttp options
type BasicHTTPOptions struct {
	UserAgent       string
	Proxy           string
	NoTLSValidation bool
	Timeout         time.Duration
	RetryOnTimeout  bool
	RetryAttempts   int
	TLSCertificate  *tls.Certificate
}

// HTTPOptions is the struct to pass in all wpkhttp options to Gobuster
type HTTPOptions struct {
	BasicHTTPOptions
	Password              string
	URL                   string
	Username              string
	Cookies               string
	Headers               []HTTPHeader
	NoCanonicalizeHeaders bool
	FollowRedirect        bool
	Method                string
}
