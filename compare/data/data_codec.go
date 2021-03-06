//go:build go1.6
// +build go1.6

// Code generated by codecgen - DO NOT EDIT.

package data

import (
	"errors"
	codec1978 "github.com/ugorji/go/codec"
	"runtime"
	"strconv"
)

const (
	// ----- content types ----
	codecSelferCcUTF82970 = 1
	codecSelferCcRAW2970  = 255
	// ----- value types used ----
	codecSelferValueTypeArray2970     = 10
	codecSelferValueTypeMap2970       = 9
	codecSelferValueTypeString2970    = 6
	codecSelferValueTypeInt2970       = 2
	codecSelferValueTypeUint2970      = 3
	codecSelferValueTypeFloat2970     = 4
	codecSelferValueTypeNil2970       = 1
	codecSelferBitsize2970            = uint8(32 << (^uint(0) >> 63))
	codecSelferDecContainerLenNil2970 = -2147483648
)

var (
	errCodecSelferOnlyMapOrArrayEncodeToStruct2970 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer2970 struct{}

func codecSelfer2970False() bool { return false }
func codecSelfer2970True() bool  { return true }

func init() {
	if codec1978.GenVersion != 25 {
		_, file, _, _ := runtime.Caller(0)
		ver := strconv.FormatInt(int64(codec1978.GenVersion), 10)
		panic(errors.New("codecgen version mismatch: current: 25, need " + ver + ". Re-generate file: " + file))
	}
}

func (Simple) codecSelferViaCodecgen() {}
func (x *Simple) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Encoder(e)
	_, _, _ = h, z, r
	if z.EncBasicHandle().CheckCircularRef {
		z.EncEncode(x)
		return
	}
	if x == nil {
		r.EncodeNil()
	} else {
		yy2arr2 := z.EncBasicHandle().StructToArray
		_ = yy2arr2
		const yyr2 bool = false // struct tag has 'toArray'
		if yyr2 || yy2arr2 {
			z.EncWriteArrayStart(6)
			z.EncWriteArrayElem()
			r.EncodeString(string(x.Title))
			z.EncWriteArrayElem()
			r.EncodeString(string(x.Text))
			z.EncWriteArrayElem()
			r.EncodeInt(int64(x.Likes))
			z.EncWriteArrayElem()
			r.EncodeInt(int64(x.Subscribers))
			z.EncWriteArrayElem()
			r.EncodeBool(bool(x.IsPublic))
			z.EncWriteArrayElem()
			r.EncodeBool(bool(x.IsCool))
			z.EncWriteArrayEnd()
		} else {
			z.EncWriteMapStart(6)
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"title\"")
			} else {
				r.EncodeString(`title`)
			}
			z.EncWriteMapElemValue()
			r.EncodeString(string(x.Title))
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"text\"")
			} else {
				r.EncodeString(`text`)
			}
			z.EncWriteMapElemValue()
			r.EncodeString(string(x.Text))
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"likes\"")
			} else {
				r.EncodeString(`likes`)
			}
			z.EncWriteMapElemValue()
			r.EncodeInt(int64(x.Likes))
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"subscribers\"")
			} else {
				r.EncodeString(`subscribers`)
			}
			z.EncWriteMapElemValue()
			r.EncodeInt(int64(x.Subscribers))
			z.EncWriteMapElemKey()
			r.EncodeString(`is_public`)
			z.EncWriteMapElemValue()
			r.EncodeBool(bool(x.IsPublic))
			z.EncWriteMapElemKey()
			r.EncodeString(`is_cool`)
			z.EncWriteMapElemValue()
			r.EncodeBool(bool(x.IsCool))
			z.EncWriteMapEnd()
		}
	}
}

