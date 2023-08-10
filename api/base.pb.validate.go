// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: base.proto

package api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Response with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Response with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResponseMultiError, or nil
// if none found.
func (m *Response) ValidateAll() error {
	return m.validate(true)
}

func (m *Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	// no validation rules for Metadata

	if len(errors) > 0 {
		return ResponseMultiError(errors)
	}

	return nil
}

// ResponseMultiError is an error wrapping multiple validation errors returned
// by Response.ValidateAll() if the designated constraints aren't met.
type ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResponseMultiError) AllErrors() []error { return m }

// ResponseValidationError is the validation error returned by
// Response.Validate if the designated constraints aren't met.
type ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponseValidationError) ErrorName() string { return "ResponseValidationError" }

// Error satisfies the builtin error interface
func (e ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponseValidationError{}

// Validate checks the field values on PageRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageRequestMultiError, or
// nil if none found.
func (m *PageRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PageRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetCurrent() <= 0 {
		err := PageRequestValidationError{
			field:  "Current",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSize() <= 0 {
		err := PageRequestValidationError{
			field:  "Size",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PageRequestMultiError(errors)
	}

	return nil
}

// PageRequestMultiError is an error wrapping multiple validation errors
// returned by PageRequest.ValidateAll() if the designated constraints aren't met.
type PageRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageRequestMultiError) AllErrors() []error { return m }

// PageRequestValidationError is the validation error returned by
// PageRequest.Validate if the designated constraints aren't met.
type PageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageRequestValidationError) ErrorName() string { return "PageRequestValidationError" }

// Error satisfies the builtin error interface
func (e PageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageRequestValidationError{}

// Validate checks the field values on PageReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageReplyMultiError, or nil
// if none found.
func (m *PageReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PageReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Current

	// no validation rules for Size

	// no validation rules for Total

	if len(errors) > 0 {
		return PageReplyMultiError(errors)
	}

	return nil
}

// PageReplyMultiError is an error wrapping multiple validation errors returned
// by PageReply.ValidateAll() if the designated constraints aren't met.
type PageReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageReplyMultiError) AllErrors() []error { return m }

// PageReplyValidationError is the validation error returned by
// PageReply.Validate if the designated constraints aren't met.
type PageReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageReplyValidationError) ErrorName() string { return "PageReplyValidationError" }

// Error satisfies the builtin error interface
func (e PageReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageReplyValidationError{}

// Validate checks the field values on Sort with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Sort) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Sort with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SortMultiError, or nil if none found.
func (m *Sort) ValidateAll() error {
	return m.validate(true)
}

func (m *Sort) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Field

	// no validation rules for Asc

	if len(errors) > 0 {
		return SortMultiError(errors)
	}

	return nil
}

// SortMultiError is an error wrapping multiple validation errors returned by
// Sort.ValidateAll() if the designated constraints aren't met.
type SortMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SortMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SortMultiError) AllErrors() []error { return m }

// SortValidationError is the validation error returned by Sort.Validate if the
// designated constraints aren't met.
type SortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SortValidationError) ErrorName() string { return "SortValidationError" }

// Error satisfies the builtin error interface
func (e SortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SortValidationError{}

// Validate checks the field values on Field with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Field) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Field with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FieldMultiError, or nil if none found.
func (m *Field) ValidateAll() error {
	return m.validate(true)
}

func (m *Field) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Field

	// no validation rules for Label

	if len(errors) > 0 {
		return FieldMultiError(errors)
	}

	return nil
}

// FieldMultiError is an error wrapping multiple validation errors returned by
// Field.ValidateAll() if the designated constraints aren't met.
type FieldMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FieldMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FieldMultiError) AllErrors() []error { return m }

// FieldValidationError is the validation error returned by Field.Validate if
// the designated constraints aren't met.
type FieldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FieldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FieldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FieldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FieldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FieldValidationError) ErrorName() string { return "FieldValidationError" }

// Error satisfies the builtin error interface
func (e FieldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sField.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FieldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FieldValidationError{}

// Validate checks the field values on ListQuery with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListQuery) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListQuery with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListQueryMultiError, or nil
// if none found.
func (m *ListQuery) ValidateAll() error {
	return m.validate(true)
}

func (m *ListQuery) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() == nil {
		err := ListQueryValidationError{
			field:  "Page",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetPage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListQueryValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListQueryValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListQueryValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSort() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListQueryValidationError{
						field:  fmt.Sprintf("Sort[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListQueryValidationError{
						field:  fmt.Sprintf("Sort[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListQueryValidationError{
					field:  fmt.Sprintf("Sort[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Keyword

	// no validation rules for StartAt

	// no validation rules for EndAt

	// no validation rules for TimeField

	if len(errors) > 0 {
		return ListQueryMultiError(errors)
	}

	return nil
}

// ListQueryMultiError is an error wrapping multiple validation errors returned
// by ListQuery.ValidateAll() if the designated constraints aren't met.
type ListQueryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListQueryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListQueryMultiError) AllErrors() []error { return m }

// ListQueryValidationError is the validation error returned by
// ListQuery.Validate if the designated constraints aren't met.
type ListQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListQueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListQueryValidationError) ErrorName() string { return "ListQueryValidationError" }

// Error satisfies the builtin error interface
func (e ListQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListQueryValidationError{}

// Validate checks the field values on ListQueryResult with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListQueryResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListQueryResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListQueryResultMultiError, or nil if none found.
func (m *ListQueryResult) ValidateAll() error {
	return m.validate(true)
}

func (m *ListQueryResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListQueryResultValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListQueryResultValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListQueryResultValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetFields() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListQueryResultValidationError{
						field:  fmt.Sprintf("Fields[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListQueryResultValidationError{
						field:  fmt.Sprintf("Fields[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListQueryResultValidationError{
					field:  fmt.Sprintf("Fields[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListQueryResultMultiError(errors)
	}

	return nil
}

// ListQueryResultMultiError is an error wrapping multiple validation errors
// returned by ListQueryResult.ValidateAll() if the designated constraints
// aren't met.
type ListQueryResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListQueryResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListQueryResultMultiError) AllErrors() []error { return m }

// ListQueryResultValidationError is the validation error returned by
// ListQueryResult.Validate if the designated constraints aren't met.
type ListQueryResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListQueryResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListQueryResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListQueryResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListQueryResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListQueryResultValidationError) ErrorName() string { return "ListQueryResultValidationError" }

// Error satisfies the builtin error interface
func (e ListQueryResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListQueryResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListQueryResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListQueryResultValidationError{}
