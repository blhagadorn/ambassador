// Code generated by protoc-gen-validate
// source: envoy/config/filter/network/tcp_proxy/v2/tcp_proxy.proto
// DO NOT EDIT!!!

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

// Validate checks the field values on TcpProxy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TcpProxy) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetStatPrefix()) < 1 {
		return TcpProxyValidationError{
			Field:  "StatPrefix",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetMetadataMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpProxyValidationError{
				Field:  "MetadataMatch",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if d := m.GetIdleTimeout(); d != nil {
		dur := *d

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return TcpProxyValidationError{
				Field:  "IdleTimeout",
				Reason: "value must be greater than 0s",
			}
		}

	}

	if v, ok := interface{}(m.GetDownstreamIdleTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpProxyValidationError{
				Field:  "DownstreamIdleTimeout",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpstreamIdleTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpProxyValidationError{
				Field:  "UpstreamIdleTimeout",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetAccessLog() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TcpProxyValidationError{
					Field:  fmt.Sprintf("AccessLog[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetDeprecatedV1()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpProxyValidationError{
				Field:  "DeprecatedV1",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if wrapper := m.GetMaxConnectAttempts(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			return TcpProxyValidationError{
				Field:  "MaxConnectAttempts",
				Reason: "value must be greater than or equal to 1",
			}
		}

	}

	switch m.ClusterSpecifier.(type) {

	case *TcpProxy_Cluster:
		// no validation rules for Cluster

	case *TcpProxy_WeightedClusters:

		if v, ok := interface{}(m.GetWeightedClusters()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TcpProxyValidationError{
					Field:  "WeightedClusters",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return TcpProxyValidationError{
			Field:  "ClusterSpecifier",
			Reason: "value is required",
		}

	}

	return nil
}

// TcpProxyValidationError is the validation error returned by
// TcpProxy.Validate if the designated constraints aren't met.
type TcpProxyValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TcpProxyValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpProxy.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TcpProxyValidationError{}

// Validate checks the field values on TcpProxy_DeprecatedV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TcpProxy_DeprecatedV1) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRoutes()) < 1 {
		return TcpProxy_DeprecatedV1ValidationError{
			Field:  "Routes",
			Reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetRoutes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TcpProxy_DeprecatedV1ValidationError{
					Field:  fmt.Sprintf("Routes[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// TcpProxy_DeprecatedV1ValidationError is the validation error returned by
// TcpProxy_DeprecatedV1.Validate if the designated constraints aren't met.
type TcpProxy_DeprecatedV1ValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TcpProxy_DeprecatedV1ValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpProxy_DeprecatedV1.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TcpProxy_DeprecatedV1ValidationError{}

// Validate checks the field values on TcpProxy_WeightedCluster with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TcpProxy_WeightedCluster) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetClusters()) < 1 {
		return TcpProxy_WeightedClusterValidationError{
			Field:  "Clusters",
			Reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetClusters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TcpProxy_WeightedClusterValidationError{
					Field:  fmt.Sprintf("Clusters[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// TcpProxy_WeightedClusterValidationError is the validation error returned by
// TcpProxy_WeightedCluster.Validate if the designated constraints aren't met.
type TcpProxy_WeightedClusterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TcpProxy_WeightedClusterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpProxy_WeightedCluster.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TcpProxy_WeightedClusterValidationError{}

// Validate checks the field values on TcpProxy_DeprecatedV1_TCPRoute with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TcpProxy_DeprecatedV1_TCPRoute) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetCluster()) < 1 {
		return TcpProxy_DeprecatedV1_TCPRouteValidationError{
			Field:  "Cluster",
			Reason: "value length must be at least 1 bytes",
		}
	}

	for idx, item := range m.GetDestinationIpList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TcpProxy_DeprecatedV1_TCPRouteValidationError{
					Field:  fmt.Sprintf("DestinationIpList[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	// no validation rules for DestinationPorts

	for idx, item := range m.GetSourceIpList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TcpProxy_DeprecatedV1_TCPRouteValidationError{
					Field:  fmt.Sprintf("SourceIpList[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	// no validation rules for SourcePorts

	return nil
}

// TcpProxy_DeprecatedV1_TCPRouteValidationError is the validation error
// returned by TcpProxy_DeprecatedV1_TCPRoute.Validate if the designated
// constraints aren't met.
type TcpProxy_DeprecatedV1_TCPRouteValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TcpProxy_DeprecatedV1_TCPRouteValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpProxy_DeprecatedV1_TCPRoute.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TcpProxy_DeprecatedV1_TCPRouteValidationError{}

// Validate checks the field values on TcpProxy_WeightedCluster_ClusterWeight
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *TcpProxy_WeightedCluster_ClusterWeight) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return TcpProxy_WeightedCluster_ClusterWeightValidationError{
			Field:  "Name",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if m.GetWeight() < 1 {
		return TcpProxy_WeightedCluster_ClusterWeightValidationError{
			Field:  "Weight",
			Reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// TcpProxy_WeightedCluster_ClusterWeightValidationError is the validation
// error returned by TcpProxy_WeightedCluster_ClusterWeight.Validate if the
// designated constraints aren't met.
type TcpProxy_WeightedCluster_ClusterWeightValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TcpProxy_WeightedCluster_ClusterWeightValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpProxy_WeightedCluster_ClusterWeight.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TcpProxy_WeightedCluster_ClusterWeightValidationError{}
