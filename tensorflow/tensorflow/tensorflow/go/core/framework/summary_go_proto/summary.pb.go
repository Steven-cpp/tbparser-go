// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: tensorboard/compat/proto/summary.proto

package summary_go_proto

import (
	summary_go_proto "github.com/google/tsl/tsl/go/core/protobuf/summary_go_proto"
	tensor_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of tensorboard/compat/proto/histogram.proto.

type HistogramProto = summary_go_proto.HistogramProto

type DataClass int32

const (
	// Unknown data class, used (implicitly) for legacy data. Will not be
	// processed by data ingestion pipelines.
	DataClass_DATA_CLASS_UNKNOWN DataClass = 0
	// Scalar time series. Each `Value` for the corresponding tag must have
	// `tensor` set to a rank-0 tensor of type `DT_FLOAT` (float32).
	DataClass_DATA_CLASS_SCALAR DataClass = 1
	// Tensor time series. Each `Value` for the corresponding tag must have
	// `tensor` set. The tensor value is arbitrary, but should be small to
	// accommodate direct storage in database backends: an upper bound of a few
	// kilobytes is a reasonable rule of thumb.
	DataClass_DATA_CLASS_TENSOR DataClass = 2
	// Blob sequence time series. Each `Value` for the corresponding tag must
	// have `tensor` set to a rank-1 tensor of bytestring dtype.
	DataClass_DATA_CLASS_BLOB_SEQUENCE DataClass = 3
)

// Enum value maps for DataClass.
var (
	DataClass_name = map[int32]string{
		0: "DATA_CLASS_UNKNOWN",
		1: "DATA_CLASS_SCALAR",
		2: "DATA_CLASS_TENSOR",
		3: "DATA_CLASS_BLOB_SEQUENCE",
	}
	DataClass_value = map[string]int32{
		"DATA_CLASS_UNKNOWN":       0,
		"DATA_CLASS_SCALAR":        1,
		"DATA_CLASS_TENSOR":        2,
		"DATA_CLASS_BLOB_SEQUENCE": 3,
	}
)

func (x DataClass) Enum() *DataClass {
	p := new(DataClass)
	*p = x
	return p
}

func (x DataClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataClass) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorboard_compat_proto_summary_proto_enumTypes[0].Descriptor()
}

func (DataClass) Type() protoreflect.EnumType {
	return &file_tensorboard_compat_proto_summary_proto_enumTypes[0]
}

func (x DataClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataClass.Descriptor instead.
func (DataClass) EnumDescriptor() ([]byte, []int) {
	return file_tensorboard_compat_proto_summary_proto_rawDescGZIP(), []int{0}
}

// Metadata associated with a series of Summary data
type SummaryDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hint on how plugins should process the data in this series.
	// Supported values include "scalar", "histogram", "image", "audio"
	TypeHint string `protobuf:"bytes,1,opt,name=type_hint,json=typeHint,proto3" json:"type_hint,omitempty"`
}

func (x *SummaryDescription) Reset() {
	*x = SummaryDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryDescription) ProtoMessage() {}

func (x *SummaryDescription) ProtoReflect() protoreflect.Message {
	mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryDescription.ProtoReflect.Descriptor instead.
func (*SummaryDescription) Descriptor() ([]byte, []int) {
	return file_tensorboard_compat_proto_summary_proto_rawDescGZIP(), []int{0}
}

func (x *SummaryDescription) GetTypeHint() string {
	if x != nil {
		return x.TypeHint
	}
	return ""
}

// A SummaryMetadata encapsulates information on which plugins are able to make
// use of a certain summary value.
type SummaryMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data that associates a summary with a certain plugin.
	PluginData *SummaryMetadata_PluginData `protobuf:"bytes,1,opt,name=plugin_data,json=pluginData,proto3" json:"plugin_data,omitempty"`
	// Display name for viewing in TensorBoard.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Longform readable description of the summary sequence. Markdown supported.
	SummaryDescription string `protobuf:"bytes,3,opt,name=summary_description,json=summaryDescription,proto3" json:"summary_description,omitempty"`
	// Class of data stored in this time series. Required for compatibility with
	// TensorBoard's generic data facilities (`DataProvider`, et al.). This value
	// imposes constraints on the dtype and shape of the corresponding tensor
	// values. See `DataClass` docs for details.
	DataClass DataClass `protobuf:"varint,4,opt,name=data_class,json=dataClass,proto3,enum=tensorboard.DataClass" json:"data_class,omitempty"`
}

