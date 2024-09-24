// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/abc/v1/word.proto

package v1

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

// Validate checks the field values on Word with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Word) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Word with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in WordMultiError, or nil if none found.
func (m *Word) ValidateAll() error {
	return m.validate(true)
}

func (m *Word) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Word

	if len(errors) > 0 {
		return WordMultiError(errors)
	}

	return nil
}

// WordMultiError is an error wrapping multiple validation errors returned by
// Word.ValidateAll() if the designated constraints aren't met.
type WordMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WordMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WordMultiError) AllErrors() []error { return m }

// WordValidationError is the validation error returned by Word.Validate if the
// designated constraints aren't met.
type WordValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WordValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WordValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WordValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WordValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WordValidationError) ErrorName() string { return "WordValidationError" }

// Error satisfies the builtin error interface
func (e WordValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWord.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WordValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WordValidationError{}

// Validate checks the field values on CreateWordRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateWordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateWordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateWordRequestMultiError, or nil if none found.
func (m *CreateWordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateWordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetWord()) < 10 {
		err := CreateWordRequestValidationError{
			field:  "Word",
			reason: "value length must be at least 10 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateWordRequestMultiError(errors)
	}

	return nil
}

// CreateWordRequestMultiError is an error wrapping multiple validation errors
// returned by CreateWordRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateWordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateWordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateWordRequestMultiError) AllErrors() []error { return m }

// CreateWordRequestValidationError is the validation error returned by
// CreateWordRequest.Validate if the designated constraints aren't met.
type CreateWordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateWordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateWordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateWordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateWordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateWordRequestValidationError) ErrorName() string {
	return "CreateWordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateWordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateWordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateWordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateWordRequestValidationError{}

// Validate checks the field values on CreateWordReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateWordReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateWordReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateWordReplyMultiError, or nil if none found.
func (m *CreateWordReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateWordReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetWord()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateWordReplyValidationError{
					field:  "Word",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateWordReplyValidationError{
					field:  "Word",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWord()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateWordReplyValidationError{
				field:  "Word",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateWordReplyMultiError(errors)
	}

	return nil
}

// CreateWordReplyMultiError is an error wrapping multiple validation errors
// returned by CreateWordReply.ValidateAll() if the designated constraints
// aren't met.
type CreateWordReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateWordReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateWordReplyMultiError) AllErrors() []error { return m }

// CreateWordReplyValidationError is the validation error returned by
// CreateWordReply.Validate if the designated constraints aren't met.
type CreateWordReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateWordReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateWordReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateWordReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateWordReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateWordReplyValidationError) ErrorName() string { return "CreateWordReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateWordReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateWordReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateWordReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateWordReplyValidationError{}

// Validate checks the field values on UpdateWordRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateWordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateWordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateWordRequestMultiError, or nil if none found.
func (m *UpdateWordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateWordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateWordRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Word

	if len(errors) > 0 {
		return UpdateWordRequestMultiError(errors)
	}

	return nil
}

// UpdateWordRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateWordRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateWordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateWordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateWordRequestMultiError) AllErrors() []error { return m }

// UpdateWordRequestValidationError is the validation error returned by
// UpdateWordRequest.Validate if the designated constraints aren't met.
type UpdateWordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateWordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateWordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateWordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateWordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateWordRequestValidationError) ErrorName() string {
	return "UpdateWordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateWordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateWordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateWordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateWordRequestValidationError{}

// Validate checks the field values on UpdateWordReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateWordReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateWordReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateWordReplyMultiError, or nil if none found.
func (m *UpdateWordReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateWordReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetWord()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateWordReplyValidationError{
					field:  "Word",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateWordReplyValidationError{
					field:  "Word",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWord()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateWordReplyValidationError{
				field:  "Word",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateWordReplyMultiError(errors)
	}

	return nil
}

// UpdateWordReplyMultiError is an error wrapping multiple validation errors
// returned by UpdateWordReply.ValidateAll() if the designated constraints
// aren't met.
type UpdateWordReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateWordReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateWordReplyMultiError) AllErrors() []error { return m }

// UpdateWordReplyValidationError is the validation error returned by
// UpdateWordReply.Validate if the designated constraints aren't met.
type UpdateWordReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateWordReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateWordReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateWordReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateWordReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateWordReplyValidationError) ErrorName() string { return "UpdateWordReplyValidationError" }

// Error satisfies the builtin error interface
func (e UpdateWordReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateWordReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateWordReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateWordReplyValidationError{}

// Validate checks the field values on DeleteWordRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteWordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteWordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteWordRequestMultiError, or nil if none found.
func (m *DeleteWordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteWordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteWordRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteWordRequestMultiError(errors)
	}

	return nil
}

// DeleteWordRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteWordRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteWordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteWordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteWordRequestMultiError) AllErrors() []error { return m }

// DeleteWordRequestValidationError is the validation error returned by
// DeleteWordRequest.Validate if the designated constraints aren't met.
type DeleteWordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteWordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteWordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteWordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteWordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteWordRequestValidationError) ErrorName() string {
	return "DeleteWordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteWordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteWordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteWordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteWordRequestValidationError{}

// Validate checks the field values on DeleteWordReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteWordReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteWordReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteWordReplyMultiError, or nil if none found.
func (m *DeleteWordReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteWordReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteWordReplyMultiError(errors)
	}

	return nil
}

// DeleteWordReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteWordReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteWordReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteWordReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteWordReplyMultiError) AllErrors() []error { return m }

// DeleteWordReplyValidationError is the validation error returned by
// DeleteWordReply.Validate if the designated constraints aren't met.
type DeleteWordReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteWordReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteWordReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteWordReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteWordReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteWordReplyValidationError) ErrorName() string { return "DeleteWordReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteWordReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteWordReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteWordReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteWordReplyValidationError{}

// Validate checks the field values on GetWordRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetWordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetWordRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetWordRequestMultiError,
// or nil if none found.
func (m *GetWordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetWordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetWordRequestMultiError(errors)
	}

	return nil
}