func (x *Simple) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	yyct2 := r.ContainerType()
	if yyct2 == codecSelferValueTypeNil2970 {
		*(x) = Simple{}
	} else if yyct2 == codecSelferValueTypeMap2970 {
		yyl2 := z.DecReadMapStart()
		if yyl2 == 0 {
		} else {
			x.codecDecodeSelfFromMap(yyl2, d)
		}
		z.DecReadMapEnd()
	} else if yyct2 == codecSelferValueTypeArray2970 {
		yyl2 := z.DecReadArrayStart()
		if yyl2 != 0 {
			x.codecDecodeSelfFromArray(yyl2, d)
		}
		z.DecReadArrayEnd()
	} else {
		panic(errCodecSelferOnlyMapOrArrayEncodeToStruct2970)
	}
}

func (x *Simple) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if z.DecCheckBreak() {
				break
			}
		}
		z.DecReadMapElemKey()
		yys3 := r.DecodeStringAsBytes()
		z.DecReadMapElemValue()
		switch string(yys3) {
		case "title":
			x.Title = (string)(z.DecStringZC(r.DecodeStringAsBytes()))
		case "text":
			x.Text = (string)(z.DecStringZC(r.DecodeStringAsBytes()))
		case "likes":
			x.Likes = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize2970))
		case "subscribers":
			x.Subscribers = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize2970))
		case "is_public":
			x.IsPublic = (bool)(r.DecodeBool())
		case "is_cool":
			x.IsCool = (bool)(r.DecodeBool())
		default:
			z.DecStructFieldNotFound(-1, string(yys3))
		} // end switch yys3
	} // end for yyj3
}

func (x *Simple) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	var yyj10 int
	var yyb10 bool
	var yyhl10 bool = l >= 0
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = z.DecCheckBreak()
	}
	if yyb10 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Title = (string)(z.DecStringZC(r.DecodeStringAsBytes()))
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = z.DecCheckBreak()
	}
	if yyb10 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Text = (string)(z.DecStringZC(r.DecodeStringAsBytes()))
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = z.DecCheckBreak()
	}
	if yyb10 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Likes = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize2970))
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = z.DecCheckBreak()
	}
	if yyb10 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Subscribers = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize2970))
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = z.DecCheckBreak()
	}
	if yyb10 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.IsPublic = (bool)(r.DecodeBool())
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = z.DecCheckBreak()
	}
	if yyb10 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.IsCool = (bool)(r.DecodeBool())
	for {
		yyj10++
		if yyhl10 {
			yyb10 = yyj10 > l
		} else {
			yyb10 = z.DecCheckBreak()
		}
		if yyb10 {
			break
		}
		z.DecReadArrayElem()
		z.DecStructFieldNotFound(yyj10-1, "")
	}
}

func (x *Simple) IsCodecEmpty() bool {
	return !(x.Title != "" || x.Text != "" || x.Likes != 0 || x.Subscribers != 0 || bool(x.IsPublic) || bool(x.IsCool) || false)
}

func (Simples) codecSelferViaCodecgen() {}
func (x Simples) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Encoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		h.encSimples((Simples)(x), e)
	} // end block: if x slice == nil
}

func (x *Simples) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	h.decSimples((*Simples)(x), d)
}

