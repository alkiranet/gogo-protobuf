package main

import (
	"strings"

	"github.com/alkiranet/gogo-protobuf/gogoproto"
	"github.com/alkiranet/gogo-protobuf/proto"
	"github.com/alkiranet/gogo-protobuf/protoc-gen-gogo/descriptor"
	"github.com/alkiranet/gogo-protobuf/protoc-gen-gogo/generator"
	"github.com/alkiranet/gogo-protobuf/vanity"
)

func lowerInitial(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}

// AcronymName preprocess the field, and set the [(gogoproto.customname) = "..."]
// if necessary, in order to avoid setting `gogoproto.customname` manually.
// The automatically assigned name should conform to Golang convention.
func AcronymName(mapping map[string]string) func(file *descriptor.FileDescriptorProto) {
	return func(file *descriptor.FileDescriptorProto) {

		f := func(field *descriptor.FieldDescriptorProto) {
			// Skip if [(gogoproto.customname) = "..."] has already been set.
			if gogoproto.IsCustomName(field) {
				return
			}
			// Skip if embedded
			if gogoproto.IsEmbed(field) {
				return
			}
			if field.OneofIndex != nil {
				return
			}

			parts := strings.Split(*field.Name, "_")
			for index, part := range parts {
				for key, value := range mapping {
					if strings.EqualFold(key, part) {
						parts[index] = lowerInitial(value)
					}
				}
			}
			convertedFieldName := strings.Join(parts, "_")

			if convertedFieldName == *field.Name {
				return
			}

			if field.Options == nil {
				field.Options = &descriptor.FieldOptions{}
			}

			fieldName := generator.CamelCase(convertedFieldName)
			if err := proto.SetExtension(field.Options, gogoproto.E_Customname, &fieldName); err != nil {
				panic(err)
			}
		}

		// Iterate through all fields in file
		vanity.ForEachFieldExcludingExtensions(file.MessageType, f)
	}
}