// GetWordRequestMultiError is an error wrapping multiple validation errors
// returned by GetWordRequest.ValidateAll() if the designated constraints
// aren't met.
type GetWordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetWordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetWordRequestMultiError) AllErrors() []error { return m }

// GetWordRequestValidationError is the validation error returned by
// GetWordRequest.Validate if the designated constraints aren't met.
type GetWordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetWordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetWordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetWordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetWordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetWordRequestValidationError) ErrorName() string { return "GetWordRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetWordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetWordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetWordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetWordRequestValidationError{}

// Validate checks the field values on GetWordReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetWordReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetWordReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetWordReplyMultiError, or
// nil if none found.
func (m *GetWordReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetWordReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetWord()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetWordReplyValidationError{
					field:  "Word",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetWordReplyValidationError{
					field:  "Word",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWord()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetWordReplyValidationError{
				field:  "Word",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetWordReplyMultiError(errors)
	}

	return nil
}

// GetWordReplyMultiError is an error wrapping multiple validation errors
// returned by GetWordReply.ValidateAll() if the designated constraints aren't met.
type GetWordReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetWordReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetWordReplyMultiError) AllErrors() []error { return m }

// GetWordReplyValidationError is the validation error returned by
// GetWordReply.Validate if the designated constraints aren't met.
type GetWordReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetWordReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetWordReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetWordReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetWordReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetWordReplyValidationError) ErrorName() string { return "GetWordReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetWordReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetWordReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetWordReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetWordReplyValidationError{}

// Validate checks the field values on ListWordRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListWordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListWordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListWordRequestMultiError, or nil if none found.
func (m *ListWordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListWordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListWordRequestMultiError(errors)
	}

	return nil
}

// ListWordRequestMultiError is an error wrapping multiple validation errors
// returned by ListWordRequest.ValidateAll() if the designated constraints
// aren't met.
type ListWordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListWordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListWordRequestMultiError) AllErrors() []error { return m }

// ListWordRequestValidationError is the validation error returned by
// ListWordRequest.Validate if the designated constraints aren't met.
type ListWordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWordRequestValidationError) ErrorName() string { return "ListWordRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListWordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWordRequestValidationError{}

// Validate checks the field values on ListWordReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListWordReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListWordReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListWordReplyMultiError, or
// nil if none found.
func (m *ListWordReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListWordReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetWords() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListWordReplyValidationError{
						field:  fmt.Sprintf("Words[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListWordReplyValidationError{
						field:  fmt.Sprintf("Words[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListWordReplyValidationError{
					field:  fmt.Sprintf("Words[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListWordReplyMultiError(errors)
	}

	return nil
}

// ListWordReplyMultiError is an error wrapping multiple validation errors
// returned by ListWordReply.ValidateAll() if the designated constraints
// aren't met.
type ListWordReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListWordReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListWordReplyMultiError) AllErrors() []error { return m }

// ListWordReplyValidationError is the validation error returned by
// ListWordReply.Validate if the designated constraints aren't met.
type ListWordReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWordReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWordReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWordReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWordReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWordReplyValidationError) ErrorName() string { return "ListWordReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListWordReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWordReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWordReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWordReplyValidationError{}