func (Medium) codecSelferViaCodecgen() {}
func (x *Medium) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Encoder(e)
	_, _, _ = h, z, r
	if z.EncBasicHandle().CheckCircularRef {
		z.EncEncode(x)
		return
	}
	if x == nil {
		r.EncodeNil()
	} else {
		yy2arr2 := z.EncBasicHandle().StructToArray
		_ = yy2arr2
		const yyr2 bool = false // struct tag has 'toArray'
		if yyr2 || yy2arr2 {
			z.EncWriteArrayStart(9)
			z.EncWriteArrayElem()
			r.EncodeString(string(x.Title))
			z.EncWriteArrayElem()
			r.EncodeString(string(x.Text))
			z.EncWriteArrayElem()
			if x.Lines == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceStringV(x.Lines, e)
			} // end block: if x.Lines slice == nil
			z.EncWriteArrayElem()
			r.EncodeInt(int64(x.Likes))
			z.EncWriteArrayElem()
			r.EncodeInt(int64(x.Subscribers))
			z.EncWriteArrayElem()
			if x.Ids == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceIntV(x.Ids, e)
			} // end block: if x.Ids slice == nil
			z.EncWriteArrayElem()
			r.EncodeBool(bool(x.IsPublic))
			z.EncWriteArrayElem()
			r.EncodeBool(bool(x.IsCool))
			z.EncWriteArrayElem()
			if x.Options == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceBoolV(x.Options, e)
			} // end block: if x.Options slice == nil
			z.EncWriteArrayEnd()
		} else {
			z.EncWriteMapStart(9)
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"title\"")
			} else {
				r.EncodeString(`title`)
			}
			z.EncWriteMapElemValue()
			r.EncodeString(string(x.Title))
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"text\"")
			} else {
				r.EncodeString(`text`)
			}
			z.EncWriteMapElemValue()
			r.EncodeString(string(x.Text))
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"lines\"")
			} else {
				r.EncodeString(`lines`)
			}
			z.EncWriteMapElemValue()
			if x.Lines == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceStringV(x.Lines, e)
			} // end block: if x.Lines slice == nil
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"likes\"")
			} else {
				r.EncodeString(`likes`)
			}
			z.EncWriteMapElemValue()
			r.EncodeInt(int64(x.Likes))
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"subscribers\"")
			} else {
				r.EncodeString(`subscribers`)
			}
			z.EncWriteMapElemValue()
			r.EncodeInt(int64(x.Subscribers))
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"ids\"")
			} else {
				r.EncodeString(`ids`)
			}
			z.EncWriteMapElemValue()
			if x.Ids == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceIntV(x.Ids, e)
			} // end block: if x.Ids slice == nil
			z.EncWriteMapElemKey()
			r.EncodeString(`is_public`)
			z.EncWriteMapElemValue()
			r.EncodeBool(bool(x.IsPublic))
			z.EncWriteMapElemKey()
			r.EncodeString(`is_cool`)
			z.EncWriteMapElemValue()
			r.EncodeBool(bool(x.IsCool))
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"options\"")
			} else {
				r.EncodeString(`options`)
			}
			z.EncWriteMapElemValue()
			if x.Options == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceBoolV(x.Options, e)
			} // end block: if x.Options slice == nil
			z.EncWriteMapEnd()
		}
	}
}

func (x *Medium) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	yyct2 := r.ContainerType()
	if yyct2 == codecSelferValueTypeNil2970 {
		*(x) = Medium{}
	} else if yyct2 == codecSelferValueTypeMap2970 {
		yyl2 := z.DecReadMapStart()
		if yyl2 == 0 {
		} else {
			x.codecDecodeSelfFromMap(yyl2, d)
		}
		z.DecReadMapEnd()
	} else if yyct2 == codecSelferValueTypeArray2970 {
		yyl2 := z.DecReadArrayStart()
		if yyl2 != 0 {
			x.codecDecodeSelfFromArray(yyl2, d)
		}
		z.DecReadArrayEnd()
	} else {
		panic(errCodecSelferOnlyMapOrArrayEncodeToStruct2970)
	}
}

func (x *Medium) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if z.DecCheckBreak() {
				break
			}
		}
		z.DecReadMapElemKey()
		yys3 := r.DecodeStringAsBytes()
		z.DecReadMapElemValue()
		switch string(yys3) {
		case "title":
			x.Title = (string)(z.DecStringZC(r.DecodeStringAsBytes()))
		case "text":
			x.Text = (string)(z.DecStringZC(r.DecodeStringAsBytes()))
		case "lines":
			z.F.DecSliceStringX(&x.Lines, d)
		case "likes":
			x.Likes = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize2970))
		case "subscribers":
			x.Subscribers = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize2970))
		case "ids":
			z.F.DecSliceIntX(&x.Ids, d)
		case "is_public":
			x.IsPublic = (bool)(r.DecodeBool())
		case "is_cool":
			x.IsCool = (bool)(r.DecodeBool())
		case "options":
			z.F.DecSliceBoolX(&x.Options, d)
		default:
			z.DecStructFieldNotFound(-1, string(yys3))
		} // end switch yys3
	} // end for yyj3
}

