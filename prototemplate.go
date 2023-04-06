package prototemplate

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

func MakeTemplate(md protoreflect.MessageDescriptor) proto.Message {
	dm := dynamicpb.NewMessage(md)
	if isWellknownType(string(md.FullName())) {
		return nil
	}
	for i := 0; i < dm.Descriptor().Fields().Len(); i++ {
		fd := dm.Descriptor().Fields().Get(i)

		var val protoreflect.Value
		switch fd.Kind() {
		case protoreflect.BoolKind:
			val = protoreflect.ValueOfBool(true)
		case protoreflect.EnumKind:
			val = protoreflect.ValueOfEnum(1)
		case protoreflect.Int32Kind:
			val = protoreflect.ValueOfInt32(1)
		case protoreflect.Sint32Kind:
			val = protoreflect.ValueOfInt32(1)
		case protoreflect.Uint32Kind:
			val = protoreflect.ValueOfInt32(1)
		case protoreflect.Int64Kind:
			val = protoreflect.ValueOfInt64(1)
		case protoreflect.Sint64Kind:
			val = protoreflect.ValueOfInt64(1)
		case protoreflect.Uint64Kind:
			val = protoreflect.ValueOfInt64(1)
		case protoreflect.Sfixed32Kind:
			val = protoreflect.ValueOfInt32(1)
		case protoreflect.Fixed32Kind:
			val = protoreflect.ValueOfFloat32(1.1)
		case protoreflect.FloatKind:
			val = protoreflect.ValueOfFloat32(1.1)
		case protoreflect.Sfixed64Kind:
			val = protoreflect.ValueOfInt64(1)
		case protoreflect.Fixed64Kind:
			val = protoreflect.ValueOfFloat64(1.1)
		case protoreflect.DoubleKind:
			val = protoreflect.ValueOfFloat64(1.1)
		case protoreflect.StringKind:
			val = protoreflect.ValueOfString("string")
		case protoreflect.BytesKind:
			val = protoreflect.ValueOfBytes([]byte(fd.JSONName()))
		case protoreflect.MessageKind:
			pr := MakeTemplate(fd.Message())
			if pr != nil {
				val = protoreflect.ValueOfMessage(pr.ProtoReflect())
			} else {
				return dm
			}
		default:
			return dm
		}
		dm.Set(fd, val)
	}
	return dm
}

// func isWellknownType returns true if type name is a well known protobuf type
func isWellknownType(name string) bool {
	switch name {
	case "google.protobuf.Any",
		"google.protobuf.Value",
		"google.protobuf.Api",
		"google.protobuf.Duration",
		"google.protobuf.Empty",
		"google.protobuf.FieldMask",
		"google.protobuf.Timestamp",
		"google.protobuf.SourceContext",
		"google.protobuf.Struct",
		"google.protobuf.Type":
		return true
	default:
		return false
	}
}