func (x *SummaryMetadata) Reset() {
	*x = SummaryMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryMetadata) ProtoMessage() {}

func (x *SummaryMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryMetadata.ProtoReflect.Descriptor instead.
func (*SummaryMetadata) Descriptor() ([]byte, []int) {
	return file_tensorboard_compat_proto_summary_proto_rawDescGZIP(), []int{1}
}

func (x *SummaryMetadata) GetPluginData() *SummaryMetadata_PluginData {
	if x != nil {
		return x.PluginData
	}
	return nil
}

func (x *SummaryMetadata) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *SummaryMetadata) GetSummaryDescription() string {
	if x != nil {
		return x.SummaryDescription
	}
	return ""
}

func (x *SummaryMetadata) GetDataClass() DataClass {
	if x != nil {
		return x.DataClass
	}
	return DataClass_DATA_CLASS_UNKNOWN
}

// A Summary is a set of named values to be displayed by the
// visualizer.
//
// Summaries are produced regularly during training, as controlled by
// the "summary_interval_secs" attribute of the training operation.
// Summaries are also produced at the end of an evaluation.
type Summary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Set of values for the summary.
	Value []*Summary_Value `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *Summary) Reset() {
	*x = Summary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary) ProtoMessage() {}

func (x *Summary) ProtoReflect() protoreflect.Message {
	mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary.ProtoReflect.Descriptor instead.
func (*Summary) Descriptor() ([]byte, []int) {
	return file_tensorboard_compat_proto_summary_proto_rawDescGZIP(), []int{2}
}

func (x *Summary) GetValue() []*Summary_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type SummaryMetadata_PluginData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the plugin this data pertains to.
	PluginName string `protobuf:"bytes,1,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
	// The content to store for the plugin. The best practice is for this to be
	// a binary serialized protocol buffer.
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SummaryMetadata_PluginData) Reset() {
	*x = SummaryMetadata_PluginData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryMetadata_PluginData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryMetadata_PluginData) ProtoMessage() {}

func (x *SummaryMetadata_PluginData) ProtoReflect() protoreflect.Message {
	mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryMetadata_PluginData.ProtoReflect.Descriptor instead.
func (*SummaryMetadata_PluginData) Descriptor() ([]byte, []int) {
	return file_tensorboard_compat_proto_summary_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SummaryMetadata_PluginData) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

func (x *SummaryMetadata_PluginData) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type Summary_Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Dimensions of the image.
	Height int32 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Width  int32 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	// Valid colorspace values are
	//
	//	1 - grayscale
	//	2 - grayscale + alpha
	//	3 - RGB
	//	4 - RGBA
	//	5 - DIGITAL_YUV
	//	6 - BGRA
	Colorspace int32 `protobuf:"varint,3,opt,name=colorspace,proto3" json:"colorspace,omitempty"`
	// Image data in encoded format.  All image formats supported by
	// image_codec::CoderUtil can be stored here.
	EncodedImageString []byte `protobuf:"bytes,4,opt,name=encoded_image_string,json=encodedImageString,proto3" json:"encoded_image_string,omitempty"`
}

func (x *Summary_Image) Reset() {
	*x = Summary_Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary_Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary_Image) ProtoMessage() {}

func (x *Summary_Image) ProtoReflect() protoreflect.Message {
	mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary_Image.ProtoReflect.Descriptor instead.
func (*Summary_Image) Descriptor() ([]byte, []int) {
	return file_tensorboard_compat_proto_summary_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Summary_Image) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Summary_Image) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Summary_Image) GetColorspace() int32 {
	if x != nil {
		return x.Colorspace
	}
	return 0
}