func (x *Medium) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	var yyj16 int
	var yyb16 bool
	var yyhl16 bool = l >= 0
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = z.DecCheckBreak()
	}
	if yyb16 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Title = (string)(z.DecStringZC(r.DecodeStringAsBytes()))
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = z.DecCheckBreak()
	}
	if yyb16 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Text = (string)(z.DecStringZC(r.DecodeStringAsBytes()))
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = z.DecCheckBreak()
	}
	if yyb16 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	z.F.DecSliceStringX(&x.Lines, d)
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = z.DecCheckBreak()
	}
	if yyb16 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Likes = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize2970))
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = z.DecCheckBreak()
	}
	if yyb16 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Subscribers = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize2970))
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = z.DecCheckBreak()
	}
	if yyb16 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	z.F.DecSliceIntX(&x.Ids, d)
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = z.DecCheckBreak()
	}
	if yyb16 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.IsPublic = (bool)(r.DecodeBool())
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = z.DecCheckBreak()
	}
	if yyb16 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.IsCool = (bool)(r.DecodeBool())
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = z.DecCheckBreak()
	}
	if yyb16 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	z.F.DecSliceBoolX(&x.Options, d)
	for {
		yyj16++
		if yyhl16 {
			yyb16 = yyj16 > l
		} else {
			yyb16 = z.DecCheckBreak()
		}
		if yyb16 {
			break
		}
		z.DecReadArrayElem()
		z.DecStructFieldNotFound(yyj16-1, "")
	}
}

func (x *Medium) IsCodecEmpty() bool {
	return !(x.Title != "" || x.Text != "" || len(x.Lines) != 0 || x.Likes != 0 || x.Subscribers != 0 || len(x.Ids) != 0 || bool(x.IsPublic) || bool(x.IsCool) || len(x.Options) != 0 || false)
}

func (Mediums) codecSelferViaCodecgen() {}
func (x Mediums) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Encoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		h.encMediums((Mediums)(x), e)
	} // end block: if x slice == nil
}

func (x *Mediums) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	h.decMediums((*Mediums)(x), d)
}

