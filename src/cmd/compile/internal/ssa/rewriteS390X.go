// Code generated from gen/S390X.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "cmd/compile/internal/types"
import "cmd/internal/obj/s390x"

func rewriteValueS390X(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		v.Op = OpS390XADDW
		return true
	case OpAdd32:
		v.Op = OpS390XADDW
		return true
	case OpAdd32F:
		v.Op = OpS390XFADDS
		return true
	case OpAdd64:
		v.Op = OpS390XADD
		return true
	case OpAdd64F:
		v.Op = OpS390XFADD
		return true
	case OpAdd8:
		v.Op = OpS390XADDW
		return true
	case OpAddPtr:
		v.Op = OpS390XADD
		return true
	case OpAddr:
		v.Op = OpS390XMOVDaddr
		return true
	case OpAnd16:
		v.Op = OpS390XANDW
		return true
	case OpAnd32:
		v.Op = OpS390XANDW
		return true
	case OpAnd64:
		v.Op = OpS390XAND
		return true
	case OpAnd8:
		v.Op = OpS390XANDW
		return true
	case OpAndB:
		v.Op = OpS390XANDW
		return true
	case OpAtomicAdd32:
		return rewriteValueS390X_OpAtomicAdd32(v)
	case OpAtomicAdd64:
		return rewriteValueS390X_OpAtomicAdd64(v)
	case OpAtomicAnd8:
		return rewriteValueS390X_OpAtomicAnd8(v)
	case OpAtomicCompareAndSwap32:
		v.Op = OpS390XLoweredAtomicCas32
		return true
	case OpAtomicCompareAndSwap64:
		v.Op = OpS390XLoweredAtomicCas64
		return true
	case OpAtomicExchange32:
		v.Op = OpS390XLoweredAtomicExchange32
		return true
	case OpAtomicExchange64:
		v.Op = OpS390XLoweredAtomicExchange64
		return true
	case OpAtomicLoad32:
		v.Op = OpS390XMOVWZatomicload
		return true
	case OpAtomicLoad64:
		v.Op = OpS390XMOVDatomicload
		return true
	case OpAtomicLoad8:
		v.Op = OpS390XMOVBZatomicload
		return true
	case OpAtomicLoadAcq32:
		v.Op = OpS390XMOVWZatomicload
		return true
	case OpAtomicLoadPtr:
		v.Op = OpS390XMOVDatomicload
		return true
	case OpAtomicOr8:
		return rewriteValueS390X_OpAtomicOr8(v)
	case OpAtomicStore32:
		return rewriteValueS390X_OpAtomicStore32(v)
	case OpAtomicStore64:
		return rewriteValueS390X_OpAtomicStore64(v)
	case OpAtomicStore8:
		return rewriteValueS390X_OpAtomicStore8(v)
	case OpAtomicStorePtrNoWB:
		return rewriteValueS390X_OpAtomicStorePtrNoWB(v)
	case OpAtomicStoreRel32:
		v.Op = OpS390XMOVWatomicstore
		return true
	case OpAvg64u:
		return rewriteValueS390X_OpAvg64u(v)
	case OpBitLen64:
		return rewriteValueS390X_OpBitLen64(v)
	case OpBswap32:
		v.Op = OpS390XMOVWBR
		return true
	case OpBswap64:
		v.Op = OpS390XMOVDBR
		return true
	case OpCeil:
		return rewriteValueS390X_OpCeil(v)
	case OpClosureCall:
		v.Op = OpS390XCALLclosure
		return true
	case OpCom16:
		v.Op = OpS390XNOTW
		return true
	case OpCom32:
		v.Op = OpS390XNOTW
		return true
	case OpCom64:
		v.Op = OpS390XNOT
		return true
	case OpCom8:
		v.Op = OpS390XNOTW
		return true
	case OpConst16:
		v.Op = OpS390XMOVDconst
		return true
	case OpConst32:
		v.Op = OpS390XMOVDconst
		return true
	case OpConst32F:
		v.Op = OpS390XFMOVSconst
		return true
	case OpConst64:
		v.Op = OpS390XMOVDconst
		return true
	case OpConst64F:
		v.Op = OpS390XFMOVDconst
		return true
	case OpConst8:
		v.Op = OpS390XMOVDconst
		return true
	case OpConstBool:
		v.Op = OpS390XMOVDconst
		return true
	case OpConstNil:
		return rewriteValueS390X_OpConstNil(v)
	case OpCtz32:
		return rewriteValueS390X_OpCtz32(v)
	case OpCtz32NonZero:
		v.Op = OpCtz32
		return true
	case OpCtz64:
		return rewriteValueS390X_OpCtz64(v)
	case OpCtz64NonZero:
		v.Op = OpCtz64
		return true
	case OpCvt32Fto32:
		v.Op = OpS390XCFEBRA
		return true
	case OpCvt32Fto64:
		v.Op = OpS390XCGEBRA
		return true
	case OpCvt32Fto64F:
		v.Op = OpS390XLDEBR
		return true
	case OpCvt32to32F:
		v.Op = OpS390XCEFBRA
		return true
	case OpCvt32to64F:
		v.Op = OpS390XCDFBRA
		return true
	case OpCvt64Fto32:
		v.Op = OpS390XCFDBRA
		return true
	case OpCvt64Fto32F:
		v.Op = OpS390XLEDBR
		return true
	case OpCvt64Fto64:
		v.Op = OpS390XCGDBRA
		return true
	case OpCvt64to32F:
		v.Op = OpS390XCEGBRA
		return true
	case OpCvt64to64F:
		v.Op = OpS390XCDGBRA
		return true
	case OpDiv16:
		return rewriteValueS390X_OpDiv16(v)
	case OpDiv16u:
		return rewriteValueS390X_OpDiv16u(v)
	case OpDiv32:
		return rewriteValueS390X_OpDiv32(v)
	case OpDiv32F:
		v.Op = OpS390XFDIVS
		return true
	case OpDiv32u:
		return rewriteValueS390X_OpDiv32u(v)
	case OpDiv64:
		return rewriteValueS390X_OpDiv64(v)
	case OpDiv64F:
		v.Op = OpS390XFDIV
		return true
	case OpDiv64u:
		v.Op = OpS390XDIVDU
		return true
	case OpDiv8:
		return rewriteValueS390X_OpDiv8(v)
	case OpDiv8u:
		return rewriteValueS390X_OpDiv8u(v)
	case OpEq16:
		return rewriteValueS390X_OpEq16(v)
	case OpEq32:
		return rewriteValueS390X_OpEq32(v)
	case OpEq32F:
		return rewriteValueS390X_OpEq32F(v)
	case OpEq64:
		return rewriteValueS390X_OpEq64(v)
	case OpEq64F:
		return rewriteValueS390X_OpEq64F(v)
	case OpEq8:
		return rewriteValueS390X_OpEq8(v)
	case OpEqB:
		return rewriteValueS390X_OpEqB(v)
	case OpEqPtr:
		return rewriteValueS390X_OpEqPtr(v)
	case OpFMA:
		return rewriteValueS390X_OpFMA(v)
	case OpFloor:
		return rewriteValueS390X_OpFloor(v)
	case OpGeq32F:
		return rewriteValueS390X_OpGeq32F(v)
	case OpGeq64F:
		return rewriteValueS390X_OpGeq64F(v)
	case OpGetCallerPC:
		v.Op = OpS390XLoweredGetCallerPC
		return true
	case OpGetCallerSP:
		v.Op = OpS390XLoweredGetCallerSP
		return true
	case OpGetClosurePtr:
		v.Op = OpS390XLoweredGetClosurePtr
		return true
	case OpGetG:
		v.Op = OpS390XLoweredGetG
		return true
	case OpGreater32F:
		return rewriteValueS390X_OpGreater32F(v)
	case OpGreater64F:
		return rewriteValueS390X_OpGreater64F(v)
	case OpHmul32:
		return rewriteValueS390X_OpHmul32(v)
	case OpHmul32u:
		return rewriteValueS390X_OpHmul32u(v)
	case OpHmul64:
		v.Op = OpS390XMULHD
		return true
	case OpHmul64u:
		v.Op = OpS390XMULHDU
		return true
	case OpITab:
		return rewriteValueS390X_OpITab(v)
	case OpInterCall:
		v.Op = OpS390XCALLinter
		return true
	case OpIsInBounds:
		return rewriteValueS390X_OpIsInBounds(v)
	case OpIsNonNil:
		return rewriteValueS390X_OpIsNonNil(v)
	case OpIsSliceInBounds:
		return rewriteValueS390X_OpIsSliceInBounds(v)
	case OpLeq16:
		return rewriteValueS390X_OpLeq16(v)
	case OpLeq16U:
		return rewriteValueS390X_OpLeq16U(v)
	case OpLeq32:
		return rewriteValueS390X_OpLeq32(v)
	case OpLeq32F:
		return rewriteValueS390X_OpLeq32F(v)
	case OpLeq32U:
		return rewriteValueS390X_OpLeq32U(v)
	case OpLeq64:
		return rewriteValueS390X_OpLeq64(v)
	case OpLeq64F:
		return rewriteValueS390X_OpLeq64F(v)
	case OpLeq64U:
		return rewriteValueS390X_OpLeq64U(v)
	case OpLeq8:
		return rewriteValueS390X_OpLeq8(v)
	case OpLeq8U:
		return rewriteValueS390X_OpLeq8U(v)
	case OpLess16:
		return rewriteValueS390X_OpLess16(v)
	case OpLess16U:
		return rewriteValueS390X_OpLess16U(v)
	case OpLess32:
		return rewriteValueS390X_OpLess32(v)
	case OpLess32F:
		return rewriteValueS390X_OpLess32F(v)
	case OpLess32U:
		return rewriteValueS390X_OpLess32U(v)
	case OpLess64:
		return rewriteValueS390X_OpLess64(v)
	case OpLess64F:
		return rewriteValueS390X_OpLess64F(v)
	case OpLess64U:
		return rewriteValueS390X_OpLess64U(v)
	case OpLess8:
		return rewriteValueS390X_OpLess8(v)
	case OpLess8U:
		return rewriteValueS390X_OpLess8U(v)
	case OpLoad:
		return rewriteValueS390X_OpLoad(v)
	case OpLocalAddr:
		return rewriteValueS390X_OpLocalAddr(v)
	case OpLsh16x16:
		return rewriteValueS390X_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValueS390X_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValueS390X_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValueS390X_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValueS390X_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValueS390X_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValueS390X_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValueS390X_OpLsh32x8(v)
	case OpLsh64x16:
		return rewriteValueS390X_OpLsh64x16(v)
	case OpLsh64x32:
		return rewriteValueS390X_OpLsh64x32(v)
	case OpLsh64x64:
		return rewriteValueS390X_OpLsh64x64(v)
	case OpLsh64x8:
		return rewriteValueS390X_OpLsh64x8(v)
	case OpLsh8x16:
		return rewriteValueS390X_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValueS390X_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValueS390X_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValueS390X_OpLsh8x8(v)
	case OpMod16:
		return rewriteValueS390X_OpMod16(v)
	case OpMod16u:
		return rewriteValueS390X_OpMod16u(v)
	case OpMod32:
		return rewriteValueS390X_OpMod32(v)
	case OpMod32u:
		return rewriteValueS390X_OpMod32u(v)
	case OpMod64:
		return rewriteValueS390X_OpMod64(v)
	case OpMod64u:
		v.Op = OpS390XMODDU
		return true
	case OpMod8:
		return rewriteValueS390X_OpMod8(v)
	case OpMod8u:
		return rewriteValueS390X_OpMod8u(v)
	case OpMove:
		return rewriteValueS390X_OpMove(v)
	case OpMul16:
		v.Op = OpS390XMULLW
		return true
	case OpMul32:
		v.Op = OpS390XMULLW
		return true
	case OpMul32F:
		v.Op = OpS390XFMULS
		return true
	case OpMul64:
		v.Op = OpS390XMULLD
		return true
	case OpMul64F:
		v.Op = OpS390XFMUL
		return true
	case OpMul64uhilo:
		v.Op = OpS390XMLGR
		return true
	case OpMul8:
		v.Op = OpS390XMULLW
		return true
	case OpNeg16:
		v.Op = OpS390XNEGW
		return true
	case OpNeg32:
		v.Op = OpS390XNEGW
		return true
	case OpNeg32F:
		v.Op = OpS390XFNEGS
		return true
	case OpNeg64:
		v.Op = OpS390XNEG
		return true
	case OpNeg64F:
		v.Op = OpS390XFNEG
		return true
	case OpNeg8:
		v.Op = OpS390XNEGW
		return true
	case OpNeq16:
		return rewriteValueS390X_OpNeq16(v)
	case OpNeq32:
		return rewriteValueS390X_OpNeq32(v)
	case OpNeq32F:
		return rewriteValueS390X_OpNeq32F(v)
	case OpNeq64:
		return rewriteValueS390X_OpNeq64(v)
	case OpNeq64F:
		return rewriteValueS390X_OpNeq64F(v)
	case OpNeq8:
		return rewriteValueS390X_OpNeq8(v)
	case OpNeqB:
		return rewriteValueS390X_OpNeqB(v)
	case OpNeqPtr:
		return rewriteValueS390X_OpNeqPtr(v)
	case OpNilCheck:
		v.Op = OpS390XLoweredNilCheck
		return true
	case OpNot:
		return rewriteValueS390X_OpNot(v)
	case OpOffPtr:
		return rewriteValueS390X_OpOffPtr(v)
	case OpOr16:
		v.Op = OpS390XORW
		return true
	case OpOr32:
		v.Op = OpS390XORW
		return true
	case OpOr64:
		v.Op = OpS390XOR
		return true
	case OpOr8:
		v.Op = OpS390XORW
		return true
	case OpOrB:
		v.Op = OpS390XORW
		return true
	case OpPanicBounds:
		return rewriteValueS390X_OpPanicBounds(v)
	case OpPopCount16:
		return rewriteValueS390X_OpPopCount16(v)
	case OpPopCount32:
		return rewriteValueS390X_OpPopCount32(v)
	case OpPopCount64:
		return rewriteValueS390X_OpPopCount64(v)
	case OpPopCount8:
		return rewriteValueS390X_OpPopCount8(v)
	case OpRotateLeft16:
		return rewriteValueS390X_OpRotateLeft16(v)
	case OpRotateLeft32:
		v.Op = OpS390XRLL
		return true
	case OpRotateLeft64:
		v.Op = OpS390XRLLG
		return true
	case OpRotateLeft8:
		return rewriteValueS390X_OpRotateLeft8(v)
	case OpRound:
		return rewriteValueS390X_OpRound(v)
	case OpRound32F:
		v.Op = OpS390XLoweredRound32F
		return true
	case OpRound64F:
		v.Op = OpS390XLoweredRound64F
		return true
	case OpRoundToEven:
		return rewriteValueS390X_OpRoundToEven(v)
	case OpRsh16Ux16:
		return rewriteValueS390X_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValueS390X_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValueS390X_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValueS390X_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValueS390X_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValueS390X_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValueS390X_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValueS390X_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValueS390X_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValueS390X_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValueS390X_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValueS390X_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValueS390X_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValueS390X_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValueS390X_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValueS390X_OpRsh32x8(v)
	case OpRsh64Ux16:
		return rewriteValueS390X_OpRsh64Ux16(v)
	case OpRsh64Ux32:
		return rewriteValueS390X_OpRsh64Ux32(v)
	case OpRsh64Ux64:
		return rewriteValueS390X_OpRsh64Ux64(v)
	case OpRsh64Ux8:
		return rewriteValueS390X_OpRsh64Ux8(v)
	case OpRsh64x16:
		return rewriteValueS390X_OpRsh64x16(v)
	case OpRsh64x32:
		return rewriteValueS390X_OpRsh64x32(v)
	case OpRsh64x64:
		return rewriteValueS390X_OpRsh64x64(v)
	case OpRsh64x8:
		return rewriteValueS390X_OpRsh64x8(v)
	case OpRsh8Ux16:
		return rewriteValueS390X_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValueS390X_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValueS390X_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValueS390X_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValueS390X_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValueS390X_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValueS390X_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValueS390X_OpRsh8x8(v)
	case OpS390XADD:
		return rewriteValueS390X_OpS390XADD(v)
	case OpS390XADDC:
		return rewriteValueS390X_OpS390XADDC(v)
	case OpS390XADDE:
		return rewriteValueS390X_OpS390XADDE(v)
	case OpS390XADDW:
		return rewriteValueS390X_OpS390XADDW(v)
	case OpS390XADDWconst:
		return rewriteValueS390X_OpS390XADDWconst(v)
	case OpS390XADDWload:
		return rewriteValueS390X_OpS390XADDWload(v)
	case OpS390XADDconst:
		return rewriteValueS390X_OpS390XADDconst(v)
	case OpS390XADDload:
		return rewriteValueS390X_OpS390XADDload(v)
	case OpS390XAND:
		return rewriteValueS390X_OpS390XAND(v)
	case OpS390XANDW:
		return rewriteValueS390X_OpS390XANDW(v)
	case OpS390XANDWconst:
		return rewriteValueS390X_OpS390XANDWconst(v)
	case OpS390XANDWload:
		return rewriteValueS390X_OpS390XANDWload(v)
	case OpS390XANDconst:
		return rewriteValueS390X_OpS390XANDconst(v)
	case OpS390XANDload:
		return rewriteValueS390X_OpS390XANDload(v)
	case OpS390XCMP:
		return rewriteValueS390X_OpS390XCMP(v)
	case OpS390XCMPU:
		return rewriteValueS390X_OpS390XCMPU(v)
	case OpS390XCMPUconst:
		return rewriteValueS390X_OpS390XCMPUconst(v)
	case OpS390XCMPW:
		return rewriteValueS390X_OpS390XCMPW(v)
	case OpS390XCMPWU:
		return rewriteValueS390X_OpS390XCMPWU(v)
	case OpS390XCMPWUconst:
		return rewriteValueS390X_OpS390XCMPWUconst(v)
	case OpS390XCMPWconst:
		return rewriteValueS390X_OpS390XCMPWconst(v)
	case OpS390XCMPconst:
		return rewriteValueS390X_OpS390XCMPconst(v)
	case OpS390XCPSDR:
		return rewriteValueS390X_OpS390XCPSDR(v)
	case OpS390XFADD:
		return rewriteValueS390X_OpS390XFADD(v)
	case OpS390XFADDS:
		return rewriteValueS390X_OpS390XFADDS(v)
	case OpS390XFMOVDload:
		return rewriteValueS390X_OpS390XFMOVDload(v)
	case OpS390XFMOVDloadidx:
		return rewriteValueS390X_OpS390XFMOVDloadidx(v)
	case OpS390XFMOVDstore:
		return rewriteValueS390X_OpS390XFMOVDstore(v)
	case OpS390XFMOVDstoreidx:
		return rewriteValueS390X_OpS390XFMOVDstoreidx(v)
	case OpS390XFMOVSload:
		return rewriteValueS390X_OpS390XFMOVSload(v)
	case OpS390XFMOVSloadidx:
		return rewriteValueS390X_OpS390XFMOVSloadidx(v)
	case OpS390XFMOVSstore:
		return rewriteValueS390X_OpS390XFMOVSstore(v)
	case OpS390XFMOVSstoreidx:
		return rewriteValueS390X_OpS390XFMOVSstoreidx(v)
	case OpS390XFNEG:
		return rewriteValueS390X_OpS390XFNEG(v)
	case OpS390XFNEGS:
		return rewriteValueS390X_OpS390XFNEGS(v)
	case OpS390XFSUB:
		return rewriteValueS390X_OpS390XFSUB(v)
	case OpS390XFSUBS:
		return rewriteValueS390X_OpS390XFSUBS(v)
	case OpS390XLDGR:
		return rewriteValueS390X_OpS390XLDGR(v)
	case OpS390XLEDBR:
		return rewriteValueS390X_OpS390XLEDBR(v)
	case OpS390XLGDR:
		return rewriteValueS390X_OpS390XLGDR(v)
	case OpS390XLOCGR:
		return rewriteValueS390X_OpS390XLOCGR(v)
	case OpS390XLoweredRound32F:
		return rewriteValueS390X_OpS390XLoweredRound32F(v)
	case OpS390XLoweredRound64F:
		return rewriteValueS390X_OpS390XLoweredRound64F(v)
	case OpS390XMOVBZload:
		return rewriteValueS390X_OpS390XMOVBZload(v)
	case OpS390XMOVBZloadidx:
		return rewriteValueS390X_OpS390XMOVBZloadidx(v)
	case OpS390XMOVBZreg:
		return rewriteValueS390X_OpS390XMOVBZreg(v)
	case OpS390XMOVBload:
		return rewriteValueS390X_OpS390XMOVBload(v)
	case OpS390XMOVBloadidx:
		return rewriteValueS390X_OpS390XMOVBloadidx(v)
	case OpS390XMOVBreg:
		return rewriteValueS390X_OpS390XMOVBreg(v)
	case OpS390XMOVBstore:
		return rewriteValueS390X_OpS390XMOVBstore(v)
	case OpS390XMOVBstoreconst:
		return rewriteValueS390X_OpS390XMOVBstoreconst(v)
	case OpS390XMOVBstoreidx:
		return rewriteValueS390X_OpS390XMOVBstoreidx(v)
	case OpS390XMOVDaddridx:
		return rewriteValueS390X_OpS390XMOVDaddridx(v)
	case OpS390XMOVDload:
		return rewriteValueS390X_OpS390XMOVDload(v)
	case OpS390XMOVDloadidx:
		return rewriteValueS390X_OpS390XMOVDloadidx(v)
	case OpS390XMOVDstore:
		return rewriteValueS390X_OpS390XMOVDstore(v)
	case OpS390XMOVDstoreconst:
		return rewriteValueS390X_OpS390XMOVDstoreconst(v)
	case OpS390XMOVDstoreidx:
		return rewriteValueS390X_OpS390XMOVDstoreidx(v)
	case OpS390XMOVHBRstore:
		return rewriteValueS390X_OpS390XMOVHBRstore(v)
	case OpS390XMOVHBRstoreidx:
		return rewriteValueS390X_OpS390XMOVHBRstoreidx(v)
	case OpS390XMOVHZload:
		return rewriteValueS390X_OpS390XMOVHZload(v)
	case OpS390XMOVHZloadidx:
		return rewriteValueS390X_OpS390XMOVHZloadidx(v)
	case OpS390XMOVHZreg:
		return rewriteValueS390X_OpS390XMOVHZreg(v)
	case OpS390XMOVHload:
		return rewriteValueS390X_OpS390XMOVHload(v)
	case OpS390XMOVHloadidx:
		return rewriteValueS390X_OpS390XMOVHloadidx(v)
	case OpS390XMOVHreg:
		return rewriteValueS390X_OpS390XMOVHreg(v)
	case OpS390XMOVHstore:
		return rewriteValueS390X_OpS390XMOVHstore(v)
	case OpS390XMOVHstoreconst:
		return rewriteValueS390X_OpS390XMOVHstoreconst(v)
	case OpS390XMOVHstoreidx:
		return rewriteValueS390X_OpS390XMOVHstoreidx(v)
	case OpS390XMOVWBRstore:
		return rewriteValueS390X_OpS390XMOVWBRstore(v)
	case OpS390XMOVWBRstoreidx:
		return rewriteValueS390X_OpS390XMOVWBRstoreidx(v)
	case OpS390XMOVWZload:
		return rewriteValueS390X_OpS390XMOVWZload(v)
	case OpS390XMOVWZloadidx:
		return rewriteValueS390X_OpS390XMOVWZloadidx(v)
	case OpS390XMOVWZreg:
		return rewriteValueS390X_OpS390XMOVWZreg(v)
	case OpS390XMOVWload:
		return rewriteValueS390X_OpS390XMOVWload(v)
	case OpS390XMOVWloadidx:
		return rewriteValueS390X_OpS390XMOVWloadidx(v)
	case OpS390XMOVWreg:
		return rewriteValueS390X_OpS390XMOVWreg(v)
	case OpS390XMOVWstore:
		return rewriteValueS390X_OpS390XMOVWstore(v)
	case OpS390XMOVWstoreconst:
		return rewriteValueS390X_OpS390XMOVWstoreconst(v)
	case OpS390XMOVWstoreidx:
		return rewriteValueS390X_OpS390XMOVWstoreidx(v)
	case OpS390XMULLD:
		return rewriteValueS390X_OpS390XMULLD(v)
	case OpS390XMULLDconst:
		return rewriteValueS390X_OpS390XMULLDconst(v)
	case OpS390XMULLDload:
		return rewriteValueS390X_OpS390XMULLDload(v)
	case OpS390XMULLW:
		return rewriteValueS390X_OpS390XMULLW(v)
	case OpS390XMULLWconst:
		return rewriteValueS390X_OpS390XMULLWconst(v)
	case OpS390XMULLWload:
		return rewriteValueS390X_OpS390XMULLWload(v)
	case OpS390XNEG:
		return rewriteValueS390X_OpS390XNEG(v)
	case OpS390XNEGW:
		return rewriteValueS390X_OpS390XNEGW(v)
	case OpS390XNOT:
		return rewriteValueS390X_OpS390XNOT(v)
	case OpS390XNOTW:
		return rewriteValueS390X_OpS390XNOTW(v)
	case OpS390XOR:
		return rewriteValueS390X_OpS390XOR(v)
	case OpS390XORW:
		return rewriteValueS390X_OpS390XORW(v)
	case OpS390XORWconst:
		return rewriteValueS390X_OpS390XORWconst(v)
	case OpS390XORWload:
		return rewriteValueS390X_OpS390XORWload(v)
	case OpS390XORconst:
		return rewriteValueS390X_OpS390XORconst(v)
	case OpS390XORload:
		return rewriteValueS390X_OpS390XORload(v)
	case OpS390XRLL:
		return rewriteValueS390X_OpS390XRLL(v)
	case OpS390XRLLG:
		return rewriteValueS390X_OpS390XRLLG(v)
	case OpS390XSLD:
		return rewriteValueS390X_OpS390XSLD(v)
	case OpS390XSLW:
		return rewriteValueS390X_OpS390XSLW(v)
	case OpS390XSRAD:
		return rewriteValueS390X_OpS390XSRAD(v)
	case OpS390XSRADconst:
		return rewriteValueS390X_OpS390XSRADconst(v)
	case OpS390XSRAW:
		return rewriteValueS390X_OpS390XSRAW(v)
	case OpS390XSRAWconst:
		return rewriteValueS390X_OpS390XSRAWconst(v)
	case OpS390XSRD:
		return rewriteValueS390X_OpS390XSRD(v)
	case OpS390XSRDconst:
		return rewriteValueS390X_OpS390XSRDconst(v)
	case OpS390XSRW:
		return rewriteValueS390X_OpS390XSRW(v)
	case OpS390XSTM2:
		return rewriteValueS390X_OpS390XSTM2(v)
	case OpS390XSTMG2:
		return rewriteValueS390X_OpS390XSTMG2(v)
	case OpS390XSUB:
		return rewriteValueS390X_OpS390XSUB(v)
	case OpS390XSUBE:
		return rewriteValueS390X_OpS390XSUBE(v)
	case OpS390XSUBW:
		return rewriteValueS390X_OpS390XSUBW(v)
	case OpS390XSUBWconst:
		return rewriteValueS390X_OpS390XSUBWconst(v)
	case OpS390XSUBWload:
		return rewriteValueS390X_OpS390XSUBWload(v)
	case OpS390XSUBconst:
		return rewriteValueS390X_OpS390XSUBconst(v)
	case OpS390XSUBload:
		return rewriteValueS390X_OpS390XSUBload(v)
	case OpS390XSumBytes2:
		return rewriteValueS390X_OpS390XSumBytes2(v)
	case OpS390XSumBytes4:
		return rewriteValueS390X_OpS390XSumBytes4(v)
	case OpS390XSumBytes8:
		return rewriteValueS390X_OpS390XSumBytes8(v)
	case OpS390XXOR:
		return rewriteValueS390X_OpS390XXOR(v)
	case OpS390XXORW:
		return rewriteValueS390X_OpS390XXORW(v)
	case OpS390XXORWconst:
		return rewriteValueS390X_OpS390XXORWconst(v)
	case OpS390XXORWload:
		return rewriteValueS390X_OpS390XXORWload(v)
	case OpS390XXORconst:
		return rewriteValueS390X_OpS390XXORconst(v)
	case OpS390XXORload:
		return rewriteValueS390X_OpS390XXORload(v)
	case OpSelect0:
		return rewriteValueS390X_OpSelect0(v)
	case OpSelect1:
		return rewriteValueS390X_OpSelect1(v)
	case OpSignExt16to32:
		v.Op = OpS390XMOVHreg
		return true
	case OpSignExt16to64:
		v.Op = OpS390XMOVHreg
		return true
	case OpSignExt32to64:
		v.Op = OpS390XMOVWreg
		return true
	case OpSignExt8to16:
		v.Op = OpS390XMOVBreg
		return true
	case OpSignExt8to32:
		v.Op = OpS390XMOVBreg
		return true
	case OpSignExt8to64:
		v.Op = OpS390XMOVBreg
		return true
	case OpSlicemask:
		return rewriteValueS390X_OpSlicemask(v)
	case OpSqrt:
		v.Op = OpS390XFSQRT
		return true
	case OpStaticCall:
		v.Op = OpS390XCALLstatic
		return true
	case OpStore:
		return rewriteValueS390X_OpStore(v)
	case OpSub16:
		v.Op = OpS390XSUBW
		return true
	case OpSub32:
		v.Op = OpS390XSUBW
		return true
	case OpSub32F:
		v.Op = OpS390XFSUBS
		return true
	case OpSub64:
		v.Op = OpS390XSUB
		return true
	case OpSub64F:
		v.Op = OpS390XFSUB
		return true
	case OpSub8:
		v.Op = OpS390XSUBW
		return true
	case OpSubPtr:
		v.Op = OpS390XSUB
		return true
	case OpTrunc:
		return rewriteValueS390X_OpTrunc(v)
	case OpTrunc16to8:
		v.Op = OpCopy
		return true
	case OpTrunc32to16:
		v.Op = OpCopy
		return true
	case OpTrunc32to8:
		v.Op = OpCopy
		return true
	case OpTrunc64to16:
		v.Op = OpCopy
		return true
	case OpTrunc64to32:
		v.Op = OpCopy
		return true
	case OpTrunc64to8:
		v.Op = OpCopy
		return true
	case OpWB:
		v.Op = OpS390XLoweredWB
		return true
	case OpXor16:
		v.Op = OpS390XXORW
		return true
	case OpXor32:
		v.Op = OpS390XXORW
		return true
	case OpXor64:
		v.Op = OpS390XXOR
		return true
	case OpXor8:
		v.Op = OpS390XXORW
		return true
	case OpZero:
		return rewriteValueS390X_OpZero(v)
	case OpZeroExt16to32:
		v.Op = OpS390XMOVHZreg
		return true
	case OpZeroExt16to64:
		v.Op = OpS390XMOVHZreg
		return true
	case OpZeroExt32to64:
		v.Op = OpS390XMOVWZreg
		return true
	case OpZeroExt8to16:
		v.Op = OpS390XMOVBZreg
		return true
	case OpZeroExt8to32:
		v.Op = OpS390XMOVBZreg
		return true
	case OpZeroExt8to64:
		v.Op = OpS390XMOVBZreg
		return true
	}
	return false
}
func rewriteValueS390X_OpAtomicAdd32(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (AtomicAdd32 ptr val mem)
	// result: (AddTupleFirst32 val (LAA ptr val mem))
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpS390XAddTupleFirst32)
		v.AddArg(val)
		v0 := b.NewValue0(v.Pos, OpS390XLAA, types.NewTuple(typ.UInt32, types.TypeMem))
		v0.AddArg(ptr)
		v0.AddArg(val)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpAtomicAdd64(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (AtomicAdd64 ptr val mem)
	// result: (AddTupleFirst64 val (LAAG ptr val mem))
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpS390XAddTupleFirst64)
		v.AddArg(val)
		v0 := b.NewValue0(v.Pos, OpS390XLAAG, types.NewTuple(typ.UInt64, types.TypeMem))
		v0.AddArg(ptr)
		v0.AddArg(val)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpAtomicAnd8(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (AtomicAnd8 ptr val mem)
	// result: (LANfloor ptr (RLL <typ.UInt32> (ORWconst <typ.UInt32> val [-1<<8]) (RXSBG <typ.UInt32> {s390x.NewRotateParams(59, 60, 3)} (MOVDconst [3<<3]) ptr)) mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpS390XLANfloor)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpS390XRLL, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpS390XORWconst, typ.UInt32)
		v1.AuxInt = -1 << 8
		v1.AddArg(val)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XRXSBG, typ.UInt32)
		v2.Aux = s390x.NewRotateParams(59, 60, 3)
		v3 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v3.AuxInt = 3 << 3
		v2.AddArg(v3)
		v2.AddArg(ptr)
		v0.AddArg(v2)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpAtomicOr8(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (AtomicOr8 ptr val mem)
	// result: (LAOfloor ptr (SLW <typ.UInt32> (MOVBZreg <typ.UInt32> val) (RXSBG <typ.UInt32> {s390x.NewRotateParams(59, 60, 3)} (MOVDconst [3<<3]) ptr)) mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpS390XLAOfloor)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpS390XSLW, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt32)
		v1.AddArg(val)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XRXSBG, typ.UInt32)
		v2.Aux = s390x.NewRotateParams(59, 60, 3)
		v3 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v3.AuxInt = 3 << 3
		v2.AddArg(v3)
		v2.AddArg(ptr)
		v0.AddArg(v2)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpAtomicStore32(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (AtomicStore32 ptr val mem)
	// result: (SYNC (MOVWatomicstore ptr val mem))
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpS390XSYNC)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWatomicstore, types.TypeMem)
		v0.AddArg(ptr)
		v0.AddArg(val)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpAtomicStore64(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (AtomicStore64 ptr val mem)
	// result: (SYNC (MOVDatomicstore ptr val mem))
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpS390XSYNC)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDatomicstore, types.TypeMem)
		v0.AddArg(ptr)
		v0.AddArg(val)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpAtomicStore8(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (AtomicStore8 ptr val mem)
	// result: (SYNC (MOVBatomicstore ptr val mem))
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpS390XSYNC)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBatomicstore, types.TypeMem)
		v0.AddArg(ptr)
		v0.AddArg(val)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpAtomicStorePtrNoWB(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (AtomicStorePtrNoWB ptr val mem)
	// result: (SYNC (MOVDatomicstore ptr val mem))
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpS390XSYNC)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDatomicstore, types.TypeMem)
		v0.AddArg(ptr)
		v0.AddArg(val)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpAvg64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Avg64u <t> x y)
	// result: (ADD (SRDconst <t> (SUB <t> x y) [1]) y)
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XADD)
		v0 := b.NewValue0(v.Pos, OpS390XSRDconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpS390XSUB, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpBitLen64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (BitLen64 x)
	// result: (SUB (MOVDconst [64]) (FLOGR x))
	for {
		x := v_0
		v.reset(OpS390XSUB)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 64
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XFLOGR, typ.UInt64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpCeil(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Ceil x)
	// result: (FIDBR [6] x)
	for {
		x := v_0
		v.reset(OpS390XFIDBR)
		v.AuxInt = 6
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// result: (MOVDconst [0])
	for {
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueS390X_OpCtz32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz32 <t> x)
	// result: (SUB (MOVDconst [64]) (FLOGR (MOVWZreg (ANDW <t> (SUBWconst <t> [1] x) (NOTW <t> x)))))
	for {
		t := v.Type
		x := v_0
		v.reset(OpS390XSUB)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 64
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XFLOGR, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XANDW, t)
		v4 := b.NewValue0(v.Pos, OpS390XSUBWconst, t)
		v4.AuxInt = 1
		v4.AddArg(x)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpS390XNOTW, t)
		v5.AddArg(x)
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpCtz64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz64 <t> x)
	// result: (SUB (MOVDconst [64]) (FLOGR (AND <t> (SUBconst <t> [1] x) (NOT <t> x))))
	for {
		t := v.Type
		x := v_0
		v.reset(OpS390XSUB)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 64
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XFLOGR, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpS390XAND, t)
		v3 := b.NewValue0(v.Pos, OpS390XSUBconst, t)
		v3.AuxInt = 1
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XNOT, t)
		v4.AddArg(x)
		v2.AddArg(v4)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpDiv16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 x y)
	// result: (DIVW (MOVHreg x) (MOVHreg y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XDIVW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpDiv16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (DIVWU (MOVHZreg x) (MOVHZreg y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XDIVWU)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpDiv32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32 x y)
	// result: (DIVW (MOVWreg x) y)
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XDIVW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpDiv32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32u x y)
	// result: (DIVWU (MOVWZreg x) y)
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XDIVWU)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpDiv64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div64 [a] x y)
	// result: (DIVD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XDIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpDiv8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (DIVW (MOVBreg x) (MOVBreg y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XDIVW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpDiv8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (DIVWU (MOVBZreg x) (MOVBZreg y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XDIVWU)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpEq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// result: (LOCGR {s390x.Equal} (MOVDconst [0]) (MOVDconst [1]) (CMPW (MOVHreg x) (MOVHreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Equal
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpEq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq32 x y)
	// result: (LOCGR {s390x.Equal} (MOVDconst [0]) (MOVDconst [1]) (CMPW x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Equal
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpEq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq32F x y)
	// result: (LOCGR {s390x.Equal} (MOVDconst [0]) (MOVDconst [1]) (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Equal
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMPS, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpEq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq64 x y)
	// result: (LOCGR {s390x.Equal} (MOVDconst [0]) (MOVDconst [1]) (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Equal
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpEq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq64F x y)
	// result: (LOCGR {s390x.Equal} (MOVDconst [0]) (MOVDconst [1]) (FCMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Equal
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMP, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpEq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// result: (LOCGR {s390x.Equal} (MOVDconst [0]) (MOVDconst [1]) (CMPW (MOVBreg x) (MOVBreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Equal
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpEqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqB x y)
	// result: (LOCGR {s390x.Equal} (MOVDconst [0]) (MOVDconst [1]) (CMPW (MOVBreg x) (MOVBreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Equal
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpEqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqPtr x y)
	// result: (LOCGR {s390x.Equal} (MOVDconst [0]) (MOVDconst [1]) (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Equal
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpFMA(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMA x y z)
	// result: (FMADD z x y)
	for {
		x := v_0
		y := v_1
		z := v_2
		v.reset(OpS390XFMADD)
		v.AddArg(z)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpFloor(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Floor x)
	// result: (FIDBR [7] x)
	for {
		x := v_0
		v.reset(OpS390XFIDBR)
		v.AuxInt = 7
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpGeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Geq32F x y)
	// result: (LOCGR {s390x.GreaterOrEqual} (MOVDconst [0]) (MOVDconst [1]) (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMPS, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpGeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Geq64F x y)
	// result: (LOCGR {s390x.GreaterOrEqual} (MOVDconst [0]) (MOVDconst [1]) (FCMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMP, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpGreater32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Greater32F x y)
	// result: (LOCGR {s390x.Greater} (MOVDconst [0]) (MOVDconst [1]) (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Greater
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMPS, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpGreater64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Greater64F x y)
	// result: (LOCGR {s390x.Greater} (MOVDconst [0]) (MOVDconst [1]) (FCMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Greater
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMP, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpHmul32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32 x y)
	// result: (SRDconst [32] (MULLD (MOVWreg x) (MOVWreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRDconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpS390XMULLD, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWreg, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XMOVWreg, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpHmul32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32u x y)
	// result: (SRDconst [32] (MULLD (MOVWZreg x) (MOVWZreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRDconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpS390XMULLD, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpITab(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ITab (Load ptr mem))
	// result: (MOVDload ptr mem)
	for {
		if v_0.Op != OpLoad {
			break
		}
		mem := v_0.Args[1]
		ptr := v_0.Args[0]
		v.reset(OpS390XMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpIsInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (IsInBounds idx len)
	// result: (LOCGR {s390x.Less} (MOVDconst [0]) (MOVDconst [1]) (CMPU idx len))
	for {
		idx := v_0
		len := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Less
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPU, types.TypeFlags)
		v2.AddArg(idx)
		v2.AddArg(len)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpIsNonNil(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (IsNonNil p)
	// result: (LOCGR {s390x.NotEqual} (MOVDconst [0]) (MOVDconst [1]) (CMPconst p [0]))
	for {
		p := v_0
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.NotEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPconst, types.TypeFlags)
		v2.AuxInt = 0
		v2.AddArg(p)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpIsSliceInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (IsSliceInBounds idx len)
	// result: (LOCGR {s390x.LessOrEqual} (MOVDconst [0]) (MOVDconst [1]) (CMPU idx len))
	for {
		idx := v_0
		len := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.LessOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPU, types.TypeFlags)
		v2.AddArg(idx)
		v2.AddArg(len)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// result: (LOCGR {s390x.LessOrEqual} (MOVDconst [0]) (MOVDconst [1]) (CMPW (MOVHreg x) (MOVHreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.LessOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// result: (LOCGR {s390x.LessOrEqual} (MOVDconst [0]) (MOVDconst [1]) (CMPWU (MOVHZreg x) (MOVHZreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.LessOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32 x y)
	// result: (LOCGR {s390x.LessOrEqual} (MOVDconst [0]) (MOVDconst [1]) (CMPW x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.LessOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32F x y)
	// result: (LOCGR {s390x.LessOrEqual} (MOVDconst [0]) (MOVDconst [1]) (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.LessOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMPS, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32U x y)
	// result: (LOCGR {s390x.LessOrEqual} (MOVDconst [0]) (MOVDconst [1]) (CMPWU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.LessOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLeq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq64 x y)
	// result: (LOCGR {s390x.LessOrEqual} (MOVDconst [0]) (MOVDconst [1]) (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.LessOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq64F x y)
	// result: (LOCGR {s390x.LessOrEqual} (MOVDconst [0]) (MOVDconst [1]) (FCMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.LessOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMP, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLeq64U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq64U x y)
	// result: (LOCGR {s390x.LessOrEqual} (MOVDconst [0]) (MOVDconst [1]) (CMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.LessOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPU, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// result: (LOCGR {s390x.LessOrEqual} (MOVDconst [0]) (MOVDconst [1]) (CMPW (MOVBreg x) (MOVBreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.LessOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// result: (LOCGR {s390x.LessOrEqual} (MOVDconst [0]) (MOVDconst [1]) (CMPWU (MOVBZreg x) (MOVBZreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.LessOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLess16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// result: (LOCGR {s390x.Less} (MOVDconst [0]) (MOVDconst [1]) (CMPW (MOVHreg x) (MOVHreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Less
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLess16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// result: (LOCGR {s390x.Less} (MOVDconst [0]) (MOVDconst [1]) (CMPWU (MOVHZreg x) (MOVHZreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Less
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLess32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32 x y)
	// result: (LOCGR {s390x.Less} (MOVDconst [0]) (MOVDconst [1]) (CMPW x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Less
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLess32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32F x y)
	// result: (LOCGR {s390x.Less} (MOVDconst [0]) (MOVDconst [1]) (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Less
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMPS, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLess32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32U x y)
	// result: (LOCGR {s390x.Less} (MOVDconst [0]) (MOVDconst [1]) (CMPWU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Less
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLess64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less64 x y)
	// result: (LOCGR {s390x.Less} (MOVDconst [0]) (MOVDconst [1]) (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Less
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLess64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less64F x y)
	// result: (LOCGR {s390x.Less} (MOVDconst [0]) (MOVDconst [1]) (FCMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Less
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMP, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLess64U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less64U x y)
	// result: (LOCGR {s390x.Less} (MOVDconst [0]) (MOVDconst [1]) (CMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Less
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPU, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLess8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// result: (LOCGR {s390x.Less} (MOVDconst [0]) (MOVDconst [1]) (CMPW (MOVBreg x) (MOVBreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Less
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLess8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// result: (LOCGR {s390x.Less} (MOVDconst [0]) (MOVDconst [1]) (CMPWU (MOVBZreg x) (MOVBZreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.Less
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLoad(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Load <t> ptr mem)
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpS390XMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitInt(t) && isSigned(t)
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpS390XMOVWload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitInt(t) && !isSigned(t)
	// result: (MOVWZload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpS390XMOVWZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is16BitInt(t) && isSigned(t)
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpS390XMOVHload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is16BitInt(t) && !isSigned(t)
	// result: (MOVHZload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpS390XMOVHZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is8BitInt(t) && isSigned(t)
	// result: (MOVBload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpS390XMOVBload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (t.IsBoolean() || (is8BitInt(t) && !isSigned(t)))
	// result: (MOVBZload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsBoolean() || (is8BitInt(t) && !isSigned(t))) {
			break
		}
		v.reset(OpS390XMOVBZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (FMOVSload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpS390XFMOVSload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (FMOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpS390XFMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpLocalAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LocalAddr {sym} base _)
	// result: (MOVDaddr {sym} base)
	for {
		sym := v.Aux
		base := v_0
		v.reset(OpS390XMOVDaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueS390X_OpLsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh16x16 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLW <t> x y) (MOVDconst [0]) (CMPWUconst (MOVHZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh16x32 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLW <t> x y) (MOVDconst [0]) (CMPWUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh16x64 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLW <t> x y) (MOVDconst [0]) (CMPUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh16x8 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLW <t> x y) (MOVDconst [0]) (CMPWUconst (MOVBZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh32x16 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLW <t> x y) (MOVDconst [0]) (CMPWUconst (MOVHZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh32x32 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLW <t> x y) (MOVDconst [0]) (CMPWUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh32x64 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLW <t> x y) (MOVDconst [0]) (CMPUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh32x8 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLW <t> x y) (MOVDconst [0]) (CMPWUconst (MOVBZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SLD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh64x16 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLD <t> x y) (MOVDconst [0]) (CMPWUconst (MOVHZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SLD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh64x32 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLD <t> x y) (MOVDconst [0]) (CMPWUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SLD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh64x64 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLD <t> x y) (MOVDconst [0]) (CMPUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SLD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh64x8 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLD <t> x y) (MOVDconst [0]) (CMPWUconst (MOVBZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh8x16 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLW <t> x y) (MOVDconst [0]) (CMPWUconst (MOVHZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh8x32 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLW <t> x y) (MOVDconst [0]) (CMPWUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh8x64 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLW <t> x y) (MOVDconst [0]) (CMPUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpLsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh8x8 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SLW <t> x y) (MOVDconst [0]) (CMPWUconst (MOVBZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpMod16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 x y)
	// result: (MODW (MOVHreg x) (MOVHreg y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XMODW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpMod16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (MODWU (MOVHZreg x) (MOVHZreg y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XMODWU)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpMod32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32 x y)
	// result: (MODW (MOVWreg x) y)
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XMODW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpMod32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32u x y)
	// result: (MODWU (MOVWZreg x) y)
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XMODWU)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpMod64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Mod64 [a] x y)
	// result: (MODD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XMODD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpMod8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (MODW (MOVBreg x) (MOVBreg y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XMODW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpMod8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (MODWU (MOVBZreg x) (MOVBZreg y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XMODWU)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpMove(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// result: (MOVBstore dst (MOVBZload src mem) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpS390XMOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZload, typ.UInt8)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// result: (MOVHstore dst (MOVHZload src mem) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpS390XMOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [4] dst src mem)
	// result: (MOVWstore dst (MOVWZload src mem) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpS390XMOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [8] dst src mem)
	// result: (MOVDstore dst (MOVDload src mem) mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpS390XMOVDstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [16] dst src mem)
	// result: (MOVDstore [8] dst (MOVDload [8] src mem) (MOVDstore dst (MOVDload src mem) mem))
	for {
		if v.AuxInt != 16 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpS390XMOVDstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [24] dst src mem)
	// result: (MOVDstore [16] dst (MOVDload [16] src mem) (MOVDstore [8] dst (MOVDload [8] src mem) (MOVDstore dst (MOVDload src mem) mem)))
	for {
		if v.AuxInt != 24 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpS390XMOVDstore)
		v.AuxInt = 16
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
		v0.AuxInt = 16
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDstore, types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
		v2.AuxInt = 8
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XMOVDstore, types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [3] dst src mem)
	// result: (MOVBstore [2] dst (MOVBZload [2] src mem) (MOVHstore dst (MOVHZload src mem) mem))
	for {
		if v.AuxInt != 3 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpS390XMOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZload, typ.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [5] dst src mem)
	// result: (MOVBstore [4] dst (MOVBZload [4] src mem) (MOVWstore dst (MOVWZload src mem) mem))
	for {
		if v.AuxInt != 5 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpS390XMOVBstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZload, typ.UInt8)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [6] dst src mem)
	// result: (MOVHstore [4] dst (MOVHZload [4] src mem) (MOVWstore dst (MOVWZload src mem) mem))
	for {
		if v.AuxInt != 6 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpS390XMOVHstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [7] dst src mem)
	// result: (MOVBstore [6] dst (MOVBZload [6] src mem) (MOVHstore [4] dst (MOVHZload [4] src mem) (MOVWstore dst (MOVWZload src mem) mem)))
	for {
		if v.AuxInt != 7 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpS390XMOVBstore)
		v.AuxInt = 6
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZload, typ.UInt8)
		v0.AuxInt = 6
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHstore, types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWstore, types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: s > 0 && s <= 256
	// result: (MVC [makeValAndOff(s, 0)] dst src mem)
	for {
		s := v.AuxInt
		dst := v_0
		src := v_1
		mem := v_2
		if !(s > 0 && s <= 256) {
			break
		}
		v.reset(OpS390XMVC)
		v.AuxInt = makeValAndOff(s, 0)
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: s > 256 && s <= 512
	// result: (MVC [makeValAndOff(s-256, 256)] dst src (MVC [makeValAndOff(256, 0)] dst src mem))
	for {
		s := v.AuxInt
		dst := v_0
		src := v_1
		mem := v_2
		if !(s > 256 && s <= 512) {
			break
		}
		v.reset(OpS390XMVC)
		v.AuxInt = makeValAndOff(s-256, 256)
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpS390XMVC, types.TypeMem)
		v0.AuxInt = makeValAndOff(256, 0)
		v0.AddArg(dst)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: s > 512 && s <= 768
	// result: (MVC [makeValAndOff(s-512, 512)] dst src (MVC [makeValAndOff(256, 256)] dst src (MVC [makeValAndOff(256, 0)] dst src mem)))
	for {
		s := v.AuxInt
		dst := v_0
		src := v_1
		mem := v_2
		if !(s > 512 && s <= 768) {
			break
		}
		v.reset(OpS390XMVC)
		v.AuxInt = makeValAndOff(s-512, 512)
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpS390XMVC, types.TypeMem)
		v0.AuxInt = makeValAndOff(256, 256)
		v0.AddArg(dst)
		v0.AddArg(src)
		v1 := b.NewValue0(v.Pos, OpS390XMVC, types.TypeMem)
		v1.AuxInt = makeValAndOff(256, 0)
		v1.AddArg(dst)
		v1.AddArg(src)
		v1.AddArg(mem)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: s > 768 && s <= 1024
	// result: (MVC [makeValAndOff(s-768, 768)] dst src (MVC [makeValAndOff(256, 512)] dst src (MVC [makeValAndOff(256, 256)] dst src (MVC [makeValAndOff(256, 0)] dst src mem))))
	for {
		s := v.AuxInt
		dst := v_0
		src := v_1
		mem := v_2
		if !(s > 768 && s <= 1024) {
			break
		}
		v.reset(OpS390XMVC)
		v.AuxInt = makeValAndOff(s-768, 768)
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpS390XMVC, types.TypeMem)
		v0.AuxInt = makeValAndOff(256, 512)
		v0.AddArg(dst)
		v0.AddArg(src)
		v1 := b.NewValue0(v.Pos, OpS390XMVC, types.TypeMem)
		v1.AuxInt = makeValAndOff(256, 256)
		v1.AddArg(dst)
		v1.AddArg(src)
		v2 := b.NewValue0(v.Pos, OpS390XMVC, types.TypeMem)
		v2.AuxInt = makeValAndOff(256, 0)
		v2.AddArg(dst)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: s > 1024
	// result: (LoweredMove [s%256] dst src (ADD <src.Type> src (MOVDconst [(s/256)*256])) mem)
	for {
		s := v.AuxInt
		dst := v_0
		src := v_1
		mem := v_2
		if !(s > 1024) {
			break
		}
		v.reset(OpS390XLoweredMove)
		v.AuxInt = s % 256
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpS390XADD, src.Type)
		v0.AddArg(src)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = (s / 256) * 256
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpNeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// result: (LOCGR {s390x.NotEqual} (MOVDconst [0]) (MOVDconst [1]) (CMPW (MOVHreg x) (MOVHreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.NotEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpNeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq32 x y)
	// result: (LOCGR {s390x.NotEqual} (MOVDconst [0]) (MOVDconst [1]) (CMPW x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.NotEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpNeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq32F x y)
	// result: (LOCGR {s390x.NotEqual} (MOVDconst [0]) (MOVDconst [1]) (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.NotEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMPS, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpNeq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq64 x y)
	// result: (LOCGR {s390x.NotEqual} (MOVDconst [0]) (MOVDconst [1]) (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.NotEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpNeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq64F x y)
	// result: (LOCGR {s390x.NotEqual} (MOVDconst [0]) (MOVDconst [1]) (FCMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.NotEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMP, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpNeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// result: (LOCGR {s390x.NotEqual} (MOVDconst [0]) (MOVDconst [1]) (CMPW (MOVBreg x) (MOVBreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.NotEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpNeqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (NeqB x y)
	// result: (LOCGR {s390x.NotEqual} (MOVDconst [0]) (MOVDconst [1]) (CMPW (MOVBreg x) (MOVBreg y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.NotEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpNeqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (NeqPtr x y)
	// result: (LOCGR {s390x.NotEqual} (MOVDconst [0]) (MOVDconst [1]) (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Aux = s390x.NotEqual
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpNot(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Not x)
	// result: (XORWconst [1] x)
	for {
		x := v_0
		v.reset(OpS390XXORWconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpOffPtr(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (OffPtr [off] ptr:(SP))
	// result: (MOVDaddr [off] ptr)
	for {
		off := v.AuxInt
		ptr := v_0
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpS390XMOVDaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// cond: is32Bit(off)
	// result: (ADDconst [off] ptr)
	for {
		off := v.AuxInt
		ptr := v_0
		if !(is32Bit(off)) {
			break
		}
		v.reset(OpS390XADDconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// result: (ADD (MOVDconst [off]) ptr)
	for {
		off := v.AuxInt
		ptr := v_0
		v.reset(OpS390XADD)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = off
		v.AddArg(v0)
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueS390X_OpPanicBounds(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 0
	// result: (LoweredPanicBoundsA [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 0) {
			break
		}
		v.reset(OpS390XLoweredPanicBoundsA)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 1
	// result: (LoweredPanicBoundsB [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 1) {
			break
		}
		v.reset(OpS390XLoweredPanicBoundsB)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 2
	// result: (LoweredPanicBoundsC [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 2) {
			break
		}
		v.reset(OpS390XLoweredPanicBoundsC)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpPopCount16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount16 x)
	// result: (MOVBZreg (SumBytes2 (POPCNT <typ.UInt16> x)))
	for {
		x := v_0
		v.reset(OpS390XMOVBZreg)
		v0 := b.NewValue0(v.Pos, OpS390XSumBytes2, typ.UInt8)
		v1 := b.NewValue0(v.Pos, OpS390XPOPCNT, typ.UInt16)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpPopCount32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount32 x)
	// result: (MOVBZreg (SumBytes4 (POPCNT <typ.UInt32> x)))
	for {
		x := v_0
		v.reset(OpS390XMOVBZreg)
		v0 := b.NewValue0(v.Pos, OpS390XSumBytes4, typ.UInt8)
		v1 := b.NewValue0(v.Pos, OpS390XPOPCNT, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpPopCount64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount64 x)
	// result: (MOVBZreg (SumBytes8 (POPCNT <typ.UInt64> x)))
	for {
		x := v_0
		v.reset(OpS390XMOVBZreg)
		v0 := b.NewValue0(v.Pos, OpS390XSumBytes8, typ.UInt8)
		v1 := b.NewValue0(v.Pos, OpS390XPOPCNT, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpPopCount8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount8 x)
	// result: (POPCNT (MOVBZreg x))
	for {
		x := v_0
		v.reset(OpS390XPOPCNT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpRotateLeft16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft16 <t> x (MOVDconst [c]))
	// result: (Or16 (Lsh16x64 <t> x (MOVDconst [c&15])) (Rsh16Ux64 <t> x (MOVDconst [-c&15])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpLsh16x64, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = c & 15
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRsh16Ux64, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v3.AuxInt = -c & 15
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueS390X_OpRotateLeft8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft8 <t> x (MOVDconst [c]))
	// result: (Or8 (Lsh8x64 <t> x (MOVDconst [c&7])) (Rsh8Ux64 <t> x (MOVDconst [-c&7])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpLsh8x64, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = c & 7
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRsh8Ux64, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v3.AuxInt = -c & 7
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueS390X_OpRound(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Round x)
	// result: (FIDBR [1] x)
	for {
		x := v_0
		v.reset(OpS390XFIDBR)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpRoundToEven(v *Value) bool {
	v_0 := v.Args[0]
	// match: (RoundToEven x)
	// result: (FIDBR [4] x)
	for {
		x := v_0
		v.reset(OpS390XFIDBR)
		v.AuxInt = 4
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpRsh16Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVHZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16Ux16 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRW <t> (MOVHZreg x) y) (MOVDconst [0]) (CMPWUconst (MOVHZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueS390X_OpRsh16Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVHZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16Ux32 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRW <t> (MOVHZreg x) y) (MOVDconst [0]) (CMPWUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueS390X_OpRsh16Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVHZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16Ux64 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRW <t> (MOVHZreg x) y) (MOVDconst [0]) (CMPUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPUconst, types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueS390X_OpRsh16Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVHZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16Ux8 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRW <t> (MOVHZreg x) y) (MOVDconst [0]) (CMPWUconst (MOVBZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueS390X_OpRsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVHreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16x16 x y)
	// result: (SRAW (MOVHreg x) (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPWUconst (MOVHZreg y) [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v1.Aux = s390x.GreaterOrEqual
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpRsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVHreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16x32 x y)
	// result: (SRAW (MOVHreg x) (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPWUconst y [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v1.Aux = s390x.GreaterOrEqual
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpRsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVHreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16x64 x y)
	// result: (SRAW (MOVHreg x) (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPUconst y [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v1.Aux = s390x.GreaterOrEqual
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPUconst, types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpRsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVHreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16x8 x y)
	// result: (SRAW (MOVHreg x) (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPWUconst (MOVBZreg y) [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v1.Aux = s390x.GreaterOrEqual
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpRsh32Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32Ux16 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRW <t> x y) (MOVDconst [0]) (CMPWUconst (MOVHZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpRsh32Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32Ux32 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRW <t> x y) (MOVDconst [0]) (CMPWUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpRsh32Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32Ux64 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRW <t> x y) (MOVDconst [0]) (CMPUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpRsh32Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32Ux8 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRW <t> x y) (MOVDconst [0]) (CMPWUconst (MOVBZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpRsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32x16 x y)
	// result: (SRAW x (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPWUconst (MOVHZreg y) [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v0.Aux = s390x.GreaterOrEqual
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpRsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh32x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32x32 x y)
	// result: (SRAW x (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPWUconst y [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v0.Aux = s390x.GreaterOrEqual
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpRsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh32x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32x64 x y)
	// result: (SRAW x (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPUconst y [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v0.Aux = s390x.GreaterOrEqual
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpRsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32x8 x y)
	// result: (SRAW x (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPWUconst (MOVBZreg y) [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v0.Aux = s390x.GreaterOrEqual
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpRsh64Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64Ux16 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRD <t> x y) (MOVDconst [0]) (CMPWUconst (MOVHZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpRsh64Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64Ux32 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRD <t> x y) (MOVDconst [0]) (CMPWUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpRsh64Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64Ux64 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRD <t> x y) (MOVDconst [0]) (CMPUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpRsh64Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64Ux8 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRD <t> x y) (MOVDconst [0]) (CMPWUconst (MOVBZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpRsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64x16 x y)
	// result: (SRAD x (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPWUconst (MOVHZreg y) [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v0.Aux = s390x.GreaterOrEqual
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpRsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh64x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64x32 x y)
	// result: (SRAD x (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPWUconst y [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v0.Aux = s390x.GreaterOrEqual
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpRsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh64x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64x64 x y)
	// result: (SRAD x (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPUconst y [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v0.Aux = s390x.GreaterOrEqual
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpRsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64x8 x y)
	// result: (SRAD x (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPWUconst (MOVBZreg y) [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v0.Aux = s390x.GreaterOrEqual
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpRsh8Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVBZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8Ux16 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRW <t> (MOVBZreg x) y) (MOVDconst [0]) (CMPWUconst (MOVHZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueS390X_OpRsh8Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVBZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8Ux32 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRW <t> (MOVBZreg x) y) (MOVDconst [0]) (CMPWUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueS390X_OpRsh8Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVBZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8Ux64 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRW <t> (MOVBZreg x) y) (MOVDconst [0]) (CMPUconst y [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPUconst, types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueS390X_OpRsh8Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVBZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8Ux8 <t> x y)
	// result: (LOCGR {s390x.GreaterOrEqual} <t> (SRW <t> (MOVBZreg x) y) (MOVDconst [0]) (CMPWUconst (MOVBZreg y) [64]))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpS390XLOCGR)
		v.Type = t
		v.Aux = s390x.GreaterOrEqual
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueS390X_OpRsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVBreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8x16 x y)
	// result: (SRAW (MOVBreg x) (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPWUconst (MOVHZreg y) [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v1.Aux = s390x.GreaterOrEqual
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpRsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVBreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8x32 x y)
	// result: (SRAW (MOVBreg x) (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPWUconst y [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v1.Aux = s390x.GreaterOrEqual
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpRsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVBreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8x64 x y)
	// result: (SRAW (MOVBreg x) (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPUconst y [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v1.Aux = s390x.GreaterOrEqual
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPUconst, types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpRsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVBreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8x8 x y)
	// result: (SRAW (MOVBreg x) (LOCGR {s390x.GreaterOrEqual} <y.Type> y (MOVDconst <y.Type> [63]) (CMPWUconst (MOVBZreg y) [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XLOCGR, y.Type)
		v1.Aux = s390x.GreaterOrEqual
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpS390XADD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADD x (MOVDconst [c]))
	// cond: is32Bit(c)
	// result: (ADDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpS390XADDconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADD (SLDconst x [c]) (SRDconst x [d]))
	// cond: d == 64-c
	// result: (RLLGconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XSLDconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			if v_1.Op != OpS390XSRDconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 64-c) {
				continue
			}
			v.reset(OpS390XRLLGconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADD idx (MOVDaddr [c] {s} ptr))
	// cond: ptr.Op != OpSB && idx.Op != OpSB
	// result: (MOVDaddridx [c] {s} ptr idx)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			idx := v_0
			if v_1.Op != OpS390XMOVDaddr {
				continue
			}
			c := v_1.AuxInt
			s := v_1.Aux
			ptr := v_1.Args[0]
			if !(ptr.Op != OpSB && idx.Op != OpSB) {
				continue
			}
			v.reset(OpS390XMOVDaddridx)
			v.AuxInt = c
			v.Aux = s
			v.AddArg(ptr)
			v.AddArg(idx)
			return true
		}
		break
	}
	// match: (ADD x (NEG y))
	// result: (SUB x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XNEG {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpS390XSUB)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADD <t> x g:(MOVDload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (ADDload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVDload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XADDload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XADDC(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADDC x (MOVDconst [c]))
	// cond: is16Bit(c)
	// result: (ADDCconst x [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1.AuxInt
			if !(is16Bit(c)) {
				continue
			}
			v.reset(OpS390XADDCconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XADDE(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADDE x y (FlagEQ))
	// result: (ADDC x y)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpS390XFlagEQ {
			break
		}
		v.reset(OpS390XADDC)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDE x y (FlagLT))
	// result: (ADDC x y)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpS390XFlagLT {
			break
		}
		v.reset(OpS390XADDC)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDE x y (Select1 (ADDCconst [-1] (Select0 (ADDE (MOVDconst [0]) (MOVDconst [0]) c)))))
	// result: (ADDE x y c)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpSelect1 {
			break
		}
		v_2_0 := v_2.Args[0]
		if v_2_0.Op != OpS390XADDCconst || v_2_0.AuxInt != -1 {
			break
		}
		v_2_0_0 := v_2_0.Args[0]
		if v_2_0_0.Op != OpSelect0 {
			break
		}
		v_2_0_0_0 := v_2_0_0.Args[0]
		if v_2_0_0_0.Op != OpS390XADDE {
			break
		}
		c := v_2_0_0_0.Args[2]
		v_2_0_0_0_0 := v_2_0_0_0.Args[0]
		if v_2_0_0_0_0.Op != OpS390XMOVDconst || v_2_0_0_0_0.AuxInt != 0 {
			break
		}
		v_2_0_0_0_1 := v_2_0_0_0.Args[1]
		if v_2_0_0_0_1.Op != OpS390XMOVDconst || v_2_0_0_0_1.AuxInt != 0 {
			break
		}
		v.reset(OpS390XADDE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(c)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XADDW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADDW x (MOVDconst [c]))
	// result: (ADDWconst [int64(int32(c))] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpS390XADDWconst)
			v.AuxInt = int64(int32(c))
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADDW (SLWconst x [c]) (SRWconst x [d]))
	// cond: d == 32-c
	// result: (RLLconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XSLWconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			if v_1.Op != OpS390XSRWconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 32-c) {
				continue
			}
			v.reset(OpS390XRLLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADDW x (NEGW y))
	// result: (SUBW x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XNEGW {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpS390XSUBW)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDW <t> x g:(MOVWload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (ADDWload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVWload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XADDWload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (ADDW <t> x g:(MOVWZload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (ADDWload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVWZload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XADDWload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XADDWconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ADDWconst [c] x)
	// cond: int32(c)==0
	// result: x
	for {
		c := v.AuxInt
		x := v_0
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ADDWconst [c] (MOVDconst [d]))
	// result: (MOVDconst [int64(int32(c+d))])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int32(c + d))
		return true
	}
	// match: (ADDWconst [c] (ADDWconst [d] x))
	// result: (ADDWconst [int64(int32(c+d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XADDWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpS390XADDWconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XADDWload(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADDWload [off1] {sym} x (ADDconst [off2] ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(off1+off2)
	// result: (ADDWload [off1+off2] {sym} x ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XADDWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (ADDWload [o1] {s1} x (MOVDaddr [o2] {s2} ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)
	// result: (ADDWload [o1+o2] {mergeSym(s1, s2)} x ptr mem)
	for {
		o1 := v.AuxInt
		s1 := v.Aux
		x := v_0
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XADDWload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XADDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ADDconst [c] (MOVDaddr [d] {s} x:(SB)))
	// cond: ((c+d)&1 == 0) && is32Bit(c+d)
	// result: (MOVDaddr [c+d] {s} x)
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		x := v_0.Args[0]
		if x.Op != OpSB || !(((c+d)&1 == 0) && is32Bit(c+d)) {
			break
		}
		v.reset(OpS390XMOVDaddr)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (MOVDaddr [d] {s} x))
	// cond: x.Op != OpSB && is20Bit(c+d)
	// result: (MOVDaddr [c+d] {s} x)
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		x := v_0.Args[0]
		if !(x.Op != OpSB && is20Bit(c+d)) {
			break
		}
		v.reset(OpS390XMOVDaddr)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (MOVDaddridx [d] {s} x y))
	// cond: is20Bit(c+d)
	// result: (MOVDaddridx [c+d] {s} x y)
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVDaddridx)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (MOVDconst [d]))
	// result: (MOVDconst [c+d])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c + d
		return true
	}
	// match: (ADDconst [c] (ADDconst [d] x))
	// cond: is32Bit(c+d)
	// result: (ADDconst [c+d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpS390XADDconst)
		v.AuxInt = c + d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XADDload(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDload <t> [off] {sym} x ptr1 (FMOVDstore [off] {sym} ptr2 y _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (ADD x (LGDR <t> y))
	for {
		t := v.Type
		off := v.AuxInt
		sym := v.Aux
		x := v_0
		ptr1 := v_1
		if v_2.Op != OpS390XFMOVDstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		y := v_2.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XADD)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpS390XLGDR, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ADDload [off1] {sym} x (ADDconst [off2] ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(off1+off2)
	// result: (ADDload [off1+off2] {sym} x ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XADDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (ADDload [o1] {s1} x (MOVDaddr [o2] {s2} ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)
	// result: (ADDload [o1+o2] {mergeSym(s1, s2)} x ptr mem)
	for {
		o1 := v.AuxInt
		s1 := v.Aux
		x := v_0
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XADDload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XAND(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (AND x (MOVDconst [c]))
	// cond: is32Bit(c) && c < 0
	// result: (ANDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c) && c < 0) {
				continue
			}
			v.reset(OpS390XANDconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (AND x (MOVDconst [c]))
	// cond: is32Bit(c) && c >= 0
	// result: (MOVWZreg (ANDWconst <typ.UInt32> [int64(int32(c))] x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c) && c >= 0) {
				continue
			}
			v.reset(OpS390XMOVWZreg)
			v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
			v0.AuxInt = int64(int32(c))
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (AND x (MOVDconst [0xFF]))
	// result: (MOVBZreg x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst || v_1.AuxInt != 0xFF {
				continue
			}
			v.reset(OpS390XMOVBZreg)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (AND x (MOVDconst [0xFFFF]))
	// result: (MOVHZreg x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst || v_1.AuxInt != 0xFFFF {
				continue
			}
			v.reset(OpS390XMOVHZreg)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (AND x (MOVDconst [0xFFFFFFFF]))
	// result: (MOVWZreg x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst || v_1.AuxInt != 0xFFFFFFFF {
				continue
			}
			v.reset(OpS390XMOVWZreg)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (AND (MOVDconst [^(-1<<63)]) (LGDR <t> x))
	// result: (LGDR <t> (LPDFR <x.Type> x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XMOVDconst || v_0.AuxInt != ^(-1<<63) || v_1.Op != OpS390XLGDR {
				continue
			}
			t := v_1.Type
			x := v_1.Args[0]
			v.reset(OpS390XLGDR)
			v.Type = t
			v0 := b.NewValue0(v.Pos, OpS390XLPDFR, x.Type)
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (AND (MOVDconst [c]) (MOVDconst [d]))
	// result: (MOVDconst [c&d])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XMOVDconst {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpS390XMOVDconst)
			v.AuxInt = c & d
			return true
		}
		break
	}
	// match: (AND x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (AND <t> x g:(MOVDload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (ANDload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVDload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XANDload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XANDW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ANDW x (MOVDconst [c]))
	// result: (ANDWconst [int64(int32(c))] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpS390XANDWconst)
			v.AuxInt = int64(int32(c))
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ANDW x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ANDW <t> x g:(MOVWload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (ANDWload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVWload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XANDWload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (ANDW <t> x g:(MOVWZload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (ANDWload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVWZload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XANDWload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XANDWconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ANDWconst [c] (ANDWconst [d] x))
	// result: (ANDWconst [c & d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XANDWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpS390XANDWconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	// match: (ANDWconst [0xFF] x)
	// result: (MOVBZreg x)
	for {
		if v.AuxInt != 0xFF {
			break
		}
		x := v_0
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}
	// match: (ANDWconst [0xFFFF] x)
	// result: (MOVHZreg x)
	for {
		if v.AuxInt != 0xFFFF {
			break
		}
		x := v_0
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}
	// match: (ANDWconst [c] _)
	// cond: int32(c)==0
	// result: (MOVDconst [0])
	for {
		c := v.AuxInt
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (ANDWconst [c] x)
	// cond: int32(c)==-1
	// result: x
	for {
		c := v.AuxInt
		x := v_0
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ANDWconst [c] (MOVDconst [d]))
	// result: (MOVDconst [c&d])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c & d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XANDWload(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ANDWload [off1] {sym} x (ADDconst [off2] ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(off1+off2)
	// result: (ANDWload [off1+off2] {sym} x ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XANDWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (ANDWload [o1] {s1} x (MOVDaddr [o2] {s2} ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)
	// result: (ANDWload [o1+o2] {mergeSym(s1, s2)} x ptr mem)
	for {
		o1 := v.AuxInt
		s1 := v.Aux
		x := v_0
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XANDWload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XANDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ANDconst [c] (ANDconst [d] x))
	// result: (ANDconst [c & d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpS390XANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [0] _)
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (ANDconst [-1] x)
	// result: x
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] (MOVDconst [d]))
	// result: (MOVDconst [c&d])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c & d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XANDload(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ANDload <t> [off] {sym} x ptr1 (FMOVDstore [off] {sym} ptr2 y _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (AND x (LGDR <t> y))
	for {
		t := v.Type
		off := v.AuxInt
		sym := v.Aux
		x := v_0
		ptr1 := v_1
		if v_2.Op != OpS390XFMOVDstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		y := v_2.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XAND)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpS390XLGDR, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ANDload [off1] {sym} x (ADDconst [off2] ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(off1+off2)
	// result: (ANDload [off1+off2] {sym} x ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XANDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (ANDload [o1] {s1} x (MOVDaddr [o2] {s2} ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)
	// result: (ANDload [o1+o2] {mergeSym(s1, s2)} x ptr mem)
	for {
		o1 := v.AuxInt
		s1 := v.Aux
		x := v_0
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XANDload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMP(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMP x (MOVDconst [c]))
	// cond: is32Bit(c)
	// result: (CMPconst x [c])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpS390XCMPconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMP (MOVDconst [c]) x)
	// cond: is32Bit(c)
	// result: (InvertFlags (CMPconst x [c]))
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpS390XInvertFlags)
		v0 := b.NewValue0(v.Pos, OpS390XCMPconst, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x y)
	// cond: x.ID > y.ID
	// result: (InvertFlags (CMP y x))
	for {
		x := v_0
		y := v_1
		if !(x.ID > y.ID) {
			break
		}
		v.reset(OpS390XInvertFlags)
		v0 := b.NewValue0(v.Pos, OpS390XCMP, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMPU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPU x (MOVDconst [c]))
	// cond: isU32Bit(c)
	// result: (CMPUconst x [int64(int32(c))])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpS390XCMPUconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (CMPU (MOVDconst [c]) x)
	// cond: isU32Bit(c)
	// result: (InvertFlags (CMPUconst x [int64(int32(c))]))
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpS390XInvertFlags)
		v0 := b.NewValue0(v.Pos, OpS390XCMPUconst, types.TypeFlags)
		v0.AuxInt = int64(int32(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMPU x y)
	// cond: x.ID > y.ID
	// result: (InvertFlags (CMPU y x))
	for {
		x := v_0
		y := v_1
		if !(x.ID > y.ID) {
			break
		}
		v.reset(OpS390XInvertFlags)
		v0 := b.NewValue0(v.Pos, OpS390XCMPU, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMPUconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (CMPUconst (MOVDconst [x]) [y])
	// cond: uint64(x)==uint64(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint64(x) == uint64(y)) {
			break
		}
		v.reset(OpS390XFlagEQ)
		return true
	}
	// match: (CMPUconst (MOVDconst [x]) [y])
	// cond: uint64(x)<uint64(y)
	// result: (FlagLT)
	for {
		y := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint64(x) < uint64(y)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}
	// match: (CMPUconst (MOVDconst [x]) [y])
	// cond: uint64(x)>uint64(y)
	// result: (FlagGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint64(x) > uint64(y)) {
			break
		}
		v.reset(OpS390XFlagGT)
		return true
	}
	// match: (CMPUconst (SRDconst _ [c]) [n])
	// cond: c > 0 && c < 64 && (1<<uint(64-c)) <= uint64(n)
	// result: (FlagLT)
	for {
		n := v.AuxInt
		if v_0.Op != OpS390XSRDconst {
			break
		}
		c := v_0.AuxInt
		if !(c > 0 && c < 64 && (1<<uint(64-c)) <= uint64(n)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}
	// match: (CMPUconst (MOVWZreg x) [c])
	// result: (CMPWUconst x [c])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPUconst x:(MOVHreg _) [c])
	// result: (CMPWUconst x [c])
	for {
		c := v.AuxInt
		x := v_0
		if x.Op != OpS390XMOVHreg {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPUconst x:(MOVHZreg _) [c])
	// result: (CMPWUconst x [c])
	for {
		c := v.AuxInt
		x := v_0
		if x.Op != OpS390XMOVHZreg {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPUconst x:(MOVBreg _) [c])
	// result: (CMPWUconst x [c])
	for {
		c := v.AuxInt
		x := v_0
		if x.Op != OpS390XMOVBreg {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPUconst x:(MOVBZreg _) [c])
	// result: (CMPWUconst x [c])
	for {
		c := v.AuxInt
		x := v_0
		if x.Op != OpS390XMOVBZreg {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPUconst (MOVWZreg x:(ANDWconst [m] _)) [c])
	// cond: int32(m) >= 0
	// result: (CMPWUconst x [c])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpS390XANDWconst {
			break
		}
		m := x.AuxInt
		if !(int32(m) >= 0) {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPUconst (MOVWreg x:(ANDWconst [m] _)) [c])
	// cond: int32(m) >= 0
	// result: (CMPWUconst x [c])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpS390XANDWconst {
			break
		}
		m := x.AuxInt
		if !(int32(m) >= 0) {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMPW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPW x (MOVDconst [c]))
	// result: (CMPWconst x [int64(int32(c))])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XCMPWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (CMPW (MOVDconst [c]) x)
	// result: (InvertFlags (CMPWconst x [int64(int32(c))]))
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpS390XInvertFlags)
		v0 := b.NewValue0(v.Pos, OpS390XCMPWconst, types.TypeFlags)
		v0.AuxInt = int64(int32(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMPW x y)
	// cond: x.ID > y.ID
	// result: (InvertFlags (CMPW y x))
	for {
		x := v_0
		y := v_1
		if !(x.ID > y.ID) {
			break
		}
		v.reset(OpS390XInvertFlags)
		v0 := b.NewValue0(v.Pos, OpS390XCMPW, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMPW x (MOVWreg y))
	// result: (CMPW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XCMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPW x (MOVWZreg y))
	// result: (CMPW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XCMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPW (MOVWreg x) y)
	// result: (CMPW x y)
	for {
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		y := v_1
		v.reset(OpS390XCMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPW (MOVWZreg x) y)
	// result: (CMPW x y)
	for {
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		y := v_1
		v.reset(OpS390XCMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMPWU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPWU x (MOVDconst [c]))
	// result: (CMPWUconst x [int64(int32(c))])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (CMPWU (MOVDconst [c]) x)
	// result: (InvertFlags (CMPWUconst x [int64(int32(c))]))
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpS390XInvertFlags)
		v0 := b.NewValue0(v.Pos, OpS390XCMPWUconst, types.TypeFlags)
		v0.AuxInt = int64(int32(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMPWU x y)
	// cond: x.ID > y.ID
	// result: (InvertFlags (CMPWU y x))
	for {
		x := v_0
		y := v_1
		if !(x.ID > y.ID) {
			break
		}
		v.reset(OpS390XInvertFlags)
		v0 := b.NewValue0(v.Pos, OpS390XCMPWU, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMPWU x (MOVWreg y))
	// result: (CMPWU x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XCMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPWU x (MOVWZreg y))
	// result: (CMPWU x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XCMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPWU (MOVWreg x) y)
	// result: (CMPWU x y)
	for {
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		y := v_1
		v.reset(OpS390XCMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPWU (MOVWZreg x) y)
	// result: (CMPWU x y)
	for {
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		y := v_1
		v.reset(OpS390XCMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMPWUconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (CMPWUconst (MOVDconst [x]) [y])
	// cond: uint32(x)==uint32(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint32(x) == uint32(y)) {
			break
		}
		v.reset(OpS390XFlagEQ)
		return true
	}
	// match: (CMPWUconst (MOVDconst [x]) [y])
	// cond: uint32(x)<uint32(y)
	// result: (FlagLT)
	for {
		y := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}
	// match: (CMPWUconst (MOVDconst [x]) [y])
	// cond: uint32(x)>uint32(y)
	// result: (FlagGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpS390XFlagGT)
		return true
	}
	// match: (CMPWUconst (MOVBZreg _) [c])
	// cond: 0xff < c
	// result: (FlagLT)
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVBZreg || !(0xff < c) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}
	// match: (CMPWUconst (MOVHZreg _) [c])
	// cond: 0xffff < c
	// result: (FlagLT)
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVHZreg || !(0xffff < c) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}
	// match: (CMPWUconst (SRWconst _ [c]) [n])
	// cond: c > 0 && c < 32 && (1<<uint(32-c)) <= uint32(n)
	// result: (FlagLT)
	for {
		n := v.AuxInt
		if v_0.Op != OpS390XSRWconst {
			break
		}
		c := v_0.AuxInt
		if !(c > 0 && c < 32 && (1<<uint(32-c)) <= uint32(n)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}
	// match: (CMPWUconst (ANDWconst _ [m]) [n])
	// cond: uint32(m) < uint32(n)
	// result: (FlagLT)
	for {
		n := v.AuxInt
		if v_0.Op != OpS390XANDWconst {
			break
		}
		m := v_0.AuxInt
		if !(uint32(m) < uint32(n)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}
	// match: (CMPWUconst (MOVWreg x) [c])
	// result: (CMPWUconst x [c])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPWUconst (MOVWZreg x) [c])
	// result: (CMPWUconst x [c])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMPWconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (CMPWconst (MOVDconst [x]) [y])
	// cond: int32(x)==int32(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(OpS390XFlagEQ)
		return true
	}
	// match: (CMPWconst (MOVDconst [x]) [y])
	// cond: int32(x)<int32(y)
	// result: (FlagLT)
	for {
		y := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}
	// match: (CMPWconst (MOVDconst [x]) [y])
	// cond: int32(x)>int32(y)
	// result: (FlagGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y)) {
			break
		}
		v.reset(OpS390XFlagGT)
		return true
	}
	// match: (CMPWconst (MOVBZreg _) [c])
	// cond: 0xff < c
	// result: (FlagLT)
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVBZreg || !(0xff < c) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}
	// match: (CMPWconst (MOVHZreg _) [c])
	// cond: 0xffff < c
	// result: (FlagLT)
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVHZreg || !(0xffff < c) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}
	// match: (CMPWconst (SRWconst _ [c]) [n])
	// cond: c > 0 && n < 0
	// result: (FlagGT)
	for {
		n := v.AuxInt
		if v_0.Op != OpS390XSRWconst {
			break
		}
		c := v_0.AuxInt
		if !(c > 0 && n < 0) {
			break
		}
		v.reset(OpS390XFlagGT)
		return true
	}
	// match: (CMPWconst (ANDWconst _ [m]) [n])
	// cond: int32(m) >= 0 && int32(m) < int32(n)
	// result: (FlagLT)
	for {
		n := v.AuxInt
		if v_0.Op != OpS390XANDWconst {
			break
		}
		m := v_0.AuxInt
		if !(int32(m) >= 0 && int32(m) < int32(n)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}
	// match: (CMPWconst x:(SRWconst _ [c]) [n])
	// cond: c > 0 && n >= 0
	// result: (CMPWUconst x [n])
	for {
		n := v.AuxInt
		x := v_0
		if x.Op != OpS390XSRWconst {
			break
		}
		c := x.AuxInt
		if !(c > 0 && n >= 0) {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = n
		v.AddArg(x)
		return true
	}
	// match: (CMPWconst (MOVWreg x) [c])
	// result: (CMPWconst x [c])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPWconst (MOVWZreg x) [c])
	// result: (CMPWconst x [c])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMPconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (CMPconst (MOVDconst [x]) [y])
	// cond: x==y
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x == y) {
			break
		}
		v.reset(OpS390XFlagEQ)
		return true
	}
	// match: (CMPconst (MOVDconst [x]) [y])
	// cond: x<y
	// result: (FlagLT)
	for {
		y := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x < y) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}
	// match: (CMPconst (MOVDconst [x]) [y])
	// cond: x>y
	// result: (FlagGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x > y) {
			break
		}
		v.reset(OpS390XFlagGT)
		return true
	}
	// match: (CMPconst (SRDconst _ [c]) [n])
	// cond: c > 0 && n < 0
	// result: (FlagGT)
	for {
		n := v.AuxInt
		if v_0.Op != OpS390XSRDconst {
			break
		}
		c := v_0.AuxInt
		if !(c > 0 && n < 0) {
			break
		}
		v.reset(OpS390XFlagGT)
		return true
	}
	// match: (CMPconst (MOVWreg x) [c])
	// result: (CMPWconst x [c])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPconst x:(MOVHreg _) [c])
	// result: (CMPWconst x [c])
	for {
		c := v.AuxInt
		x := v_0
		if x.Op != OpS390XMOVHreg {
			break
		}
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPconst x:(MOVHZreg _) [c])
	// result: (CMPWconst x [c])
	for {
		c := v.AuxInt
		x := v_0
		if x.Op != OpS390XMOVHZreg {
			break
		}
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPconst x:(MOVBreg _) [c])
	// result: (CMPWconst x [c])
	for {
		c := v.AuxInt
		x := v_0
		if x.Op != OpS390XMOVBreg {
			break
		}
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPconst x:(MOVBZreg _) [c])
	// result: (CMPWconst x [c])
	for {
		c := v.AuxInt
		x := v_0
		if x.Op != OpS390XMOVBZreg {
			break
		}
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPconst (MOVWZreg x:(ANDWconst [m] _)) [c])
	// cond: int32(m) >= 0 && c >= 0
	// result: (CMPWUconst x [c])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpS390XANDWconst {
			break
		}
		m := x.AuxInt
		if !(int32(m) >= 0 && c >= 0) {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPconst (MOVWreg x:(ANDWconst [m] _)) [c])
	// cond: int32(m) >= 0 && c >= 0
	// result: (CMPWUconst x [c])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpS390XANDWconst {
			break
		}
		m := x.AuxInt
		if !(int32(m) >= 0 && c >= 0) {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPconst x:(SRDconst _ [c]) [n])
	// cond: c > 0 && n >= 0
	// result: (CMPUconst x [n])
	for {
		n := v.AuxInt
		x := v_0
		if x.Op != OpS390XSRDconst {
			break
		}
		c := x.AuxInt
		if !(c > 0 && n >= 0) {
			break
		}
		v.reset(OpS390XCMPUconst)
		v.AuxInt = n
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCPSDR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CPSDR y (FMOVDconst [c]))
	// cond: c & -1<<63 == 0
	// result: (LPDFR y)
	for {
		y := v_0
		if v_1.Op != OpS390XFMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c&-1<<63 == 0) {
			break
		}
		v.reset(OpS390XLPDFR)
		v.AddArg(y)
		return true
	}
	// match: (CPSDR y (FMOVDconst [c]))
	// cond: c & -1<<63 != 0
	// result: (LNDFR y)
	for {
		y := v_0
		if v_1.Op != OpS390XFMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c&-1<<63 != 0) {
			break
		}
		v.reset(OpS390XLNDFR)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFADD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FADD (FMUL y z) x)
	// result: (FMADD x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XFMUL {
				continue
			}
			z := v_0.Args[1]
			y := v_0.Args[0]
			x := v_1
			v.reset(OpS390XFMADD)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XFADDS(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FADDS (FMULS y z) x)
	// result: (FMADDS x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XFMULS {
				continue
			}
			z := v_0.Args[1]
			y := v_0.Args[0]
			x := v_1
			v.reset(OpS390XFMADDS)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVDload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVDload [off] {sym} ptr1 (MOVDstore [off] {sym} ptr2 x _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (LDGR x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr1 := v_0
		if v_1.Op != OpS390XMOVDstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XLDGR)
		v.AddArg(x)
		return true
	}
	// match: (FMOVDload [off] {sym} ptr1 (FMOVDstore [off] {sym} ptr2 x _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr1 := v_0
		if v_1.Op != OpS390XFMOVDstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (FMOVDload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: is20Bit(off1+off2)
	// result: (FMOVDload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XFMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVDload [off1] {sym1} (MOVDaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (FMOVDload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVDload [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (FMOVDloadidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVDloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVDload [off] {sym} (ADD ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (FMOVDloadidx [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			mem := v_1
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XFMOVDloadidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVDloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVDloadidx [c] {sym} (ADDconst [d] ptr) idx mem)
	// cond: is20Bit(c+d)
	// result: (FMOVDloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v_1
		mem := v_2
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVDloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVDloadidx [c] {sym} ptr (ADDconst [d] idx) mem)
	// cond: is20Bit(c+d)
	// result: (FMOVDloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v_2
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVDloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVDstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVDstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: is20Bit(off1+off2)
	// result: (FMOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XFMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVDstore [off1] {sym1} (MOVDaddr [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (FMOVDstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVDstore [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (FMOVDstoreidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVDstoreidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVDstore [off] {sym} (ADD ptr idx) val mem)
	// cond: ptr.Op != OpSB
	// result: (FMOVDstoreidx [off] {sym} ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			val := v_1
			mem := v_2
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XFMOVDstoreidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVDstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVDstoreidx [c] {sym} (ADDconst [d] ptr) idx val mem)
	// cond: is20Bit(c+d)
	// result: (FMOVDstoreidx [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v_1
		val := v_2
		mem := v_3
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVDstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVDstoreidx [c] {sym} ptr (ADDconst [d] idx) val mem)
	// cond: is20Bit(c+d)
	// result: (FMOVDstoreidx [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v_2
		mem := v_3
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVDstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVSload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVSload [off] {sym} ptr1 (FMOVSstore [off] {sym} ptr2 x _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr1 := v_0
		if v_1.Op != OpS390XFMOVSstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (FMOVSload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: is20Bit(off1+off2)
	// result: (FMOVSload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XFMOVSload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVSload [off1] {sym1} (MOVDaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (FMOVSload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVSload [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (FMOVSloadidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVSloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVSload [off] {sym} (ADD ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (FMOVSloadidx [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			mem := v_1
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XFMOVSloadidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVSloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVSloadidx [c] {sym} (ADDconst [d] ptr) idx mem)
	// cond: is20Bit(c+d)
	// result: (FMOVSloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v_1
		mem := v_2
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVSloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVSloadidx [c] {sym} ptr (ADDconst [d] idx) mem)
	// cond: is20Bit(c+d)
	// result: (FMOVSloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v_2
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVSloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVSstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVSstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: is20Bit(off1+off2)
	// result: (FMOVSstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XFMOVSstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVSstore [off1] {sym1} (MOVDaddr [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (FMOVSstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVSstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVSstore [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (FMOVSstoreidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVSstoreidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVSstore [off] {sym} (ADD ptr idx) val mem)
	// cond: ptr.Op != OpSB
	// result: (FMOVSstoreidx [off] {sym} ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			val := v_1
			mem := v_2
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XFMOVSstoreidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVSstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVSstoreidx [c] {sym} (ADDconst [d] ptr) idx val mem)
	// cond: is20Bit(c+d)
	// result: (FMOVSstoreidx [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v_1
		val := v_2
		mem := v_3
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVSstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVSstoreidx [c] {sym} ptr (ADDconst [d] idx) val mem)
	// cond: is20Bit(c+d)
	// result: (FMOVSstoreidx [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v_2
		mem := v_3
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVSstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFNEG(v *Value) bool {
	v_0 := v.Args[0]
	// match: (FNEG (LPDFR x))
	// result: (LNDFR x)
	for {
		if v_0.Op != OpS390XLPDFR {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XLNDFR)
		v.AddArg(x)
		return true
	}
	// match: (FNEG (LNDFR x))
	// result: (LPDFR x)
	for {
		if v_0.Op != OpS390XLNDFR {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XLPDFR)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFNEGS(v *Value) bool {
	v_0 := v.Args[0]
	// match: (FNEGS (LPDFR x))
	// result: (LNDFR x)
	for {
		if v_0.Op != OpS390XLPDFR {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XLNDFR)
		v.AddArg(x)
		return true
	}
	// match: (FNEGS (LNDFR x))
	// result: (LPDFR x)
	for {
		if v_0.Op != OpS390XLNDFR {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XLPDFR)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFSUB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FSUB (FMUL y z) x)
	// result: (FMSUB x y z)
	for {
		if v_0.Op != OpS390XFMUL {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpS390XFMSUB)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFSUBS(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FSUBS (FMULS y z) x)
	// result: (FMSUBS x y z)
	for {
		if v_0.Op != OpS390XFMULS {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpS390XFMSUBS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XLDGR(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (LDGR <t> (SRDconst [1] (SLDconst [1] x)))
	// result: (LPDFR (LDGR <t> x))
	for {
		t := v.Type
		if v_0.Op != OpS390XSRDconst || v_0.AuxInt != 1 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XSLDconst || v_0_0.AuxInt != 1 {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpS390XLPDFR)
		v0 := b.NewValue0(v.Pos, OpS390XLDGR, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (LDGR <t> (AND (MOVDconst [^(-1<<63)]) x))
	// result: (LPDFR (LDGR <t> x))
	for {
		t := v.Type
		if v_0.Op != OpS390XAND {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpS390XMOVDconst || v_0_0.AuxInt != ^(-1<<63) {
				continue
			}
			x := v_0_1
			v.reset(OpS390XLPDFR)
			v0 := b.NewValue0(v.Pos, OpS390XLDGR, t)
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (LDGR <t> (OR (MOVDconst [-1<<63]) x))
	// result: (LNDFR (LDGR <t> x))
	for {
		t := v.Type
		if v_0.Op != OpS390XOR {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpS390XMOVDconst || v_0_0.AuxInt != -1<<63 {
				continue
			}
			x := v_0_1
			v.reset(OpS390XLNDFR)
			v0 := b.NewValue0(v.Pos, OpS390XLDGR, t)
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (LDGR <t> x:(ORload <t1> [off] {sym} (MOVDconst [-1<<63]) ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (LNDFR <t> (LDGR <t> (MOVDload <t1> [off] {sym} ptr mem)))
	for {
		t := v.Type
		x := v_0
		if x.Op != OpS390XORload {
			break
		}
		t1 := x.Type
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[2]
		x_0 := x.Args[0]
		if x_0.Op != OpS390XMOVDconst || x_0.AuxInt != -1<<63 {
			break
		}
		ptr := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpS390XLNDFR, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(x.Pos, OpS390XLDGR, t)
		v2 := b.NewValue0(x.Pos, OpS390XMOVDload, t1)
		v2.AuxInt = off
		v2.Aux = sym
		v2.AddArg(ptr)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		return true
	}
	// match: (LDGR (LGDR x))
	// result: x
	for {
		if v_0.Op != OpS390XLGDR {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XLEDBR(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LEDBR (LPDFR (LDEBR x)))
	// result: (LPDFR x)
	for {
		if v_0.Op != OpS390XLPDFR {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XLDEBR {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpS390XLPDFR)
		v.AddArg(x)
		return true
	}
	// match: (LEDBR (LNDFR (LDEBR x)))
	// result: (LNDFR x)
	for {
		if v_0.Op != OpS390XLNDFR {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XLDEBR {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpS390XLNDFR)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XLGDR(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LGDR (LDGR x))
	// result: x
	for {
		if v_0.Op != OpS390XLDGR {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XLOCGR(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LOCGR {c} x y (InvertFlags cmp))
	// result: (LOCGR {c.(s390x.CCMask).ReverseComparison()} x y cmp)
	for {
		c := v.Aux
		x := v_0
		y := v_1
		if v_2.Op != OpS390XInvertFlags {
			break
		}
		cmp := v_2.Args[0]
		v.reset(OpS390XLOCGR)
		v.Aux = c.(s390x.CCMask).ReverseComparison()
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cmp)
		return true
	}
	// match: (LOCGR {c} _ x (FlagEQ))
	// cond: c.(s390x.CCMask) & s390x.Equal != 0
	// result: x
	for {
		c := v.Aux
		x := v_1
		if v_2.Op != OpS390XFlagEQ || !(c.(s390x.CCMask)&s390x.Equal != 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (LOCGR {c} _ x (FlagLT))
	// cond: c.(s390x.CCMask) & s390x.Less != 0
	// result: x
	for {
		c := v.Aux
		x := v_1
		if v_2.Op != OpS390XFlagLT || !(c.(s390x.CCMask)&s390x.Less != 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (LOCGR {c} _ x (FlagGT))
	// cond: c.(s390x.CCMask) & s390x.Greater != 0
	// result: x
	for {
		c := v.Aux
		x := v_1
		if v_2.Op != OpS390XFlagGT || !(c.(s390x.CCMask)&s390x.Greater != 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (LOCGR {c} _ x (FlagOV))
	// cond: c.(s390x.CCMask) & s390x.Unordered != 0
	// result: x
	for {
		c := v.Aux
		x := v_1
		if v_2.Op != OpS390XFlagOV || !(c.(s390x.CCMask)&s390x.Unordered != 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (LOCGR {c} x _ (FlagEQ))
	// cond: c.(s390x.CCMask) & s390x.Equal == 0
	// result: x
	for {
		c := v.Aux
		x := v_0
		if v_2.Op != OpS390XFlagEQ || !(c.(s390x.CCMask)&s390x.Equal == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (LOCGR {c} x _ (FlagLT))
	// cond: c.(s390x.CCMask) & s390x.Less == 0
	// result: x
	for {
		c := v.Aux
		x := v_0
		if v_2.Op != OpS390XFlagLT || !(c.(s390x.CCMask)&s390x.Less == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (LOCGR {c} x _ (FlagGT))
	// cond: c.(s390x.CCMask) & s390x.Greater == 0
	// result: x
	for {
		c := v.Aux
		x := v_0
		if v_2.Op != OpS390XFlagGT || !(c.(s390x.CCMask)&s390x.Greater == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (LOCGR {c} x _ (FlagOV))
	// cond: c.(s390x.CCMask) & s390x.Unordered == 0
	// result: x
	for {
		c := v.Aux
		x := v_0
		if v_2.Op != OpS390XFlagOV || !(c.(s390x.CCMask)&s390x.Unordered == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XLoweredRound32F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LoweredRound32F x:(FMOVSconst))
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XFMOVSconst {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XLoweredRound64F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LoweredRound64F x:(FMOVDconst))
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XFMOVDconst {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBZload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBZload [off] {sym} ptr1 (MOVBstore [off] {sym} ptr2 x _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (MOVBZreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr1 := v_0
		if v_1.Op != OpS390XMOVBstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBZload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: is20Bit(off1+off2)
	// result: (MOVBZload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVBZload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBZload [off1] {sym1} (MOVDaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBZload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVBZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBZload [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBZloadidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVBZloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBZload [off] {sym} (ADD ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (MOVBZloadidx [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			mem := v_1
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XMOVBZloadidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBZloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBZloadidx [c] {sym} (ADDconst [d] ptr) idx mem)
	// cond: is20Bit(c+d)
	// result: (MOVBZloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XADDconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v_1
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVBZloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVBZloadidx [c] {sym} ptr (ADDconst [d] idx) mem)
	// cond: is20Bit(c+d)
	// result: (MOVBZloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			ptr := v_0
			if v_1.Op != OpS390XADDconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVBZloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBZreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVBZreg e:(MOVBreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVBreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg e:(MOVHreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVHreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg e:(MOVWreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVWreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg e:(MOVBZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVBZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg e:(MOVHZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVHZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg e:(MOVWZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVWZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg x:(MOVBZload _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 1)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBZload {
			break
		}
		_ = x.Args[1]
		if !(!x.Type.IsSigned() || x.Type.Size() > 1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg x:(MOVBZloadidx _ _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 1)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBZloadidx {
			break
		}
		_ = x.Args[2]
		if !(!x.Type.IsSigned() || x.Type.Size() > 1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg <t> x:(MOVBload [o] {s} p mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBZload <t> [o] {s} p mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpS390XMOVBload {
			break
		}
		o := x.AuxInt
		s := x.Aux
		mem := x.Args[1]
		p := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpS390XMOVBZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = o
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBZreg <t> x:(MOVBloadidx [o] {s} p i mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBZloadidx <t> [o] {s} p i mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpS390XMOVBloadidx {
			break
		}
		o := x.AuxInt
		s := x.Aux
		mem := x.Args[2]
		p := x.Args[0]
		i := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = o
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(i)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBZreg x:(Arg <t>))
	// cond: !t.IsSigned() && t.Size() == 1
	// result: x
	for {
		x := v_0
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !(!t.IsSigned() && t.Size() == 1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg (MOVDconst [c]))
	// result: (MOVDconst [int64( uint8(c))])
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(uint8(c))
		return true
	}
	// match: (MOVBZreg x:(LOCGR (MOVDconst [c]) (MOVDconst [d]) _))
	// cond: int64(uint8(c)) == c && int64(uint8(d)) == d && (!x.Type.IsSigned() || x.Type.Size() > 1)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XLOCGR {
			break
		}
		_ = x.Args[2]
		x_0 := x.Args[0]
		if x_0.Op != OpS390XMOVDconst {
			break
		}
		c := x_0.AuxInt
		x_1 := x.Args[1]
		if x_1.Op != OpS390XMOVDconst {
			break
		}
		d := x_1.AuxInt
		if !(int64(uint8(c)) == c && int64(uint8(d)) == d && (!x.Type.IsSigned() || x.Type.Size() > 1)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg (ANDWconst [m] x))
	// result: (MOVWZreg (ANDWconst <typ.UInt32> [int64( uint8(m))] x))
	for {
		if v_0.Op != OpS390XANDWconst {
			break
		}
		m := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpS390XMOVWZreg)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = int64(uint8(m))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBload [off] {sym} ptr1 (MOVBstore [off] {sym} ptr2 x _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (MOVBreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr1 := v_0
		if v_1.Op != OpS390XMOVBstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: is20Bit(off1+off2)
	// result: (MOVBload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off1] {sym1} (MOVDaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBloadidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVBloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off] {sym} (ADD ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (MOVBloadidx [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			mem := v_1
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XMOVBloadidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBloadidx [c] {sym} (ADDconst [d] ptr) idx mem)
	// cond: is20Bit(c+d)
	// result: (MOVBloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XADDconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v_1
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVBloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVBloadidx [c] {sym} ptr (ADDconst [d] idx) mem)
	// cond: is20Bit(c+d)
	// result: (MOVBloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			ptr := v_0
			if v_1.Op != OpS390XADDconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVBloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVBreg e:(MOVBreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVBreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg e:(MOVHreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVHreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg e:(MOVWreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVWreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg e:(MOVBZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVBZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg e:(MOVHZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVHZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg e:(MOVWZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVWZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg x:(MOVBload _ _))
	// cond: (x.Type.IsSigned() || x.Type.Size() == 8)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBload {
			break
		}
		_ = x.Args[1]
		if !(x.Type.IsSigned() || x.Type.Size() == 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg x:(MOVBloadidx _ _ _))
	// cond: (x.Type.IsSigned() || x.Type.Size() == 8)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBloadidx {
			break
		}
		_ = x.Args[2]
		if !(x.Type.IsSigned() || x.Type.Size() == 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg <t> x:(MOVBZload [o] {s} p mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBload <t> [o] {s} p mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpS390XMOVBZload {
			break
		}
		o := x.AuxInt
		s := x.Aux
		mem := x.Args[1]
		p := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpS390XMOVBload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = o
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBreg <t> x:(MOVBZloadidx [o] {s} p i mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBloadidx <t> [o] {s} p i mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpS390XMOVBZloadidx {
			break
		}
		o := x.AuxInt
		s := x.Aux
		mem := x.Args[2]
		p := x.Args[0]
		i := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVBloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = o
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(i)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBreg x:(Arg <t>))
	// cond: t.IsSigned() && t.Size() == 1
	// result: x
	for {
		x := v_0
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !(t.IsSigned() && t.Size() == 1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (MOVDconst [c]))
	// result: (MOVDconst [int64( int8(c))])
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int8(c))
		return true
	}
	// match: (MOVBreg (ANDWconst [m] x))
	// cond: int8(m) >= 0
	// result: (MOVWZreg (ANDWconst <typ.UInt32> [int64( uint8(m))] x))
	for {
		if v_0.Op != OpS390XANDWconst {
			break
		}
		m := v_0.AuxInt
		x := v_0.Args[0]
		if !(int8(m) >= 0) {
			break
		}
		v.reset(OpS390XMOVWZreg)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = int64(uint8(m))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstore [off] {sym} ptr (MOVBreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpS390XMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBZreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpS390XMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: is20Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVDconst [c]) mem)
	// cond: is20Bit(off) && ptr.Op != OpSB
	// result: (MOVBstoreconst [makeValAndOff(int64(int8(c)),off)] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		if !(is20Bit(off) && ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVBstoreconst)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (MOVDaddr [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBstoreidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVBstoreidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} (ADD ptr idx) val mem)
	// cond: ptr.Op != OpSB
	// result: (MOVBstoreidx [off] {sym} ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			val := v_1
			mem := v_2
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XMOVBstoreidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVBstore [i] {s} p w x:(MOVBstore [i-1] {s} p (SRDconst [8] w) mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVHstore [i-1] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w := v_1
		x := v_2
		if x.Op != OpS390XMOVBstore || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRDconst || x_1.AuxInt != 8 || w != x_1.Args[0] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p w0:(SRDconst [j] w) x:(MOVBstore [i-1] {s} p (SRDconst [j+8] w) mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVHstore [i-1] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w0 := v_1
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v_2
		if x.Op != OpS390XMOVBstore || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRDconst || x_1.AuxInt != j+8 || w != x_1.Args[0] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p w x:(MOVBstore [i-1] {s} p (SRWconst [8] w) mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVHstore [i-1] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w := v_1
		x := v_2
		if x.Op != OpS390XMOVBstore || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRWconst || x_1.AuxInt != 8 || w != x_1.Args[0] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p w0:(SRWconst [j] w) x:(MOVBstore [i-1] {s} p (SRWconst [j+8] w) mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVHstore [i-1] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w0 := v_1
		if w0.Op != OpS390XSRWconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v_2
		if x.Op != OpS390XMOVBstore || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRWconst || x_1.AuxInt != j+8 || w != x_1.Args[0] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p (SRDconst [8] w) x:(MOVBstore [i-1] {s} p w mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVHBRstore [i-1] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpS390XSRDconst || v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x := v_2
		if x.Op != OpS390XMOVBstore || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || w != x.Args[1] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p (SRDconst [j] w) x:(MOVBstore [i-1] {s} p w0:(SRDconst [j-8] w) mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVHBRstore [i-1] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpS390XSRDconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v_2
		if x.Op != OpS390XMOVBstore || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpS390XSRDconst || w0.AuxInt != j-8 || w != w0.Args[0] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p (SRWconst [8] w) x:(MOVBstore [i-1] {s} p w mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVHBRstore [i-1] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpS390XSRWconst || v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x := v_2
		if x.Op != OpS390XMOVBstore || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || w != x.Args[1] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p (SRWconst [j] w) x:(MOVBstore [i-1] {s} p w0:(SRWconst [j-8] w) mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVHBRstore [i-1] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpS390XSRWconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v_2
		if x.Op != OpS390XMOVBstore || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpS390XSRWconst || w0.AuxInt != j-8 || w != w0.Args[0] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBstoreconst(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstoreconst [sc] {s} (ADDconst [off] ptr) mem)
	// cond: is20Bit(ValAndOff(sc).Off()+off)
	// result: (MOVBstoreconst [ValAndOff(sc).add(off)] {s} ptr mem)
	for {
		sc := v.AuxInt
		s := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is20Bit(ValAndOff(sc).Off() + off)) {
			break
		}
		v.reset(OpS390XMOVBstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreconst [sc] {sym1} (MOVDaddr [off] {sym2} ptr) mem)
	// cond: ptr.Op != OpSB && canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)
	// result: (MOVBstoreconst [ValAndOff(sc).add(off)] {mergeSym(sym1, sym2)} ptr mem)
	for {
		sc := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(ptr.Op != OpSB && canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpS390XMOVBstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreconst [c] {s} p x:(MOVBstoreconst [a] {s} p mem))
	// cond: p.Op != OpSB && x.Uses == 1 && ValAndOff(a).Off() + 1 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVHstoreconst [makeValAndOff(ValAndOff(c).Val()&0xff | ValAndOff(a).Val()<<8, ValAndOff(a).Off())] {s} p mem)
	for {
		c := v.AuxInt
		s := v.Aux
		p := v_0
		x := v_1
		if x.Op != OpS390XMOVBstoreconst {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[1]
		if p != x.Args[0] || !(p.Op != OpSB && x.Uses == 1 && ValAndOff(a).Off()+1 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreconst)
		v.AuxInt = makeValAndOff(ValAndOff(c).Val()&0xff|ValAndOff(a).Val()<<8, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstoreidx [c] {sym} (ADDconst [d] ptr) idx val mem)
	// cond: is20Bit(c+d)
	// result: (MOVBstoreidx [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XADDconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v_1
			val := v_2
			mem := v_3
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVBstoreidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVBstoreidx [c] {sym} ptr (ADDconst [d] idx) val mem)
	// cond: is20Bit(c+d)
	// result: (MOVBstoreidx [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			ptr := v_0
			if v_1.Op != OpS390XADDconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			val := v_2
			mem := v_3
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVBstoreidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVBstoreidx [i] {s} p idx w x:(MOVBstoreidx [i-1] {s} p idx (SRDconst [8] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVHstoreidx [i-1] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			w := v_2
			x := v_3
			if x.Op != OpS390XMOVBstoreidx || x.AuxInt != i-1 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				x_2 := x.Args[2]
				if x_2.Op != OpS390XSRDconst || x_2.AuxInt != 8 || w != x_2.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVHstoreidx)
				v.AuxInt = i - 1
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVBstoreidx [i] {s} p idx w0:(SRDconst [j] w) x:(MOVBstoreidx [i-1] {s} p idx (SRDconst [j+8] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVHstoreidx [i-1] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			w0 := v_2
			if w0.Op != OpS390XSRDconst {
				continue
			}
			j := w0.AuxInt
			w := w0.Args[0]
			x := v_3
			if x.Op != OpS390XMOVBstoreidx || x.AuxInt != i-1 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				x_2 := x.Args[2]
				if x_2.Op != OpS390XSRDconst || x_2.AuxInt != j+8 || w != x_2.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVHstoreidx)
				v.AuxInt = i - 1
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w0)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVBstoreidx [i] {s} p idx w x:(MOVBstoreidx [i-1] {s} p idx (SRWconst [8] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVHstoreidx [i-1] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			w := v_2
			x := v_3
			if x.Op != OpS390XMOVBstoreidx || x.AuxInt != i-1 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				x_2 := x.Args[2]
				if x_2.Op != OpS390XSRWconst || x_2.AuxInt != 8 || w != x_2.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVHstoreidx)
				v.AuxInt = i - 1
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVBstoreidx [i] {s} p idx w0:(SRWconst [j] w) x:(MOVBstoreidx [i-1] {s} p idx (SRWconst [j+8] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVHstoreidx [i-1] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			w0 := v_2
			if w0.Op != OpS390XSRWconst {
				continue
			}
			j := w0.AuxInt
			w := w0.Args[0]
			x := v_3
			if x.Op != OpS390XMOVBstoreidx || x.AuxInt != i-1 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				x_2 := x.Args[2]
				if x_2.Op != OpS390XSRWconst || x_2.AuxInt != j+8 || w != x_2.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVHstoreidx)
				v.AuxInt = i - 1
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w0)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVBstoreidx [i] {s} p idx (SRDconst [8] w) x:(MOVBstoreidx [i-1] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVHBRstoreidx [i-1] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			if v_2.Op != OpS390XSRDconst || v_2.AuxInt != 8 {
				continue
			}
			w := v_2.Args[0]
			x := v_3
			if x.Op != OpS390XMOVBstoreidx || x.AuxInt != i-1 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVHBRstoreidx)
				v.AuxInt = i - 1
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVBstoreidx [i] {s} p idx (SRDconst [j] w) x:(MOVBstoreidx [i-1] {s} p idx w0:(SRDconst [j-8] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVHBRstoreidx [i-1] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			if v_2.Op != OpS390XSRDconst {
				continue
			}
			j := v_2.AuxInt
			w := v_2.Args[0]
			x := v_3
			if x.Op != OpS390XMOVBstoreidx || x.AuxInt != i-1 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				w0 := x.Args[2]
				if w0.Op != OpS390XSRDconst || w0.AuxInt != j-8 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVHBRstoreidx)
				v.AuxInt = i - 1
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w0)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVBstoreidx [i] {s} p idx (SRWconst [8] w) x:(MOVBstoreidx [i-1] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVHBRstoreidx [i-1] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			if v_2.Op != OpS390XSRWconst || v_2.AuxInt != 8 {
				continue
			}
			w := v_2.Args[0]
			x := v_3
			if x.Op != OpS390XMOVBstoreidx || x.AuxInt != i-1 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVHBRstoreidx)
				v.AuxInt = i - 1
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVBstoreidx [i] {s} p idx (SRWconst [j] w) x:(MOVBstoreidx [i-1] {s} p idx w0:(SRWconst [j-8] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVHBRstoreidx [i-1] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			if v_2.Op != OpS390XSRWconst {
				continue
			}
			j := v_2.AuxInt
			w := v_2.Args[0]
			x := v_3
			if x.Op != OpS390XMOVBstoreidx || x.AuxInt != i-1 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				w0 := x.Args[2]
				if w0.Op != OpS390XSRWconst || w0.AuxInt != j-8 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVHBRstoreidx)
				v.AuxInt = i - 1
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w0)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDaddridx(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDaddridx [c] {s} (ADDconst [d] x) y)
	// cond: is20Bit(c+d) && x.Op != OpSB
	// result: (MOVDaddridx [c+d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		y := v_1
		if !(is20Bit(c+d) && x.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDaddridx)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (MOVDaddridx [c] {s} x (ADDconst [d] y))
	// cond: is20Bit(c+d) && y.Op != OpSB
	// result: (MOVDaddridx [c+d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		x := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		y := v_1.Args[0]
		if !(is20Bit(c+d) && y.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDaddridx)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (MOVDaddridx [off1] {sym1} (MOVDaddr [off2] {sym2} x) y)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB
	// result: (MOVDaddridx [off1+off2] {mergeSym(sym1,sym2)} x y)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		y := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDaddridx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (MOVDaddridx [off1] {sym1} x (MOVDaddr [off2] {sym2} y))
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && y.Op != OpSB
	// result: (MOVDaddridx [off1+off2] {mergeSym(sym1,sym2)} x y)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		x := v_0
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		y := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && y.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDaddridx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDload [off] {sym} ptr1 (MOVDstore [off] {sym} ptr2 x _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr1 := v_0
		if v_1.Op != OpS390XMOVDstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVDload [off] {sym} ptr1 (FMOVDstore [off] {sym} ptr2 x _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (LGDR x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr1 := v_0
		if v_1.Op != OpS390XFMOVDstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XLGDR)
		v.AddArg(x)
		return true
	}
	// match: (MOVDload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: is20Bit(off1+off2)
	// result: (MOVDload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off1] {sym1} (MOVDaddr <t> [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%8 == 0 && (off1+off2)%8 == 0))
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%8 == 0 && (off1+off2)%8 == 0))) {
			break
		}
		v.reset(OpS390XMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVDloadidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVDloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off] {sym} (ADD ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (MOVDloadidx [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			mem := v_1
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XMOVDloadidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDloadidx [c] {sym} (ADDconst [d] ptr) idx mem)
	// cond: is20Bit(c+d)
	// result: (MOVDloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XADDconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v_1
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVDloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVDloadidx [c] {sym} ptr (ADDconst [d] idx) mem)
	// cond: is20Bit(c+d)
	// result: (MOVDloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			ptr := v_0
			if v_1.Op != OpS390XADDconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVDloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: is20Bit(off1+off2)
	// result: (MOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off] {sym} ptr (MOVDconst [c]) mem)
	// cond: is16Bit(c) && isU12Bit(off) && ptr.Op != OpSB
	// result: (MOVDstoreconst [makeValAndOff(c,off)] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		if !(is16Bit(c) && isU12Bit(off) && ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDstoreconst)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off1] {sym1} (MOVDaddr <t> [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%8 == 0 && (off1+off2)%8 == 0))
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%8 == 0 && (off1+off2)%8 == 0))) {
			break
		}
		v.reset(OpS390XMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVDstoreidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off] {sym} (ADD ptr idx) val mem)
	// cond: ptr.Op != OpSB
	// result: (MOVDstoreidx [off] {sym} ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			val := v_1
			mem := v_2
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XMOVDstoreidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVDstore [i] {s} p w1 x:(MOVDstore [i-8] {s} p w0 mem))
	// cond: p.Op != OpSB && x.Uses == 1 && is20Bit(i-8) && clobber(x)
	// result: (STMG2 [i-8] {s} p w0 w1 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w1 := v_1
		x := v_2
		if x.Op != OpS390XMOVDstore || x.AuxInt != i-8 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if !(p.Op != OpSB && x.Uses == 1 && is20Bit(i-8) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTMG2)
		v.AuxInt = i - 8
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [i] {s} p w2 x:(STMG2 [i-16] {s} p w0 w1 mem))
	// cond: x.Uses == 1 && is20Bit(i-16) && clobber(x)
	// result: (STMG3 [i-16] {s} p w0 w1 w2 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w2 := v_1
		x := v_2
		if x.Op != OpS390XSTMG2 || x.AuxInt != i-16 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		w1 := x.Args[2]
		if !(x.Uses == 1 && is20Bit(i-16) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTMG3)
		v.AuxInt = i - 16
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(w2)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [i] {s} p w3 x:(STMG3 [i-24] {s} p w0 w1 w2 mem))
	// cond: x.Uses == 1 && is20Bit(i-24) && clobber(x)
	// result: (STMG4 [i-24] {s} p w0 w1 w2 w3 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w3 := v_1
		x := v_2
		if x.Op != OpS390XSTMG3 || x.AuxInt != i-24 || x.Aux != s {
			break
		}
		mem := x.Args[4]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		w1 := x.Args[2]
		w2 := x.Args[3]
		if !(x.Uses == 1 && is20Bit(i-24) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTMG4)
		v.AuxInt = i - 24
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(w2)
		v.AddArg(w3)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDstoreconst(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstoreconst [sc] {s} (ADDconst [off] ptr) mem)
	// cond: isU12Bit(ValAndOff(sc).Off()+off)
	// result: (MOVDstoreconst [ValAndOff(sc).add(off)] {s} ptr mem)
	for {
		sc := v.AuxInt
		s := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU12Bit(ValAndOff(sc).Off() + off)) {
			break
		}
		v.reset(OpS390XMOVDstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstoreconst [sc] {sym1} (MOVDaddr [off] {sym2} ptr) mem)
	// cond: ptr.Op != OpSB && canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)
	// result: (MOVDstoreconst [ValAndOff(sc).add(off)] {mergeSym(sym1, sym2)} ptr mem)
	for {
		sc := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(ptr.Op != OpSB && canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpS390XMOVDstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstoreidx [c] {sym} (ADDconst [d] ptr) idx val mem)
	// cond: is20Bit(c+d)
	// result: (MOVDstoreidx [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XADDconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v_1
			val := v_2
			mem := v_3
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVDstoreidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVDstoreidx [c] {sym} ptr (ADDconst [d] idx) val mem)
	// cond: is20Bit(c+d)
	// result: (MOVDstoreidx [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			ptr := v_0
			if v_1.Op != OpS390XADDconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			val := v_2
			mem := v_3
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVDstoreidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHBRstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHBRstore [i] {s} p (SRDconst [16] w) x:(MOVHBRstore [i-2] {s} p w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWBRstore [i-2] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpS390XSRDconst || v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x := v_2
		if x.Op != OpS390XMOVHBRstore || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || w != x.Args[1] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHBRstore [i] {s} p (SRDconst [j] w) x:(MOVHBRstore [i-2] {s} p w0:(SRDconst [j-16] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWBRstore [i-2] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpS390XSRDconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v_2
		if x.Op != OpS390XMOVHBRstore || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpS390XSRDconst || w0.AuxInt != j-16 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHBRstore [i] {s} p (SRWconst [16] w) x:(MOVHBRstore [i-2] {s} p w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWBRstore [i-2] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpS390XSRWconst || v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x := v_2
		if x.Op != OpS390XMOVHBRstore || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || w != x.Args[1] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHBRstore [i] {s} p (SRWconst [j] w) x:(MOVHBRstore [i-2] {s} p w0:(SRWconst [j-16] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWBRstore [i-2] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpS390XSRWconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v_2
		if x.Op != OpS390XMOVHBRstore || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpS390XSRWconst || w0.AuxInt != j-16 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHBRstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHBRstoreidx [i] {s} p idx (SRDconst [16] w) x:(MOVHBRstoreidx [i-2] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWBRstoreidx [i-2] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			if v_2.Op != OpS390XSRDconst || v_2.AuxInt != 16 {
				continue
			}
			w := v_2.Args[0]
			x := v_3
			if x.Op != OpS390XMOVHBRstoreidx || x.AuxInt != i-2 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVWBRstoreidx)
				v.AuxInt = i - 2
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVHBRstoreidx [i] {s} p idx (SRDconst [j] w) x:(MOVHBRstoreidx [i-2] {s} p idx w0:(SRDconst [j-16] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWBRstoreidx [i-2] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			if v_2.Op != OpS390XSRDconst {
				continue
			}
			j := v_2.AuxInt
			w := v_2.Args[0]
			x := v_3
			if x.Op != OpS390XMOVHBRstoreidx || x.AuxInt != i-2 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				w0 := x.Args[2]
				if w0.Op != OpS390XSRDconst || w0.AuxInt != j-16 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVWBRstoreidx)
				v.AuxInt = i - 2
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w0)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVHBRstoreidx [i] {s} p idx (SRWconst [16] w) x:(MOVHBRstoreidx [i-2] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWBRstoreidx [i-2] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			if v_2.Op != OpS390XSRWconst || v_2.AuxInt != 16 {
				continue
			}
			w := v_2.Args[0]
			x := v_3
			if x.Op != OpS390XMOVHBRstoreidx || x.AuxInt != i-2 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVWBRstoreidx)
				v.AuxInt = i - 2
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVHBRstoreidx [i] {s} p idx (SRWconst [j] w) x:(MOVHBRstoreidx [i-2] {s} p idx w0:(SRWconst [j-16] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWBRstoreidx [i-2] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			if v_2.Op != OpS390XSRWconst {
				continue
			}
			j := v_2.AuxInt
			w := v_2.Args[0]
			x := v_3
			if x.Op != OpS390XMOVHBRstoreidx || x.AuxInt != i-2 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				w0 := x.Args[2]
				if w0.Op != OpS390XSRWconst || w0.AuxInt != j-16 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVWBRstoreidx)
				v.AuxInt = i - 2
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w0)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHZload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHZload [off] {sym} ptr1 (MOVHstore [off] {sym} ptr2 x _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (MOVHZreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr1 := v_0
		if v_1.Op != OpS390XMOVHstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHZload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: is20Bit(off1+off2)
	// result: (MOVHZload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVHZload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHZload [off1] {sym1} (MOVDaddr <t> [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%2 == 0 && (off1+off2)%2 == 0))
	// result: (MOVHZload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%2 == 0 && (off1+off2)%2 == 0))) {
			break
		}
		v.reset(OpS390XMOVHZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHZload [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVHZloadidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVHZloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHZload [off] {sym} (ADD ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (MOVHZloadidx [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			mem := v_1
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XMOVHZloadidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHZloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHZloadidx [c] {sym} (ADDconst [d] ptr) idx mem)
	// cond: is20Bit(c+d)
	// result: (MOVHZloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XADDconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v_1
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVHZloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVHZloadidx [c] {sym} ptr (ADDconst [d] idx) mem)
	// cond: is20Bit(c+d)
	// result: (MOVHZloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			ptr := v_0
			if v_1.Op != OpS390XADDconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVHZloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHZreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVHZreg e:(MOVBZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVBZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg e:(MOVHreg x))
	// cond: clobberIfDead(e)
	// result: (MOVHZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVHreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg e:(MOVWreg x))
	// cond: clobberIfDead(e)
	// result: (MOVHZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVWreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg e:(MOVHZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVHZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVHZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg e:(MOVWZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVHZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVWZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg x:(MOVBZload _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 1)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBZload {
			break
		}
		_ = x.Args[1]
		if !(!x.Type.IsSigned() || x.Type.Size() > 1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg x:(MOVBZloadidx _ _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 1)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBZloadidx {
			break
		}
		_ = x.Args[2]
		if !(!x.Type.IsSigned() || x.Type.Size() > 1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg x:(MOVHZload _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 2)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVHZload {
			break
		}
		_ = x.Args[1]
		if !(!x.Type.IsSigned() || x.Type.Size() > 2) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg x:(MOVHZloadidx _ _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 2)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVHZloadidx {
			break
		}
		_ = x.Args[2]
		if !(!x.Type.IsSigned() || x.Type.Size() > 2) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg <t> x:(MOVHload [o] {s} p mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVHZload <t> [o] {s} p mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpS390XMOVHload {
			break
		}
		o := x.AuxInt
		s := x.Aux
		mem := x.Args[1]
		p := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpS390XMOVHZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = o
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVHZreg <t> x:(MOVHloadidx [o] {s} p i mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVHZloadidx <t> [o] {s} p i mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpS390XMOVHloadidx {
			break
		}
		o := x.AuxInt
		s := x.Aux
		mem := x.Args[2]
		p := x.Args[0]
		i := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = o
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(i)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVHZreg x:(Arg <t>))
	// cond: !t.IsSigned() && t.Size() <= 2
	// result: x
	for {
		x := v_0
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !(!t.IsSigned() && t.Size() <= 2) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg (MOVDconst [c]))
	// result: (MOVDconst [int64(uint16(c))])
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(uint16(c))
		return true
	}
	// match: (MOVHZreg (ANDWconst [m] x))
	// result: (MOVWZreg (ANDWconst <typ.UInt32> [int64(uint16(m))] x))
	for {
		if v_0.Op != OpS390XANDWconst {
			break
		}
		m := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpS390XMOVWZreg)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = int64(uint16(m))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHload [off] {sym} ptr1 (MOVHstore [off] {sym} ptr2 x _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (MOVHreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr1 := v_0
		if v_1.Op != OpS390XMOVHstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: is20Bit(off1+off2)
	// result: (MOVHload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off1] {sym1} (MOVDaddr <t> [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%2 == 0 && (off1+off2)%2 == 0))
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%2 == 0 && (off1+off2)%2 == 0))) {
			break
		}
		v.reset(OpS390XMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVHloadidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVHloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off] {sym} (ADD ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (MOVHloadidx [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			mem := v_1
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XMOVHloadidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHloadidx [c] {sym} (ADDconst [d] ptr) idx mem)
	// cond: is20Bit(c+d)
	// result: (MOVHloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XADDconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v_1
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVHloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVHloadidx [c] {sym} ptr (ADDconst [d] idx) mem)
	// cond: is20Bit(c+d)
	// result: (MOVHloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			ptr := v_0
			if v_1.Op != OpS390XADDconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVHloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVHreg e:(MOVBreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVBreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg e:(MOVHreg x))
	// cond: clobberIfDead(e)
	// result: (MOVHreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVHreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg e:(MOVWreg x))
	// cond: clobberIfDead(e)
	// result: (MOVHreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVWreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg e:(MOVHZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVHreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVHZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg e:(MOVWZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVHreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVWZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBload _ _))
	// cond: (x.Type.IsSigned() || x.Type.Size() == 8)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBload {
			break
		}
		_ = x.Args[1]
		if !(x.Type.IsSigned() || x.Type.Size() == 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBloadidx _ _ _))
	// cond: (x.Type.IsSigned() || x.Type.Size() == 8)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBloadidx {
			break
		}
		_ = x.Args[2]
		if !(x.Type.IsSigned() || x.Type.Size() == 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHload _ _))
	// cond: (x.Type.IsSigned() || x.Type.Size() == 8)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVHload {
			break
		}
		_ = x.Args[1]
		if !(x.Type.IsSigned() || x.Type.Size() == 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHloadidx _ _ _))
	// cond: (x.Type.IsSigned() || x.Type.Size() == 8)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVHloadidx {
			break
		}
		_ = x.Args[2]
		if !(x.Type.IsSigned() || x.Type.Size() == 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBZload _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 1)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBZload {
			break
		}
		_ = x.Args[1]
		if !(!x.Type.IsSigned() || x.Type.Size() > 1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBZloadidx _ _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 1)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBZloadidx {
			break
		}
		_ = x.Args[2]
		if !(!x.Type.IsSigned() || x.Type.Size() > 1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg <t> x:(MOVHZload [o] {s} p mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVHload <t> [o] {s} p mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpS390XMOVHZload {
			break
		}
		o := x.AuxInt
		s := x.Aux
		mem := x.Args[1]
		p := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpS390XMOVHload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = o
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVHreg <t> x:(MOVHZloadidx [o] {s} p i mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVHloadidx <t> [o] {s} p i mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpS390XMOVHZloadidx {
			break
		}
		o := x.AuxInt
		s := x.Aux
		mem := x.Args[2]
		p := x.Args[0]
		i := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVHloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = o
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(i)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVHreg x:(Arg <t>))
	// cond: t.IsSigned() && t.Size() <= 2
	// result: x
	for {
		x := v_0
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !(t.IsSigned() && t.Size() <= 2) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (MOVDconst [c]))
	// result: (MOVDconst [int64(int16(c))])
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int16(c))
		return true
	}
	// match: (MOVHreg (ANDWconst [m] x))
	// cond: int16(m) >= 0
	// result: (MOVWZreg (ANDWconst <typ.UInt32> [int64(uint16(m))] x))
	for {
		if v_0.Op != OpS390XANDWconst {
			break
		}
		m := v_0.AuxInt
		x := v_0.Args[0]
		if !(int16(m) >= 0) {
			break
		}
		v.reset(OpS390XMOVWZreg)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = int64(uint16(m))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstore [off] {sym} ptr (MOVHreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpS390XMOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHZreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpS390XMOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: is20Bit(off1+off2)
	// result: (MOVHstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVDconst [c]) mem)
	// cond: isU12Bit(off) && ptr.Op != OpSB
	// result: (MOVHstoreconst [makeValAndOff(int64(int16(c)),off)] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		if !(isU12Bit(off) && ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVHstoreconst)
		v.AuxInt = makeValAndOff(int64(int16(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off1] {sym1} (MOVDaddr <t> [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%2 == 0 && (off1+off2)%2 == 0))
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%2 == 0 && (off1+off2)%2 == 0))) {
			break
		}
		v.reset(OpS390XMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVHstoreidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} (ADD ptr idx) val mem)
	// cond: ptr.Op != OpSB
	// result: (MOVHstoreidx [off] {sym} ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			val := v_1
			mem := v_2
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XMOVHstoreidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVHstore [i] {s} p w x:(MOVHstore [i-2] {s} p (SRDconst [16] w) mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVWstore [i-2] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w := v_1
		x := v_2
		if x.Op != OpS390XMOVHstore || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRDconst || x_1.AuxInt != 16 || w != x_1.Args[0] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [i] {s} p w0:(SRDconst [j] w) x:(MOVHstore [i-2] {s} p (SRDconst [j+16] w) mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVWstore [i-2] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w0 := v_1
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v_2
		if x.Op != OpS390XMOVHstore || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRDconst || x_1.AuxInt != j+16 || w != x_1.Args[0] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [i] {s} p w x:(MOVHstore [i-2] {s} p (SRWconst [16] w) mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVWstore [i-2] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w := v_1
		x := v_2
		if x.Op != OpS390XMOVHstore || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRWconst || x_1.AuxInt != 16 || w != x_1.Args[0] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [i] {s} p w0:(SRWconst [j] w) x:(MOVHstore [i-2] {s} p (SRWconst [j+16] w) mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVWstore [i-2] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w0 := v_1
		if w0.Op != OpS390XSRWconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v_2
		if x.Op != OpS390XMOVHstore || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRWconst || x_1.AuxInt != j+16 || w != x_1.Args[0] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHstoreconst(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVHstoreconst [sc] {s} (ADDconst [off] ptr) mem)
	// cond: isU12Bit(ValAndOff(sc).Off()+off)
	// result: (MOVHstoreconst [ValAndOff(sc).add(off)] {s} ptr mem)
	for {
		sc := v.AuxInt
		s := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU12Bit(ValAndOff(sc).Off() + off)) {
			break
		}
		v.reset(OpS390XMOVHstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstoreconst [sc] {sym1} (MOVDaddr [off] {sym2} ptr) mem)
	// cond: ptr.Op != OpSB && canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)
	// result: (MOVHstoreconst [ValAndOff(sc).add(off)] {mergeSym(sym1, sym2)} ptr mem)
	for {
		sc := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(ptr.Op != OpSB && canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpS390XMOVHstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstoreconst [c] {s} p x:(MOVHstoreconst [a] {s} p mem))
	// cond: p.Op != OpSB && x.Uses == 1 && ValAndOff(a).Off() + 2 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVWstore [ValAndOff(a).Off()] {s} p (MOVDconst [int64(int32(ValAndOff(c).Val()&0xffff | ValAndOff(a).Val()<<16))]) mem)
	for {
		c := v.AuxInt
		s := v.Aux
		p := v_0
		x := v_1
		if x.Op != OpS390XMOVHstoreconst {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[1]
		if p != x.Args[0] || !(p.Op != OpSB && x.Uses == 1 && ValAndOff(a).Off()+2 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = ValAndOff(a).Off()
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(x.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = int64(int32(ValAndOff(c).Val()&0xffff | ValAndOff(a).Val()<<16))
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstoreidx [c] {sym} (ADDconst [d] ptr) idx val mem)
	// cond: is20Bit(c+d)
	// result: (MOVHstoreidx [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XADDconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v_1
			val := v_2
			mem := v_3
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVHstoreidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVHstoreidx [c] {sym} ptr (ADDconst [d] idx) val mem)
	// cond: is20Bit(c+d)
	// result: (MOVHstoreidx [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			ptr := v_0
			if v_1.Op != OpS390XADDconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			val := v_2
			mem := v_3
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVHstoreidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVHstoreidx [i] {s} p idx w x:(MOVHstoreidx [i-2] {s} p idx (SRDconst [16] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstoreidx [i-2] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			w := v_2
			x := v_3
			if x.Op != OpS390XMOVHstoreidx || x.AuxInt != i-2 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				x_2 := x.Args[2]
				if x_2.Op != OpS390XSRDconst || x_2.AuxInt != 16 || w != x_2.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVWstoreidx)
				v.AuxInt = i - 2
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVHstoreidx [i] {s} p idx w0:(SRDconst [j] w) x:(MOVHstoreidx [i-2] {s} p idx (SRDconst [j+16] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstoreidx [i-2] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			w0 := v_2
			if w0.Op != OpS390XSRDconst {
				continue
			}
			j := w0.AuxInt
			w := w0.Args[0]
			x := v_3
			if x.Op != OpS390XMOVHstoreidx || x.AuxInt != i-2 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				x_2 := x.Args[2]
				if x_2.Op != OpS390XSRDconst || x_2.AuxInt != j+16 || w != x_2.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVWstoreidx)
				v.AuxInt = i - 2
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w0)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVHstoreidx [i] {s} p idx w x:(MOVHstoreidx [i-2] {s} p idx (SRWconst [16] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstoreidx [i-2] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			w := v_2
			x := v_3
			if x.Op != OpS390XMOVHstoreidx || x.AuxInt != i-2 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				x_2 := x.Args[2]
				if x_2.Op != OpS390XSRWconst || x_2.AuxInt != 16 || w != x_2.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVWstoreidx)
				v.AuxInt = i - 2
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVHstoreidx [i] {s} p idx w0:(SRWconst [j] w) x:(MOVHstoreidx [i-2] {s} p idx (SRWconst [j+16] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstoreidx [i-2] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			w0 := v_2
			if w0.Op != OpS390XSRWconst {
				continue
			}
			j := w0.AuxInt
			w := w0.Args[0]
			x := v_3
			if x.Op != OpS390XMOVHstoreidx || x.AuxInt != i-2 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				x_2 := x.Args[2]
				if x_2.Op != OpS390XSRWconst || x_2.AuxInt != j+16 || w != x_2.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVWstoreidx)
				v.AuxInt = i - 2
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w0)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWBRstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWBRstore [i] {s} p (SRDconst [32] w) x:(MOVWBRstore [i-4] {s} p w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVDBRstore [i-4] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpS390XSRDconst || v_1.AuxInt != 32 {
			break
		}
		w := v_1.Args[0]
		x := v_2
		if x.Op != OpS390XMOVWBRstore || x.AuxInt != i-4 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || w != x.Args[1] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDBRstore)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWBRstore [i] {s} p (SRDconst [j] w) x:(MOVWBRstore [i-4] {s} p w0:(SRDconst [j-32] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVDBRstore [i-4] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpS390XSRDconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v_2
		if x.Op != OpS390XMOVWBRstore || x.AuxInt != i-4 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpS390XSRDconst || w0.AuxInt != j-32 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDBRstore)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWBRstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWBRstoreidx [i] {s} p idx (SRDconst [32] w) x:(MOVWBRstoreidx [i-4] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVDBRstoreidx [i-4] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			if v_2.Op != OpS390XSRDconst || v_2.AuxInt != 32 {
				continue
			}
			w := v_2.Args[0]
			x := v_3
			if x.Op != OpS390XMOVWBRstoreidx || x.AuxInt != i-4 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVDBRstoreidx)
				v.AuxInt = i - 4
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVWBRstoreidx [i] {s} p idx (SRDconst [j] w) x:(MOVWBRstoreidx [i-4] {s} p idx w0:(SRDconst [j-32] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVDBRstoreidx [i-4] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			if v_2.Op != OpS390XSRDconst {
				continue
			}
			j := v_2.AuxInt
			w := v_2.Args[0]
			x := v_3
			if x.Op != OpS390XMOVWBRstoreidx || x.AuxInt != i-4 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				w0 := x.Args[2]
				if w0.Op != OpS390XSRDconst || w0.AuxInt != j-32 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVDBRstoreidx)
				v.AuxInt = i - 4
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w0)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWZload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWZload [off] {sym} ptr1 (MOVWstore [off] {sym} ptr2 x _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (MOVWZreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr1 := v_0
		if v_1.Op != OpS390XMOVWstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMOVWZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWZload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: is20Bit(off1+off2)
	// result: (MOVWZload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVWZload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWZload [off1] {sym1} (MOVDaddr <t> [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%4 == 0 && (off1+off2)%4 == 0))
	// result: (MOVWZload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%4 == 0 && (off1+off2)%4 == 0))) {
			break
		}
		v.reset(OpS390XMOVWZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWZload [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWZloadidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVWZloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWZload [off] {sym} (ADD ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (MOVWZloadidx [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			mem := v_1
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XMOVWZloadidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWZloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWZloadidx [c] {sym} (ADDconst [d] ptr) idx mem)
	// cond: is20Bit(c+d)
	// result: (MOVWZloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XADDconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v_1
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVWZloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVWZloadidx [c] {sym} ptr (ADDconst [d] idx) mem)
	// cond: is20Bit(c+d)
	// result: (MOVWZloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			ptr := v_0
			if v_1.Op != OpS390XADDconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVWZloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWZreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVWZreg e:(MOVBZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVBZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg e:(MOVHZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVHZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVHZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg e:(MOVWreg x))
	// cond: clobberIfDead(e)
	// result: (MOVWZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVWreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVWZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg e:(MOVWZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVWZreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVWZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVWZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg x:(MOVBZload _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 1)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBZload {
			break
		}
		_ = x.Args[1]
		if !(!x.Type.IsSigned() || x.Type.Size() > 1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg x:(MOVBZloadidx _ _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 1)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBZloadidx {
			break
		}
		_ = x.Args[2]
		if !(!x.Type.IsSigned() || x.Type.Size() > 1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg x:(MOVHZload _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 2)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVHZload {
			break
		}
		_ = x.Args[1]
		if !(!x.Type.IsSigned() || x.Type.Size() > 2) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg x:(MOVHZloadidx _ _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 2)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVHZloadidx {
			break
		}
		_ = x.Args[2]
		if !(!x.Type.IsSigned() || x.Type.Size() > 2) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg x:(MOVWZload _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 4)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVWZload {
			break
		}
		_ = x.Args[1]
		if !(!x.Type.IsSigned() || x.Type.Size() > 4) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg x:(MOVWZloadidx _ _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 4)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVWZloadidx {
			break
		}
		_ = x.Args[2]
		if !(!x.Type.IsSigned() || x.Type.Size() > 4) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg <t> x:(MOVWload [o] {s} p mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVWZload <t> [o] {s} p mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpS390XMOVWload {
			break
		}
		o := x.AuxInt
		s := x.Aux
		mem := x.Args[1]
		p := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpS390XMOVWZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = o
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVWZreg <t> x:(MOVWloadidx [o] {s} p i mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVWZloadidx <t> [o] {s} p i mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpS390XMOVWloadidx {
			break
		}
		o := x.AuxInt
		s := x.Aux
		mem := x.Args[2]
		p := x.Args[0]
		i := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = o
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(i)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVWZreg x:(Arg <t>))
	// cond: !t.IsSigned() && t.Size() <= 4
	// result: x
	for {
		x := v_0
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !(!t.IsSigned() && t.Size() <= 4) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg (MOVDconst [c]))
	// result: (MOVDconst [int64(uint32(c))])
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(uint32(c))
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWload [off] {sym} ptr1 (MOVWstore [off] {sym} ptr2 x _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (MOVWreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr1 := v_0
		if v_1.Op != OpS390XMOVWstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: is20Bit(off1+off2)
	// result: (MOVWload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (MOVDaddr <t> [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%4 == 0 && (off1+off2)%4 == 0))
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%4 == 0 && (off1+off2)%4 == 0))) {
			break
		}
		v.reset(OpS390XMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWloadidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVWloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off] {sym} (ADD ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (MOVWloadidx [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			mem := v_1
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XMOVWloadidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWloadidx [c] {sym} (ADDconst [d] ptr) idx mem)
	// cond: is20Bit(c+d)
	// result: (MOVWloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XADDconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v_1
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVWloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVWloadidx [c] {sym} ptr (ADDconst [d] idx) mem)
	// cond: is20Bit(c+d)
	// result: (MOVWloadidx [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			ptr := v_0
			if v_1.Op != OpS390XADDconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			mem := v_2
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVWloadidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVWreg e:(MOVBreg x))
	// cond: clobberIfDead(e)
	// result: (MOVBreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVBreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg e:(MOVHreg x))
	// cond: clobberIfDead(e)
	// result: (MOVHreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVHreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg e:(MOVWreg x))
	// cond: clobberIfDead(e)
	// result: (MOVWreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVWreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg e:(MOVWZreg x))
	// cond: clobberIfDead(e)
	// result: (MOVWreg x)
	for {
		e := v_0
		if e.Op != OpS390XMOVWZreg {
			break
		}
		x := e.Args[0]
		if !(clobberIfDead(e)) {
			break
		}
		v.reset(OpS390XMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVBload _ _))
	// cond: (x.Type.IsSigned() || x.Type.Size() == 8)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBload {
			break
		}
		_ = x.Args[1]
		if !(x.Type.IsSigned() || x.Type.Size() == 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVBloadidx _ _ _))
	// cond: (x.Type.IsSigned() || x.Type.Size() == 8)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBloadidx {
			break
		}
		_ = x.Args[2]
		if !(x.Type.IsSigned() || x.Type.Size() == 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVHload _ _))
	// cond: (x.Type.IsSigned() || x.Type.Size() == 8)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVHload {
			break
		}
		_ = x.Args[1]
		if !(x.Type.IsSigned() || x.Type.Size() == 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVHloadidx _ _ _))
	// cond: (x.Type.IsSigned() || x.Type.Size() == 8)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVHloadidx {
			break
		}
		_ = x.Args[2]
		if !(x.Type.IsSigned() || x.Type.Size() == 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVWload _ _))
	// cond: (x.Type.IsSigned() || x.Type.Size() == 8)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVWload {
			break
		}
		_ = x.Args[1]
		if !(x.Type.IsSigned() || x.Type.Size() == 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVWloadidx _ _ _))
	// cond: (x.Type.IsSigned() || x.Type.Size() == 8)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVWloadidx {
			break
		}
		_ = x.Args[2]
		if !(x.Type.IsSigned() || x.Type.Size() == 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVBZload _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 1)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBZload {
			break
		}
		_ = x.Args[1]
		if !(!x.Type.IsSigned() || x.Type.Size() > 1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVBZloadidx _ _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 1)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVBZloadidx {
			break
		}
		_ = x.Args[2]
		if !(!x.Type.IsSigned() || x.Type.Size() > 1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVHZload _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 2)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVHZload {
			break
		}
		_ = x.Args[1]
		if !(!x.Type.IsSigned() || x.Type.Size() > 2) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVHZloadidx _ _ _))
	// cond: (!x.Type.IsSigned() || x.Type.Size() > 2)
	// result: x
	for {
		x := v_0
		if x.Op != OpS390XMOVHZloadidx {
			break
		}
		_ = x.Args[2]
		if !(!x.Type.IsSigned() || x.Type.Size() > 2) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg <t> x:(MOVWZload [o] {s} p mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVWload <t> [o] {s} p mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpS390XMOVWZload {
			break
		}
		o := x.AuxInt
		s := x.Aux
		mem := x.Args[1]
		p := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpS390XMOVWload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = o
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVWreg <t> x:(MOVWZloadidx [o] {s} p i mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVWloadidx <t> [o] {s} p i mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpS390XMOVWZloadidx {
			break
		}
		o := x.AuxInt
		s := x.Aux
		mem := x.Args[2]
		p := x.Args[0]
		i := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVWloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = o
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(i)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVWreg x:(Arg <t>))
	// cond: t.IsSigned() && t.Size() <= 4
	// result: x
	for {
		x := v_0
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !(t.IsSigned() && t.Size() <= 4) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg (MOVDconst [c]))
	// result: (MOVDconst [int64(int32(c))])
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int32(c))
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstore [off] {sym} ptr (MOVWreg x) mem)
	// result: (MOVWstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpS390XMOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWZreg x) mem)
	// result: (MOVWstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpS390XMOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: is20Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVDconst [c]) mem)
	// cond: is16Bit(c) && isU12Bit(off) && ptr.Op != OpSB
	// result: (MOVWstoreconst [makeValAndOff(int64(int32(c)),off)] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		if !(is16Bit(c) && isU12Bit(off) && ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVWstoreconst)
		v.AuxInt = makeValAndOff(int64(int32(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (MOVDaddr <t> [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%4 == 0 && (off1+off2)%4 == 0))
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem().Alignment()%4 == 0 && (off1+off2)%4 == 0))) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (MOVDaddridx [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWstoreidx [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} (ADD ptr idx) val mem)
	// cond: ptr.Op != OpSB
	// result: (MOVWstoreidx [off] {sym} ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			ptr := v_0_0
			idx := v_0_1
			val := v_1
			mem := v_2
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpS390XMOVWstoreidx)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVWstore [i] {s} p (SRDconst [32] w) x:(MOVWstore [i-4] {s} p w mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVDstore [i-4] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpS390XSRDconst || v_1.AuxInt != 32 {
			break
		}
		w := v_1.Args[0]
		x := v_2
		if x.Op != OpS390XMOVWstore || x.AuxInt != i-4 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || w != x.Args[1] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstore)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [i] {s} p w0:(SRDconst [j] w) x:(MOVWstore [i-4] {s} p (SRDconst [j+32] w) mem))
	// cond: p.Op != OpSB && x.Uses == 1 && clobber(x)
	// result: (MOVDstore [i-4] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w0 := v_1
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v_2
		if x.Op != OpS390XMOVWstore || x.AuxInt != i-4 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRDconst || x_1.AuxInt != j+32 || w != x_1.Args[0] || !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstore)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [i] {s} p w1 x:(MOVWstore [i-4] {s} p w0 mem))
	// cond: p.Op != OpSB && x.Uses == 1 && is20Bit(i-4) && clobber(x)
	// result: (STM2 [i-4] {s} p w0 w1 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w1 := v_1
		x := v_2
		if x.Op != OpS390XMOVWstore || x.AuxInt != i-4 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if !(p.Op != OpSB && x.Uses == 1 && is20Bit(i-4) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTM2)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [i] {s} p w2 x:(STM2 [i-8] {s} p w0 w1 mem))
	// cond: x.Uses == 1 && is20Bit(i-8) && clobber(x)
	// result: (STM3 [i-8] {s} p w0 w1 w2 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w2 := v_1
		x := v_2
		if x.Op != OpS390XSTM2 || x.AuxInt != i-8 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		w1 := x.Args[2]
		if !(x.Uses == 1 && is20Bit(i-8) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTM3)
		v.AuxInt = i - 8
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(w2)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [i] {s} p w3 x:(STM3 [i-12] {s} p w0 w1 w2 mem))
	// cond: x.Uses == 1 && is20Bit(i-12) && clobber(x)
	// result: (STM4 [i-12] {s} p w0 w1 w2 w3 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w3 := v_1
		x := v_2
		if x.Op != OpS390XSTM3 || x.AuxInt != i-12 || x.Aux != s {
			break
		}
		mem := x.Args[4]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		w1 := x.Args[2]
		w2 := x.Args[3]
		if !(x.Uses == 1 && is20Bit(i-12) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTM4)
		v.AuxInt = i - 12
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(w2)
		v.AddArg(w3)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWstoreconst(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVWstoreconst [sc] {s} (ADDconst [off] ptr) mem)
	// cond: isU12Bit(ValAndOff(sc).Off()+off)
	// result: (MOVWstoreconst [ValAndOff(sc).add(off)] {s} ptr mem)
	for {
		sc := v.AuxInt
		s := v.Aux
		if v_0.Op != OpS390XADDconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU12Bit(ValAndOff(sc).Off() + off)) {
			break
		}
		v.reset(OpS390XMOVWstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconst [sc] {sym1} (MOVDaddr [off] {sym2} ptr) mem)
	// cond: ptr.Op != OpSB && canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)
	// result: (MOVWstoreconst [ValAndOff(sc).add(off)] {mergeSym(sym1, sym2)} ptr mem)
	for {
		sc := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(ptr.Op != OpSB && canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpS390XMOVWstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconst [c] {s} p x:(MOVWstoreconst [a] {s} p mem))
	// cond: p.Op != OpSB && x.Uses == 1 && ValAndOff(a).Off() + 4 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVDstore [ValAndOff(a).Off()] {s} p (MOVDconst [ValAndOff(c).Val()&0xffffffff | ValAndOff(a).Val()<<32]) mem)
	for {
		c := v.AuxInt
		s := v.Aux
		p := v_0
		x := v_1
		if x.Op != OpS390XMOVWstoreconst {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[1]
		if p != x.Args[0] || !(p.Op != OpSB && x.Uses == 1 && ValAndOff(a).Off()+4 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstore)
		v.AuxInt = ValAndOff(a).Off()
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(x.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = ValAndOff(c).Val()&0xffffffff | ValAndOff(a).Val()<<32
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstoreidx [c] {sym} (ADDconst [d] ptr) idx val mem)
	// cond: is20Bit(c+d)
	// result: (MOVWstoreidx [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XADDconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v_1
			val := v_2
			mem := v_3
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVWstoreidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVWstoreidx [c] {sym} ptr (ADDconst [d] idx) val mem)
	// cond: is20Bit(c+d)
	// result: (MOVWstoreidx [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			ptr := v_0
			if v_1.Op != OpS390XADDconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			val := v_2
			mem := v_3
			if !(is20Bit(c + d)) {
				continue
			}
			v.reset(OpS390XMOVWstoreidx)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVWstoreidx [i] {s} p idx w x:(MOVWstoreidx [i-4] {s} p idx (SRDconst [32] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVDstoreidx [i-4] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			w := v_2
			x := v_3
			if x.Op != OpS390XMOVWstoreidx || x.AuxInt != i-4 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				x_2 := x.Args[2]
				if x_2.Op != OpS390XSRDconst || x_2.AuxInt != 32 || w != x_2.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVDstoreidx)
				v.AuxInt = i - 4
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (MOVWstoreidx [i] {s} p idx w0:(SRDconst [j] w) x:(MOVWstoreidx [i-4] {s} p idx (SRDconst [j+32] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVDstoreidx [i-4] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			p := v_0
			idx := v_1
			w0 := v_2
			if w0.Op != OpS390XSRDconst {
				continue
			}
			j := w0.AuxInt
			w := w0.Args[0]
			x := v_3
			if x.Op != OpS390XMOVWstoreidx || x.AuxInt != i-4 || x.Aux != s {
				continue
			}
			mem := x.Args[3]
			x_0 := x.Args[0]
			x_1 := x.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x_0, x_1 = _i1+1, x_1, x_0 {
				if p != x_0 || idx != x_1 {
					continue
				}
				x_2 := x.Args[2]
				if x_2.Op != OpS390XSRDconst || x_2.AuxInt != j+32 || w != x_2.Args[0] || !(x.Uses == 1 && clobber(x)) {
					continue
				}
				v.reset(OpS390XMOVDstoreidx)
				v.AuxInt = i - 4
				v.Aux = s
				v.AddArg(p)
				v.AddArg(idx)
				v.AddArg(w0)
				v.AddArg(mem)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMULLD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MULLD x (MOVDconst [c]))
	// cond: is32Bit(c)
	// result: (MULLDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpS390XMULLDconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (MULLD <t> x g:(MOVDload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (MULLDload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVDload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XMULLDload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMULLDconst(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MULLDconst [-1] x)
	// result: (NEG x)
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v_0
		v.reset(OpS390XNEG)
		v.AddArg(x)
		return true
	}
	// match: (MULLDconst [0] _)
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (MULLDconst [1] x)
	// result: x
	for {
		if v.AuxInt != 1 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MULLDconst [c] x)
	// cond: isPowerOfTwo(c)
	// result: (SLDconst [log2(c)] x)
	for {
		c := v.AuxInt
		x := v_0
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpS390XSLDconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (MULLDconst [c] x)
	// cond: isPowerOfTwo(c+1) && c >= 15
	// result: (SUB (SLDconst <v.Type> [log2(c+1)] x) x)
	for {
		c := v.AuxInt
		x := v_0
		if !(isPowerOfTwo(c+1) && c >= 15) {
			break
		}
		v.reset(OpS390XSUB)
		v0 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULLDconst [c] x)
	// cond: isPowerOfTwo(c-1) && c >= 17
	// result: (ADD (SLDconst <v.Type> [log2(c-1)] x) x)
	for {
		c := v.AuxInt
		x := v_0
		if !(isPowerOfTwo(c-1) && c >= 17) {
			break
		}
		v.reset(OpS390XADD)
		v0 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULLDconst [c] (MOVDconst [d]))
	// result: (MOVDconst [c*d])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c * d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMULLDload(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (MULLDload <t> [off] {sym} x ptr1 (FMOVDstore [off] {sym} ptr2 y _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (MULLD x (LGDR <t> y))
	for {
		t := v.Type
		off := v.AuxInt
		sym := v.Aux
		x := v_0
		ptr1 := v_1
		if v_2.Op != OpS390XFMOVDstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		y := v_2.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMULLD)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpS390XLGDR, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (MULLDload [off1] {sym} x (ADDconst [off2] ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(off1+off2)
	// result: (MULLDload [off1+off2] {sym} x ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XMULLDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MULLDload [o1] {s1} x (MOVDaddr [o2] {s2} ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)
	// result: (MULLDload [o1+o2] {mergeSym(s1, s2)} x ptr mem)
	for {
		o1 := v.AuxInt
		s1 := v.Aux
		x := v_0
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XMULLDload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMULLW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MULLW x (MOVDconst [c]))
	// result: (MULLWconst [int64(int32(c))] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpS390XMULLWconst)
			v.AuxInt = int64(int32(c))
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (MULLW <t> x g:(MOVWload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (MULLWload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVWload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XMULLWload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MULLW <t> x g:(MOVWZload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (MULLWload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVWZload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XMULLWload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XMULLWconst(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MULLWconst [-1] x)
	// result: (NEGW x)
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v_0
		v.reset(OpS390XNEGW)
		v.AddArg(x)
		return true
	}
	// match: (MULLWconst [0] _)
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (MULLWconst [1] x)
	// result: x
	for {
		if v.AuxInt != 1 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MULLWconst [c] x)
	// cond: isPowerOfTwo(c)
	// result: (SLWconst [log2(c)] x)
	for {
		c := v.AuxInt
		x := v_0
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpS390XSLWconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (MULLWconst [c] x)
	// cond: isPowerOfTwo(c+1) && c >= 15
	// result: (SUBW (SLWconst <v.Type> [log2(c+1)] x) x)
	for {
		c := v.AuxInt
		x := v_0
		if !(isPowerOfTwo(c+1) && c >= 15) {
			break
		}
		v.reset(OpS390XSUBW)
		v0 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULLWconst [c] x)
	// cond: isPowerOfTwo(c-1) && c >= 17
	// result: (ADDW (SLWconst <v.Type> [log2(c-1)] x) x)
	for {
		c := v.AuxInt
		x := v_0
		if !(isPowerOfTwo(c-1) && c >= 17) {
			break
		}
		v.reset(OpS390XADDW)
		v0 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULLWconst [c] (MOVDconst [d]))
	// result: (MOVDconst [int64(int32(c*d))])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int32(c * d))
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMULLWload(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MULLWload [off1] {sym} x (ADDconst [off2] ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(off1+off2)
	// result: (MULLWload [off1+off2] {sym} x ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XMULLWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MULLWload [o1] {s1} x (MOVDaddr [o2] {s2} ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)
	// result: (MULLWload [o1+o2] {mergeSym(s1, s2)} x ptr mem)
	for {
		o1 := v.AuxInt
		s1 := v.Aux
		x := v_0
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XMULLWload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XNEG(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NEG (MOVDconst [c]))
	// result: (MOVDconst [-c])
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = -c
		return true
	}
	// match: (NEG (ADDconst [c] (NEG x)))
	// cond: c != -(1<<31)
	// result: (ADDconst [-c] x)
	for {
		if v_0.Op != OpS390XADDconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XNEG {
			break
		}
		x := v_0_0.Args[0]
		if !(c != -(1 << 31)) {
			break
		}
		v.reset(OpS390XADDconst)
		v.AuxInt = -c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XNEGW(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NEGW (MOVDconst [c]))
	// result: (MOVDconst [int64(int32(-c))])
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int32(-c))
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XNOT(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (NOT x)
	// result: (XOR (MOVDconst [-1]) x)
	for {
		x := v_0
		v.reset(OpS390XXOR)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = -1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpS390XNOTW(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NOTW x)
	// result: (XORWconst [-1] x)
	for {
		x := v_0
		v.reset(OpS390XXORWconst)
		v.AuxInt = -1
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpS390XOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (OR x (MOVDconst [c]))
	// cond: isU32Bit(c)
	// result: (ORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1.AuxInt
			if !(isU32Bit(c)) {
				continue
			}
			v.reset(OpS390XORconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: ( OR (SLDconst x [c]) (SRDconst x [d]))
	// cond: d == 64-c
	// result: (RLLGconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XSLDconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			if v_1.Op != OpS390XSRDconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 64-c) {
				continue
			}
			v.reset(OpS390XRLLGconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (OR (MOVDconst [-1<<63]) (LGDR <t> x))
	// result: (LGDR <t> (LNDFR <x.Type> x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XMOVDconst || v_0.AuxInt != -1<<63 || v_1.Op != OpS390XLGDR {
				continue
			}
			t := v_1.Type
			x := v_1.Args[0]
			v.reset(OpS390XLGDR)
			v.Type = t
			v0 := b.NewValue0(v.Pos, OpS390XLNDFR, x.Type)
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (OR (SLDconst [63] (SRDconst [63] (LGDR x))) (LGDR (LPDFR <t> y)))
	// result: (LGDR (CPSDR <t> y x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XSLDconst || v_0.AuxInt != 63 {
				continue
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpS390XSRDconst || v_0_0.AuxInt != 63 {
				continue
			}
			v_0_0_0 := v_0_0.Args[0]
			if v_0_0_0.Op != OpS390XLGDR {
				continue
			}
			x := v_0_0_0.Args[0]
			if v_1.Op != OpS390XLGDR {
				continue
			}
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpS390XLPDFR {
				continue
			}
			t := v_1_0.Type
			y := v_1_0.Args[0]
			v.reset(OpS390XLGDR)
			v0 := b.NewValue0(v.Pos, OpS390XCPSDR, t)
			v0.AddArg(y)
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (OR (SLDconst [63] (SRDconst [63] (LGDR x))) (MOVDconst [c]))
	// cond: c & -1<<63 == 0
	// result: (LGDR (CPSDR <x.Type> (FMOVDconst <x.Type> [c]) x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XSLDconst || v_0.AuxInt != 63 {
				continue
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpS390XSRDconst || v_0_0.AuxInt != 63 {
				continue
			}
			v_0_0_0 := v_0_0.Args[0]
			if v_0_0_0.Op != OpS390XLGDR {
				continue
			}
			x := v_0_0_0.Args[0]
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1.AuxInt
			if !(c&-1<<63 == 0) {
				continue
			}
			v.reset(OpS390XLGDR)
			v0 := b.NewValue0(v.Pos, OpS390XCPSDR, x.Type)
			v1 := b.NewValue0(v.Pos, OpS390XFMOVDconst, x.Type)
			v1.AuxInt = c
			v0.AddArg(v1)
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (OR (AND (MOVDconst [-1<<63]) (LGDR x)) (LGDR (LPDFR <t> y)))
	// result: (LGDR (CPSDR <t> y x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XAND {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				if v_0_0.Op != OpS390XMOVDconst || v_0_0.AuxInt != -1<<63 || v_0_1.Op != OpS390XLGDR {
					continue
				}
				x := v_0_1.Args[0]
				if v_1.Op != OpS390XLGDR {
					continue
				}
				v_1_0 := v_1.Args[0]
				if v_1_0.Op != OpS390XLPDFR {
					continue
				}
				t := v_1_0.Type
				y := v_1_0.Args[0]
				v.reset(OpS390XLGDR)
				v0 := b.NewValue0(v.Pos, OpS390XCPSDR, t)
				v0.AddArg(y)
				v0.AddArg(x)
				v.AddArg(v0)
				return true
			}
		}
		break
	}
	// match: (OR (AND (MOVDconst [-1<<63]) (LGDR x)) (MOVDconst [c]))
	// cond: c & -1<<63 == 0
	// result: (LGDR (CPSDR <x.Type> (FMOVDconst <x.Type> [c]) x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XAND {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				if v_0_0.Op != OpS390XMOVDconst || v_0_0.AuxInt != -1<<63 || v_0_1.Op != OpS390XLGDR {
					continue
				}
				x := v_0_1.Args[0]
				if v_1.Op != OpS390XMOVDconst {
					continue
				}
				c := v_1.AuxInt
				if !(c&-1<<63 == 0) {
					continue
				}
				v.reset(OpS390XLGDR)
				v0 := b.NewValue0(v.Pos, OpS390XCPSDR, x.Type)
				v1 := b.NewValue0(v.Pos, OpS390XFMOVDconst, x.Type)
				v1.AuxInt = c
				v0.AddArg(v1)
				v0.AddArg(x)
				v.AddArg(v0)
				return true
			}
		}
		break
	}
	// match: (OR (MOVDconst [c]) (MOVDconst [d]))
	// result: (MOVDconst [c|d])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XMOVDconst {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpS390XMOVDconst)
			v.AuxInt = c | d
			return true
		}
		break
	}
	// match: (OR x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (OR <t> x g:(MOVDload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (ORload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVDload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XORload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (OR x1:(MOVBZload [i1] {s} p mem) sh:(SLDconst [8] x0:(MOVBZload [i0] {s} p mem)))
	// cond: i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVHZload [i0] {s} p mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x1 := v_0
			if x1.Op != OpS390XMOVBZload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			sh := v_1
			if sh.Op != OpS390XSLDconst || sh.AuxInt != 8 {
				continue
			}
			x0 := sh.Args[0]
			if x0.Op != OpS390XMOVBZload {
				continue
			}
			i0 := x0.AuxInt
			if x0.Aux != s {
				continue
			}
			_ = x0.Args[1]
			if p != x0.Args[0] || mem != x0.Args[1] || !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x0.Pos, OpS390XMOVHZload, typ.UInt16)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (OR x1:(MOVHZload [i1] {s} p mem) sh:(SLDconst [16] x0:(MOVHZload [i0] {s} p mem)))
	// cond: i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVWZload [i0] {s} p mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x1 := v_0
			if x1.Op != OpS390XMOVHZload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			sh := v_1
			if sh.Op != OpS390XSLDconst || sh.AuxInt != 16 {
				continue
			}
			x0 := sh.Args[0]
			if x0.Op != OpS390XMOVHZload {
				continue
			}
			i0 := x0.AuxInt
			if x0.Aux != s {
				continue
			}
			_ = x0.Args[1]
			if p != x0.Args[0] || mem != x0.Args[1] || !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x0.Pos, OpS390XMOVWZload, typ.UInt32)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (OR x1:(MOVWZload [i1] {s} p mem) sh:(SLDconst [32] x0:(MOVWZload [i0] {s} p mem)))
	// cond: i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVDload [i0] {s} p mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x1 := v_0
			if x1.Op != OpS390XMOVWZload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			sh := v_1
			if sh.Op != OpS390XSLDconst || sh.AuxInt != 32 {
				continue
			}
			x0 := sh.Args[0]
			if x0.Op != OpS390XMOVWZload {
				continue
			}
			i0 := x0.AuxInt
			if x0.Aux != s {
				continue
			}
			_ = x0.Args[1]
			if p != x0.Args[0] || mem != x0.Args[1] || !(i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x0.Pos, OpS390XMOVDload, typ.UInt64)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (OR s0:(SLDconst [j0] x0:(MOVBZload [i0] {s} p mem)) or:(OR s1:(SLDconst [j1] x1:(MOVBZload [i1] {s} p mem)) y))
	// cond: i1 == i0+1 && j1 == j0-8 && j1 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0, x1, s0, s1, or)
	// result: @mergePoint(b,x0,x1,y) (OR <v.Type> (SLDconst <v.Type> [j1] (MOVHZload [i0] {s} p mem)) y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s0 := v_0
			if s0.Op != OpS390XSLDconst {
				continue
			}
			j0 := s0.AuxInt
			x0 := s0.Args[0]
			if x0.Op != OpS390XMOVBZload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			or := v_1
			if or.Op != OpS390XOR {
				continue
			}
			_ = or.Args[1]
			or_0 := or.Args[0]
			or_1 := or.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, or_0, or_1 = _i1+1, or_1, or_0 {
				s1 := or_0
				if s1.Op != OpS390XSLDconst {
					continue
				}
				j1 := s1.AuxInt
				x1 := s1.Args[0]
				if x1.Op != OpS390XMOVBZload {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				y := or_1
				if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0, x1, s0, s1, or)) {
					continue
				}
				b = mergePoint(b, x0, x1, y)
				v0 := b.NewValue0(x1.Pos, OpS390XOR, v.Type)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x1.Pos, OpS390XSLDconst, v.Type)
				v1.AuxInt = j1
				v2 := b.NewValue0(x1.Pos, OpS390XMOVHZload, typ.UInt16)
				v2.AuxInt = i0
				v2.Aux = s
				v2.AddArg(p)
				v2.AddArg(mem)
				v1.AddArg(v2)
				v0.AddArg(v1)
				v0.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (OR s0:(SLDconst [j0] x0:(MOVHZload [i0] {s} p mem)) or:(OR s1:(SLDconst [j1] x1:(MOVHZload [i1] {s} p mem)) y))
	// cond: i1 == i0+2 && j1 == j0-16 && j1 % 32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0, x1, s0, s1, or)
	// result: @mergePoint(b,x0,x1,y) (OR <v.Type> (SLDconst <v.Type> [j1] (MOVWZload [i0] {s} p mem)) y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s0 := v_0
			if s0.Op != OpS390XSLDconst {
				continue
			}
			j0 := s0.AuxInt
			x0 := s0.Args[0]
			if x0.Op != OpS390XMOVHZload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			or := v_1
			if or.Op != OpS390XOR {
				continue
			}
			_ = or.Args[1]
			or_0 := or.Args[0]
			or_1 := or.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, or_0, or_1 = _i1+1, or_1, or_0 {
				s1 := or_0
				if s1.Op != OpS390XSLDconst {
					continue
				}
				j1 := s1.AuxInt
				x1 := s1.Args[0]
				if x1.Op != OpS390XMOVHZload {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				y := or_1
				if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0, x1, s0, s1, or)) {
					continue
				}
				b = mergePoint(b, x0, x1, y)
				v0 := b.NewValue0(x1.Pos, OpS390XOR, v.Type)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x1.Pos, OpS390XSLDconst, v.Type)
				v1.AuxInt = j1
				v2 := b.NewValue0(x1.Pos, OpS390XMOVWZload, typ.UInt32)
				v2.AuxInt = i0
				v2.Aux = s
				v2.AddArg(p)
				v2.AddArg(mem)
				v1.AddArg(v2)
				v0.AddArg(v1)
				v0.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (OR x1:(MOVBZloadidx [i1] {s} p idx mem) sh:(SLDconst [8] x0:(MOVBZloadidx [i0] {s} p idx mem)))
	// cond: i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVHZloadidx [i0] {s} p idx mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x1 := v_0
			if x1.Op != OpS390XMOVBZloadidx {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			x1_0 := x1.Args[0]
			x1_1 := x1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x1_0, x1_1 = _i1+1, x1_1, x1_0 {
				p := x1_0
				idx := x1_1
				sh := v_1
				if sh.Op != OpS390XSLDconst || sh.AuxInt != 8 {
					continue
				}
				x0 := sh.Args[0]
				if x0.Op != OpS390XMOVBZloadidx {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[2]
				x0_0 := x0.Args[0]
				x0_1 := x0.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, x0_0, x0_1 = _i2+1, x0_1, x0_0 {
					if p != x0_0 || idx != x0_1 || mem != x0.Args[2] || !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = i0
					v0.Aux = s
					v0.AddArg(p)
					v0.AddArg(idx)
					v0.AddArg(mem)
					return true
				}
			}
		}
		break
	}
	// match: (OR x1:(MOVHZloadidx [i1] {s} p idx mem) sh:(SLDconst [16] x0:(MOVHZloadidx [i0] {s} p idx mem)))
	// cond: i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVWZloadidx [i0] {s} p idx mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x1 := v_0
			if x1.Op != OpS390XMOVHZloadidx {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			x1_0 := x1.Args[0]
			x1_1 := x1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x1_0, x1_1 = _i1+1, x1_1, x1_0 {
				p := x1_0
				idx := x1_1
				sh := v_1
				if sh.Op != OpS390XSLDconst || sh.AuxInt != 16 {
					continue
				}
				x0 := sh.Args[0]
				if x0.Op != OpS390XMOVHZloadidx {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[2]
				x0_0 := x0.Args[0]
				x0_1 := x0.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, x0_0, x0_1 = _i2+1, x0_1, x0_0 {
					if p != x0_0 || idx != x0_1 || mem != x0.Args[2] || !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = i0
					v0.Aux = s
					v0.AddArg(p)
					v0.AddArg(idx)
					v0.AddArg(mem)
					return true
				}
			}
		}
		break
	}
	// match: (OR x1:(MOVWZloadidx [i1] {s} p idx mem) sh:(SLDconst [32] x0:(MOVWZloadidx [i0] {s} p idx mem)))
	// cond: i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVDloadidx [i0] {s} p idx mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x1 := v_0
			if x1.Op != OpS390XMOVWZloadidx {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			x1_0 := x1.Args[0]
			x1_1 := x1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x1_0, x1_1 = _i1+1, x1_1, x1_0 {
				p := x1_0
				idx := x1_1
				sh := v_1
				if sh.Op != OpS390XSLDconst || sh.AuxInt != 32 {
					continue
				}
				x0 := sh.Args[0]
				if x0.Op != OpS390XMOVWZloadidx {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[2]
				x0_0 := x0.Args[0]
				x0_1 := x0.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, x0_0, x0_1 = _i2+1, x0_1, x0_0 {
					if p != x0_0 || idx != x0_1 || mem != x0.Args[2] || !(i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpS390XMOVDloadidx, typ.UInt64)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = i0
					v0.Aux = s
					v0.AddArg(p)
					v0.AddArg(idx)
					v0.AddArg(mem)
					return true
				}
			}
		}
		break
	}
	// match: (OR s0:(SLDconst [j0] x0:(MOVBZloadidx [i0] {s} p idx mem)) or:(OR s1:(SLDconst [j1] x1:(MOVBZloadidx [i1] {s} p idx mem)) y))
	// cond: i1 == i0+1 && j1 == j0-8 && j1 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0, x1, s0, s1, or)
	// result: @mergePoint(b,x0,x1,y) (OR <v.Type> (SLDconst <v.Type> [j1] (MOVHZloadidx [i0] {s} p idx mem)) y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s0 := v_0
			if s0.Op != OpS390XSLDconst {
				continue
			}
			j0 := s0.AuxInt
			x0 := s0.Args[0]
			if x0.Op != OpS390XMOVBZloadidx {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			x0_0 := x0.Args[0]
			x0_1 := x0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x0_0, x0_1 = _i1+1, x0_1, x0_0 {
				p := x0_0
				idx := x0_1
				or := v_1
				if or.Op != OpS390XOR {
					continue
				}
				_ = or.Args[1]
				or_0 := or.Args[0]
				or_1 := or.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, or_0, or_1 = _i2+1, or_1, or_0 {
					s1 := or_0
					if s1.Op != OpS390XSLDconst {
						continue
					}
					j1 := s1.AuxInt
					x1 := s1.Args[0]
					if x1.Op != OpS390XMOVBZloadidx {
						continue
					}
					i1 := x1.AuxInt
					if x1.Aux != s {
						continue
					}
					_ = x1.Args[2]
					x1_0 := x1.Args[0]
					x1_1 := x1.Args[1]
					for _i3 := 0; _i3 <= 1; _i3, x1_0, x1_1 = _i3+1, x1_1, x1_0 {
						if p != x1_0 || idx != x1_1 || mem != x1.Args[2] {
							continue
						}
						y := or_1
						if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0, x1, s0, s1, or)) {
							continue
						}
						b = mergePoint(b, x0, x1, y)
						v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
						v1.AuxInt = j1
						v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
						v2.AuxInt = i0
						v2.Aux = s
						v2.AddArg(p)
						v2.AddArg(idx)
						v2.AddArg(mem)
						v1.AddArg(v2)
						v0.AddArg(v1)
						v0.AddArg(y)
						return true
					}
				}
			}
		}
		break
	}
	// match: (OR s0:(SLDconst [j0] x0:(MOVHZloadidx [i0] {s} p idx mem)) or:(OR s1:(SLDconst [j1] x1:(MOVHZloadidx [i1] {s} p idx mem)) y))
	// cond: i1 == i0+2 && j1 == j0-16 && j1 % 32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0, x1, s0, s1, or)
	// result: @mergePoint(b,x0,x1,y) (OR <v.Type> (SLDconst <v.Type> [j1] (MOVWZloadidx [i0] {s} p idx mem)) y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s0 := v_0
			if s0.Op != OpS390XSLDconst {
				continue
			}
			j0 := s0.AuxInt
			x0 := s0.Args[0]
			if x0.Op != OpS390XMOVHZloadidx {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			x0_0 := x0.Args[0]
			x0_1 := x0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x0_0, x0_1 = _i1+1, x0_1, x0_0 {
				p := x0_0
				idx := x0_1
				or := v_1
				if or.Op != OpS390XOR {
					continue
				}
				_ = or.Args[1]
				or_0 := or.Args[0]
				or_1 := or.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, or_0, or_1 = _i2+1, or_1, or_0 {
					s1 := or_0
					if s1.Op != OpS390XSLDconst {
						continue
					}
					j1 := s1.AuxInt
					x1 := s1.Args[0]
					if x1.Op != OpS390XMOVHZloadidx {
						continue
					}
					i1 := x1.AuxInt
					if x1.Aux != s {
						continue
					}
					_ = x1.Args[2]
					x1_0 := x1.Args[0]
					x1_1 := x1.Args[1]
					for _i3 := 0; _i3 <= 1; _i3, x1_0, x1_1 = _i3+1, x1_1, x1_0 {
						if p != x1_0 || idx != x1_1 || mem != x1.Args[2] {
							continue
						}
						y := or_1
						if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0, x1, s0, s1, or)) {
							continue
						}
						b = mergePoint(b, x0, x1, y)
						v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
						v1.AuxInt = j1
						v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
						v2.AuxInt = i0
						v2.Aux = s
						v2.AddArg(p)
						v2.AddArg(idx)
						v2.AddArg(mem)
						v1.AddArg(v2)
						v0.AddArg(v1)
						v0.AddArg(y)
						return true
					}
				}
			}
		}
		break
	}
	// match: (OR x0:(MOVBZload [i0] {s} p mem) sh:(SLDconst [8] x1:(MOVBZload [i1] {s} p mem)))
	// cond: p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVHZreg (MOVHBRload [i0] {s} p mem))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x0 := v_0
			if x0.Op != OpS390XMOVBZload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			sh := v_1
			if sh.Op != OpS390XSLDconst || sh.AuxInt != 8 {
				continue
			}
			x1 := sh.Args[0]
			if x1.Op != OpS390XMOVBZload {
				continue
			}
			i1 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpS390XMOVHZreg, typ.UInt64)
			v.reset(OpCopy)
			v.AddArg(v0)
			v1 := b.NewValue0(x1.Pos, OpS390XMOVHBRload, typ.UInt16)
			v1.AuxInt = i0
			v1.Aux = s
			v1.AddArg(p)
			v1.AddArg(mem)
			v0.AddArg(v1)
			return true
		}
		break
	}
	// match: (OR r0:(MOVHZreg x0:(MOVHBRload [i0] {s} p mem)) sh:(SLDconst [16] r1:(MOVHZreg x1:(MOVHBRload [i1] {s} p mem))))
	// cond: i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, r0, r1, sh)
	// result: @mergePoint(b,x0,x1) (MOVWZreg (MOVWBRload [i0] {s} p mem))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			r0 := v_0
			if r0.Op != OpS390XMOVHZreg {
				continue
			}
			x0 := r0.Args[0]
			if x0.Op != OpS390XMOVHBRload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			sh := v_1
			if sh.Op != OpS390XSLDconst || sh.AuxInt != 16 {
				continue
			}
			r1 := sh.Args[0]
			if r1.Op != OpS390XMOVHZreg {
				continue
			}
			x1 := r1.Args[0]
			if x1.Op != OpS390XMOVHBRload {
				continue
			}
			i1 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, r0, r1, sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpS390XMOVWZreg, typ.UInt64)
			v.reset(OpCopy)
			v.AddArg(v0)
			v1 := b.NewValue0(x1.Pos, OpS390XMOVWBRload, typ.UInt32)
			v1.AuxInt = i0
			v1.Aux = s
			v1.AddArg(p)
			v1.AddArg(mem)
			v0.AddArg(v1)
			return true
		}
		break
	}
	// match: (OR r0:(MOVWZreg x0:(MOVWBRload [i0] {s} p mem)) sh:(SLDconst [32] r1:(MOVWZreg x1:(MOVWBRload [i1] {s} p mem))))
	// cond: i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, r0, r1, sh)
	// result: @mergePoint(b,x0,x1) (MOVDBRload [i0] {s} p mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			r0 := v_0
			if r0.Op != OpS390XMOVWZreg {
				continue
			}
			x0 := r0.Args[0]
			if x0.Op != OpS390XMOVWBRload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			sh := v_1
			if sh.Op != OpS390XSLDconst || sh.AuxInt != 32 {
				continue
			}
			r1 := sh.Args[0]
			if r1.Op != OpS390XMOVWZreg {
				continue
			}
			x1 := r1.Args[0]
			if x1.Op != OpS390XMOVWBRload {
				continue
			}
			i1 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, r0, r1, sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpS390XMOVDBRload, typ.UInt64)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (OR s1:(SLDconst [j1] x1:(MOVBZload [i1] {s} p mem)) or:(OR s0:(SLDconst [j0] x0:(MOVBZload [i0] {s} p mem)) y))
	// cond: p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0, x1, s0, s1, or)
	// result: @mergePoint(b,x0,x1,y) (OR <v.Type> (SLDconst <v.Type> [j0] (MOVHZreg (MOVHBRload [i0] {s} p mem))) y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s1 := v_0
			if s1.Op != OpS390XSLDconst {
				continue
			}
			j1 := s1.AuxInt
			x1 := s1.Args[0]
			if x1.Op != OpS390XMOVBZload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			or := v_1
			if or.Op != OpS390XOR {
				continue
			}
			_ = or.Args[1]
			or_0 := or.Args[0]
			or_1 := or.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, or_0, or_1 = _i1+1, or_1, or_0 {
				s0 := or_0
				if s0.Op != OpS390XSLDconst {
					continue
				}
				j0 := s0.AuxInt
				x0 := s0.Args[0]
				if x0.Op != OpS390XMOVBZload {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[1]
				if p != x0.Args[0] || mem != x0.Args[1] {
					continue
				}
				y := or_1
				if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0, x1, s0, s1, or)) {
					continue
				}
				b = mergePoint(b, x0, x1, y)
				v0 := b.NewValue0(x0.Pos, OpS390XOR, v.Type)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x0.Pos, OpS390XSLDconst, v.Type)
				v1.AuxInt = j0
				v2 := b.NewValue0(x0.Pos, OpS390XMOVHZreg, typ.UInt64)
				v3 := b.NewValue0(x0.Pos, OpS390XMOVHBRload, typ.UInt16)
				v3.AuxInt = i0
				v3.Aux = s
				v3.AddArg(p)
				v3.AddArg(mem)
				v2.AddArg(v3)
				v1.AddArg(v2)
				v0.AddArg(v1)
				v0.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (OR s1:(SLDconst [j1] r1:(MOVHZreg x1:(MOVHBRload [i1] {s} p mem))) or:(OR s0:(SLDconst [j0] r0:(MOVHZreg x0:(MOVHBRload [i0] {s} p mem))) y))
	// cond: i1 == i0+2 && j1 == j0+16 && j0 % 32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0, x1, r0, r1, s0, s1, or)
	// result: @mergePoint(b,x0,x1,y) (OR <v.Type> (SLDconst <v.Type> [j0] (MOVWZreg (MOVWBRload [i0] {s} p mem))) y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s1 := v_0
			if s1.Op != OpS390XSLDconst {
				continue
			}
			j1 := s1.AuxInt
			r1 := s1.Args[0]
			if r1.Op != OpS390XMOVHZreg {
				continue
			}
			x1 := r1.Args[0]
			if x1.Op != OpS390XMOVHBRload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			or := v_1
			if or.Op != OpS390XOR {
				continue
			}
			_ = or.Args[1]
			or_0 := or.Args[0]
			or_1 := or.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, or_0, or_1 = _i1+1, or_1, or_0 {
				s0 := or_0
				if s0.Op != OpS390XSLDconst {
					continue
				}
				j0 := s0.AuxInt
				r0 := s0.Args[0]
				if r0.Op != OpS390XMOVHZreg {
					continue
				}
				x0 := r0.Args[0]
				if x0.Op != OpS390XMOVHBRload {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[1]
				if p != x0.Args[0] || mem != x0.Args[1] {
					continue
				}
				y := or_1
				if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0, x1, r0, r1, s0, s1, or)) {
					continue
				}
				b = mergePoint(b, x0, x1, y)
				v0 := b.NewValue0(x0.Pos, OpS390XOR, v.Type)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x0.Pos, OpS390XSLDconst, v.Type)
				v1.AuxInt = j0
				v2 := b.NewValue0(x0.Pos, OpS390XMOVWZreg, typ.UInt64)
				v3 := b.NewValue0(x0.Pos, OpS390XMOVWBRload, typ.UInt32)
				v3.AuxInt = i0
				v3.Aux = s
				v3.AddArg(p)
				v3.AddArg(mem)
				v2.AddArg(v3)
				v1.AddArg(v2)
				v0.AddArg(v1)
				v0.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (OR x0:(MOVBZloadidx [i0] {s} p idx mem) sh:(SLDconst [8] x1:(MOVBZloadidx [i1] {s} p idx mem)))
	// cond: p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVHZreg (MOVHBRloadidx [i0] {s} p idx mem))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x0 := v_0
			if x0.Op != OpS390XMOVBZloadidx {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			x0_0 := x0.Args[0]
			x0_1 := x0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x0_0, x0_1 = _i1+1, x0_1, x0_0 {
				p := x0_0
				idx := x0_1
				sh := v_1
				if sh.Op != OpS390XSLDconst || sh.AuxInt != 8 {
					continue
				}
				x1 := sh.Args[0]
				if x1.Op != OpS390XMOVBZloadidx {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[2]
				x1_0 := x1.Args[0]
				x1_1 := x1.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, x1_0, x1_1 = _i2+1, x1_1, x1_0 {
					if p != x1_0 || idx != x1_1 || mem != x1.Args[2] || !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
					v.reset(OpCopy)
					v.AddArg(v0)
					v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
					v1.AuxInt = i0
					v1.Aux = s
					v1.AddArg(p)
					v1.AddArg(idx)
					v1.AddArg(mem)
					v0.AddArg(v1)
					return true
				}
			}
		}
		break
	}
	// match: (OR r0:(MOVHZreg x0:(MOVHBRloadidx [i0] {s} p idx mem)) sh:(SLDconst [16] r1:(MOVHZreg x1:(MOVHBRloadidx [i1] {s} p idx mem))))
	// cond: i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, r0, r1, sh)
	// result: @mergePoint(b,x0,x1) (MOVWZreg (MOVWBRloadidx [i0] {s} p idx mem))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			r0 := v_0
			if r0.Op != OpS390XMOVHZreg {
				continue
			}
			x0 := r0.Args[0]
			if x0.Op != OpS390XMOVHBRloadidx {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			x0_0 := x0.Args[0]
			x0_1 := x0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x0_0, x0_1 = _i1+1, x0_1, x0_0 {
				p := x0_0
				idx := x0_1
				sh := v_1
				if sh.Op != OpS390XSLDconst || sh.AuxInt != 16 {
					continue
				}
				r1 := sh.Args[0]
				if r1.Op != OpS390XMOVHZreg {
					continue
				}
				x1 := r1.Args[0]
				if x1.Op != OpS390XMOVHBRloadidx {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[2]
				x1_0 := x1.Args[0]
				x1_1 := x1.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, x1_0, x1_1 = _i2+1, x1_1, x1_0 {
					if p != x1_0 || idx != x1_1 || mem != x1.Args[2] || !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, r0, r1, sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
					v.reset(OpCopy)
					v.AddArg(v0)
					v1 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
					v1.AuxInt = i0
					v1.Aux = s
					v1.AddArg(p)
					v1.AddArg(idx)
					v1.AddArg(mem)
					v0.AddArg(v1)
					return true
				}
			}
		}
		break
	}
	// match: (OR r0:(MOVWZreg x0:(MOVWBRloadidx [i0] {s} p idx mem)) sh:(SLDconst [32] r1:(MOVWZreg x1:(MOVWBRloadidx [i1] {s} p idx mem))))
	// cond: i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, r0, r1, sh)
	// result: @mergePoint(b,x0,x1) (MOVDBRloadidx [i0] {s} p idx mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			r0 := v_0
			if r0.Op != OpS390XMOVWZreg {
				continue
			}
			x0 := r0.Args[0]
			if x0.Op != OpS390XMOVWBRloadidx {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			x0_0 := x0.Args[0]
			x0_1 := x0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x0_0, x0_1 = _i1+1, x0_1, x0_0 {
				p := x0_0
				idx := x0_1
				sh := v_1
				if sh.Op != OpS390XSLDconst || sh.AuxInt != 32 {
					continue
				}
				r1 := sh.Args[0]
				if r1.Op != OpS390XMOVWZreg {
					continue
				}
				x1 := r1.Args[0]
				if x1.Op != OpS390XMOVWBRloadidx {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[2]
				x1_0 := x1.Args[0]
				x1_1 := x1.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, x1_0, x1_1 = _i2+1, x1_1, x1_0 {
					if p != x1_0 || idx != x1_1 || mem != x1.Args[2] || !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, r0, r1, sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpS390XMOVDBRloadidx, typ.Int64)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = i0
					v0.Aux = s
					v0.AddArg(p)
					v0.AddArg(idx)
					v0.AddArg(mem)
					return true
				}
			}
		}
		break
	}
	// match: (OR s1:(SLDconst [j1] x1:(MOVBZloadidx [i1] {s} p idx mem)) or:(OR s0:(SLDconst [j0] x0:(MOVBZloadidx [i0] {s} p idx mem)) y))
	// cond: p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0, x1, s0, s1, or)
	// result: @mergePoint(b,x0,x1,y) (OR <v.Type> (SLDconst <v.Type> [j0] (MOVHZreg (MOVHBRloadidx [i0] {s} p idx mem))) y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s1 := v_0
			if s1.Op != OpS390XSLDconst {
				continue
			}
			j1 := s1.AuxInt
			x1 := s1.Args[0]
			if x1.Op != OpS390XMOVBZloadidx {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			x1_0 := x1.Args[0]
			x1_1 := x1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x1_0, x1_1 = _i1+1, x1_1, x1_0 {
				p := x1_0
				idx := x1_1
				or := v_1
				if or.Op != OpS390XOR {
					continue
				}
				_ = or.Args[1]
				or_0 := or.Args[0]
				or_1 := or.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, or_0, or_1 = _i2+1, or_1, or_0 {
					s0 := or_0
					if s0.Op != OpS390XSLDconst {
						continue
					}
					j0 := s0.AuxInt
					x0 := s0.Args[0]
					if x0.Op != OpS390XMOVBZloadidx {
						continue
					}
					i0 := x0.AuxInt
					if x0.Aux != s {
						continue
					}
					_ = x0.Args[2]
					x0_0 := x0.Args[0]
					x0_1 := x0.Args[1]
					for _i3 := 0; _i3 <= 1; _i3, x0_0, x0_1 = _i3+1, x0_1, x0_0 {
						if p != x0_0 || idx != x0_1 || mem != x0.Args[2] {
							continue
						}
						y := or_1
						if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0, x1, s0, s1, or)) {
							continue
						}
						b = mergePoint(b, x0, x1, y)
						v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
						v1.AuxInt = j0
						v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
						v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
						v3.AuxInt = i0
						v3.Aux = s
						v3.AddArg(p)
						v3.AddArg(idx)
						v3.AddArg(mem)
						v2.AddArg(v3)
						v1.AddArg(v2)
						v0.AddArg(v1)
						v0.AddArg(y)
						return true
					}
				}
			}
		}
		break
	}
	// match: (OR s1:(SLDconst [j1] r1:(MOVHZreg x1:(MOVHBRloadidx [i1] {s} p idx mem))) or:(OR s0:(SLDconst [j0] r0:(MOVHZreg x0:(MOVHBRloadidx [i0] {s} p idx mem))) y))
	// cond: i1 == i0+2 && j1 == j0+16 && j0 % 32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0, x1, r0, r1, s0, s1, or)
	// result: @mergePoint(b,x0,x1,y) (OR <v.Type> (SLDconst <v.Type> [j0] (MOVWZreg (MOVWBRloadidx [i0] {s} p idx mem))) y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s1 := v_0
			if s1.Op != OpS390XSLDconst {
				continue
			}
			j1 := s1.AuxInt
			r1 := s1.Args[0]
			if r1.Op != OpS390XMOVHZreg {
				continue
			}
			x1 := r1.Args[0]
			if x1.Op != OpS390XMOVHBRloadidx {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			x1_0 := x1.Args[0]
			x1_1 := x1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x1_0, x1_1 = _i1+1, x1_1, x1_0 {
				p := x1_0
				idx := x1_1
				or := v_1
				if or.Op != OpS390XOR {
					continue
				}
				_ = or.Args[1]
				or_0 := or.Args[0]
				or_1 := or.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, or_0, or_1 = _i2+1, or_1, or_0 {
					s0 := or_0
					if s0.Op != OpS390XSLDconst {
						continue
					}
					j0 := s0.AuxInt
					r0 := s0.Args[0]
					if r0.Op != OpS390XMOVHZreg {
						continue
					}
					x0 := r0.Args[0]
					if x0.Op != OpS390XMOVHBRloadidx {
						continue
					}
					i0 := x0.AuxInt
					if x0.Aux != s {
						continue
					}
					_ = x0.Args[2]
					x0_0 := x0.Args[0]
					x0_1 := x0.Args[1]
					for _i3 := 0; _i3 <= 1; _i3, x0_0, x0_1 = _i3+1, x0_1, x0_0 {
						if p != x0_0 || idx != x0_1 || mem != x0.Args[2] {
							continue
						}
						y := or_1
						if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0, x1, r0, r1, s0, s1, or)) {
							continue
						}
						b = mergePoint(b, x0, x1, y)
						v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
						v1.AuxInt = j0
						v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
						v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
						v3.AuxInt = i0
						v3.Aux = s
						v3.AddArg(p)
						v3.AddArg(idx)
						v3.AddArg(mem)
						v2.AddArg(v3)
						v1.AddArg(v2)
						v0.AddArg(v1)
						v0.AddArg(y)
						return true
					}
				}
			}
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XORW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ORW x (MOVDconst [c]))
	// result: (ORWconst [int64(int32(c))] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpS390XORWconst)
			v.AuxInt = int64(int32(c))
			v.AddArg(x)
			return true
		}
		break
	}
	// match: ( ORW (SLWconst x [c]) (SRWconst x [d]))
	// cond: d == 32-c
	// result: (RLLconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XSLWconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			if v_1.Op != OpS390XSRWconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 32-c) {
				continue
			}
			v.reset(OpS390XRLLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ORW x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ORW <t> x g:(MOVWload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (ORWload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVWload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XORWload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (ORW <t> x g:(MOVWZload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (ORWload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVWZload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XORWload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (ORW x1:(MOVBZload [i1] {s} p mem) sh:(SLWconst [8] x0:(MOVBZload [i0] {s} p mem)))
	// cond: i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVHZload [i0] {s} p mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x1 := v_0
			if x1.Op != OpS390XMOVBZload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			sh := v_1
			if sh.Op != OpS390XSLWconst || sh.AuxInt != 8 {
				continue
			}
			x0 := sh.Args[0]
			if x0.Op != OpS390XMOVBZload {
				continue
			}
			i0 := x0.AuxInt
			if x0.Aux != s {
				continue
			}
			_ = x0.Args[1]
			if p != x0.Args[0] || mem != x0.Args[1] || !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x0.Pos, OpS390XMOVHZload, typ.UInt16)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (ORW x1:(MOVHZload [i1] {s} p mem) sh:(SLWconst [16] x0:(MOVHZload [i0] {s} p mem)))
	// cond: i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVWZload [i0] {s} p mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x1 := v_0
			if x1.Op != OpS390XMOVHZload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			sh := v_1
			if sh.Op != OpS390XSLWconst || sh.AuxInt != 16 {
				continue
			}
			x0 := sh.Args[0]
			if x0.Op != OpS390XMOVHZload {
				continue
			}
			i0 := x0.AuxInt
			if x0.Aux != s {
				continue
			}
			_ = x0.Args[1]
			if p != x0.Args[0] || mem != x0.Args[1] || !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x0.Pos, OpS390XMOVWZload, typ.UInt32)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (ORW s0:(SLWconst [j0] x0:(MOVBZload [i0] {s} p mem)) or:(ORW s1:(SLWconst [j1] x1:(MOVBZload [i1] {s} p mem)) y))
	// cond: i1 == i0+1 && j1 == j0-8 && j1 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0, x1, s0, s1, or)
	// result: @mergePoint(b,x0,x1,y) (ORW <v.Type> (SLWconst <v.Type> [j1] (MOVHZload [i0] {s} p mem)) y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s0 := v_0
			if s0.Op != OpS390XSLWconst {
				continue
			}
			j0 := s0.AuxInt
			x0 := s0.Args[0]
			if x0.Op != OpS390XMOVBZload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			or := v_1
			if or.Op != OpS390XORW {
				continue
			}
			_ = or.Args[1]
			or_0 := or.Args[0]
			or_1 := or.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, or_0, or_1 = _i1+1, or_1, or_0 {
				s1 := or_0
				if s1.Op != OpS390XSLWconst {
					continue
				}
				j1 := s1.AuxInt
				x1 := s1.Args[0]
				if x1.Op != OpS390XMOVBZload {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				y := or_1
				if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0, x1, s0, s1, or)) {
					continue
				}
				b = mergePoint(b, x0, x1, y)
				v0 := b.NewValue0(x1.Pos, OpS390XORW, v.Type)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x1.Pos, OpS390XSLWconst, v.Type)
				v1.AuxInt = j1
				v2 := b.NewValue0(x1.Pos, OpS390XMOVHZload, typ.UInt16)
				v2.AuxInt = i0
				v2.Aux = s
				v2.AddArg(p)
				v2.AddArg(mem)
				v1.AddArg(v2)
				v0.AddArg(v1)
				v0.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORW x1:(MOVBZloadidx [i1] {s} p idx mem) sh:(SLWconst [8] x0:(MOVBZloadidx [i0] {s} p idx mem)))
	// cond: i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVHZloadidx [i0] {s} p idx mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x1 := v_0
			if x1.Op != OpS390XMOVBZloadidx {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			x1_0 := x1.Args[0]
			x1_1 := x1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x1_0, x1_1 = _i1+1, x1_1, x1_0 {
				p := x1_0
				idx := x1_1
				sh := v_1
				if sh.Op != OpS390XSLWconst || sh.AuxInt != 8 {
					continue
				}
				x0 := sh.Args[0]
				if x0.Op != OpS390XMOVBZloadidx {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[2]
				x0_0 := x0.Args[0]
				x0_1 := x0.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, x0_0, x0_1 = _i2+1, x0_1, x0_0 {
					if p != x0_0 || idx != x0_1 || mem != x0.Args[2] || !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = i0
					v0.Aux = s
					v0.AddArg(p)
					v0.AddArg(idx)
					v0.AddArg(mem)
					return true
				}
			}
		}
		break
	}
	// match: (ORW x1:(MOVHZloadidx [i1] {s} p idx mem) sh:(SLWconst [16] x0:(MOVHZloadidx [i0] {s} p idx mem)))
	// cond: i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVWZloadidx [i0] {s} p idx mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x1 := v_0
			if x1.Op != OpS390XMOVHZloadidx {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			x1_0 := x1.Args[0]
			x1_1 := x1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x1_0, x1_1 = _i1+1, x1_1, x1_0 {
				p := x1_0
				idx := x1_1
				sh := v_1
				if sh.Op != OpS390XSLWconst || sh.AuxInt != 16 {
					continue
				}
				x0 := sh.Args[0]
				if x0.Op != OpS390XMOVHZloadidx {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[2]
				x0_0 := x0.Args[0]
				x0_1 := x0.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, x0_0, x0_1 = _i2+1, x0_1, x0_0 {
					if p != x0_0 || idx != x0_1 || mem != x0.Args[2] || !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = i0
					v0.Aux = s
					v0.AddArg(p)
					v0.AddArg(idx)
					v0.AddArg(mem)
					return true
				}
			}
		}
		break
	}
	// match: (ORW s0:(SLWconst [j0] x0:(MOVBZloadidx [i0] {s} p idx mem)) or:(ORW s1:(SLWconst [j1] x1:(MOVBZloadidx [i1] {s} p idx mem)) y))
	// cond: i1 == i0+1 && j1 == j0-8 && j1 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0, x1, s0, s1, or)
	// result: @mergePoint(b,x0,x1,y) (ORW <v.Type> (SLWconst <v.Type> [j1] (MOVHZloadidx [i0] {s} p idx mem)) y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s0 := v_0
			if s0.Op != OpS390XSLWconst {
				continue
			}
			j0 := s0.AuxInt
			x0 := s0.Args[0]
			if x0.Op != OpS390XMOVBZloadidx {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			x0_0 := x0.Args[0]
			x0_1 := x0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x0_0, x0_1 = _i1+1, x0_1, x0_0 {
				p := x0_0
				idx := x0_1
				or := v_1
				if or.Op != OpS390XORW {
					continue
				}
				_ = or.Args[1]
				or_0 := or.Args[0]
				or_1 := or.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, or_0, or_1 = _i2+1, or_1, or_0 {
					s1 := or_0
					if s1.Op != OpS390XSLWconst {
						continue
					}
					j1 := s1.AuxInt
					x1 := s1.Args[0]
					if x1.Op != OpS390XMOVBZloadidx {
						continue
					}
					i1 := x1.AuxInt
					if x1.Aux != s {
						continue
					}
					_ = x1.Args[2]
					x1_0 := x1.Args[0]
					x1_1 := x1.Args[1]
					for _i3 := 0; _i3 <= 1; _i3, x1_0, x1_1 = _i3+1, x1_1, x1_0 {
						if p != x1_0 || idx != x1_1 || mem != x1.Args[2] {
							continue
						}
						y := or_1
						if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0, x1, s0, s1, or)) {
							continue
						}
						b = mergePoint(b, x0, x1, y)
						v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
						v1.AuxInt = j1
						v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
						v2.AuxInt = i0
						v2.Aux = s
						v2.AddArg(p)
						v2.AddArg(idx)
						v2.AddArg(mem)
						v1.AddArg(v2)
						v0.AddArg(v1)
						v0.AddArg(y)
						return true
					}
				}
			}
		}
		break
	}
	// match: (ORW x0:(MOVBZload [i0] {s} p mem) sh:(SLWconst [8] x1:(MOVBZload [i1] {s} p mem)))
	// cond: p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVHZreg (MOVHBRload [i0] {s} p mem))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x0 := v_0
			if x0.Op != OpS390XMOVBZload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			sh := v_1
			if sh.Op != OpS390XSLWconst || sh.AuxInt != 8 {
				continue
			}
			x1 := sh.Args[0]
			if x1.Op != OpS390XMOVBZload {
				continue
			}
			i1 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpS390XMOVHZreg, typ.UInt64)
			v.reset(OpCopy)
			v.AddArg(v0)
			v1 := b.NewValue0(x1.Pos, OpS390XMOVHBRload, typ.UInt16)
			v1.AuxInt = i0
			v1.Aux = s
			v1.AddArg(p)
			v1.AddArg(mem)
			v0.AddArg(v1)
			return true
		}
		break
	}
	// match: (ORW r0:(MOVHZreg x0:(MOVHBRload [i0] {s} p mem)) sh:(SLWconst [16] r1:(MOVHZreg x1:(MOVHBRload [i1] {s} p mem))))
	// cond: i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, r0, r1, sh)
	// result: @mergePoint(b,x0,x1) (MOVWBRload [i0] {s} p mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			r0 := v_0
			if r0.Op != OpS390XMOVHZreg {
				continue
			}
			x0 := r0.Args[0]
			if x0.Op != OpS390XMOVHBRload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			sh := v_1
			if sh.Op != OpS390XSLWconst || sh.AuxInt != 16 {
				continue
			}
			r1 := sh.Args[0]
			if r1.Op != OpS390XMOVHZreg {
				continue
			}
			x1 := r1.Args[0]
			if x1.Op != OpS390XMOVHBRload {
				continue
			}
			i1 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, r0, r1, sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpS390XMOVWBRload, typ.UInt32)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (ORW s1:(SLWconst [j1] x1:(MOVBZload [i1] {s} p mem)) or:(ORW s0:(SLWconst [j0] x0:(MOVBZload [i0] {s} p mem)) y))
	// cond: p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0, x1, s0, s1, or)
	// result: @mergePoint(b,x0,x1,y) (ORW <v.Type> (SLWconst <v.Type> [j0] (MOVHZreg (MOVHBRload [i0] {s} p mem))) y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s1 := v_0
			if s1.Op != OpS390XSLWconst {
				continue
			}
			j1 := s1.AuxInt
			x1 := s1.Args[0]
			if x1.Op != OpS390XMOVBZload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			or := v_1
			if or.Op != OpS390XORW {
				continue
			}
			_ = or.Args[1]
			or_0 := or.Args[0]
			or_1 := or.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, or_0, or_1 = _i1+1, or_1, or_0 {
				s0 := or_0
				if s0.Op != OpS390XSLWconst {
					continue
				}
				j0 := s0.AuxInt
				x0 := s0.Args[0]
				if x0.Op != OpS390XMOVBZload {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[1]
				if p != x0.Args[0] || mem != x0.Args[1] {
					continue
				}
				y := or_1
				if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0, x1, s0, s1, or)) {
					continue
				}
				b = mergePoint(b, x0, x1, y)
				v0 := b.NewValue0(x0.Pos, OpS390XORW, v.Type)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x0.Pos, OpS390XSLWconst, v.Type)
				v1.AuxInt = j0
				v2 := b.NewValue0(x0.Pos, OpS390XMOVHZreg, typ.UInt64)
				v3 := b.NewValue0(x0.Pos, OpS390XMOVHBRload, typ.UInt16)
				v3.AuxInt = i0
				v3.Aux = s
				v3.AddArg(p)
				v3.AddArg(mem)
				v2.AddArg(v3)
				v1.AddArg(v2)
				v0.AddArg(v1)
				v0.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORW x0:(MOVBZloadidx [i0] {s} p idx mem) sh:(SLWconst [8] x1:(MOVBZloadidx [i1] {s} p idx mem)))
	// cond: p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, sh)
	// result: @mergePoint(b,x0,x1) (MOVHZreg (MOVHBRloadidx [i0] {s} p idx mem))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x0 := v_0
			if x0.Op != OpS390XMOVBZloadidx {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			x0_0 := x0.Args[0]
			x0_1 := x0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x0_0, x0_1 = _i1+1, x0_1, x0_0 {
				p := x0_0
				idx := x0_1
				sh := v_1
				if sh.Op != OpS390XSLWconst || sh.AuxInt != 8 {
					continue
				}
				x1 := sh.Args[0]
				if x1.Op != OpS390XMOVBZloadidx {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[2]
				x1_0 := x1.Args[0]
				x1_1 := x1.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, x1_0, x1_1 = _i2+1, x1_1, x1_0 {
					if p != x1_0 || idx != x1_1 || mem != x1.Args[2] || !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
					v.reset(OpCopy)
					v.AddArg(v0)
					v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
					v1.AuxInt = i0
					v1.Aux = s
					v1.AddArg(p)
					v1.AddArg(idx)
					v1.AddArg(mem)
					v0.AddArg(v1)
					return true
				}
			}
		}
		break
	}
	// match: (ORW r0:(MOVHZreg x0:(MOVHBRloadidx [i0] {s} p idx mem)) sh:(SLWconst [16] r1:(MOVHZreg x1:(MOVHBRloadidx [i1] {s} p idx mem))))
	// cond: i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0, x1, r0, r1, sh)
	// result: @mergePoint(b,x0,x1) (MOVWBRloadidx [i0] {s} p idx mem)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			r0 := v_0
			if r0.Op != OpS390XMOVHZreg {
				continue
			}
			x0 := r0.Args[0]
			if x0.Op != OpS390XMOVHBRloadidx {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			x0_0 := x0.Args[0]
			x0_1 := x0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x0_0, x0_1 = _i1+1, x0_1, x0_0 {
				p := x0_0
				idx := x0_1
				sh := v_1
				if sh.Op != OpS390XSLWconst || sh.AuxInt != 16 {
					continue
				}
				r1 := sh.Args[0]
				if r1.Op != OpS390XMOVHZreg {
					continue
				}
				x1 := r1.Args[0]
				if x1.Op != OpS390XMOVHBRloadidx {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[2]
				x1_0 := x1.Args[0]
				x1_1 := x1.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, x1_0, x1_1 = _i2+1, x1_1, x1_0 {
					if p != x1_0 || idx != x1_1 || mem != x1.Args[2] || !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0, x1, r0, r1, sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = i0
					v0.Aux = s
					v0.AddArg(p)
					v0.AddArg(idx)
					v0.AddArg(mem)
					return true
				}
			}
		}
		break
	}
	// match: (ORW s1:(SLWconst [j1] x1:(MOVBZloadidx [i1] {s} p idx mem)) or:(ORW s0:(SLWconst [j0] x0:(MOVBZloadidx [i0] {s} p idx mem)) y))
	// cond: p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0, x1, s0, s1, or)
	// result: @mergePoint(b,x0,x1,y) (ORW <v.Type> (SLWconst <v.Type> [j0] (MOVHZreg (MOVHBRloadidx [i0] {s} p idx mem))) y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s1 := v_0
			if s1.Op != OpS390XSLWconst {
				continue
			}
			j1 := s1.AuxInt
			x1 := s1.Args[0]
			if x1.Op != OpS390XMOVBZloadidx {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			x1_0 := x1.Args[0]
			x1_1 := x1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, x1_0, x1_1 = _i1+1, x1_1, x1_0 {
				p := x1_0
				idx := x1_1
				or := v_1
				if or.Op != OpS390XORW {
					continue
				}
				_ = or.Args[1]
				or_0 := or.Args[0]
				or_1 := or.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, or_0, or_1 = _i2+1, or_1, or_0 {
					s0 := or_0
					if s0.Op != OpS390XSLWconst {
						continue
					}
					j0 := s0.AuxInt
					x0 := s0.Args[0]
					if x0.Op != OpS390XMOVBZloadidx {
						continue
					}
					i0 := x0.AuxInt
					if x0.Aux != s {
						continue
					}
					_ = x0.Args[2]
					x0_0 := x0.Args[0]
					x0_1 := x0.Args[1]
					for _i3 := 0; _i3 <= 1; _i3, x0_0, x0_1 = _i3+1, x0_1, x0_0 {
						if p != x0_0 || idx != x0_1 || mem != x0.Args[2] {
							continue
						}
						y := or_1
						if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0, x1, s0, s1, or)) {
							continue
						}
						b = mergePoint(b, x0, x1, y)
						v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
						v1.AuxInt = j0
						v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
						v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
						v3.AuxInt = i0
						v3.Aux = s
						v3.AddArg(p)
						v3.AddArg(idx)
						v3.AddArg(mem)
						v2.AddArg(v3)
						v1.AddArg(v2)
						v0.AddArg(v1)
						v0.AddArg(y)
						return true
					}
				}
			}
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XORWconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ORWconst [c] x)
	// cond: int32(c)==0
	// result: x
	for {
		c := v.AuxInt
		x := v_0
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ORWconst [c] _)
	// cond: int32(c)==-1
	// result: (MOVDconst [-1])
	for {
		c := v.AuxInt
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = -1
		return true
	}
	// match: (ORWconst [c] (MOVDconst [d]))
	// result: (MOVDconst [c|d])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c | d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORWload(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ORWload [off1] {sym} x (ADDconst [off2] ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(off1+off2)
	// result: (ORWload [off1+off2] {sym} x ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XORWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (ORWload [o1] {s1} x (MOVDaddr [o2] {s2} ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)
	// result: (ORWload [o1+o2] {mergeSym(s1, s2)} x ptr mem)
	for {
		o1 := v.AuxInt
		s1 := v.Aux
		x := v_0
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XORWload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ORconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ORconst [-1] _)
	// result: (MOVDconst [-1])
	for {
		if v.AuxInt != -1 {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = -1
		return true
	}
	// match: (ORconst [c] (MOVDconst [d]))
	// result: (MOVDconst [c|d])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c | d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORload(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ORload <t> [off] {sym} x ptr1 (FMOVDstore [off] {sym} ptr2 y _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (OR x (LGDR <t> y))
	for {
		t := v.Type
		off := v.AuxInt
		sym := v.Aux
		x := v_0
		ptr1 := v_1
		if v_2.Op != OpS390XFMOVDstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		y := v_2.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XOR)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpS390XLGDR, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ORload [off1] {sym} x (ADDconst [off2] ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(off1+off2)
	// result: (ORload [off1+off2] {sym} x ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XORload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (ORload [o1] {s1} x (MOVDaddr [o2] {s2} ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)
	// result: (ORload [o1+o2] {mergeSym(s1, s2)} x ptr mem)
	for {
		o1 := v.AuxInt
		s1 := v.Aux
		x := v_0
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XORload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XRLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (RLL x (MOVDconst [c]))
	// result: (RLLconst x [c&31])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XRLLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XRLLG(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (RLLG x (MOVDconst [c]))
	// result: (RLLGconst x [c&63])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XRLLGconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSLD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SLD x (MOVDconst [c]))
	// result: (SLDconst x [c&63])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSLDconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (SLD x (AND (MOVDconst [c]) y))
	// result: (SLD x (ANDWconst <typ.UInt32> [c&63] y))
	for {
		x := v_0
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			if v_1_0.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1_0.AuxInt
			y := v_1_1
			v.reset(OpS390XSLD)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
			v0.AuxInt = c & 63
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SLD x (ANDWconst [c] y))
	// cond: c&63 == 63
	// result: (SLD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XANDWconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SLD x (MOVWreg y))
	// result: (SLD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SLD x (MOVHreg y))
	// result: (SLD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SLD x (MOVBreg y))
	// result: (SLD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SLD x (MOVWZreg y))
	// result: (SLD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SLD x (MOVHZreg y))
	// result: (SLD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SLD x (MOVBZreg y))
	// result: (SLD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSLW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SLW x (MOVDconst [c]))
	// result: (SLWconst x [c&63])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSLWconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (SLW x (AND (MOVDconst [c]) y))
	// result: (SLW x (ANDWconst <typ.UInt32> [c&63] y))
	for {
		x := v_0
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			if v_1_0.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1_0.AuxInt
			y := v_1_1
			v.reset(OpS390XSLW)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
			v0.AuxInt = c & 63
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SLW x (ANDWconst [c] y))
	// cond: c&63 == 63
	// result: (SLW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XANDWconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SLW x (MOVWreg y))
	// result: (SLW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SLW x (MOVHreg y))
	// result: (SLW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SLW x (MOVBreg y))
	// result: (SLW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SLW x (MOVWZreg y))
	// result: (SLW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SLW x (MOVHZreg y))
	// result: (SLW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SLW x (MOVBZreg y))
	// result: (SLW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRAD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SRAD x (MOVDconst [c]))
	// result: (SRADconst x [c&63])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSRADconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (SRAD x (AND (MOVDconst [c]) y))
	// result: (SRAD x (ANDWconst <typ.UInt32> [c&63] y))
	for {
		x := v_0
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			if v_1_0.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1_0.AuxInt
			y := v_1_1
			v.reset(OpS390XSRAD)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
			v0.AuxInt = c & 63
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SRAD x (ANDWconst [c] y))
	// cond: c&63 == 63
	// result: (SRAD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XANDWconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAD x (MOVWreg y))
	// result: (SRAD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAD x (MOVHreg y))
	// result: (SRAD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAD x (MOVBreg y))
	// result: (SRAD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAD x (MOVWZreg y))
	// result: (SRAD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAD x (MOVHZreg y))
	// result: (SRAD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAD x (MOVBZreg y))
	// result: (SRAD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRADconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRADconst [c] (MOVDconst [d]))
	// result: (MOVDconst [d>>uint64(c)])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = d >> uint64(c)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRAW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SRAW x (MOVDconst [c]))
	// result: (SRAWconst x [c&63])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSRAWconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (SRAW x (AND (MOVDconst [c]) y))
	// result: (SRAW x (ANDWconst <typ.UInt32> [c&63] y))
	for {
		x := v_0
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			if v_1_0.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1_0.AuxInt
			y := v_1_1
			v.reset(OpS390XSRAW)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
			v0.AuxInt = c & 63
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SRAW x (ANDWconst [c] y))
	// cond: c&63 == 63
	// result: (SRAW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XANDWconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAW x (MOVWreg y))
	// result: (SRAW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAW x (MOVHreg y))
	// result: (SRAW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAW x (MOVBreg y))
	// result: (SRAW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAW x (MOVWZreg y))
	// result: (SRAW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAW x (MOVHZreg y))
	// result: (SRAW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAW x (MOVBZreg y))
	// result: (SRAW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRAWconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRAWconst [c] (MOVDconst [d]))
	// result: (MOVDconst [int64(int32(d))>>uint64(c)])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int32(d)) >> uint64(c)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SRD x (MOVDconst [c]))
	// result: (SRDconst x [c&63])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSRDconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (SRD x (AND (MOVDconst [c]) y))
	// result: (SRD x (ANDWconst <typ.UInt32> [c&63] y))
	for {
		x := v_0
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			if v_1_0.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1_0.AuxInt
			y := v_1_1
			v.reset(OpS390XSRD)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
			v0.AuxInt = c & 63
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SRD x (ANDWconst [c] y))
	// cond: c&63 == 63
	// result: (SRD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XANDWconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRD x (MOVWreg y))
	// result: (SRD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRD x (MOVHreg y))
	// result: (SRD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRD x (MOVBreg y))
	// result: (SRD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRD x (MOVWZreg y))
	// result: (SRD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRD x (MOVHZreg y))
	// result: (SRD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRD x (MOVBZreg y))
	// result: (SRD x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRDconst(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (SRDconst [1] (SLDconst [1] (LGDR <t> x)))
	// result: (LGDR <t> (LPDFR <x.Type> x))
	for {
		if v.AuxInt != 1 || v_0.Op != OpS390XSLDconst || v_0.AuxInt != 1 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XLGDR {
			break
		}
		t := v_0_0.Type
		x := v_0_0.Args[0]
		v.reset(OpS390XLGDR)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XLPDFR, x.Type)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SRW x (MOVDconst [c]))
	// result: (SRWconst x [c&63])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSRWconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (SRW x (AND (MOVDconst [c]) y))
	// result: (SRW x (ANDWconst <typ.UInt32> [c&63] y))
	for {
		x := v_0
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			if v_1_0.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1_0.AuxInt
			y := v_1_1
			v.reset(OpS390XSRW)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
			v0.AuxInt = c & 63
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SRW x (ANDWconst [c] y))
	// cond: c&63 == 63
	// result: (SRW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XANDWconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRW x (MOVWreg y))
	// result: (SRW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRW x (MOVHreg y))
	// result: (SRW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRW x (MOVBreg y))
	// result: (SRW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRW x (MOVWZreg y))
	// result: (SRW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRW x (MOVHZreg y))
	// result: (SRW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRW x (MOVBZreg y))
	// result: (SRW x y)
	for {
		x := v_0
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSTM2(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (STM2 [i] {s} p w2 w3 x:(STM2 [i-8] {s} p w0 w1 mem))
	// cond: x.Uses == 1 && is20Bit(i-8) && clobber(x)
	// result: (STM4 [i-8] {s} p w0 w1 w2 w3 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w2 := v_1
		w3 := v_2
		x := v_3
		if x.Op != OpS390XSTM2 || x.AuxInt != i-8 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		w1 := x.Args[2]
		if !(x.Uses == 1 && is20Bit(i-8) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTM4)
		v.AuxInt = i - 8
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(w2)
		v.AddArg(w3)
		v.AddArg(mem)
		return true
	}
	// match: (STM2 [i] {s} p (SRDconst [32] x) x mem)
	// result: (MOVDstore [i] {s} p x mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpS390XSRDconst || v_1.AuxInt != 32 {
			break
		}
		x := v_1.Args[0]
		if x != v_2 {
			break
		}
		mem := v_3
		v.reset(OpS390XMOVDstore)
		v.AuxInt = i
		v.Aux = s
		v.AddArg(p)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSTMG2(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (STMG2 [i] {s} p w2 w3 x:(STMG2 [i-16] {s} p w0 w1 mem))
	// cond: x.Uses == 1 && is20Bit(i-16) && clobber(x)
	// result: (STMG4 [i-16] {s} p w0 w1 w2 w3 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		p := v_0
		w2 := v_1
		w3 := v_2
		x := v_3
		if x.Op != OpS390XSTMG2 || x.AuxInt != i-16 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		w1 := x.Args[2]
		if !(x.Uses == 1 && is20Bit(i-16) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTMG4)
		v.AuxInt = i - 16
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(w2)
		v.AddArg(w3)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSUB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUB x (MOVDconst [c]))
	// cond: is32Bit(c)
	// result: (SUBconst x [c])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpS390XSUBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SUB (MOVDconst [c]) x)
	// cond: is32Bit(c)
	// result: (NEG (SUBconst <v.Type> x [c]))
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpS390XNEG)
		v0 := b.NewValue0(v.Pos, OpS390XSUBconst, v.Type)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUB x x)
	// result: (MOVDconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (SUB <t> x g:(MOVDload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (SUBload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		x := v_0
		g := v_1
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		mem := g.Args[1]
		ptr := g.Args[0]
		if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XSUBload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSUBE(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUBE x y (FlagGT))
	// result: (SUBC x y)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpS390XFlagGT {
			break
		}
		v.reset(OpS390XSUBC)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBE x y (FlagOV))
	// result: (SUBC x y)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpS390XFlagOV {
			break
		}
		v.reset(OpS390XSUBC)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBE x y (Select1 (SUBC (MOVDconst [0]) (NEG (Select0 (SUBE (MOVDconst [0]) (MOVDconst [0]) c))))))
	// result: (SUBE x y c)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpSelect1 {
			break
		}
		v_2_0 := v_2.Args[0]
		if v_2_0.Op != OpS390XSUBC {
			break
		}
		_ = v_2_0.Args[1]
		v_2_0_0 := v_2_0.Args[0]
		if v_2_0_0.Op != OpS390XMOVDconst || v_2_0_0.AuxInt != 0 {
			break
		}
		v_2_0_1 := v_2_0.Args[1]
		if v_2_0_1.Op != OpS390XNEG {
			break
		}
		v_2_0_1_0 := v_2_0_1.Args[0]
		if v_2_0_1_0.Op != OpSelect0 {
			break
		}
		v_2_0_1_0_0 := v_2_0_1_0.Args[0]
		if v_2_0_1_0_0.Op != OpS390XSUBE {
			break
		}
		c := v_2_0_1_0_0.Args[2]
		v_2_0_1_0_0_0 := v_2_0_1_0_0.Args[0]
		if v_2_0_1_0_0_0.Op != OpS390XMOVDconst || v_2_0_1_0_0_0.AuxInt != 0 {
			break
		}
		v_2_0_1_0_0_1 := v_2_0_1_0_0.Args[1]
		if v_2_0_1_0_0_1.Op != OpS390XMOVDconst || v_2_0_1_0_0_1.AuxInt != 0 {
			break
		}
		v.reset(OpS390XSUBE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(c)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSUBW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBW x (MOVDconst [c]))
	// result: (SUBWconst x [int64(int32(c))])
	for {
		x := v_0
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSUBWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (SUBW (MOVDconst [c]) x)
	// result: (NEGW (SUBWconst <v.Type> x [int64(int32(c))]))
	for {
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpS390XNEGW)
		v0 := b.NewValue0(v.Pos, OpS390XSUBWconst, v.Type)
		v0.AuxInt = int64(int32(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBW x x)
	// result: (MOVDconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (SUBW <t> x g:(MOVWload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (SUBWload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		x := v_0
		g := v_1
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		mem := g.Args[1]
		ptr := g.Args[0]
		if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XSUBWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (SUBW <t> x g:(MOVWZload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (SUBWload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		x := v_0
		g := v_1
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		mem := g.Args[1]
		ptr := g.Args[0]
		if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XSUBWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSUBWconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SUBWconst [c] x)
	// cond: int32(c) == 0
	// result: x
	for {
		c := v.AuxInt
		x := v_0
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (SUBWconst [c] x)
	// result: (ADDWconst [int64(int32(-c))] x)
	for {
		c := v.AuxInt
		x := v_0
		v.reset(OpS390XADDWconst)
		v.AuxInt = int64(int32(-c))
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpS390XSUBWload(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUBWload [off1] {sym} x (ADDconst [off2] ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(off1+off2)
	// result: (SUBWload [off1+off2] {sym} x ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XSUBWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (SUBWload [o1] {s1} x (MOVDaddr [o2] {s2} ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)
	// result: (SUBWload [o1+o2] {mergeSym(s1, s2)} x ptr mem)
	for {
		o1 := v.AuxInt
		s1 := v.Aux
		x := v_0
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XSUBWload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSUBconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SUBconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] x)
	// cond: c != -(1<<31)
	// result: (ADDconst [-c] x)
	for {
		c := v.AuxInt
		x := v_0
		if !(c != -(1 << 31)) {
			break
		}
		v.reset(OpS390XADDconst)
		v.AuxInt = -c
		v.AddArg(x)
		return true
	}
	// match: (SUBconst (MOVDconst [d]) [c])
	// result: (MOVDconst [d-c])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = d - c
		return true
	}
	// match: (SUBconst (SUBconst x [d]) [c])
	// cond: is32Bit(-c-d)
	// result: (ADDconst [-c-d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(-c - d)) {
			break
		}
		v.reset(OpS390XADDconst)
		v.AuxInt = -c - d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSUBload(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBload <t> [off] {sym} x ptr1 (FMOVDstore [off] {sym} ptr2 y _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (SUB x (LGDR <t> y))
	for {
		t := v.Type
		off := v.AuxInt
		sym := v.Aux
		x := v_0
		ptr1 := v_1
		if v_2.Op != OpS390XFMOVDstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		y := v_2.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XSUB)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpS390XLGDR, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SUBload [off1] {sym} x (ADDconst [off2] ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(off1+off2)
	// result: (SUBload [off1+off2] {sym} x ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XSUBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (SUBload [o1] {s1} x (MOVDaddr [o2] {s2} ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)
	// result: (SUBload [o1+o2] {mergeSym(s1, s2)} x ptr mem)
	for {
		o1 := v.AuxInt
		s1 := v.Aux
		x := v_0
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XSUBload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSumBytes2(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SumBytes2 x)
	// result: (ADDW (SRWconst <typ.UInt8> x [8]) x)
	for {
		x := v_0
		v.reset(OpS390XADDW)
		v0 := b.NewValue0(v.Pos, OpS390XSRWconst, typ.UInt8)
		v0.AuxInt = 8
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpS390XSumBytes4(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SumBytes4 x)
	// result: (SumBytes2 (ADDW <typ.UInt16> (SRWconst <typ.UInt16> x [16]) x))
	for {
		x := v_0
		v.reset(OpS390XSumBytes2)
		v0 := b.NewValue0(v.Pos, OpS390XADDW, typ.UInt16)
		v1 := b.NewValue0(v.Pos, OpS390XSRWconst, typ.UInt16)
		v1.AuxInt = 16
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpS390XSumBytes8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SumBytes8 x)
	// result: (SumBytes4 (ADDW <typ.UInt32> (SRDconst <typ.UInt32> x [32]) x))
	for {
		x := v_0
		v.reset(OpS390XSumBytes4)
		v0 := b.NewValue0(v.Pos, OpS390XADDW, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpS390XSRDconst, typ.UInt32)
		v1.AuxInt = 32
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpS390XXOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (XOR x (MOVDconst [c]))
	// cond: isU32Bit(c)
	// result: (XORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1.AuxInt
			if !(isU32Bit(c)) {
				continue
			}
			v.reset(OpS390XXORconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XOR (SLDconst x [c]) (SRDconst x [d]))
	// cond: d == 64-c
	// result: (RLLGconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XSLDconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			if v_1.Op != OpS390XSRDconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 64-c) {
				continue
			}
			v.reset(OpS390XRLLGconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XOR (MOVDconst [c]) (MOVDconst [d]))
	// result: (MOVDconst [c^d])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XMOVDconst {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpS390XMOVDconst)
			v.AuxInt = c ^ d
			return true
		}
		break
	}
	// match: (XOR x x)
	// result: (MOVDconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (XOR <t> x g:(MOVDload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (XORload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVDload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XXORload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XXORW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (XORW x (MOVDconst [c]))
	// result: (XORWconst [int64(int32(c))] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpS390XMOVDconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpS390XXORWconst)
			v.AuxInt = int64(int32(c))
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XORW (SLWconst x [c]) (SRWconst x [d]))
	// cond: d == 32-c
	// result: (RLLconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpS390XSLWconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			if v_1.Op != OpS390XSRWconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 32-c) {
				continue
			}
			v.reset(OpS390XRLLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XORW x x)
	// result: (MOVDconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (XORW <t> x g:(MOVWload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (XORWload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVWload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XXORWload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (XORW <t> x g:(MOVWZload [off] {sym} ptr mem))
	// cond: ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)
	// result: (XORWload <t> [off] {sym} x ptr mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			g := v_1
			if g.Op != OpS390XMOVWZload {
				continue
			}
			off := g.AuxInt
			sym := g.Aux
			mem := g.Args[1]
			ptr := g.Args[0]
			if !(ptr.Op != OpSB && is20Bit(off) && canMergeLoadClobber(v, g, x) && clobber(g)) {
				continue
			}
			v.reset(OpS390XXORWload)
			v.Type = t
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueS390X_OpS390XXORWconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (XORWconst [c] x)
	// cond: int32(c)==0
	// result: x
	for {
		c := v.AuxInt
		x := v_0
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (XORWconst [c] (MOVDconst [d]))
	// result: (MOVDconst [c^d])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c ^ d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XXORWload(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (XORWload [off1] {sym} x (ADDconst [off2] ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(off1+off2)
	// result: (XORWload [off1+off2] {sym} x ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XXORWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (XORWload [o1] {s1} x (MOVDaddr [o2] {s2} ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)
	// result: (XORWload [o1+o2] {mergeSym(s1, s2)} x ptr mem)
	for {
		o1 := v.AuxInt
		s1 := v.Aux
		x := v_0
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XXORWload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XXORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (XORconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (XORconst [c] (MOVDconst [d]))
	// result: (MOVDconst [c^d])
	for {
		c := v.AuxInt
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c ^ d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XXORload(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (XORload <t> [off] {sym} x ptr1 (FMOVDstore [off] {sym} ptr2 y _))
	// cond: isSamePtr(ptr1, ptr2)
	// result: (XOR x (LGDR <t> y))
	for {
		t := v.Type
		off := v.AuxInt
		sym := v.Aux
		x := v_0
		ptr1 := v_1
		if v_2.Op != OpS390XFMOVDstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		y := v_2.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XXOR)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpS390XLGDR, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (XORload [off1] {sym} x (ADDconst [off2] ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(off1+off2)
	// result: (XORload [off1+off2] {sym} x ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v_0
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XXORload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (XORload [o1] {s1} x (MOVDaddr [o2] {s2} ptr) mem)
	// cond: ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)
	// result: (XORload [o1+o2] {mergeSym(s1, s2)} x ptr mem)
	for {
		o1 := v.AuxInt
		s1 := v.Aux
		x := v_0
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v_2
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XXORload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpSelect0(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Select0 (Add64carry x y c))
	// result: (Select0 <typ.UInt64> (ADDE x y (Select1 <types.TypeFlags> (ADDCconst c [-1]))))
	for {
		if v_0.Op != OpAdd64carry {
			break
		}
		c := v_0.Args[2]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpSelect0)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpS390XADDE, types.NewTuple(typ.UInt64, types.TypeFlags))
		v0.AddArg(x)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v2 := b.NewValue0(v.Pos, OpS390XADDCconst, types.NewTuple(typ.UInt64, types.TypeFlags))
		v2.AuxInt = -1
		v2.AddArg(c)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Select0 (Sub64borrow x y c))
	// result: (Select0 <typ.UInt64> (SUBE x y (Select1 <types.TypeFlags> (SUBC (MOVDconst [0]) c))))
	for {
		if v_0.Op != OpSub64borrow {
			break
		}
		c := v_0.Args[2]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpSelect0)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpS390XSUBE, types.NewTuple(typ.UInt64, types.TypeFlags))
		v0.AddArg(x)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v2 := b.NewValue0(v.Pos, OpS390XSUBC, types.NewTuple(typ.UInt64, types.TypeFlags))
		v3 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v3.AuxInt = 0
		v2.AddArg(v3)
		v2.AddArg(c)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Select0 <t> (AddTupleFirst32 val tuple))
	// result: (ADDW val (Select0 <t> tuple))
	for {
		t := v.Type
		if v_0.Op != OpS390XAddTupleFirst32 {
			break
		}
		tuple := v_0.Args[1]
		val := v_0.Args[0]
		v.reset(OpS390XADDW)
		v.AddArg(val)
		v0 := b.NewValue0(v.Pos, OpSelect0, t)
		v0.AddArg(tuple)
		v.AddArg(v0)
		return true
	}
	// match: (Select0 <t> (AddTupleFirst64 val tuple))
	// result: (ADD val (Select0 <t> tuple))
	for {
		t := v.Type
		if v_0.Op != OpS390XAddTupleFirst64 {
			break
		}
		tuple := v_0.Args[1]
		val := v_0.Args[0]
		v.reset(OpS390XADD)
		v.AddArg(val)
		v0 := b.NewValue0(v.Pos, OpSelect0, t)
		v0.AddArg(tuple)
		v.AddArg(v0)
		return true
	}
	// match: (Select0 (ADDCconst (MOVDconst [c]) [d]))
	// result: (MOVDconst [c+d])
	for {
		if v_0.Op != OpS390XADDCconst {
			break
		}
		d := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c + d
		return true
	}
	// match: (Select0 (SUBC (MOVDconst [c]) (MOVDconst [d])))
	// result: (MOVDconst [c-d])
	for {
		if v_0.Op != OpS390XSUBC {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpS390XMOVDconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c - d
		return true
	}
	return false
}
func rewriteValueS390X_OpSelect1(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Select1 (Add64carry x y c))
	// result: (Select0 <typ.UInt64> (ADDE (MOVDconst [0]) (MOVDconst [0]) (Select1 <types.TypeFlags> (ADDE x y (Select1 <types.TypeFlags> (ADDCconst c [-1]))))))
	for {
		if v_0.Op != OpAdd64carry {
			break
		}
		c := v_0.Args[2]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpSelect0)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpS390XADDE, types.NewTuple(typ.UInt64, types.TypeFlags))
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v0.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v4 := b.NewValue0(v.Pos, OpS390XADDE, types.NewTuple(typ.UInt64, types.TypeFlags))
		v4.AddArg(x)
		v4.AddArg(y)
		v5 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v6 := b.NewValue0(v.Pos, OpS390XADDCconst, types.NewTuple(typ.UInt64, types.TypeFlags))
		v6.AuxInt = -1
		v6.AddArg(c)
		v5.AddArg(v6)
		v4.AddArg(v5)
		v3.AddArg(v4)
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
	// match: (Select1 (Sub64borrow x y c))
	// result: (NEG (Select0 <typ.UInt64> (SUBE (MOVDconst [0]) (MOVDconst [0]) (Select1 <types.TypeFlags> (SUBE x y (Select1 <types.TypeFlags> (SUBC (MOVDconst [0]) c)))))))
	for {
		if v_0.Op != OpSub64borrow {
			break
		}
		c := v_0.Args[2]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpS390XNEG)
		v0 := b.NewValue0(v.Pos, OpSelect0, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpS390XSUBE, types.NewTuple(typ.UInt64, types.TypeFlags))
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v3.AuxInt = 0
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v5 := b.NewValue0(v.Pos, OpS390XSUBE, types.NewTuple(typ.UInt64, types.TypeFlags))
		v5.AddArg(x)
		v5.AddArg(y)
		v6 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v7 := b.NewValue0(v.Pos, OpS390XSUBC, types.NewTuple(typ.UInt64, types.TypeFlags))
		v8 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v8.AuxInt = 0
		v7.AddArg(v8)
		v7.AddArg(c)
		v6.AddArg(v7)
		v5.AddArg(v6)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Select1 (AddTupleFirst32 _ tuple))
	// result: (Select1 tuple)
	for {
		if v_0.Op != OpS390XAddTupleFirst32 {
			break
		}
		tuple := v_0.Args[1]
		v.reset(OpSelect1)
		v.AddArg(tuple)
		return true
	}
	// match: (Select1 (AddTupleFirst64 _ tuple))
	// result: (Select1 tuple)
	for {
		if v_0.Op != OpS390XAddTupleFirst64 {
			break
		}
		tuple := v_0.Args[1]
		v.reset(OpSelect1)
		v.AddArg(tuple)
		return true
	}
	// match: (Select1 (ADDCconst (MOVDconst [c]) [d]))
	// cond: uint64(c+d) >= uint64(c) && c+d == 0
	// result: (FlagEQ)
	for {
		if v_0.Op != OpS390XADDCconst {
			break
		}
		d := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0_0.AuxInt
		if !(uint64(c+d) >= uint64(c) && c+d == 0) {
			break
		}
		v.reset(OpS390XFlagEQ)
		return true
	}
	// match: (Select1 (ADDCconst (MOVDconst [c]) [d]))
	// cond: uint64(c+d) >= uint64(c) && c+d != 0
	// result: (FlagLT)
	for {
		if v_0.Op != OpS390XADDCconst {
			break
		}
		d := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0_0.AuxInt
		if !(uint64(c+d) >= uint64(c) && c+d != 0) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}
	// match: (Select1 (SUBC (MOVDconst [c]) (MOVDconst [d])))
	// cond: uint64(d) <= uint64(c) && c-d == 0
	// result: (FlagGT)
	for {
		if v_0.Op != OpS390XSUBC {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpS390XMOVDconst {
			break
		}
		d := v_0_1.AuxInt
		if !(uint64(d) <= uint64(c) && c-d == 0) {
			break
		}
		v.reset(OpS390XFlagGT)
		return true
	}
	// match: (Select1 (SUBC (MOVDconst [c]) (MOVDconst [d])))
	// cond: uint64(d) <= uint64(c) && c-d != 0
	// result: (FlagOV)
	for {
		if v_0.Op != OpS390XSUBC {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpS390XMOVDconst {
			break
		}
		d := v_0_1.AuxInt
		if !(uint64(d) <= uint64(c) && c-d != 0) {
			break
		}
		v.reset(OpS390XFlagOV)
		return true
	}
	return false
}
func rewriteValueS390X_OpSlicemask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Slicemask <t> x)
	// result: (SRADconst (NEG <t> x) [63])
	for {
		t := v.Type
		x := v_0
		v.reset(OpS390XSRADconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpS390XNEG, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpStore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)
	// result: (FMOVDstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpS390XFMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)
	// result: (FMOVSstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpS390XFMOVSstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8
	// result: (MOVDstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 8) {
			break
		}
		v.reset(OpS390XMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4
	// result: (MOVWstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 4) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 2
	// result: (MOVHstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 2) {
			break
		}
		v.reset(OpS390XMOVHstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 1) {
			break
		}
		v.reset(OpS390XMOVBstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpTrunc(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc x)
	// result: (FIDBR [5] x)
	for {
		x := v_0
		v.reset(OpS390XFIDBR)
		v.AuxInt = 5
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpZero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Zero [0] _ mem)
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v_1
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Zero [1] destptr mem)
	// result: (MOVBstoreconst [0] destptr mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpS390XMOVBstoreconst)
		v.AuxInt = 0
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] destptr mem)
	// result: (MOVHstoreconst [0] destptr mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpS390XMOVHstoreconst)
		v.AuxInt = 0
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [4] destptr mem)
	// result: (MOVWstoreconst [0] destptr mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpS390XMOVWstoreconst)
		v.AuxInt = 0
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [8] destptr mem)
	// result: (MOVDstoreconst [0] destptr mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpS390XMOVDstoreconst)
		v.AuxInt = 0
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [3] destptr mem)
	// result: (MOVBstoreconst [makeValAndOff(0,2)] destptr (MOVHstoreconst [0] destptr mem))
	for {
		if v.AuxInt != 3 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpS390XMOVBstoreconst)
		v.AuxInt = makeValAndOff(0, 2)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHstoreconst, types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [5] destptr mem)
	// result: (MOVBstoreconst [makeValAndOff(0,4)] destptr (MOVWstoreconst [0] destptr mem))
	for {
		if v.AuxInt != 5 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpS390XMOVBstoreconst)
		v.AuxInt = makeValAndOff(0, 4)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWstoreconst, types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [6] destptr mem)
	// result: (MOVHstoreconst [makeValAndOff(0,4)] destptr (MOVWstoreconst [0] destptr mem))
	for {
		if v.AuxInt != 6 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpS390XMOVHstoreconst)
		v.AuxInt = makeValAndOff(0, 4)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWstoreconst, types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [7] destptr mem)
	// result: (MOVWstoreconst [makeValAndOff(0,3)] destptr (MOVWstoreconst [0] destptr mem))
	for {
		if v.AuxInt != 7 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpS390XMOVWstoreconst)
		v.AuxInt = makeValAndOff(0, 3)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWstoreconst, types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [s] destptr mem)
	// cond: s > 0 && s <= 1024
	// result: (CLEAR [makeValAndOff(s, 0)] destptr mem)
	for {
		s := v.AuxInt
		destptr := v_0
		mem := v_1
		if !(s > 0 && s <= 1024) {
			break
		}
		v.reset(OpS390XCLEAR)
		v.AuxInt = makeValAndOff(s, 0)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [s] destptr mem)
	// cond: s > 1024
	// result: (LoweredZero [s%256] destptr (ADDconst <destptr.Type> destptr [(s/256)*256]) mem)
	for {
		s := v.AuxInt
		destptr := v_0
		mem := v_1
		if !(s > 1024) {
			break
		}
		v.reset(OpS390XLoweredZero)
		v.AuxInt = s % 256
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpS390XADDconst, destptr.Type)
		v0.AuxInt = (s / 256) * 256
		v0.AddArg(destptr)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteBlockS390X(b *Block) bool {
	typ := &b.Func.Config.Types
	switch b.Kind {
	case BlockS390XBRC:
		// match: (BRC {c} (CMP x y) yes no)
		// result: (CGRJ {c.(s390x.CCMask)&^s390x.Unordered} x y yes no)
		for b.Controls[0].Op == OpS390XCMP {
			v_0 := b.Controls[0]
			y := v_0.Args[1]
			x := v_0.Args[0]
			c := b.Aux
			b.Reset(BlockS390XCGRJ)
			b.AddControl(x)
			b.AddControl(y)
			b.Aux = c.(s390x.CCMask) &^ s390x.Unordered
			return true
		}
		// match: (BRC {c} (CMPW x y) yes no)
		// result: (CRJ {c.(s390x.CCMask)&^s390x.Unordered} x y yes no)
		for b.Controls[0].Op == OpS390XCMPW {
			v_0 := b.Controls[0]
			y := v_0.Args[1]
			x := v_0.Args[0]
			c := b.Aux
			b.Reset(BlockS390XCRJ)
			b.AddControl(x)
			b.AddControl(y)
			b.Aux = c.(s390x.CCMask) &^ s390x.Unordered
			return true
		}
		// match: (BRC {c} (CMPU x y) yes no)
		// result: (CLGRJ {c.(s390x.CCMask)&^s390x.Unordered} x y yes no)
		for b.Controls[0].Op == OpS390XCMPU {
			v_0 := b.Controls[0]
			y := v_0.Args[1]
			x := v_0.Args[0]
			c := b.Aux
			b.Reset(BlockS390XCLGRJ)
			b.AddControl(x)
			b.AddControl(y)
			b.Aux = c.(s390x.CCMask) &^ s390x.Unordered
			return true
		}
		// match: (BRC {c} (CMPWU x y) yes no)
		// result: (CLRJ {c.(s390x.CCMask)&^s390x.Unordered} x y yes no)
		for b.Controls[0].Op == OpS390XCMPWU {
			v_0 := b.Controls[0]
			y := v_0.Args[1]
			x := v_0.Args[0]
			c := b.Aux
			b.Reset(BlockS390XCLRJ)
			b.AddControl(x)
			b.AddControl(y)
			b.Aux = c.(s390x.CCMask) &^ s390x.Unordered
			return true
		}
		// match: (BRC {c} (CMPconst x [y]) yes no)
		// cond: is8Bit(y)
		// result: (CGIJ {c.(s390x.CCMask)&^s390x.Unordered} x [int64(int8(y))] yes no)
		for b.Controls[0].Op == OpS390XCMPconst {
			v_0 := b.Controls[0]
			y := v_0.AuxInt
			x := v_0.Args[0]
			c := b.Aux
			if !(is8Bit(y)) {
				break
			}
			b.Reset(BlockS390XCGIJ)
			b.AddControl(x)
			b.AuxInt = int64(int8(y))
			b.Aux = c.(s390x.CCMask) &^ s390x.Unordered
			return true
		}
		// match: (BRC {c} (CMPWconst x [y]) yes no)
		// cond: is8Bit(y)
		// result: (CIJ {c.(s390x.CCMask)&^s390x.Unordered} x [int64(int8(y))] yes no)
		for b.Controls[0].Op == OpS390XCMPWconst {
			v_0 := b.Controls[0]
			y := v_0.AuxInt
			x := v_0.Args[0]
			c := b.Aux
			if !(is8Bit(y)) {
				break
			}
			b.Reset(BlockS390XCIJ)
			b.AddControl(x)
			b.AuxInt = int64(int8(y))
			b.Aux = c.(s390x.CCMask) &^ s390x.Unordered
			return true
		}
		// match: (BRC {c} (CMPUconst x [y]) yes no)
		// cond: isU8Bit(y)
		// result: (CLGIJ {c.(s390x.CCMask)&^s390x.Unordered} x [int64(int8(y))] yes no)
		for b.Controls[0].Op == OpS390XCMPUconst {
			v_0 := b.Controls[0]
			y := v_0.AuxInt
			x := v_0.Args[0]
			c := b.Aux
			if !(isU8Bit(y)) {
				break
			}
			b.Reset(BlockS390XCLGIJ)
			b.AddControl(x)
			b.AuxInt = int64(int8(y))
			b.Aux = c.(s390x.CCMask) &^ s390x.Unordered
			return true
		}
		// match: (BRC {c} (CMPWUconst x [y]) yes no)
		// cond: isU8Bit(y)
		// result: (CLIJ {c.(s390x.CCMask)&^s390x.Unordered} x [int64(int8(y))] yes no)
		for b.Controls[0].Op == OpS390XCMPWUconst {
			v_0 := b.Controls[0]
			y := v_0.AuxInt
			x := v_0.Args[0]
			c := b.Aux
			if !(isU8Bit(y)) {
				break
			}
			b.Reset(BlockS390XCLIJ)
			b.AddControl(x)
			b.AuxInt = int64(int8(y))
			b.Aux = c.(s390x.CCMask) &^ s390x.Unordered
			return true
		}
		// match: (BRC {s390x.Less} (CMPconst x [ 128]) yes no)
		// result: (CGIJ {s390x.LessOrEqual} x [ 127] yes no)
		for b.Controls[0].Op == OpS390XCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 128 {
				break
			}
			x := v_0.Args[0]
			if b.Aux != s390x.Less {
				break
			}
			b.Reset(BlockS390XCGIJ)
			b.AddControl(x)
			b.AuxInt = 127
			b.Aux = s390x.LessOrEqual
			return true
		}
		// match: (BRC {s390x.Less} (CMPWconst x [ 128]) yes no)
		// result: (CIJ {s390x.LessOrEqual} x [ 127] yes no)
		for b.Controls[0].Op == OpS390XCMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 128 {
				break
			}
			x := v_0.Args[0]
			if b.Aux != s390x.Less {
				break
			}
			b.Reset(BlockS390XCIJ)
			b.AddControl(x)
			b.AuxInt = 127
			b.Aux = s390x.LessOrEqual
			return true
		}
		// match: (BRC {s390x.LessOrEqual} (CMPconst x [-129]) yes no)
		// result: (CGIJ {s390x.Less} x [-128] yes no)
		for b.Controls[0].Op == OpS390XCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != -129 {
				break
			}
			x := v_0.Args[0]
			if b.Aux != s390x.LessOrEqual {
				break
			}
			b.Reset(BlockS390XCGIJ)
			b.AddControl(x)
			b.AuxInt = -128
			b.Aux = s390x.Less
			return true
		}
		// match: (BRC {s390x.LessOrEqual} (CMPWconst x [-129]) yes no)
		// result: (CIJ {s390x.Less} x [-128] yes no)
		for b.Controls[0].Op == OpS390XCMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != -129 {
				break
			}
			x := v_0.Args[0]
			if b.Aux != s390x.LessOrEqual {
				break
			}
			b.Reset(BlockS390XCIJ)
			b.AddControl(x)
			b.AuxInt = -128
			b.Aux = s390x.Less
			return true
		}
		// match: (BRC {s390x.Greater} (CMPconst x [-129]) yes no)
		// result: (CGIJ {s390x.GreaterOrEqual} x [-128] yes no)
		for b.Controls[0].Op == OpS390XCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != -129 {
				break
			}
			x := v_0.Args[0]
			if b.Aux != s390x.Greater {
				break
			}
			b.Reset(BlockS390XCGIJ)
			b.AddControl(x)
			b.AuxInt = -128
			b.Aux = s390x.GreaterOrEqual
			return true
		}
		// match: (BRC {s390x.Greater} (CMPWconst x [-129]) yes no)
		// result: (CIJ {s390x.GreaterOrEqual} x [-128] yes no)
		for b.Controls[0].Op == OpS390XCMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != -129 {
				break
			}
			x := v_0.Args[0]
			if b.Aux != s390x.Greater {
				break
			}
			b.Reset(BlockS390XCIJ)
			b.AddControl(x)
			b.AuxInt = -128
			b.Aux = s390x.GreaterOrEqual
			return true
		}
		// match: (BRC {s390x.GreaterOrEqual} (CMPconst x [ 128]) yes no)
		// result: (CGIJ {s390x.Greater} x [ 127] yes no)
		for b.Controls[0].Op == OpS390XCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 128 {
				break
			}
			x := v_0.Args[0]
			if b.Aux != s390x.GreaterOrEqual {
				break
			}
			b.Reset(BlockS390XCGIJ)
			b.AddControl(x)
			b.AuxInt = 127
			b.Aux = s390x.Greater
			return true
		}
		// match: (BRC {s390x.GreaterOrEqual} (CMPWconst x [ 128]) yes no)
		// result: (CIJ {s390x.Greater} x [ 127] yes no)
		for b.Controls[0].Op == OpS390XCMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 128 {
				break
			}
			x := v_0.Args[0]
			if b.Aux != s390x.GreaterOrEqual {
				break
			}
			b.Reset(BlockS390XCIJ)
			b.AddControl(x)
			b.AuxInt = 127
			b.Aux = s390x.Greater
			return true
		}
		// match: (BRC {s390x.Less} (CMPWUconst x [256]) yes no)
		// result: (CLIJ {s390x.LessOrEqual} x [-1] yes no)
		for b.Controls[0].Op == OpS390XCMPWUconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 256 {
				break
			}
			x := v_0.Args[0]
			if b.Aux != s390x.Less {
				break
			}
			b.Reset(BlockS390XCLIJ)
			b.AddControl(x)
			b.AuxInt = -1
			b.Aux = s390x.LessOrEqual
			return true
		}
		// match: (BRC {s390x.Less} (CMPUconst x [256]) yes no)
		// result: (CLGIJ {s390x.LessOrEqual} x [-1] yes no)
		for b.Controls[0].Op == OpS390XCMPUconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 256 {
				break
			}
			x := v_0.Args[0]
			if b.Aux != s390x.Less {
				break
			}
			b.Reset(BlockS390XCLGIJ)
			b.AddControl(x)
			b.AuxInt = -1
			b.Aux = s390x.LessOrEqual
			return true
		}
		// match: (BRC {s390x.GreaterOrEqual} (CMPWUconst x [256]) yes no)
		// result: (CLIJ {s390x.Greater} x [-1] yes no)
		for b.Controls[0].Op == OpS390XCMPWUconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 256 {
				break
			}
			x := v_0.Args[0]
			if b.Aux != s390x.GreaterOrEqual {
				break
			}
			b.Reset(BlockS390XCLIJ)
			b.AddControl(x)
			b.AuxInt = -1
			b.Aux = s390x.Greater
			return true
		}
		// match: (BRC {s390x.GreaterOrEqual} (CMPUconst x [256]) yes no)
		// result: (CLGIJ {s390x.Greater} x [-1] yes no)
		for b.Controls[0].Op == OpS390XCMPUconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 256 {
				break
			}
			x := v_0.Args[0]
			if b.Aux != s390x.GreaterOrEqual {
				break
			}
			b.Reset(BlockS390XCLGIJ)
			b.AddControl(x)
			b.AuxInt = -1
			b.Aux = s390x.Greater
			return true
		}
		// match: (BRC {c} (InvertFlags cmp) yes no)
		// result: (BRC {c.(s390x.CCMask).ReverseComparison()} cmp yes no)
		for b.Controls[0].Op == OpS390XInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			c := b.Aux
			b.Reset(BlockS390XBRC)
			b.AddControl(cmp)
			b.Aux = c.(s390x.CCMask).ReverseComparison()
			return true
		}
		// match: (BRC {c} (FlagEQ) yes no)
		// cond: c.(s390x.CCMask) & s390x.Equal != 0
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XFlagEQ {
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Equal != 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (BRC {c} (FlagLT) yes no)
		// cond: c.(s390x.CCMask) & s390x.Less != 0
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XFlagLT {
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Less != 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (BRC {c} (FlagGT) yes no)
		// cond: c.(s390x.CCMask) & s390x.Greater != 0
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XFlagGT {
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Greater != 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (BRC {c} (FlagOV) yes no)
		// cond: c.(s390x.CCMask) & s390x.Unordered != 0
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XFlagOV {
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Unordered != 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (BRC {c} (FlagEQ) yes no)
		// cond: c.(s390x.CCMask) & s390x.Equal == 0
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XFlagEQ {
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Equal == 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (BRC {c} (FlagLT) yes no)
		// cond: c.(s390x.CCMask) & s390x.Less == 0
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XFlagLT {
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Less == 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (BRC {c} (FlagGT) yes no)
		// cond: c.(s390x.CCMask) & s390x.Greater == 0
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XFlagGT {
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Greater == 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (BRC {c} (FlagOV) yes no)
		// cond: c.(s390x.CCMask) & s390x.Unordered == 0
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XFlagOV {
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Unordered == 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockS390XCGIJ:
		// match: (CGIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Equal != 0 && int64(x) == int64( int8(y))
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Equal != 0 && int64(x) == int64(int8(y))) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CGIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Less != 0 && int64(x) < int64( int8(y))
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Less != 0 && int64(x) < int64(int8(y))) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CGIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Greater != 0 && int64(x) > int64( int8(y))
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Greater != 0 && int64(x) > int64(int8(y))) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CGIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Equal == 0 && int64(x) == int64( int8(y))
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Equal == 0 && int64(x) == int64(int8(y))) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (CGIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Less == 0 && int64(x) < int64( int8(y))
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Less == 0 && int64(x) < int64(int8(y))) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (CGIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Greater == 0 && int64(x) > int64( int8(y))
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Greater == 0 && int64(x) > int64(int8(y))) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockS390XCGRJ:
		// match: (CGRJ {c} x (MOVDconst [y]) yes no)
		// cond: is8Bit(y)
		// result: (CGIJ {c} x [int64(int8(y))] yes no)
		for b.Controls[1].Op == OpS390XMOVDconst {
			x := b.Controls[0]
			v_1 := b.Controls[1]
			y := v_1.AuxInt
			c := b.Aux
			if !(is8Bit(y)) {
				break
			}
			b.Reset(BlockS390XCGIJ)
			b.AddControl(x)
			b.AuxInt = int64(int8(y))
			b.Aux = c
			return true
		}
		// match: (CGRJ {c} (MOVDconst [x]) y yes no)
		// cond: is8Bit(x)
		// result: (CGIJ {c.(s390x.CCMask).ReverseComparison()} y [int64(int8(x))] yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.Controls[1]
			c := b.Aux
			if !(is8Bit(x)) {
				break
			}
			b.Reset(BlockS390XCGIJ)
			b.AddControl(y)
			b.AuxInt = int64(int8(x))
			b.Aux = c.(s390x.CCMask).ReverseComparison()
			return true
		}
		// match: (CGRJ {c} x (MOVDconst [y]) yes no)
		// cond: !is8Bit(y) && is32Bit(y)
		// result: (BRC {c} (CMPconst x [int64(int32(y))]) yes no)
		for b.Controls[1].Op == OpS390XMOVDconst {
			x := b.Controls[0]
			v_1 := b.Controls[1]
			y := v_1.AuxInt
			c := b.Aux
			if !(!is8Bit(y) && is32Bit(y)) {
				break
			}
			b.Reset(BlockS390XBRC)
			v0 := b.NewValue0(x.Pos, OpS390XCMPconst, types.TypeFlags)
			v0.AuxInt = int64(int32(y))
			v0.AddArg(x)
			b.AddControl(v0)
			b.Aux = c
			return true
		}
		// match: (CGRJ {c} (MOVDconst [x]) y yes no)
		// cond: !is8Bit(x) && is32Bit(x)
		// result: (BRC {c.(s390x.CCMask).ReverseComparison()} (CMPconst y [int64(int32(x))]) yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.Controls[1]
			c := b.Aux
			if !(!is8Bit(x) && is32Bit(x)) {
				break
			}
			b.Reset(BlockS390XBRC)
			v0 := b.NewValue0(v_0.Pos, OpS390XCMPconst, types.TypeFlags)
			v0.AuxInt = int64(int32(x))
			v0.AddArg(y)
			b.AddControl(v0)
			b.Aux = c.(s390x.CCMask).ReverseComparison()
			return true
		}
		// match: (CGRJ {c} x y yes no)
		// cond: x == y && c.(s390x.CCMask)&s390x.Equal != 0
		// result: (First yes no)
		for {
			x := b.Controls[0]
			y := b.Controls[1]
			c := b.Aux
			if !(x == y && c.(s390x.CCMask)&s390x.Equal != 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CGRJ {c} x y yes no)
		// cond: x == y && c.(s390x.CCMask)&s390x.Equal == 0
		// result: (First no yes)
		for {
			x := b.Controls[0]
			y := b.Controls[1]
			c := b.Aux
			if !(x == y && c.(s390x.CCMask)&s390x.Equal == 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockS390XCIJ:
		// match: (CIJ {c} (MOVWreg x) [y] yes no)
		// result: (CIJ {c} x [y] yes no)
		for b.Controls[0].Op == OpS390XMOVWreg {
			v_0 := b.Controls[0]
			x := v_0.Args[0]
			y := b.AuxInt
			c := b.Aux
			b.Reset(BlockS390XCIJ)
			b.AddControl(x)
			b.AuxInt = y
			b.Aux = c
			return true
		}
		// match: (CIJ {c} (MOVWZreg x) [y] yes no)
		// result: (CIJ {c} x [y] yes no)
		for b.Controls[0].Op == OpS390XMOVWZreg {
			v_0 := b.Controls[0]
			x := v_0.Args[0]
			y := b.AuxInt
			c := b.Aux
			b.Reset(BlockS390XCIJ)
			b.AddControl(x)
			b.AuxInt = y
			b.Aux = c
			return true
		}
		// match: (CIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Equal != 0 && int32(x) == int32( int8(y))
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Equal != 0 && int32(x) == int32(int8(y))) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Less != 0 && int32(x) < int32( int8(y))
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Less != 0 && int32(x) < int32(int8(y))) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Greater != 0 && int32(x) > int32( int8(y))
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Greater != 0 && int32(x) > int32(int8(y))) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Equal == 0 && int32(x) == int32( int8(y))
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Equal == 0 && int32(x) == int32(int8(y))) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (CIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Less == 0 && int32(x) < int32( int8(y))
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Less == 0 && int32(x) < int32(int8(y))) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (CIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Greater == 0 && int32(x) > int32( int8(y))
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Greater == 0 && int32(x) > int32(int8(y))) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockS390XCLGIJ:
		// match: (CLGIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Equal != 0 && uint64(x) == uint64(uint8(y))
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Equal != 0 && uint64(x) == uint64(uint8(y))) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CLGIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Less != 0 && uint64(x) < uint64(uint8(y))
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Less != 0 && uint64(x) < uint64(uint8(y))) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CLGIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Greater != 0 && uint64(x) > uint64(uint8(y))
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Greater != 0 && uint64(x) > uint64(uint8(y))) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CLGIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Equal == 0 && uint64(x) == uint64(uint8(y))
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Equal == 0 && uint64(x) == uint64(uint8(y))) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (CLGIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Less == 0 && uint64(x) < uint64(uint8(y))
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Less == 0 && uint64(x) < uint64(uint8(y))) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (CLGIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Greater == 0 && uint64(x) > uint64(uint8(y))
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Greater == 0 && uint64(x) > uint64(uint8(y))) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (CLGIJ {s390x.GreaterOrEqual} _ [0] yes no)
		// result: (First yes no)
		for {
			if b.AuxInt != 0 || b.Aux != s390x.GreaterOrEqual {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CLGIJ {s390x.Less} _ [0] yes no)
		// result: (First no yes)
		for {
			if b.AuxInt != 0 || b.Aux != s390x.Less {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockS390XCLGRJ:
		// match: (CLGRJ {c} x (MOVDconst [y]) yes no)
		// cond: isU8Bit(y)
		// result: (CLGIJ {c} x [int64(int8(y))] yes no)
		for b.Controls[1].Op == OpS390XMOVDconst {
			x := b.Controls[0]
			v_1 := b.Controls[1]
			y := v_1.AuxInt
			c := b.Aux
			if !(isU8Bit(y)) {
				break
			}
			b.Reset(BlockS390XCLGIJ)
			b.AddControl(x)
			b.AuxInt = int64(int8(y))
			b.Aux = c
			return true
		}
		// match: (CLGRJ {c} (MOVDconst [x]) y yes no)
		// cond: isU8Bit(x)
		// result: (CLGIJ {c.(s390x.CCMask).ReverseComparison()} y [int64(int8(x))] yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.Controls[1]
			c := b.Aux
			if !(isU8Bit(x)) {
				break
			}
			b.Reset(BlockS390XCLGIJ)
			b.AddControl(y)
			b.AuxInt = int64(int8(x))
			b.Aux = c.(s390x.CCMask).ReverseComparison()
			return true
		}
		// match: (CLGRJ {c} x (MOVDconst [y]) yes no)
		// cond: !isU8Bit(y) && isU32Bit(y)
		// result: (BRC {c} (CMPUconst x [int64(int32(y))]) yes no)
		for b.Controls[1].Op == OpS390XMOVDconst {
			x := b.Controls[0]
			v_1 := b.Controls[1]
			y := v_1.AuxInt
			c := b.Aux
			if !(!isU8Bit(y) && isU32Bit(y)) {
				break
			}
			b.Reset(BlockS390XBRC)
			v0 := b.NewValue0(x.Pos, OpS390XCMPUconst, types.TypeFlags)
			v0.AuxInt = int64(int32(y))
			v0.AddArg(x)
			b.AddControl(v0)
			b.Aux = c
			return true
		}
		// match: (CLGRJ {c} (MOVDconst [x]) y yes no)
		// cond: !isU8Bit(x) && isU32Bit(x)
		// result: (BRC {c.(s390x.CCMask).ReverseComparison()} (CMPUconst y [int64(int32(x))]) yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.Controls[1]
			c := b.Aux
			if !(!isU8Bit(x) && isU32Bit(x)) {
				break
			}
			b.Reset(BlockS390XBRC)
			v0 := b.NewValue0(v_0.Pos, OpS390XCMPUconst, types.TypeFlags)
			v0.AuxInt = int64(int32(x))
			v0.AddArg(y)
			b.AddControl(v0)
			b.Aux = c.(s390x.CCMask).ReverseComparison()
			return true
		}
		// match: (CLGRJ {c} x y yes no)
		// cond: x == y && c.(s390x.CCMask)&s390x.Equal != 0
		// result: (First yes no)
		for {
			x := b.Controls[0]
			y := b.Controls[1]
			c := b.Aux
			if !(x == y && c.(s390x.CCMask)&s390x.Equal != 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CLGRJ {c} x y yes no)
		// cond: x == y && c.(s390x.CCMask)&s390x.Equal == 0
		// result: (First no yes)
		for {
			x := b.Controls[0]
			y := b.Controls[1]
			c := b.Aux
			if !(x == y && c.(s390x.CCMask)&s390x.Equal == 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockS390XCLIJ:
		// match: (CLIJ {s390x.LessOrGreater} (LOCGR {d} (MOVDconst [0]) (MOVDconst [x]) cmp) [0] yes no)
		// cond: int32(x) != 0
		// result: (BRC {d} cmp yes no)
		for b.Controls[0].Op == OpS390XLOCGR {
			v_0 := b.Controls[0]
			d := v_0.Aux
			cmp := v_0.Args[2]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpS390XMOVDconst || v_0_0.AuxInt != 0 {
				break
			}
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpS390XMOVDconst {
				break
			}
			x := v_0_1.AuxInt
			if b.AuxInt != 0 || b.Aux != s390x.LessOrGreater || !(int32(x) != 0) {
				break
			}
			b.Reset(BlockS390XBRC)
			b.AddControl(cmp)
			b.Aux = d
			return true
		}
		// match: (CLIJ {c} (MOVWreg x) [y] yes no)
		// result: (CLIJ {c} x [y] yes no)
		for b.Controls[0].Op == OpS390XMOVWreg {
			v_0 := b.Controls[0]
			x := v_0.Args[0]
			y := b.AuxInt
			c := b.Aux
			b.Reset(BlockS390XCLIJ)
			b.AddControl(x)
			b.AuxInt = y
			b.Aux = c
			return true
		}
		// match: (CLIJ {c} (MOVWZreg x) [y] yes no)
		// result: (CLIJ {c} x [y] yes no)
		for b.Controls[0].Op == OpS390XMOVWZreg {
			v_0 := b.Controls[0]
			x := v_0.Args[0]
			y := b.AuxInt
			c := b.Aux
			b.Reset(BlockS390XCLIJ)
			b.AddControl(x)
			b.AuxInt = y
			b.Aux = c
			return true
		}
		// match: (CLIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Equal != 0 && uint32(x) == uint32(uint8(y))
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Equal != 0 && uint32(x) == uint32(uint8(y))) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CLIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Less != 0 && uint32(x) < uint32(uint8(y))
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Less != 0 && uint32(x) < uint32(uint8(y))) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CLIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Greater != 0 && uint32(x) > uint32(uint8(y))
		// result: (First yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Greater != 0 && uint32(x) > uint32(uint8(y))) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CLIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Equal == 0 && uint32(x) == uint32(uint8(y))
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Equal == 0 && uint32(x) == uint32(uint8(y))) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (CLIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Less == 0 && uint32(x) < uint32(uint8(y))
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Less == 0 && uint32(x) < uint32(uint8(y))) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (CLIJ {c} (MOVDconst [x]) [y] yes no)
		// cond: c.(s390x.CCMask)&s390x.Greater == 0 && uint32(x) > uint32(uint8(y))
		// result: (First no yes)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.AuxInt
			c := b.Aux
			if !(c.(s390x.CCMask)&s390x.Greater == 0 && uint32(x) > uint32(uint8(y))) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (CLIJ {s390x.GreaterOrEqual} _ [0] yes no)
		// result: (First yes no)
		for {
			if b.AuxInt != 0 || b.Aux != s390x.GreaterOrEqual {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CLIJ {s390x.Less} _ [0] yes no)
		// result: (First no yes)
		for {
			if b.AuxInt != 0 || b.Aux != s390x.Less {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockS390XCLRJ:
		// match: (CLRJ {c} x (MOVDconst [y]) yes no)
		// cond: isU8Bit(y)
		// result: (CLIJ {c} x [int64(int8(y))] yes no)
		for b.Controls[1].Op == OpS390XMOVDconst {
			x := b.Controls[0]
			v_1 := b.Controls[1]
			y := v_1.AuxInt
			c := b.Aux
			if !(isU8Bit(y)) {
				break
			}
			b.Reset(BlockS390XCLIJ)
			b.AddControl(x)
			b.AuxInt = int64(int8(y))
			b.Aux = c
			return true
		}
		// match: (CLRJ {c} (MOVDconst [x]) y yes no)
		// cond: isU8Bit(x)
		// result: (CLIJ {c.(s390x.CCMask).ReverseComparison()} y [int64(int8(x))] yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.Controls[1]
			c := b.Aux
			if !(isU8Bit(x)) {
				break
			}
			b.Reset(BlockS390XCLIJ)
			b.AddControl(y)
			b.AuxInt = int64(int8(x))
			b.Aux = c.(s390x.CCMask).ReverseComparison()
			return true
		}
		// match: (CLRJ {c} x (MOVDconst [y]) yes no)
		// cond: !isU8Bit(y) && isU32Bit(y)
		// result: (BRC {c} (CMPWUconst x [int64(int32(y))]) yes no)
		for b.Controls[1].Op == OpS390XMOVDconst {
			x := b.Controls[0]
			v_1 := b.Controls[1]
			y := v_1.AuxInt
			c := b.Aux
			if !(!isU8Bit(y) && isU32Bit(y)) {
				break
			}
			b.Reset(BlockS390XBRC)
			v0 := b.NewValue0(x.Pos, OpS390XCMPWUconst, types.TypeFlags)
			v0.AuxInt = int64(int32(y))
			v0.AddArg(x)
			b.AddControl(v0)
			b.Aux = c
			return true
		}
		// match: (CLRJ {c} (MOVDconst [x]) y yes no)
		// cond: !isU8Bit(x) && isU32Bit(x)
		// result: (BRC {c.(s390x.CCMask).ReverseComparison()} (CMPWUconst y [int64(int32(x))]) yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.Controls[1]
			c := b.Aux
			if !(!isU8Bit(x) && isU32Bit(x)) {
				break
			}
			b.Reset(BlockS390XBRC)
			v0 := b.NewValue0(v_0.Pos, OpS390XCMPWUconst, types.TypeFlags)
			v0.AuxInt = int64(int32(x))
			v0.AddArg(y)
			b.AddControl(v0)
			b.Aux = c.(s390x.CCMask).ReverseComparison()
			return true
		}
		// match: (CLRJ {c} x y yes no)
		// cond: x == y && c.(s390x.CCMask)&s390x.Equal != 0
		// result: (First yes no)
		for {
			x := b.Controls[0]
			y := b.Controls[1]
			c := b.Aux
			if !(x == y && c.(s390x.CCMask)&s390x.Equal != 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CLRJ {c} x y yes no)
		// cond: x == y && c.(s390x.CCMask)&s390x.Equal == 0
		// result: (First no yes)
		for {
			x := b.Controls[0]
			y := b.Controls[1]
			c := b.Aux
			if !(x == y && c.(s390x.CCMask)&s390x.Equal == 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockS390XCRJ:
		// match: (CRJ {c} x (MOVDconst [y]) yes no)
		// cond: is8Bit(y)
		// result: (CIJ {c} x [int64(int8(y))] yes no)
		for b.Controls[1].Op == OpS390XMOVDconst {
			x := b.Controls[0]
			v_1 := b.Controls[1]
			y := v_1.AuxInt
			c := b.Aux
			if !(is8Bit(y)) {
				break
			}
			b.Reset(BlockS390XCIJ)
			b.AddControl(x)
			b.AuxInt = int64(int8(y))
			b.Aux = c
			return true
		}
		// match: (CRJ {c} (MOVDconst [x]) y yes no)
		// cond: is8Bit(x)
		// result: (CIJ {c.(s390x.CCMask).ReverseComparison()} y [int64(int8(x))] yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.Controls[1]
			c := b.Aux
			if !(is8Bit(x)) {
				break
			}
			b.Reset(BlockS390XCIJ)
			b.AddControl(y)
			b.AuxInt = int64(int8(x))
			b.Aux = c.(s390x.CCMask).ReverseComparison()
			return true
		}
		// match: (CRJ {c} x (MOVDconst [y]) yes no)
		// cond: !is8Bit(y) && is32Bit(y)
		// result: (BRC {c} (CMPWconst x [int64(int32(y))]) yes no)
		for b.Controls[1].Op == OpS390XMOVDconst {
			x := b.Controls[0]
			v_1 := b.Controls[1]
			y := v_1.AuxInt
			c := b.Aux
			if !(!is8Bit(y) && is32Bit(y)) {
				break
			}
			b.Reset(BlockS390XBRC)
			v0 := b.NewValue0(x.Pos, OpS390XCMPWconst, types.TypeFlags)
			v0.AuxInt = int64(int32(y))
			v0.AddArg(x)
			b.AddControl(v0)
			b.Aux = c
			return true
		}
		// match: (CRJ {c} (MOVDconst [x]) y yes no)
		// cond: !is8Bit(x) && is32Bit(x)
		// result: (BRC {c.(s390x.CCMask).ReverseComparison()} (CMPWconst y [int64(int32(x))]) yes no)
		for b.Controls[0].Op == OpS390XMOVDconst {
			v_0 := b.Controls[0]
			x := v_0.AuxInt
			y := b.Controls[1]
			c := b.Aux
			if !(!is8Bit(x) && is32Bit(x)) {
				break
			}
			b.Reset(BlockS390XBRC)
			v0 := b.NewValue0(v_0.Pos, OpS390XCMPWconst, types.TypeFlags)
			v0.AuxInt = int64(int32(x))
			v0.AddArg(y)
			b.AddControl(v0)
			b.Aux = c.(s390x.CCMask).ReverseComparison()
			return true
		}
		// match: (CRJ {c} x y yes no)
		// cond: x == y && c.(s390x.CCMask)&s390x.Equal != 0
		// result: (First yes no)
		for {
			x := b.Controls[0]
			y := b.Controls[1]
			c := b.Aux
			if !(x == y && c.(s390x.CCMask)&s390x.Equal != 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (CRJ {c} x y yes no)
		// cond: x == y && c.(s390x.CCMask)&s390x.Equal == 0
		// result: (First no yes)
		for {
			x := b.Controls[0]
			y := b.Controls[1]
			c := b.Aux
			if !(x == y && c.(s390x.CCMask)&s390x.Equal == 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockIf:
		// match: (If cond yes no)
		// result: (CLIJ {s390x.LessOrGreater} (MOVBZreg <typ.Bool> cond) [0] yes no)
		for {
			cond := b.Controls[0]
			b.Reset(BlockS390XCLIJ)
			v0 := b.NewValue0(cond.Pos, OpS390XMOVBZreg, typ.Bool)
			v0.AddArg(cond)
			b.AddControl(v0)
			b.AuxInt = 0
			b.Aux = s390x.LessOrGreater
			return true
		}
	}
	return false
}
