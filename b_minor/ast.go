package main

type Decl struct{
	name string
	typ *Type

}

const (
	VOID int = iota
	INT
	BOOL
	CHAR
	STRING
	ARRAY
	FUNCTION
)

type Type struct {
	kind int
	subtype *Type
	params *param_list
}