func (Heavy) codecSelferViaCodecgen() {}
func (x *Heavy) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Encoder(e)
	_, _, _ = h, z, r
	if z.EncBasicHandle().CheckCircularRef {
		z.EncEncode(x)
		return
	}
	if x == nil {
		r.EncodeNil()
	} else {
		yy2arr2 := z.EncBasicHandle().StructToArray
		_ = yy2arr2
		const yyr2 bool = false // struct tag has 'toArray'
		if yyr2 || yy2arr2 {
			z.EncWriteArrayStart(10)
			z.EncWriteArrayElem()
			r.EncodeString(string(x.Title))
			z.EncWriteArrayElem()
			r.EncodeString(string(x.Text))
			z.EncWriteArrayElem()
			if x.Lines == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceStringV(x.Lines, e)
			} // end block: if x.Lines slice == nil
			z.EncWriteArrayElem()
			r.EncodeInt(int64(x.Likes))
			z.EncWriteArrayElem()
			r.EncodeInt(int64(x.Subscribers))
			z.EncWriteArrayElem()
			if x.Ids == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceIntV(x.Ids, e)
			} // end block: if x.Ids slice == nil
			z.EncWriteArrayElem()
			r.EncodeBool(bool(x.IsPublic))
			z.EncWriteArrayElem()
			r.EncodeBool(bool(x.IsCool))
			z.EncWriteArrayElem()
			if x.Options == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceBoolV(x.Options, e)
			} // end block: if x.Options slice == nil
			z.EncWriteArrayElem()
			if x.MediumElements == nil {
				r.EncodeNil()
			} else {
				h.encSliceMedium(([]Medium)(x.MediumElements), e)
			} // end block: if x.MediumElements slice == nil
			z.EncWriteArrayEnd()
		} else {
			z.EncWriteMapStart(10)
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"title\"")
			} else {
				r.EncodeString(`title`)
			}
			z.EncWriteMapElemValue()
			r.EncodeString(string(x.Title))
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"text\"")
			} else {
				r.EncodeString(`text`)
			}
			z.EncWriteMapElemValue()
			r.EncodeString(string(x.Text))
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"lines\"")
			} else {
				r.EncodeString(`lines`)
			}
			z.EncWriteMapElemValue()
			if x.Lines == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceStringV(x.Lines, e)
			} // end block: if x.Lines slice == nil
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"likes\"")
			} else {
				r.EncodeString(`likes`)
			}
			z.EncWriteMapElemValue()
			r.EncodeInt(int64(x.Likes))
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"subscribers\"")
			} else {
				r.EncodeString(`subscribers`)
			}
			z.EncWriteMapElemValue()
			r.EncodeInt(int64(x.Subscribers))
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"ids\"")
			} else {
				r.EncodeString(`ids`)
			}
			z.EncWriteMapElemValue()
			if x.Ids == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceIntV(x.Ids, e)
			} // end block: if x.Ids slice == nil
			z.EncWriteMapElemKey()
			r.EncodeString(`is_public`)
			z.EncWriteMapElemValue()
			r.EncodeBool(bool(x.IsPublic))
			z.EncWriteMapElemKey()
			r.EncodeString(`is_cool`)
			z.EncWriteMapElemValue()
			r.EncodeBool(bool(x.IsCool))
			z.EncWriteMapElemKey()
			if z.IsJSONHandle() {
				z.WriteStr("\"options\"")
			} else {
				r.EncodeString(`options`)
			}
			z.EncWriteMapElemValue()
			if x.Options == nil {
				r.EncodeNil()
			} else {
				z.F.EncSliceBoolV(x.Options, e)
			} // end block: if x.Options slice == nil
			z.EncWriteMapElemKey()
			r.EncodeString(`medium_elements`)
			z.EncWriteMapElemValue()
			if x.MediumElements == nil {
				r.EncodeNil()
			} else {
				h.encSliceMedium(([]Medium)(x.MediumElements), e)
			} // end block: if x.MediumElements slice == nil
			z.EncWriteMapEnd()
		}
	}
}

func (x *Heavy) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	yyct2 := r.ContainerType()
	if yyct2 == codecSelferValueTypeNil2970 {
		*(x) = Heavy{}
	} else if yyct2 == codecSelferValueTypeMap2970 {
		yyl2 := z.DecReadMapStart()
		if yyl2 == 0 {
		} else {
			x.codecDecodeSelfFromMap(yyl2, d)
		}
		z.DecReadMapEnd()
	} else if yyct2 == codecSelferValueTypeArray2970 {
		yyl2 := z.DecReadArrayStart()
		if yyl2 != 0 {
			x.codecDecodeSelfFromArray(yyl2, d)
		}
		z.DecReadArrayEnd()
	} else {
		panic(errCodecSelferOnlyMapOrArrayEncodeToStruct2970)
	}
}

