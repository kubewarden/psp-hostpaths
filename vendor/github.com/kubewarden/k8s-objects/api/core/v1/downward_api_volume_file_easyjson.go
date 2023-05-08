// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
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

func easyjson72afaf59DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *DownwardAPIVolumeFile) {
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
		case "fieldRef":
			if in.IsNull() {
				in.Skip()
				out.FieldRef = nil
			} else {
				if out.FieldRef == nil {
					out.FieldRef = new(ObjectFieldSelector)
				}
				easyjson72afaf59DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.FieldRef)
			}
		case "mode":
			out.Mode = int32(in.Int32())
		case "path":
			if in.IsNull() {
				in.Skip()
				out.Path = nil
			} else {
				if out.Path == nil {
					out.Path = new(string)
				}
				*out.Path = string(in.String())
			}
		case "resourceFieldRef":
			if in.IsNull() {
				in.Skip()
				out.ResourceFieldRef = nil
			} else {
				if out.ResourceFieldRef == nil {
					out.ResourceFieldRef = new(ResourceFieldSelector)
				}
				easyjson72afaf59DecodeGithubComKubewardenK8sObjectsApiCoreV12(in, out.ResourceFieldRef)
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
func easyjson72afaf59EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in DownwardAPIVolumeFile) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FieldRef != nil {
		const prefix string = ",\"fieldRef\":"
		first = false
		out.RawString(prefix[1:])
		easyjson72afaf59EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.FieldRef)
	}
	if in.Mode != 0 {
		const prefix string = ",\"mode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Mode))
	}
	{
		const prefix string = ",\"path\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Path == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Path))
		}
	}
	if in.ResourceFieldRef != nil {
		const prefix string = ",\"resourceFieldRef\":"
		out.RawString(prefix)
		easyjson72afaf59EncodeGithubComKubewardenK8sObjectsApiCoreV12(out, *in.ResourceFieldRef)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DownwardAPIVolumeFile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson72afaf59EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DownwardAPIVolumeFile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson72afaf59EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DownwardAPIVolumeFile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson72afaf59DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DownwardAPIVolumeFile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson72afaf59DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjson72afaf59DecodeGithubComKubewardenK8sObjectsApiCoreV12(in *jlexer.Lexer, out *ResourceFieldSelector) {
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
		case "containerName":
			out.ContainerName = string(in.String())
		case "divisor":
			if in.IsNull() {
				in.Skip()
				out.Divisor = nil
			} else {
				if out.Divisor == nil {
					out.Divisor = new(resource.Quantity)
				}
				*out.Divisor = resource.Quantity(in.String())
			}
		case "resource":
			if in.IsNull() {
				in.Skip()
				out.Resource = nil
			} else {
				if out.Resource == nil {
					out.Resource = new(string)
				}
				*out.Resource = string(in.String())
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
func easyjson72afaf59EncodeGithubComKubewardenK8sObjectsApiCoreV12(out *jwriter.Writer, in ResourceFieldSelector) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ContainerName != "" {
		const prefix string = ",\"containerName\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ContainerName))
	}
	if in.Divisor != nil {
		const prefix string = ",\"divisor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Divisor))
	}
	{
		const prefix string = ",\"resource\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Resource == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Resource))
		}
	}
	out.RawByte('}')
}
func easyjson72afaf59DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *ObjectFieldSelector) {
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
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "fieldPath":
			if in.IsNull() {
				in.Skip()
				out.FieldPath = nil
			} else {
				if out.FieldPath == nil {
					out.FieldPath = new(string)
				}
				*out.FieldPath = string(in.String())
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
func easyjson72afaf59EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in ObjectFieldSelector) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	{
		const prefix string = ",\"fieldPath\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.FieldPath == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.FieldPath))
		}
	}
	out.RawByte('}')
}