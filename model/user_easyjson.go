// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson9e1087fdDecodeAppConstructorBackendModel(in *jlexer.Lexer, out *UserDataJwt) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Sub = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "email":
			out.Email = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9e1087fdEncodeAppConstructorBackendModel(out *jwriter.Writer, in UserDataJwt) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Sub))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserDataJwt) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeAppConstructorBackendModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserDataJwt) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeAppConstructorBackendModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserDataJwt) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeAppConstructorBackendModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserDataJwt) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeAppConstructorBackendModel(l, v)
}
func easyjson9e1087fdDecodeAppConstructorBackendModel1(in *jlexer.Lexer, out *UserClaims) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "aud":
			out.Audience = string(in.String())
		case "exp":
			out.ExpiresAt = int64(in.Int64())
		case "jti":
			out.Id = string(in.String())
		case "iat":
			out.IssuedAt = int64(in.Int64())
		case "iss":
			out.Issuer = string(in.String())
		case "nbf":
			out.NotBefore = int64(in.Int64())
		case "sub":
			out.Subject = string(in.String())
		case "id":
			out.Sub = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "email":
			out.Email = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9e1087fdEncodeAppConstructorBackendModel1(out *jwriter.Writer, in UserClaims) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Audience != "" {
		const prefix string = ",\"aud\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Audience))
	}
	if in.ExpiresAt != 0 {
		const prefix string = ",\"exp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ExpiresAt))
	}
	if in.Id != "" {
		const prefix string = ",\"jti\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Id))
	}
	if in.IssuedAt != 0 {
		const prefix string = ",\"iat\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.IssuedAt))
	}
	if in.Issuer != "" {
		const prefix string = ",\"iss\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Issuer))
	}
	if in.NotBefore != 0 {
		const prefix string = ",\"nbf\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.NotBefore))
	}
	if in.Subject != "" {
		const prefix string = ",\"sub\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Subject))
	}
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Sub))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserClaims) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeAppConstructorBackendModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserClaims) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeAppConstructorBackendModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserClaims) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeAppConstructorBackendModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserClaims) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeAppConstructorBackendModel1(l, v)
}
func easyjson9e1087fdDecodeAppConstructorBackendModel2(in *jlexer.Lexer, out *Response) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "widgetsCount":
			out.WidgetsCount = int(in.Int())
		case "projects":
			if in.IsNull() {
				in.Skip()
				out.Projects = nil
			} else {
				in.Delim('[')
				if out.Projects == nil {
					if !in.IsDelim(']') {
						out.Projects = make([]Project, 0, 1)
					} else {
						out.Projects = []Project{}
					}
				} else {
					out.Projects = (out.Projects)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Project
					(v1).UnmarshalEasyJSON(in)
					out.Projects = append(out.Projects, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9e1087fdEncodeAppConstructorBackendModel2(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"widgetsCount\":"
		out.RawString(prefix[1:])
		out.Int(int(in.WidgetsCount))
	}
	{
		const prefix string = ",\"projects\":"
		out.RawString(prefix)
		if in.Projects == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Projects {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeAppConstructorBackendModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeAppConstructorBackendModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeAppConstructorBackendModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeAppConstructorBackendModel2(l, v)
}
func easyjson9e1087fdDecodeAppConstructorBackendModel3(in *jlexer.Lexer, out *RequestBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "widgetsCount":
			out.WidgetsCount = int(in.Int())
		case "project":
			(out.Project).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9e1087fdEncodeAppConstructorBackendModel3(out *jwriter.Writer, in RequestBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"widgetsCount\":"
		out.RawString(prefix[1:])
		out.Int(int(in.WidgetsCount))
	}
	{
		const prefix string = ",\"project\":"
		out.RawString(prefix)
		(in.Project).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeAppConstructorBackendModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeAppConstructorBackendModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeAppConstructorBackendModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeAppConstructorBackendModel3(l, v)
}
func easyjson9e1087fdDecodeAppConstructorBackendModel4(in *jlexer.Lexer, out *Project) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "app":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.App).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9e1087fdEncodeAppConstructorBackendModel4(out *jwriter.Writer, in Project) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"app\":"
		out.RawString(prefix)
		out.Raw((in.App).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Project) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeAppConstructorBackendModel4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Project) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeAppConstructorBackendModel4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Project) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeAppConstructorBackendModel4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Project) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeAppConstructorBackendModel4(l, v)
}