func (x *Heavy) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if z.DecCheckBreak() {
				break
			}
		}
		z.DecReadMapElemKey()
		yys3 := r.DecodeStringAsBytes()
		z.DecReadMapElemValue()
		switch string(yys3) {
		case "title":
			x.Title = (string)(z.DecStringZC(r.DecodeStringAsBytes()))
		case "text":
			x.Text = (string)(z.DecStringZC(r.DecodeStringAsBytes()))
		case "lines":
			z.F.DecSliceStringX(&x.Lines, d)
		case "likes":
			x.Likes = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize2970))
		case "subscribers":
			x.Subscribers = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize2970))
		case "ids":
			z.F.DecSliceIntX(&x.Ids, d)
		case "is_public":
			x.IsPublic = (bool)(r.DecodeBool())
		case "is_cool":
			x.IsCool = (bool)(r.DecodeBool())
		case "options":
			z.F.DecSliceBoolX(&x.Options, d)
		case "medium_elements":
			h.decSliceMedium((*[]Medium)(&x.MediumElements), d)
		default:
			z.DecStructFieldNotFound(-1, string(yys3))
		} // end switch yys3
	} // end for yyj3
}

func (x *Heavy) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	var yyj18 int
	var yyb18 bool
	var yyhl18 bool = l >= 0
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = z.DecCheckBreak()
	}
	if yyb18 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Title = (string)(z.DecStringZC(r.DecodeStringAsBytes()))
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = z.DecCheckBreak()
	}
	if yyb18 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Text = (string)(z.DecStringZC(r.DecodeStringAsBytes()))
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = z.DecCheckBreak()
	}
	if yyb18 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	z.F.DecSliceStringX(&x.Lines, d)
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = z.DecCheckBreak()
	}
	if yyb18 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Likes = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize2970))
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = z.DecCheckBreak()
	}
	if yyb18 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Subscribers = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize2970))
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = z.DecCheckBreak()
	}
	if yyb18 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	z.F.DecSliceIntX(&x.Ids, d)
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = z.DecCheckBreak()
	}
	if yyb18 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.IsPublic = (bool)(r.DecodeBool())
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = z.DecCheckBreak()
	}
	if yyb18 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.IsCool = (bool)(r.DecodeBool())
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = z.DecCheckBreak()
	}
	if yyb18 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	z.F.DecSliceBoolX(&x.Options, d)
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = z.DecCheckBreak()
	}
	if yyb18 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	h.decSliceMedium((*[]Medium)(&x.MediumElements), d)
	for {
		yyj18++
		if yyhl18 {
			yyb18 = yyj18 > l
		} else {
			yyb18 = z.DecCheckBreak()
		}
		if yyb18 {
			break
		}
		z.DecReadArrayElem()
		z.DecStructFieldNotFound(yyj18-1, "")
	}
}

func (x *Heavy) IsCodecEmpty() bool {
	return !(x.Title != "" || x.Text != "" || len(x.Lines) != 0 || x.Likes != 0 || x.Subscribers != 0 || len(x.Ids) != 0 || bool(x.IsPublic) || bool(x.IsCool) || len(x.Options) != 0 || len(x.MediumElements) != 0 || false)
}

func (Heavies) codecSelferViaCodecgen() {}
func (x Heavies) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Encoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		h.encHeavies((Heavies)(x), e)
	} // end block: if x slice == nil
}

func (x *Heavies) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	h.decHeavies((*Heavies)(x), d)
}

func (x codecSelfer2970) encSimples(v Simples, e *codec1978.Encoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Encoder(e)
	_, _, _ = h, z, r
	if v == nil {
		r.EncodeNil()
		return
	}
	z.EncWriteArrayStart(len(v))
	for yyv1 := range v {
		z.EncWriteArrayElem()
		yy2 := &v[yyv1]
		if yyxt3 := z.Extension(yy2); yyxt3 != nil {
			z.EncExtension(yy2, yyxt3)
		} else {
			yy2.CodecEncodeSelf(e)
		}
	}
	z.EncWriteArrayEnd()
}

