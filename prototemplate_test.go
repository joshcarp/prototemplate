package prototemplate

import (
	"google.golang.org/protobuf/encoding/protojson"
	"prototemplate/gen/test/simple"

	"testing"
)

func TestMakeTemplate(t *testing.T) {
	msg := &simple.Message1{}
	newmsg := MakeTemplate(msg.ProtoReflect().Descriptor())
	jsonmsg, err := protojson.Marshal(newmsg)
	if err != nil {
		t.Fail()
	}
	t.Logf(string(jsonmsg))
	//dm := dynamicpb.NewMessage(md)

}
