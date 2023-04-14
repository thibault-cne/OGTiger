package ttype

type TypeID int

const (
	Int TypeID = iota
	String
	Record
	AnyRecord
	Array
	Type
	Function
	NoReturn
)

type RecordField struct {
	Name string
	Type *TigerType
}

type FunctionParameter struct {
	Name string
	Type *TigerType
}

type TigerType struct {
	ID TypeID

	// For records
	Fields []*RecordField

	// For arrays
	ElementType *TigerType

	// For functions
	Parameters []*FunctionParameter
	ReturnType *TigerType
}

func (t *TigerType) Equals(other *TigerType) bool {
	if t.ID != other.ID {
		return false
	}

	switch t.ID {
	case Record:
		// If either type is AnyRecord, they are equal
		if t.ID == AnyRecord || other.ID == AnyRecord {
			return true
		}

		if other.ID != Record {
			return false
		}

		// Check that the number of fields is the same
		if len(t.Fields) != len(other.Fields) {
			return false
		}

		// It doesn't matter which order the fields are in
		for i := range t.Fields {
			found := false
			for j := range other.Fields {
				if t.Fields[i].Name == other.Fields[j].Name && !t.Fields[i].Type.Equals(other.Fields[j].Type) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	case Array:
		return t.ElementType.Equals(other.ElementType)
	}

	return true
}

func (t *TigerType) SizeInStack() int {
	switch t.ID {
	case Int:
		return 4
	case String:
		return 4
	case Record:
		return 4
	case AnyRecord:
		return 4
	case Array:
		return 4
	case Type:
		return 0
	case Function:
		return 0
	}
	return 0
}

func NewTigerType(id TypeID) *TigerType {
	return &TigerType{
		ID: id,
	}
}

func NewRecordType(fields []*RecordField) *TigerType {
	return &TigerType{
		ID:     Record,
		Fields: fields,
	}
}

func NewArrayType(elementType *TigerType) *TigerType {
	return &TigerType{
		ID:          Array,
		ElementType: elementType,
	}
}

func NewFunctionType(returnType *TigerType, parameters []*FunctionParameter) *TigerType {
	return &TigerType{
		ID:         Function,
		Parameters: parameters,
		ReturnType: returnType,
	}
}
