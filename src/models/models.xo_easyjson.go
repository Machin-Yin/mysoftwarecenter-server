// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson91a2f4f5DecodeModels(in *jlexer.Lexer, out *ScUser) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = uint(in.Uint())
		case "user_name":
			out.UserName = string(in.String())
		case "avatar_url":
			out.AvatarURL = string(in.String())
		case "mail":
			out.Mail = string(in.String())
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
func easyjson91a2f4f5EncodeModels(out *jwriter.Writer, in ScUser) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ID\":")
		out.Uint(uint(in.ID))
	}
	if in.UserName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"user_name\":")
		out.String(string(in.UserName))
	}
	if in.AvatarURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"avatar_url\":")
		out.String(string(in.AvatarURL))
	}
	if in.Mail != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"mail\":")
		out.String(string(in.Mail))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ScUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson91a2f4f5EncodeModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ScUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson91a2f4f5EncodeModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ScUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson91a2f4f5DecodeModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ScUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson91a2f4f5DecodeModels(l, v)
}
func easyjson91a2f4f5DecodeModels1(in *jlexer.Lexer, out *ScScreenImage) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = uint(in.Uint())
		case "product_ID":
			out.ProductID = uint(in.Uint())
		case "release_ID":
			out.ReleaseID = uint(in.Uint())
		case "image_url":
			out.ImageURL = string(in.String())
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
func easyjson91a2f4f5EncodeModels1(out *jwriter.Writer, in ScScreenImage) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ID\":")
		out.Uint(uint(in.ID))
	}
	if in.ProductID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"product_ID\":")
		out.Uint(uint(in.ProductID))
	}
	if in.ReleaseID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"release_ID\":")
		out.Uint(uint(in.ReleaseID))
	}
	if in.ImageURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"image_url\":")
		out.String(string(in.ImageURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ScScreenImage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson91a2f4f5EncodeModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ScScreenImage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson91a2f4f5EncodeModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ScScreenImage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson91a2f4f5DecodeModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ScScreenImage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson91a2f4f5DecodeModels1(l, v)
}
func easyjson91a2f4f5DecodeModels2(in *jlexer.Lexer, out *ScRelease) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = uint(in.Uint())
		case "product_ID":
			out.ProductID = uint(in.Uint())
		case "product_name":
			out.ProductName = string(in.String())
		case "version":
			out.Version = string(in.String())
		case "icon_url":
			out.IconURL = string(in.String())
		case "download_url":
			out.DownloadURL = string(in.String())
		case "changelog":
			out.Changelog = string(in.String())
		case "package_size":
			out.PackageSize = uint(in.Uint())
		case "package_type":
			out.PackageType = int8(in.Int8())
		case "release_grade":
			out.ReleaseGrade = uint(in.Uint())
		case "grade_count":
			out.GradeCount = float32(in.Float32())
		case "release_date":
			out.ReleaseDate = uint(in.Uint())
		case "executable_file":
			out.ExecutableFile = string(in.String())
		case "package_name":
			out.PackageName = string(in.String())
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
func easyjson91a2f4f5EncodeModels2(out *jwriter.Writer, in ScRelease) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ID\":")
		out.Uint(uint(in.ID))
	}
	if in.ProductID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"product_ID\":")
		out.Uint(uint(in.ProductID))
	}
	if in.ProductName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"product_name\":")
		out.String(string(in.ProductName))
	}
	if in.Version != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"version\":")
		out.String(string(in.Version))
	}
	if in.IconURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"icon_url\":")
		out.String(string(in.IconURL))
	}
	if in.DownloadURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"download_url\":")
		out.String(string(in.DownloadURL))
	}
	if in.Changelog != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"changelog\":")
		out.String(string(in.Changelog))
	}
	if in.PackageSize != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"package_size\":")
		out.Uint(uint(in.PackageSize))
	}
	if in.PackageType != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"package_type\":")
		out.Int8(int8(in.PackageType))
	}
	if in.ReleaseGrade != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"release_grade\":")
		out.Uint(uint(in.ReleaseGrade))
	}
	if in.GradeCount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"grade_count\":")
		out.Float32(float32(in.GradeCount))
	}
	if in.ReleaseDate != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"release_date\":")
		out.Uint(uint(in.ReleaseDate))
	}
	if in.ExecutableFile != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"executable_file\":")
		out.String(string(in.ExecutableFile))
	}
	if in.PackageName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"package_name\":")
		out.String(string(in.PackageName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ScRelease) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson91a2f4f5EncodeModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ScRelease) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson91a2f4f5EncodeModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ScRelease) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson91a2f4f5DecodeModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ScRelease) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson91a2f4f5DecodeModels2(l, v)
}
func easyjson91a2f4f5DecodeModels3(in *jlexer.Lexer, out *ScRecommend) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = uint(in.Uint())
		case "priority":
			out.Priority = int8(in.Int8())
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
func easyjson91a2f4f5EncodeModels3(out *jwriter.Writer, in ScRecommend) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ID\":")
		out.Uint(uint(in.ID))
	}
	if in.Priority != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"priority\":")
		out.Int8(int8(in.Priority))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ScRecommend) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson91a2f4f5EncodeModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ScRecommend) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson91a2f4f5EncodeModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ScRecommend) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson91a2f4f5DecodeModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ScRecommend) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson91a2f4f5DecodeModels3(l, v)
}
func easyjson91a2f4f5DecodeModels4(in *jlexer.Lexer, out *ScProduct) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = uint(in.Uint())
		case "category_ID":
			out.CategoryID = uint(in.Uint())
		case "release_ID":
			out.ReleaseID = uint(in.Uint())
		case "product_name":
			out.ProductName = string(in.String())
		case "vendor_name":
			out.VendorName = string(in.String())
		case "icon_url":
			out.IconURL = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "product_description":
			out.ProductDescription = string(in.String())
		case "product_grade":
			out.ProductGrade = uint(in.Uint())
		case "grade_count":
			out.GradeCount = float32(in.Float32())
		case "executable_file":
			out.ExecutableFile = string(in.String())
		case "package_name":
			out.PackageName = string(in.String())
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
func easyjson91a2f4f5EncodeModels4(out *jwriter.Writer, in ScProduct) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ID\":")
		out.Uint(uint(in.ID))
	}
	if in.CategoryID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"category_ID\":")
		out.Uint(uint(in.CategoryID))
	}
	if in.ReleaseID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"release_ID\":")
		out.Uint(uint(in.ReleaseID))
	}
	if in.ProductName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"product_name\":")
		out.String(string(in.ProductName))
	}
	if in.VendorName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"vendor_name\":")
		out.String(string(in.VendorName))
	}
	if in.IconURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"icon_url\":")
		out.String(string(in.IconURL))
	}
	if in.URL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"url\":")
		out.String(string(in.URL))
	}
	if in.ProductDescription != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"product_description\":")
		out.String(string(in.ProductDescription))
	}
	if in.ProductGrade != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"product_grade\":")
		out.Uint(uint(in.ProductGrade))
	}
	if in.GradeCount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"grade_count\":")
		out.Float32(float32(in.GradeCount))
	}
	if in.ExecutableFile != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"executable_file\":")
		out.String(string(in.ExecutableFile))
	}
	if in.PackageName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"package_name\":")
		out.String(string(in.PackageName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ScProduct) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson91a2f4f5EncodeModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ScProduct) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson91a2f4f5EncodeModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ScProduct) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson91a2f4f5DecodeModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ScProduct) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson91a2f4f5DecodeModels4(l, v)
}
func easyjson91a2f4f5DecodeModels5(in *jlexer.Lexer, out *ScComment) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = uint(in.Uint())
		case "product_ID":
			out.ProductID = uint(in.Uint())
		case "release_ID":
			out.ReleaseID = uint(in.Uint())
		case "user_ID":
			out.UserID = uint(in.Uint())
		case "comment_text":
			out.CommentText = string(in.String())
		case "comment_grade":
			out.CommentGrade = int8(in.Int8())
		case "comment_date":
			out.CommentDate = uint(in.Uint())
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
func easyjson91a2f4f5EncodeModels5(out *jwriter.Writer, in ScComment) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ID\":")
		out.Uint(uint(in.ID))
	}
	if in.ProductID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"product_ID\":")
		out.Uint(uint(in.ProductID))
	}
	if in.ReleaseID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"release_ID\":")
		out.Uint(uint(in.ReleaseID))
	}
	if in.UserID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"user_ID\":")
		out.Uint(uint(in.UserID))
	}
	if in.CommentText != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"comment_text\":")
		out.String(string(in.CommentText))
	}
	if in.CommentGrade != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"comment_grade\":")
		out.Int8(int8(in.CommentGrade))
	}
	if in.CommentDate != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"comment_date\":")
		out.Uint(uint(in.CommentDate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ScComment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson91a2f4f5EncodeModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ScComment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson91a2f4f5EncodeModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ScComment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson91a2f4f5DecodeModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ScComment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson91a2f4f5DecodeModels5(l, v)
}
func easyjson91a2f4f5DecodeModels6(in *jlexer.Lexer, out *ScCategory) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = uint(in.Uint())
		case "category_name":
			out.CategoryName = string(in.String())
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
func easyjson91a2f4f5EncodeModels6(out *jwriter.Writer, in ScCategory) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ID\":")
		out.Uint(uint(in.ID))
	}
	if in.CategoryName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"category_name\":")
		out.String(string(in.CategoryName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ScCategory) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson91a2f4f5EncodeModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ScCategory) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson91a2f4f5EncodeModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ScCategory) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson91a2f4f5DecodeModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ScCategory) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson91a2f4f5DecodeModels6(l, v)
}
func easyjson91a2f4f5DecodeModels7(in *jlexer.Lexer, out *ScBanner) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = uint(in.Uint())
		case "priority":
			out.Priority = int8(in.Int8())
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
func easyjson91a2f4f5EncodeModels7(out *jwriter.Writer, in ScBanner) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ID\":")
		out.Uint(uint(in.ID))
	}
	if in.Priority != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"priority\":")
		out.Int8(int8(in.Priority))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ScBanner) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson91a2f4f5EncodeModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ScBanner) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson91a2f4f5EncodeModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ScBanner) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson91a2f4f5DecodeModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ScBanner) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson91a2f4f5DecodeModels7(l, v)
}