func (x *Summary_Image) GetEncodedImageString() []byte {
	if x != nil {
		return x.EncodedImageString
	}
	return nil
}

type Summary_Audio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sample rate of the audio in Hz.
	SampleRate float32 `protobuf:"fixed32,1,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
	// Number of channels of audio.
	NumChannels int64 `protobuf:"varint,2,opt,name=num_channels,json=numChannels,proto3" json:"num_channels,omitempty"`
	// Length of the audio in frames (samples per channel).
	LengthFrames int64 `protobuf:"varint,3,opt,name=length_frames,json=lengthFrames,proto3" json:"length_frames,omitempty"`
	// Encoded audio data and its associated RFC 2045 content type (e.g.
	// "audio/wav").
	EncodedAudioString []byte `protobuf:"bytes,4,opt,name=encoded_audio_string,json=encodedAudioString,proto3" json:"encoded_audio_string,omitempty"`
	ContentType        string `protobuf:"bytes,5,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *Summary_Audio) Reset() {
	*x = Summary_Audio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary_Audio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary_Audio) ProtoMessage() {}

func (x *Summary_Audio) ProtoReflect() protoreflect.Message {
	mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary_Audio.ProtoReflect.Descriptor instead.
func (*Summary_Audio) Descriptor() ([]byte, []int) {
	return file_tensorboard_compat_proto_summary_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Summary_Audio) GetSampleRate() float32 {
	if x != nil {
		return x.SampleRate
	}
	return 0
}

func (x *Summary_Audio) GetNumChannels() int64 {
	if x != nil {
		return x.NumChannels
	}
	return 0
}

func (x *Summary_Audio) GetLengthFrames() int64 {
	if x != nil {
		return x.LengthFrames
	}
	return 0
}

func (x *Summary_Audio) GetEncodedAudioString() []byte {
	if x != nil {
		return x.EncodedAudioString
	}
	return nil
}

func (x *Summary_Audio) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

type Summary_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field is deprecated and will not be set.
	NodeName string `protobuf:"bytes,7,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Tag name for the data. Used by TensorBoard plugins to organize data. Tags
	// are often organized by scope (which contains slashes to convey
	// hierarchy). For example: foo/bar/0
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Contains metadata on the summary value such as which plugins may use it.
	// Take note that many summary values may lack a metadata field. This is
	// because the FileWriter only keeps a metadata object on the first summary
	// value with a certain tag for each tag. TensorBoard then remembers which
	// tags are associated with which plugins. This saves space.
	Metadata *SummaryMetadata `protobuf:"bytes,9,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Value associated with the tag.
	//
	// Types that are assignable to Value:
	//
	//	*Summary_Value_SimpleValue
	//	*Summary_Value_ObsoleteOldStyleHistogram
	//	*Summary_Value_Image
	//	*Summary_Value_Histo
	//	*Summary_Value_Audio
	//	*Summary_Value_Tensor
	Value isSummary_Value_Value `protobuf_oneof:"value"`
}

func (x *Summary_Value) Reset() {
	*x = Summary_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary_Value) ProtoMessage() {}

func (x *Summary_Value) ProtoReflect() protoreflect.Message {
	mi := &file_tensorboard_compat_proto_summary_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary_Value.ProtoReflect.Descriptor instead.
func (*Summary_Value) Descriptor() ([]byte, []int) {
	return file_tensorboard_compat_proto_summary_proto_rawDescGZIP(), []int{2, 2}
}

func (x *Summary_Value) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *Summary_Value) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Summary_Value) GetMetadata() *SummaryMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (m *Summary_Value) GetValue() isSummary_Value_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Summary_Value) GetSimpleValue() float32 {
	if x, ok := x.GetValue().(*Summary_Value_SimpleValue); ok {
		return x.SimpleValue
	}
	return 0
}

func (x *Summary_Value) GetObsoleteOldStyleHistogram() []byte {
	if x, ok := x.GetValue().(*Summary_Value_ObsoleteOldStyleHistogram); ok {
		return x.ObsoleteOldStyleHistogram
	}
	return nil
}

