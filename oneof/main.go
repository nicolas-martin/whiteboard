package main

import (
	"fmt"
	"log"
	"os"
	types "whiteboard/oneof/gen/go"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

type Test struct{}

func main() {
	s := LoadFixture("other.json")

	protoOptions := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	var resp types.ExtrinsicsResponse
	kind := resp.Extrinsics.Events[0].Data[0].GetKind()
	switch kind.(type) {
	case *structpb.Value_BoolValue:
		fmt.Println("bool")
	case *structpb.Value_StringValue:
		fmt.Println("string")
	case *structpb.Value_NumberValue:
		fmt.Println("number")
	case *structpb.Value_StructValue:
		fmt.Println("struct")
	case *structpb.Value_ListValue:
		fmt.Println("list")
	default:
		fmt.Println("unknown")
	}

	// json.Unmarshal([]byte(s), resp)
	err := protoOptions.Unmarshal([]byte(s), &resp)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := protojson.Marshal(&resp)
	fmt.Println(string(b))

}

func LoadFixture(fileName string) string {
	content, err := os.ReadFile(fmt.Sprintf("test_data/%s", fileName))
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
