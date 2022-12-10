package event

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func eventDataToItem(event *Event) map[string]types.AttributeValue {
	attributeMap := make(map[string]types.AttributeValue)

	attributeMap["ID"] = getAttributeValueMemberType(reflect.ValueOf(event.ID))
	attributeMap["Name"] = getAttributeValueMemberType(reflect.ValueOf(event.Name))
	attributeMap["CreatedAt"] = getAttributeValueMemberType(reflect.ValueOf(event.CreatedAt))
	attributeMap["Data"] = getAttributeValueMemberType(reflect.ValueOf(event.Data()))

	return attributeMap
}

func getAttributeValueMemberType(val reflect.Value) types.AttributeValue {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", val.Interface())}
	case reflect.Bool:
		return &types.AttributeValueMemberBOOL{Value: val.Bool()}
	//case reflect.Map:
	case reflect.Interface:
		return &types.AttributeValueMemberB{Value: val.Bytes()}
	case reflect.Struct:
		bytes, err := json.Marshal(val.Interface())
		if err != nil {
			panic("fail marshal struct type in event")
		}

		return &types.AttributeValueMemberB{Value: bytes}
	case reflect.Slice, reflect.Array:
		return &types.AttributeValueMemberSS{Value: val.Interface().([]string)}
	default:
		return &types.AttributeValueMemberS{Value: val.String()}
	}
}