func (x *Summary_Value) GetImage() *Summary_Image {
	if x, ok := x.GetValue().(*Summary_Value_Image); ok {
		return x.Image
	}
	return nil
}

func (x *Summary_Value) GetHisto() *summary_go_proto.HistogramProto {
	if x, ok := x.GetValue().(*Summary_Value_Histo); ok {
		return x.Histo
	}
	return nil
}

func (x *Summary_Value) GetAudio() *Summary_Audio {
	if x, ok := x.GetValue().(*Summary_Value_Audio); ok {
		return x.Audio
	}
	return nil
}

func (x *Summary_Value) GetTensor() *tensor_go_proto.TensorProto {
	if x, ok := x.GetValue().(*Summary_Value_Tensor); ok {
		return x.Tensor
	}
	return nil
}

type isSummary_Value_Value interface {
	isSummary_Value_Value()
}

type Summary_Value_SimpleValue struct {
	SimpleValue float32 `protobuf:"fixed32,2,opt,name=simple_value,json=simpleValue,proto3,oneof"`
}

type Summary_Value_ObsoleteOldStyleHistogram struct {
	ObsoleteOldStyleHistogram []byte `protobuf:"bytes,3,opt,name=obsolete_old_style_histogram,json=obsoleteOldStyleHistogram,proto3,oneof"`
}

type Summary_Value_Image struct {
	Image *Summary_Image `protobuf:"bytes,4,opt,name=image,proto3,oneof"`
}

type Summary_Value_Histo struct {
	Histo *summary_go_proto.HistogramProto `protobuf:"bytes,5,opt,name=histo,proto3,oneof"`
}

type Summary_Value_Audio struct {
	Audio *Summary_Audio `protobuf:"bytes,6,opt,name=audio,proto3,oneof"`
}

type Summary_Value_Tensor struct {
	Tensor *tensor_go_proto.TensorProto `protobuf:"bytes,8,opt,name=tensor,proto3,oneof"`
}

func (*Summary_Value_SimpleValue) isSummary_Value_Value() {}

func (*Summary_Value_ObsoleteOldStyleHistogram) isSummary_Value_Value() {}

func (*Summary_Value_Image) isSummary_Value_Value() {}

func (*Summary_Value_Histo) isSummary_Value_Value() {}

func (*Summary_Value_Audio) isSummary_Value_Value() {}

func (*Summary_Value_Tensor) isSummary_Value_Value() {}

var File_tensorboard_compat_proto_summary_proto protoreflect.FileDescriptor

var file_tensorboard_compat_proto_summary_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x12, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x79, 0x70, 0x65, 0x48, 0x69, 0x6e, 0x74, 0x22, 0xaf, 0x02, 0x0a, 0x0f, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x48, 0x0a,
	0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x1a, 0x47, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xc2, 0x06, 0x0a, 0x07,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x87, 0x01, 0x0a, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x12, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x1a, 0xc5, 0x01, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x64, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x41, 0x75, 0x64,
	0x69, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x1a, 0xb2, 0x03, 0x0a, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x74, 0x61, 0x67, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23,
	0x0a, 0x0c, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x41, 0x0a, 0x1c, 0x6f, 0x62, 0x73, 0x6f, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x19, 0x6f, 0x62, 0x73,
	0x6f, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x6c, 0x64, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x12,
	0x32, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x61, 0x75,
	0x64, 0x69, 0x6f, 0x12, 0x32, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x00, 0x52,
	0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2a, 0x6f, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x12, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x43, 0x4c,
	0x41, 0x53, 0x53, 0x5f, 0x53, 0x43, 0x41, 0x4c, 0x41, 0x52, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x54, 0x45, 0x4e, 0x53, 0x4f,
	0x52, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x43, 0x4c, 0x41, 0x53,
	0x53, 0x5f, 0x42, 0x4c, 0x4f, 0x42, 0x5f, 0x53, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x45, 0x10,
	0x03, 0x42, 0x7e, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x0d, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x4e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01,
	0x01, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorboard_compat_proto_summary_proto_rawDescOnce sync.Once
	file_tensorboard_compat_proto_summary_proto_rawDescData = file_tensorboard_compat_proto_summary_proto_rawDesc
)