func (x codecSelfer2970) decSimples(v *Simples, d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyh1.IsNil {
		if yyv1 != nil {
			yyv1 = nil
			yyc1 = true
		}
	} else if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []Simple{}
			yyc1 = true
		} else if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} else {
		yyhl1 := yyl1 > 0
		var yyrl1 int
		_ = yyrl1
		if yyhl1 {
			if yyl1 > cap(yyv1) {
				yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 56)
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} else {
					yyv1 = make([]Simple, yyrl1)
				}
				yyc1 = true
			} else if yyl1 != len(yyv1) {
				yyv1 = yyv1[:yyl1]
				yyc1 = true
			}
		}
		var yyj1 int
		for yyj1 = 0; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || z.DecCheckBreak()); yyj1++ { // bounds-check-elimination
			if yyj1 == 0 && yyv1 == nil {
				if yyhl1 {
					yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 56)
				} else {
					yyrl1 = 8
				}
				yyv1 = make([]Simple, yyrl1)
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)
			var yydb1 bool
			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, Simple{})
				yyc1 = true
			}
			if yydb1 {
				z.DecSwallow()
			} else {
				if yyxt3 := z.Extension(yyv1[yyj1]); yyxt3 != nil {
					z.DecExtension(&yyv1[yyj1], yyxt3)
				} else {
					yyv1[yyj1].CodecDecodeSelf(d)
				}
			}
		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} else if yyj1 == 0 && yyv1 == nil {
			yyv1 = make([]Simple, 0)
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}
}

func (x codecSelfer2970) encMediums(v Mediums, e *codec1978.Encoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Encoder(e)
	_, _, _ = h, z, r
	if v == nil {
		r.EncodeNil()
		return
	}
	z.EncWriteArrayStart(len(v))
	for yyv1 := range v {
		z.EncWriteArrayElem()
		yy2 := &v[yyv1]
		if yyxt3 := z.Extension(yy2); yyxt3 != nil {
			z.EncExtension(yy2, yyxt3)
		} else {
			yy2.CodecEncodeSelf(e)
		}
	}
	z.EncWriteArrayEnd()
}

func (x codecSelfer2970) decMediums(v *Mediums, d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyh1.IsNil {
		if yyv1 != nil {
			yyv1 = nil
			yyc1 = true
		}
	} else if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []Medium{}
			yyc1 = true
		} else if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} else {
		yyhl1 := yyl1 > 0
		var yyrl1 int
		_ = yyrl1
		if yyhl1 {
			if yyl1 > cap(yyv1) {
				yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 128)
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} else {
					yyv1 = make([]Medium, yyrl1)
				}
				yyc1 = true
			} else if yyl1 != len(yyv1) {
				yyv1 = yyv1[:yyl1]
				yyc1 = true
			}
		}
		var yyj1 int
		for yyj1 = 0; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || z.DecCheckBreak()); yyj1++ { // bounds-check-elimination
			if yyj1 == 0 && yyv1 == nil {
				if yyhl1 {
					yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 128)
				} else {
					yyrl1 = 8
				}
				yyv1 = make([]Medium, yyrl1)
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)
			var yydb1 bool
			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, Medium{})
				yyc1 = true
			}
			if yydb1 {
				z.DecSwallow()
			} else {
				if yyxt3 := z.Extension(yyv1[yyj1]); yyxt3 != nil {
					z.DecExtension(&yyv1[yyj1], yyxt3)
				} else {
					yyv1[yyj1].CodecDecodeSelf(d)
				}
			}
		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} else if yyj1 == 0 && yyv1 == nil {
			yyv1 = make([]Medium, 0)
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}
}

func (x codecSelfer2970) encSliceMedium(v []Medium, e *codec1978.Encoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Encoder(e)
	_, _, _ = h, z, r
	if v == nil {
		r.EncodeNil()
		return
	}
	z.EncWriteArrayStart(len(v))
	for yyv1 := range v {
		z.EncWriteArrayElem()
		yy2 := &v[yyv1]
		if yyxt3 := z.Extension(yy2); yyxt3 != nil {
			z.EncExtension(yy2, yyxt3)
		} else {
			yy2.CodecEncodeSelf(e)
		}
	}
	z.EncWriteArrayEnd()
}

