// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: app.proto

package pb

import (
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

type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Props   []*App_Props   `protobuf:"bytes,1,rep,name=props,proto3" json:"props,omitempty"`
	Screens []*App_Screens `protobuf:"bytes,2,rep,name=screens,proto3" json:"screens,omitempty"`
}

func (x *App) Reset() {
	*x = App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{0}
}

func (x *App) GetProps() []*App_Props {
	if x != nil {
		return x.Props
	}
	return nil
}

func (x *App) GetScreens() []*App_Screens {
	if x != nil {
		return x.Screens
	}
	return nil
}

type App_Props struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *App_Props) Reset() {
	*x = App_Props{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App_Props) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App_Props) ProtoMessage() {}

func (x *App_Props) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App_Props.ProtoReflect.Descriptor instead.
func (*App_Props) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{0, 0}
}

func (x *App_Props) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *App_Props) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *App_Props) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *App_Props) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type App_Widgets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Props []*App_Props `protobuf:"bytes,2,rep,name=props,proto3" json:"props,omitempty"`
	Id    uint32       `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *App_Widgets) Reset() {
	*x = App_Widgets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App_Widgets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App_Widgets) ProtoMessage() {}

func (x *App_Widgets) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App_Widgets.ProtoReflect.Descriptor instead.
func (*App_Widgets) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{0, 1}
}

func (x *App_Widgets) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *App_Widgets) GetProps() []*App_Props {
	if x != nil {
		return x.Props
	}
	return nil
}

func (x *App_Widgets) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type App_Screens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type    string         `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Props   []*App_Props   `protobuf:"bytes,3,rep,name=props,proto3" json:"props,omitempty"`
	Widgets []*App_Widgets `protobuf:"bytes,4,rep,name=widgets,proto3" json:"widgets,omitempty"`
}

func (x *App_Screens) Reset() {
	*x = App_Screens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App_Screens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App_Screens) ProtoMessage() {}

func (x *App_Screens) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App_Screens.ProtoReflect.Descriptor instead.
func (*App_Screens) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{0, 2}
}

func (x *App_Screens) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *App_Screens) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *App_Screens) GetProps() []*App_Props {
	if x != nil {
		return x.Props
	}
	return nil
}

func (x *App_Screens) GetWidgets() []*App_Widgets {
	if x != nil {
		return x.Widgets
	}
	return nil
}

var File_app_proto protoreflect.FileDescriptor

var file_app_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x61, 0x73,
	0x6b, 0x22, 0x8a, 0x03, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x72, 0x6f,
	0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x41, 0x70, 0x70, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73,
	0x12, 0x2b, 0x0a, 0x07, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x53, 0x63, 0x72,
	0x65, 0x65, 0x6e, 0x73, 0x52, 0x07, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x1a, 0x55, 0x0a,
	0x05, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x54, 0x0a, 0x07, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x73, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x81, 0x01, 0x0a, 0x07, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x72,
	0x6f, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x41, 0x70, 0x70, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x70,
	0x73, 0x12, 0x2b, 0x0a, 0x07, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x57, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x73, 0x52, 0x07, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x42, 0x05,
	0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_proto_rawDescOnce sync.Once
	file_app_proto_rawDescData = file_app_proto_rawDesc
)

func file_app_proto_rawDescGZIP() []byte {
	file_app_proto_rawDescOnce.Do(func() {
		file_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_proto_rawDescData)
	})
	return file_app_proto_rawDescData
}

var file_app_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_app_proto_goTypes = []interface{}{
	(*App)(nil),         // 0: task.App
	(*App_Props)(nil),   // 1: task.App.Props
	(*App_Widgets)(nil), // 2: task.App.Widgets
	(*App_Screens)(nil), // 3: task.App.Screens
}
var file_app_proto_depIdxs = []int32{
	1, // 0: task.App.props:type_name -> task.App.Props
	3, // 1: task.App.screens:type_name -> task.App.Screens
	1, // 2: task.App.Widgets.props:type_name -> task.App.Props
	1, // 3: task.App.Screens.props:type_name -> task.App.Props
	2, // 4: task.App.Screens.widgets:type_name -> task.App.Widgets
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_app_proto_init() }
func file_app_proto_init() {
	if File_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App); i {
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
		file_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App_Props); i {
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
		file_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App_Widgets); i {
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
		file_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App_Screens); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_proto_goTypes,
		DependencyIndexes: file_app_proto_depIdxs,
		MessageInfos:      file_app_proto_msgTypes,
	}.Build()
	File_app_proto = out.File
	file_app_proto_rawDesc = nil
	file_app_proto_goTypes = nil
	file_app_proto_depIdxs = nil
}