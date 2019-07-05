// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/network/redis_proxy/v2/redis_proxy.proto

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on RedisProxy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RedisProxy) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetStatPrefix()) < 1 {
		return RedisProxyValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for Cluster

	if m.GetSettings() == nil {
		return RedisProxyValidationError{
			field:  "Settings",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetSettings()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RedisProxyValidationError{
					field:  "Settings",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for LatencyInMicros

	{
		tmp := m.GetPrefixRoutes()

		if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RedisProxyValidationError{
					field:  "PrefixRoutes",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetDownstreamAuthPassword()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RedisProxyValidationError{
					field:  "DownstreamAuthPassword",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// RedisProxyValidationError is the validation error returned by
// RedisProxy.Validate if the designated constraints aren't met.
type RedisProxyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedisProxyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedisProxyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedisProxyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedisProxyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedisProxyValidationError) ErrorName() string { return "RedisProxyValidationError" }

// Error satisfies the builtin error interface
func (e RedisProxyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedisProxy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedisProxyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedisProxyValidationError{}

// Validate checks the field values on RedisProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RedisProtocolOptions) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetAuthPassword()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RedisProtocolOptionsValidationError{
					field:  "AuthPassword",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// RedisProtocolOptionsValidationError is the validation error returned by
// RedisProtocolOptions.Validate if the designated constraints aren't met.
type RedisProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedisProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedisProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedisProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedisProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedisProtocolOptionsValidationError) ErrorName() string {
	return "RedisProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e RedisProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedisProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedisProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedisProtocolOptionsValidationError{}

// Validate checks the field values on RedisProxy_ConnPoolSettings with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RedisProxy_ConnPoolSettings) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOpTimeout() == nil {
		return RedisProxy_ConnPoolSettingsValidationError{
			field:  "OpTimeout",
			reason: "value is required",
		}
	}

	// no validation rules for EnableHashtagging

	// no validation rules for EnableRedirection

	// no validation rules for MaxBufferSizeBeforeFlush

	{
		tmp := m.GetBufferFlushTimeout()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RedisProxy_ConnPoolSettingsValidationError{
					field:  "BufferFlushTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// RedisProxy_ConnPoolSettingsValidationError is the validation error returned
// by RedisProxy_ConnPoolSettings.Validate if the designated constraints
// aren't met.
type RedisProxy_ConnPoolSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedisProxy_ConnPoolSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedisProxy_ConnPoolSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedisProxy_ConnPoolSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedisProxy_ConnPoolSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedisProxy_ConnPoolSettingsValidationError) ErrorName() string {
	return "RedisProxy_ConnPoolSettingsValidationError"
}

// Error satisfies the builtin error interface
func (e RedisProxy_ConnPoolSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedisProxy_ConnPoolSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedisProxy_ConnPoolSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedisProxy_ConnPoolSettingsValidationError{}

// Validate checks the field values on RedisProxy_PrefixRoutes with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RedisProxy_PrefixRoutes) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRoutes() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return RedisProxy_PrefixRoutesValidationError{
						field:  fmt.Sprintf("Routes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	// no validation rules for CaseInsensitive

	// no validation rules for CatchAllCluster

	return nil
}

// RedisProxy_PrefixRoutesValidationError is the validation error returned by
// RedisProxy_PrefixRoutes.Validate if the designated constraints aren't met.
type RedisProxy_PrefixRoutesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedisProxy_PrefixRoutesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedisProxy_PrefixRoutesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedisProxy_PrefixRoutesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedisProxy_PrefixRoutesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedisProxy_PrefixRoutesValidationError) ErrorName() string {
	return "RedisProxy_PrefixRoutesValidationError"
}

// Error satisfies the builtin error interface
func (e RedisProxy_PrefixRoutesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedisProxy_PrefixRoutes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedisProxy_PrefixRoutesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedisProxy_PrefixRoutesValidationError{}

// Validate checks the field values on RedisProxy_PrefixRoutes_Route with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RedisProxy_PrefixRoutes_Route) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetPrefix()) < 1 {
		return RedisProxy_PrefixRoutes_RouteValidationError{
			field:  "Prefix",
			reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for RemovePrefix

	if len(m.GetCluster()) < 1 {
		return RedisProxy_PrefixRoutes_RouteValidationError{
			field:  "Cluster",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// RedisProxy_PrefixRoutes_RouteValidationError is the validation error
// returned by RedisProxy_PrefixRoutes_Route.Validate if the designated
// constraints aren't met.
type RedisProxy_PrefixRoutes_RouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedisProxy_PrefixRoutes_RouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedisProxy_PrefixRoutes_RouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedisProxy_PrefixRoutes_RouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedisProxy_PrefixRoutes_RouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedisProxy_PrefixRoutes_RouteValidationError) ErrorName() string {
	return "RedisProxy_PrefixRoutes_RouteValidationError"
}

// Error satisfies the builtin error interface
func (e RedisProxy_PrefixRoutes_RouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedisProxy_PrefixRoutes_Route.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedisProxy_PrefixRoutes_RouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedisProxy_PrefixRoutes_RouteValidationError{}