func (x codecSelfer2970) decSliceMedium(v *[]Medium, d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyh1.IsNil {
		if yyv1 != nil {
			yyv1 = nil
			yyc1 = true
		}
	} else if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []Medium{}
			yyc1 = true
		} else if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} else {
		yyhl1 := yyl1 > 0
		var yyrl1 int
		_ = yyrl1
		if yyhl1 {
			if yyl1 > cap(yyv1) {
				yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 128)
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} else {
					yyv1 = make([]Medium, yyrl1)
				}
				yyc1 = true
			} else if yyl1 != len(yyv1) {
				yyv1 = yyv1[:yyl1]
				yyc1 = true
			}
		}
		var yyj1 int
		for yyj1 = 0; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || z.DecCheckBreak()); yyj1++ { // bounds-check-elimination
			if yyj1 == 0 && yyv1 == nil {
				if yyhl1 {
					yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 128)
				} else {
					yyrl1 = 8
				}
				yyv1 = make([]Medium, yyrl1)
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)
			var yydb1 bool
			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, Medium{})
				yyc1 = true
			}
			if yydb1 {
				z.DecSwallow()
			} else {
				if yyxt3 := z.Extension(yyv1[yyj1]); yyxt3 != nil {
					z.DecExtension(&yyv1[yyj1], yyxt3)
				} else {
					yyv1[yyj1].CodecDecodeSelf(d)
				}
			}
		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} else if yyj1 == 0 && yyv1 == nil {
			yyv1 = make([]Medium, 0)
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}
}

func (x codecSelfer2970) encHeavies(v Heavies, e *codec1978.Encoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Encoder(e)
	_, _, _ = h, z, r
	if v == nil {
		r.EncodeNil()
		return
	}
	z.EncWriteArrayStart(len(v))
	for yyv1 := range v {
		z.EncWriteArrayElem()
		yy2 := &v[yyv1]
		if yyxt3 := z.Extension(yy2); yyxt3 != nil {
			z.EncExtension(yy2, yyxt3)
		} else {
			yy2.CodecEncodeSelf(e)
		}
	}
	z.EncWriteArrayEnd()
}

func (x codecSelfer2970) decHeavies(v *Heavies, d *codec1978.Decoder) {
	var h codecSelfer2970
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyh1.IsNil {
		if yyv1 != nil {
			yyv1 = nil
			yyc1 = true
		}
	} else if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []Heavy{}
			yyc1 = true
		} else if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} else {
		yyhl1 := yyl1 > 0
		var yyrl1 int
		_ = yyrl1
		if yyhl1 {
			if yyl1 > cap(yyv1) {
				yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 152)
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} else {
					yyv1 = make([]Heavy, yyrl1)
				}
				yyc1 = true
			} else if yyl1 != len(yyv1) {
				yyv1 = yyv1[:yyl1]
				yyc1 = true
			}
		}
		var yyj1 int
		for yyj1 = 0; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || z.DecCheckBreak()); yyj1++ { // bounds-check-elimination
			if yyj1 == 0 && yyv1 == nil {
				if yyhl1 {
					yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 152)
				} else {
					yyrl1 = 8
				}
				yyv1 = make([]Heavy, yyrl1)
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)
			var yydb1 bool
			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, Heavy{})
				yyc1 = true
			}
			if yydb1 {
				z.DecSwallow()
			} else {
				if yyxt3 := z.Extension(yyv1[yyj1]); yyxt3 != nil {
					z.DecExtension(&yyv1[yyj1], yyxt3)
				} else {
					yyv1[yyj1].CodecDecodeSelf(d)
				}
			}
		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} else if yyj1 == 0 && yyv1 == nil {
			yyv1 = make([]Heavy, 0)
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}
}
