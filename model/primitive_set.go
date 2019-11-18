package model

type PrimitiveFunctor interface {
	DrawArrays(mode int, first int, count uint64)

	DrawElements8(mode int, count uint64, indices []uint8)

	DrawElements16(mode int, count uint64, indices []uint16)

	DrawElements32(mode int, count uint64, indices []uint32)
}

type PrimitiveIndexFunctor interface {
	DrawArrays(mode int, first int, count uint64)

	DrawElements8(mode int, count uint64, indices []uint8)

	DrawElements16(mode int, count uint64, indices []uint16)

	DrawElements32(mode int, count uint64, indices []uint32)
}

const (
	POINTS                        = GLPOINTS
	LINES                         = GLLINES
	LINESTRIP                     = GLLINESTRIP
	LINELOOP                      = GLLINELOOP
	TRIANGLES                     = GLTRIANGLES
	TRIANGLESTRIP                 = GLTRIANGLESTRIP
	TRIANGLEFAN                   = GLTRIANGLEFAN
	QUADS                         = GLQUADS
	QUADSTRIP                     = GLQUADSTRIP
	POLYGON                       = GLPOLYGON
	LINESADJACENCY                = GLLINESADJACENCY
	LINESTRIPADJACENCY            = GLLINESTRIPADJACENCY
	TRIANGLESADJACENCY            = GLTRIANGLESADJACENCY
	TRIANGLESTRIPADJACENCY        = GLTRIANGLESTRIPADJACENCY
	PATCHES                       = GLPATCHES
	PRIMITIVESETT          string = "osg::PrimitiveSet"
)

type PrimitiveSet struct {
	BufferData
	PrimitiveType int
	NumInstances  int
	Mode          uint
}

func NewPrimitiveSet() PrimitiveSet {
	bf := NewBufferData()
	bf.Type = PRIMITIVESETT
	return PrimitiveSet{BufferData: bf, NumInstances: 0, Mode: 0}
}

const (
	DRAWARRAYT          string = "osg::DrawArrays"
	DRAWARRAYLENGHTT    string = "osg::DrawArrayLengths"
	DRWAELEMENTSUBYTET  string = "osg::DrawElementsUByte"
	DRWAELEMENTSUSHORTT string = "osg::DrawElementsUShort"
	DRWAELEMENTSUINTT   string = "osg::DrawElementsUInt"
)

type DrawArrays struct {
	PrimitiveSet
	First int
	Count uint64
}

func (d *DrawArrays) Accept(functor interface{}) {
	switch f := functor.(type) {
	case PrimitiveFunctor:
	case PrimitiveIndexFunctor:
		f.DrawArrays(int(d.Mode), d.First, d.Count)
	}
}

func NewDrawArrays() DrawArrays {
	p := NewPrimitiveSet()
	p.Type = DRAWARRAYT
	return DrawArrays{PrimitiveSet: p, First: 0, Count: 0}
}

type DrawArrayLengths struct {
	PrimitiveSet
	Data  []uint64
	First int
}

func NewDrawArrayLengths() DrawArrayLengths {
	p := NewPrimitiveSet()
	p.Type = DRAWARRAYLENGHTT
	return DrawArrayLengths{PrimitiveSet: p, First: 0}
}

func (dal *DrawArrayLengths) Accept(functor interface{}) {
	switch f := functor.(type) {
	case PrimitiveFunctor:
	case PrimitiveIndexFunctor:
		for _, v := range dal.Data {
			f.DrawArrays(int(dal.Mode), dal.First, v)
		}
	}
}

func (dal *DrawArrayLengths) GetNumPrimitives() int {
	l := len(dal.Data)
	switch dal.Mode {
	case POINTS:
		return l
	case LINES:
		return l / 2
	case TRIANGLES:
		return l / 3
	case QUADS:
		return l / 4
	case LINESTRIP:
	case LINELOOP:
	case TRIANGLESTRIP:
	case TRIANGLEFAN:
	case QUADSTRIP:
	case PATCHES:
	case POLYGON:
		return l
	}
	return 0
}

type DrawElementsUByte struct {
	PrimitiveSet
	Data []uint8
}

func (dw *DrawElementsUByte) Size() uint64 {
	return uint64(len(dw.Data))
}

func (dw *DrawElementsUByte) ResizeElements(size uint64) {
	dw.Data = make([]uint8, size, size)
}

func (dw *DrawElementsUByte) AddElement(e uint8) {
	dw.Data = append(dw.Data, e)
}

func (dw *DrawElementsUByte) Accept(functor interface{}) {
	l := len(dw.Data)
	if l == 0 {
		return
	}
	switch f := functor.(type) {
	case PrimitiveFunctor:
	case PrimitiveIndexFunctor:
		f.DrawElements8(int(dw.Mode), uint64(l), dw.Data)
	}
}

func NewDrawElementsUByte() DrawElementsUByte {
	p := NewPrimitiveSet()
	p.Type = DRWAELEMENTSUBYTET
	return DrawElementsUByte{PrimitiveSet: p}
}

type DrawElementsUShort struct {
	PrimitiveSet
	Data []uint16
}

func (dw *DrawElementsUShort) Size() uint64 {
	return uint64(len(dw.Data))
}

func (dw *DrawElementsUShort) ResizeElements(size uint64) {
	dw.Data = make([]uint16, size, size)
}

func (dw *DrawElementsUShort) AddElement(e uint16) {
	dw.Data = append(dw.Data, e)
}

func (dw *DrawElementsUShort) Accept(functor interface{}) {
	l := len(dw.Data)
	if l == 0 {
		return
	}
	switch f := functor.(type) {
	case PrimitiveFunctor:
	case PrimitiveIndexFunctor:
		f.DrawElements16(int(dw.Mode), uint64(l), dw.Data)
	}
}

func NewDrawElementsUShort() DrawElementsUShort {
	p := NewPrimitiveSet()
	p.Type = DRWAELEMENTSUSHORTT
	return DrawElementsUShort{PrimitiveSet: p}
}

type DrawElementsUInt struct {
	PrimitiveSet
	Data []uint32
}

func (dw *DrawElementsUInt) Size() uint64 {
	return uint64(len(dw.Data))
}

func (dw *DrawElementsUInt) ResizeElements(size uint64) {
	dw.Data = make([]uint32, size, size)
}

func (dw *DrawElementsUInt) AddElement(e uint32) {
	dw.Data = append(dw.Data, e)
}

func (dw *DrawElementsUInt) Accept(functor interface{}) {
	l := len(dw.Data)
	if l == 0 {
		return
	}
	switch f := functor.(type) {
	case PrimitiveFunctor:
	case PrimitiveIndexFunctor:
		f.DrawElements32(int(dw.Mode), uint64(l), dw.Data)
	}
}

func NewDrawElementsUInt() DrawElementsUInt {
	p := NewPrimitiveSet()
	p.Type = DRWAELEMENTSUINTT
	return DrawElementsUInt{PrimitiveSet: p}
}
