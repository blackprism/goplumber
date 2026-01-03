package pipewire

type PODType uint32

const (
	PODTypeNone PODType = 1 + iota
	PODTypeBool
	PODTypeID
	PODTypeInt
	PODTypeLong
	PODTypeFloat
	PODTypeDouble
	PODTypeString
	PODTypeBytes
	PODTypeRectangle
	PODTypeFraction
	PODTypeBitmap
	PODTypeArray
	PODTypeStruct
	PODTypeObject
	PODTypeSequence
	PODTypePointer
	PODTypeFD
	PODTypeChoice
	PODTypePOD
)
