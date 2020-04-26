// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: article.proto

package articlepb

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _article_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AddArticleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AddArticleRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		return AddArticleRequestValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetContent()) < 1 {
		return AddArticleRequestValidationError{
			field:  "Content",
			reason: "value length must be at least 1 runes",
		}
	}

	if len(m.GetAuthor()) > 256 {
		return AddArticleRequestValidationError{
			field:  "Author",
			reason: "value length must be at most 256 bytes",
		}
	}

	return nil
}

// AddArticleRequestValidationError is the validation error returned by
// AddArticleRequest.Validate if the designated constraints aren't met.
type AddArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddArticleRequestValidationError) ErrorName() string {
	return "AddArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddArticleRequestValidationError{}

// Validate checks the field values on AddArticleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AddArticleResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddArticleResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AddArticleResponseValidationError is the validation error returned by
// AddArticleResponse.Validate if the designated constraints aren't met.
type AddArticleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddArticleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddArticleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddArticleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddArticleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddArticleResponseValidationError) ErrorName() string {
	return "AddArticleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddArticleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddArticleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddArticleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddArticleResponseValidationError{}

// Validate checks the field values on GetArticleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetArticleRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetArticleRequestValidationError is the validation error returned by
// GetArticleRequest.Validate if the designated constraints aren't met.
type GetArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetArticleRequestValidationError) ErrorName() string {
	return "GetArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetArticleRequestValidationError{}

// Validate checks the field values on GetArticleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetArticleResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetArticleResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetArticleResponseValidationError is the validation error returned by
// GetArticleResponse.Validate if the designated constraints aren't met.
type GetArticleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetArticleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetArticleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetArticleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetArticleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetArticleResponseValidationError) ErrorName() string {
	return "GetArticleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetArticleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetArticleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetArticleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetArticleResponseValidationError{}

// Validate checks the field values on ListArticlesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListArticlesRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListArticlesRequestValidationError is the validation error returned by
// ListArticlesRequest.Validate if the designated constraints aren't met.
type ListArticlesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListArticlesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListArticlesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListArticlesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListArticlesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListArticlesRequestValidationError) ErrorName() string {
	return "ListArticlesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListArticlesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListArticlesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListArticlesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListArticlesRequestValidationError{}

// Validate checks the field values on ListArticlesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListArticlesResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	// no validation rules for Message

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListArticlesResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListArticlesResponseValidationError is the validation error returned by
// ListArticlesResponse.Validate if the designated constraints aren't met.
type ListArticlesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListArticlesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListArticlesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListArticlesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListArticlesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListArticlesResponseValidationError) ErrorName() string {
	return "ListArticlesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListArticlesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListArticlesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListArticlesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListArticlesResponseValidationError{}

// Validate checks the field values on Article with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Article) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Content

	// no validation rules for Author

	return nil
}

// ArticleValidationError is the validation error returned by Article.Validate
// if the designated constraints aren't met.
type ArticleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleValidationError) ErrorName() string { return "ArticleValidationError" }

// Error satisfies the builtin error interface
func (e ArticleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticle.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleValidationError{}

// Validate checks the field values on AddArticleResponse_ArticleId with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AddArticleResponse_ArticleId) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// AddArticleResponse_ArticleIdValidationError is the validation error returned
// by AddArticleResponse_ArticleId.Validate if the designated constraints
// aren't met.
type AddArticleResponse_ArticleIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddArticleResponse_ArticleIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddArticleResponse_ArticleIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddArticleResponse_ArticleIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddArticleResponse_ArticleIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddArticleResponse_ArticleIdValidationError) ErrorName() string {
	return "AddArticleResponse_ArticleIdValidationError"
}

// Error satisfies the builtin error interface
func (e AddArticleResponse_ArticleIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddArticleResponse_ArticleId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddArticleResponse_ArticleIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddArticleResponse_ArticleIdValidationError{}