func file_tensorboard_compat_proto_summary_proto_rawDescGZIP() []byte {
	file_tensorboard_compat_proto_summary_proto_rawDescOnce.Do(func() {
		file_tensorboard_compat_proto_summary_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorboard_compat_proto_summary_proto_rawDescData)
	})
	return file_tensorboard_compat_proto_summary_proto_rawDescData
}

var file_tensorboard_compat_proto_summary_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorboard_compat_proto_summary_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tensorboard_compat_proto_summary_proto_goTypes = []interface{}{
	(DataClass)(0),                          // 0: tensorboard.DataClass
	(*SummaryDescription)(nil),              // 1: tensorboard.SummaryDescription
	(*SummaryMetadata)(nil),                 // 2: tensorboard.SummaryMetadata
	(*Summary)(nil),                         // 3: tensorboard.Summary
	(*SummaryMetadata_PluginData)(nil),      // 4: tensorboard.SummaryMetadata.PluginData
	(*Summary_Image)(nil),                   // 5: tensorboard.Summary.Image
	(*Summary_Audio)(nil),                   // 6: tensorboard.Summary.Audio
	(*Summary_Value)(nil),                   // 7: tensorboard.Summary.Value
	(*summary_go_proto.HistogramProto)(nil), // 8: tensorboard.HistogramProto
	(*tensor_go_proto.TensorProto)(nil),     // 9: tensorboard.TensorProto
}
var file_tensorboard_compat_proto_summary_proto_depIdxs = []int32{
	4, // 0: tensorboard.SummaryMetadata.plugin_data:type_name -> tensorboard.SummaryMetadata.PluginData
	0, // 1: tensorboard.SummaryMetadata.data_class:type_name -> tensorboard.DataClass
	7, // 2: tensorboard.Summary.value:type_name -> tensorboard.Summary.Value
	2, // 3: tensorboard.Summary.Value.metadata:type_name -> tensorboard.SummaryMetadata
	5, // 4: tensorboard.Summary.Value.image:type_name -> tensorboard.Summary.Image
	8, // 5: tensorboard.Summary.Value.histo:type_name -> tensorboard.HistogramProto
	6, // 6: tensorboard.Summary.Value.audio:type_name -> tensorboard.Summary.Audio
	9, // 7: tensorboard.Summary.Value.tensor:type_name -> tensorboard.TensorProto
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_tensorboard_compat_proto_summary_proto_init() }
func file_tensorboard_compat_proto_summary_proto_init() {
	if File_tensorboard_compat_proto_summary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorboard_compat_proto_summary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryDescription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorboard_compat_proto_summary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorboard_compat_proto_summary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summary); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorboard_compat_proto_summary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryMetadata_PluginData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorboard_compat_proto_summary_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summary_Image); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorboard_compat_proto_summary_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summary_Audio); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorboard_compat_proto_summary_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summary_Value); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_tensorboard_compat_proto_summary_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Summary_Value_SimpleValue)(nil),
		(*Summary_Value_ObsoleteOldStyleHistogram)(nil),
		(*Summary_Value_Image)(nil),
		(*Summary_Value_Histo)(nil),
		(*Summary_Value_Audio)(nil),
		(*Summary_Value_Tensor)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorboard_compat_proto_summary_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorboard_compat_proto_summary_proto_goTypes,
		DependencyIndexes: file_tensorboard_compat_proto_summary_proto_depIdxs,
		EnumInfos:         file_tensorboard_compat_proto_summary_proto_enumTypes,
		MessageInfos:      file_tensorboard_compat_proto_summary_proto_msgTypes,
	}.Build()
	File_tensorboard_compat_proto_summary_proto = out.File
	file_tensorboard_compat_proto_summary_proto_rawDesc = nil
	file_tensorboard_compat_proto_summary_proto_goTypes = nil
	file_tensorboard_compat_proto_summary_proto_depIdxs = nil
}
