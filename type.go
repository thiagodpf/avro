package avro

import (
	"fmt"

	"github.com/rogpeppe/gogen-avro/v7/parser"
	"github.com/rogpeppe/gogen-avro/v7/resolver"
	"github.com/rogpeppe/gogen-avro/v7/schema"
)

// Type represents an Avro schema type.
type Type struct {
	schema   string
	avroType schema.AvroType
}

// ParseType parses an Avro schema in the format defined by the Avro
// specification at https://avro.apache.org/docs/current/spec.html.
func ParseType(s string) (*Type, error) {
	ns := parser.NewNamespace(false)
	avroType, err := ns.TypeForSchema([]byte(s))
	if err != nil {
		return nil, fmt.Errorf("invalid schema %q: %v", s, err)
	}
	for _, def := range ns.Roots {
		if err := resolver.ResolveDefinition(def, ns.Definitions); err != nil {
			return nil, fmt.Errorf("cannot resolve references in schema\n%s\n: %v", s, err)
		}
	}
	return &Type{
		schema:   s,
		avroType: avroType,
	}, nil
}

func (t *Type) String() string {
	return t.schema
}

// name returns the fully qualified Avro name for the type,
// or the empty string if it's not a definition.
func (t *Type) name() string {
	ref, ok := t.avroType.(*schema.Reference)
	if !ok {
		return ""
	}
	return ref.TypeName.String()
}
