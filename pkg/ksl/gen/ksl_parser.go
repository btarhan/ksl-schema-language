// Code generated from /home/kooshy/Projects/ksl-schema-language/pkg/ksl/ksl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ksl

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type kslParser struct {
	*antlr.BaseParser
}

var KslParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kslParserInit() {
	staticData := &KslParserStaticData
	staticData.LiteralNames = []string{
		"", "'version'", "", "'.'", "'module'", "", "'public'", "'internal'",
		"'private'", "'type'", "'relation'", "'import'", "'extension'", "",
		"'AtMostOne'", "'ExactlyOne'", "'AtLeastOne'", "'Any'", "'as'", "'and'",
		"'or'", "'unless'", "'allow_duplicates'", "':'", "'{'", "'}'", "'@'",
		"'('", "')'", "'['", "']'", "'$'", "'`'", "'''", "','",
	}
	staticData.SymbolicNames = []string{
		"", "VERSION", "VERSIONNUM", "RESOLVE", "MODULE", "ACCESS", "PUBLIC",
		"INTERNAL", "PRIVATE", "TYPE", "RELATION", "IMPORT", "EXTENSION", "CARDINALITY",
		"ATMOSTONE", "EXACTLYONE", "ATLEASTONE", "ANY", "AS", "AND", "OR", "UNLESS",
		"ALLOW_DUPLICATES", "EXPAND", "LBRACE", "RBRACE", "EXTENSION_CALL",
		"LPAREN", "RPAREN", "LSQUARE", "RSQARE", "VARREF", "TEMPLATE_DELIM",
		"STRING_DELIM", "ARG_DELIM", "NAME", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"file", "version", "module", "import_stmt", "declaration", "typeExpr",
		"extensionParam", "extensionParams", "extensionReference", "relation",
		"relationBody", "paramNames", "extension", "dynamicType", "dynamicRelation",
		"dynamicName", "dynamicBody",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 37, 235, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 5, 0, 38, 8, 0, 10, 0, 12, 0, 41, 9, 0,
		1, 0, 4, 0, 44, 8, 0, 11, 0, 12, 0, 45, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 59, 8, 4, 1, 5, 3, 5, 62, 8, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 68, 8, 5, 10, 5, 12, 5, 71, 9, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 84, 8, 7,
		10, 7, 12, 7, 87, 9, 7, 1, 8, 1, 8, 1, 8, 3, 8, 92, 8, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 5, 9, 100, 8, 9, 10, 9, 12, 9, 103, 9, 9, 1, 9, 3,
		9, 106, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10,
		116, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 3, 10, 128, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 5, 10, 139, 8, 10, 10, 10, 12, 10, 142, 9, 10, 1,
		11, 1, 11, 1, 11, 3, 11, 147, 8, 11, 1, 12, 3, 12, 150, 8, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 4, 12, 159, 8, 12, 11, 12, 12, 12,
		160, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 169, 8, 13, 10, 13,
		12, 13, 172, 9, 13, 1, 13, 1, 13, 1, 14, 3, 14, 177, 8, 14, 1, 14, 3, 14,
		180, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 4, 15, 194, 8, 15, 11, 15, 12, 15, 195, 1, 15,
		1, 15, 3, 15, 200, 8, 15, 1, 16, 1, 16, 1, 16, 3, 16, 205, 8, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 3, 16, 219, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 5, 16, 230, 8, 16, 10, 16, 12, 16, 233, 9, 16, 1, 16, 0,
		2, 20, 32, 17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 0, 1, 1, 0, 33, 33, 249, 0, 34, 1, 0, 0, 0, 2, 47, 1, 0, 0, 0, 4, 50,
		1, 0, 0, 0, 6, 53, 1, 0, 0, 0, 8, 58, 1, 0, 0, 0, 10, 61, 1, 0, 0, 0, 12,
		74, 1, 0, 0, 0, 14, 80, 1, 0, 0, 0, 16, 88, 1, 0, 0, 0, 18, 101, 1, 0,
		0, 0, 20, 127, 1, 0, 0, 0, 22, 143, 1, 0, 0, 0, 24, 149, 1, 0, 0, 0, 26,
		164, 1, 0, 0, 0, 28, 176, 1, 0, 0, 0, 30, 199, 1, 0, 0, 0, 32, 218, 1,
		0, 0, 0, 34, 35, 3, 2, 1, 0, 35, 39, 3, 4, 2, 0, 36, 38, 3, 6, 3, 0, 37,
		36, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0,
		0, 40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 44, 3, 8, 4, 0, 43, 42,
		1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0,
		46, 1, 1, 0, 0, 0, 47, 48, 5, 1, 0, 0, 48, 49, 5, 2, 0, 0, 49, 3, 1, 0,
		0, 0, 50, 51, 5, 4, 0, 0, 51, 52, 5, 35, 0, 0, 52, 5, 1, 0, 0, 0, 53, 54,
		5, 11, 0, 0, 54, 55, 5, 35, 0, 0, 55, 7, 1, 0, 0, 0, 56, 59, 3, 10, 5,
		0, 57, 59, 3, 24, 12, 0, 58, 56, 1, 0, 0, 0, 58, 57, 1, 0, 0, 0, 59, 9,
		1, 0, 0, 0, 60, 62, 5, 5, 0, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0,
		62, 63, 1, 0, 0, 0, 63, 64, 5, 9, 0, 0, 64, 65, 5, 35, 0, 0, 65, 69, 5,
		24, 0, 0, 66, 68, 3, 18, 9, 0, 67, 66, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0,
		69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 72, 1, 0, 0, 0, 71, 69, 1,
		0, 0, 0, 72, 73, 5, 25, 0, 0, 73, 11, 1, 0, 0, 0, 74, 75, 5, 35, 0, 0,
		75, 76, 5, 23, 0, 0, 76, 77, 5, 33, 0, 0, 77, 78, 8, 0, 0, 0, 78, 79, 5,
		33, 0, 0, 79, 13, 1, 0, 0, 0, 80, 85, 3, 12, 6, 0, 81, 82, 5, 34, 0, 0,
		82, 84, 3, 12, 6, 0, 83, 81, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1,
		0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 15, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88,
		91, 5, 26, 0, 0, 89, 90, 5, 35, 0, 0, 90, 92, 5, 3, 0, 0, 91, 89, 1, 0,
		0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 5, 35, 0, 0, 94,
		95, 5, 27, 0, 0, 95, 96, 3, 14, 7, 0, 96, 97, 5, 28, 0, 0, 97, 17, 1, 0,
		0, 0, 98, 100, 3, 16, 8, 0, 99, 98, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0, 101,
		99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1,
		0, 0, 0, 104, 106, 5, 5, 0, 0, 105, 104, 1, 0, 0, 0, 105, 106, 1, 0, 0,
		0, 106, 107, 1, 0, 0, 0, 107, 108, 5, 10, 0, 0, 108, 109, 5, 35, 0, 0,
		109, 110, 5, 23, 0, 0, 110, 111, 3, 20, 10, 0, 111, 19, 1, 0, 0, 0, 112,
		113, 6, 10, -1, 0, 113, 115, 5, 29, 0, 0, 114, 116, 5, 13, 0, 0, 115, 114,
		1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 5, 35,
		0, 0, 118, 128, 5, 30, 0, 0, 119, 128, 5, 35, 0, 0, 120, 121, 5, 35, 0,
		0, 121, 122, 5, 3, 0, 0, 122, 128, 5, 35, 0, 0, 123, 124, 5, 27, 0, 0,
		124, 125, 3, 20, 10, 0, 125, 126, 5, 28, 0, 0, 126, 128, 1, 0, 0, 0, 127,
		112, 1, 0, 0, 0, 127, 119, 1, 0, 0, 0, 127, 120, 1, 0, 0, 0, 127, 123,
		1, 0, 0, 0, 128, 140, 1, 0, 0, 0, 129, 130, 10, 3, 0, 0, 130, 131, 5, 19,
		0, 0, 131, 139, 3, 20, 10, 4, 132, 133, 10, 2, 0, 0, 133, 134, 5, 20, 0,
		0, 134, 139, 3, 20, 10, 3, 135, 136, 10, 1, 0, 0, 136, 137, 5, 21, 0, 0,
		137, 139, 3, 20, 10, 2, 138, 129, 1, 0, 0, 0, 138, 132, 1, 0, 0, 0, 138,
		135, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141,
		1, 0, 0, 0, 141, 21, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 146, 5, 35,
		0, 0, 144, 145, 5, 34, 0, 0, 145, 147, 5, 35, 0, 0, 146, 144, 1, 0, 0,
		0, 146, 147, 1, 0, 0, 0, 147, 23, 1, 0, 0, 0, 148, 150, 5, 5, 0, 0, 149,
		148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152,
		5, 12, 0, 0, 152, 153, 5, 35, 0, 0, 153, 154, 5, 27, 0, 0, 154, 155, 3,
		22, 11, 0, 155, 156, 5, 28, 0, 0, 156, 158, 5, 24, 0, 0, 157, 159, 3, 26,
		13, 0, 158, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0,
		160, 161, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 5, 25, 0, 0, 163,
		25, 1, 0, 0, 0, 164, 165, 5, 9, 0, 0, 165, 166, 3, 30, 15, 0, 166, 170,
		5, 24, 0, 0, 167, 169, 3, 28, 14, 0, 168, 167, 1, 0, 0, 0, 169, 172, 1,
		0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 173, 1, 0, 0,
		0, 172, 170, 1, 0, 0, 0, 173, 174, 5, 25, 0, 0, 174, 27, 1, 0, 0, 0, 175,
		177, 5, 22, 0, 0, 176, 175, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 179,
		1, 0, 0, 0, 178, 180, 5, 5, 0, 0, 179, 178, 1, 0, 0, 0, 179, 180, 1, 0,
		0, 0, 180, 181, 1, 0, 0, 0, 181, 182, 5, 10, 0, 0, 182, 183, 3, 30, 15,
		0, 183, 184, 5, 23, 0, 0, 184, 185, 3, 32, 16, 0, 185, 29, 1, 0, 0, 0,
		186, 200, 5, 35, 0, 0, 187, 188, 5, 31, 0, 0, 188, 189, 5, 24, 0, 0, 189,
		190, 5, 35, 0, 0, 190, 200, 5, 25, 0, 0, 191, 193, 5, 32, 0, 0, 192, 194,
		3, 30, 15, 0, 193, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 193, 1,
		0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 198, 5, 32, 0,
		0, 198, 200, 1, 0, 0, 0, 199, 186, 1, 0, 0, 0, 199, 187, 1, 0, 0, 0, 199,
		191, 1, 0, 0, 0, 200, 31, 1, 0, 0, 0, 201, 202, 6, 16, -1, 0, 202, 204,
		5, 29, 0, 0, 203, 205, 5, 13, 0, 0, 204, 203, 1, 0, 0, 0, 204, 205, 1,
		0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207, 3, 30, 15, 0, 207, 208, 5, 30,
		0, 0, 208, 219, 1, 0, 0, 0, 209, 219, 3, 30, 15, 0, 210, 211, 3, 30, 15,
		0, 211, 212, 5, 3, 0, 0, 212, 213, 3, 30, 15, 0, 213, 219, 1, 0, 0, 0,
		214, 215, 5, 27, 0, 0, 215, 216, 3, 32, 16, 0, 216, 217, 5, 28, 0, 0, 217,
		219, 1, 0, 0, 0, 218, 201, 1, 0, 0, 0, 218, 209, 1, 0, 0, 0, 218, 210,
		1, 0, 0, 0, 218, 214, 1, 0, 0, 0, 219, 231, 1, 0, 0, 0, 220, 221, 10, 3,
		0, 0, 221, 222, 5, 19, 0, 0, 222, 230, 3, 32, 16, 4, 223, 224, 10, 2, 0,
		0, 224, 225, 5, 20, 0, 0, 225, 230, 3, 32, 16, 3, 226, 227, 10, 1, 0, 0,
		227, 228, 5, 21, 0, 0, 228, 230, 3, 32, 16, 2, 229, 220, 1, 0, 0, 0, 229,
		223, 1, 0, 0, 0, 229, 226, 1, 0, 0, 0, 230, 233, 1, 0, 0, 0, 231, 229,
		1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 33, 1, 0, 0, 0, 233, 231, 1, 0,
		0, 0, 25, 39, 45, 58, 61, 69, 85, 91, 101, 105, 115, 127, 138, 140, 146,
		149, 160, 170, 176, 179, 195, 199, 204, 218, 229, 231,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// kslParserInit initializes any static state used to implement kslParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewkslParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func KslParserInit() {
	staticData := &KslParserStaticData
	staticData.once.Do(kslParserInit)
}

// NewkslParser produces a new parser instance for the optional input antlr.TokenStream.
func NewkslParser(input antlr.TokenStream) *kslParser {
	KslParserInit()
	this := new(kslParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &KslParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ksl.g4"

	return this
}

// kslParser tokens.
const (
	kslParserEOF              = antlr.TokenEOF
	kslParserVERSION          = 1
	kslParserVERSIONNUM       = 2
	kslParserRESOLVE          = 3
	kslParserMODULE           = 4
	kslParserACCESS           = 5
	kslParserPUBLIC           = 6
	kslParserINTERNAL         = 7
	kslParserPRIVATE          = 8
	kslParserTYPE             = 9
	kslParserRELATION         = 10
	kslParserIMPORT           = 11
	kslParserEXTENSION        = 12
	kslParserCARDINALITY      = 13
	kslParserATMOSTONE        = 14
	kslParserEXACTLYONE       = 15
	kslParserATLEASTONE       = 16
	kslParserANY              = 17
	kslParserAS               = 18
	kslParserAND              = 19
	kslParserOR               = 20
	kslParserUNLESS           = 21
	kslParserALLOW_DUPLICATES = 22
	kslParserEXPAND           = 23
	kslParserLBRACE           = 24
	kslParserRBRACE           = 25
	kslParserEXTENSION_CALL   = 26
	kslParserLPAREN           = 27
	kslParserRPAREN           = 28
	kslParserLSQUARE          = 29
	kslParserRSQARE           = 30
	kslParserVARREF           = 31
	kslParserTEMPLATE_DELIM   = 32
	kslParserSTRING_DELIM     = 33
	kslParserARG_DELIM        = 34
	kslParserNAME             = 35
	kslParserCOMMENT          = 36
	kslParserWS               = 37
)

// kslParser rules.
const (
	kslParserRULE_file               = 0
	kslParserRULE_version            = 1
	kslParserRULE_module             = 2
	kslParserRULE_import_stmt        = 3
	kslParserRULE_declaration        = 4
	kslParserRULE_typeExpr           = 5
	kslParserRULE_extensionParam     = 6
	kslParserRULE_extensionParams    = 7
	kslParserRULE_extensionReference = 8
	kslParserRULE_relation           = 9
	kslParserRULE_relationBody       = 10
	kslParserRULE_paramNames         = 11
	kslParserRULE_extension          = 12
	kslParserRULE_dynamicType        = 13
	kslParserRULE_dynamicRelation    = 14
	kslParserRULE_dynamicName        = 15
	kslParserRULE_dynamicBody        = 16
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Version() IVersionContext
	Module() IModuleContext
	AllImport_stmt() []IImport_stmtContext
	Import_stmt(i int) IImport_stmtContext
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_file
	return p
}

func InitEmptyFileContext(p *FileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_file
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) Version() IVersionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVersionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVersionContext)
}

func (s *FileContext) Module() IModuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModuleContext)
}

func (s *FileContext) AllImport_stmt() []IImport_stmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImport_stmtContext); ok {
			len++
		}
	}

	tst := make([]IImport_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImport_stmtContext); ok {
			tst[i] = t.(IImport_stmtContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) Import_stmt(i int) IImport_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImport_stmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImport_stmtContext)
}

func (s *FileContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitFile(s)
	}
}

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, kslParserRULE_file)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Version()
	}
	{
		p.SetState(35)
		p.Module()
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserIMPORT {
		{
			p.SetState(36)
			p.Import_stmt()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4640) != 0) {
		{
			p.SetState(42)
			p.Declaration()
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVersionContext is an interface to support dynamic dispatch.
type IVersionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VERSION() antlr.TerminalNode
	VERSIONNUM() antlr.TerminalNode

	// IsVersionContext differentiates from other interfaces.
	IsVersionContext()
}

type VersionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionContext() *VersionContext {
	var p = new(VersionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_version
	return p
}

func InitEmptyVersionContext(p *VersionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_version
}

func (*VersionContext) IsVersionContext() {}

func NewVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionContext {
	var p = new(VersionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_version

	return p
}

func (s *VersionContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionContext) VERSION() antlr.TerminalNode {
	return s.GetToken(kslParserVERSION, 0)
}

func (s *VersionContext) VERSIONNUM() antlr.TerminalNode {
	return s.GetToken(kslParserVERSIONNUM, 0)
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterVersion(s)
	}
}

func (s *VersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitVersion(s)
	}
}

func (s *VersionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitVersion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) Version() (localctx IVersionContext) {
	localctx = NewVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, kslParserRULE_version)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(kslParserVERSION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(kslParserVERSIONNUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModuleContext is an interface to support dynamic dispatch.
type IModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MODULE() antlr.TerminalNode
	NAME() antlr.TerminalNode

	// IsModuleContext differentiates from other interfaces.
	IsModuleContext()
}

type ModuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_module
	return p
}

func InitEmptyModuleContext(p *ModuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_module
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) MODULE() antlr.TerminalNode {
	return s.GetToken(kslParserMODULE, 0)
}

func (s *ModuleContext) NAME() antlr.TerminalNode {
	return s.GetToken(kslParserNAME, 0)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitModule(s)
	}
}

func (s *ModuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitModule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, kslParserRULE_module)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(kslParserMODULE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(kslParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImport_stmtContext is an interface to support dynamic dispatch.
type IImport_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	NAME() antlr.TerminalNode

	// IsImport_stmtContext differentiates from other interfaces.
	IsImport_stmtContext()
}

type Import_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_stmtContext() *Import_stmtContext {
	var p = new(Import_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_import_stmt
	return p
}

func InitEmptyImport_stmtContext(p *Import_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_import_stmt
}

func (*Import_stmtContext) IsImport_stmtContext() {}

func NewImport_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_stmtContext {
	var p = new(Import_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_import_stmt

	return p
}

func (s *Import_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Import_stmtContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(kslParserIMPORT, 0)
}

func (s *Import_stmtContext) NAME() antlr.TerminalNode {
	return s.GetToken(kslParserNAME, 0)
}

func (s *Import_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterImport_stmt(s)
	}
}

func (s *Import_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitImport_stmt(s)
	}
}

func (s *Import_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitImport_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) Import_stmt() (localctx IImport_stmtContext) {
	localctx = NewImport_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, kslParserRULE_import_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(kslParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(kslParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeExpr() ITypeExprContext
	Extension() IExtensionContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *DeclarationContext) Extension() IExtensionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtensionContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, kslParserRULE_declaration)
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.TypeExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Extension()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeExprContext is an interface to support dynamic dispatch.
type ITypeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	NAME() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	ACCESS() antlr.TerminalNode
	AllRelation() []IRelationContext
	Relation(i int) IRelationContext

	// IsTypeExprContext differentiates from other interfaces.
	IsTypeExprContext()
}

type TypeExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprContext() *TypeExprContext {
	var p = new(TypeExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_typeExpr
	return p
}

func InitEmptyTypeExprContext(p *TypeExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_typeExpr
}

func (*TypeExprContext) IsTypeExprContext() {}

func NewTypeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprContext {
	var p = new(TypeExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_typeExpr

	return p
}

func (s *TypeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExprContext) TYPE() antlr.TerminalNode {
	return s.GetToken(kslParserTYPE, 0)
}

func (s *TypeExprContext) NAME() antlr.TerminalNode {
	return s.GetToken(kslParserNAME, 0)
}

func (s *TypeExprContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserLBRACE, 0)
}

func (s *TypeExprContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserRBRACE, 0)
}

func (s *TypeExprContext) ACCESS() antlr.TerminalNode {
	return s.GetToken(kslParserACCESS, 0)
}

func (s *TypeExprContext) AllRelation() []IRelationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationContext); ok {
			len++
		}
	}

	tst := make([]IRelationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationContext); ok {
			tst[i] = t.(IRelationContext)
			i++
		}
	}

	return tst
}

func (s *TypeExprContext) Relation(i int) IRelationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *TypeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterTypeExpr(s)
	}
}

func (s *TypeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitTypeExpr(s)
	}
}

func (s *TypeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitTypeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) TypeExpr() (localctx ITypeExprContext) {
	localctx = NewTypeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, kslParserRULE_typeExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(60)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(63)
		p.Match(kslParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(kslParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Match(kslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67109920) != 0 {
		{
			p.SetState(66)
			p.Relation()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
		p.Match(kslParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExtensionParamContext is an interface to support dynamic dispatch.
type IExtensionParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	EXPAND() antlr.TerminalNode
	AllSTRING_DELIM() []antlr.TerminalNode
	STRING_DELIM(i int) antlr.TerminalNode

	// IsExtensionParamContext differentiates from other interfaces.
	IsExtensionParamContext()
}

type ExtensionParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionParamContext() *ExtensionParamContext {
	var p = new(ExtensionParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extensionParam
	return p
}

func InitEmptyExtensionParamContext(p *ExtensionParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extensionParam
}

func (*ExtensionParamContext) IsExtensionParamContext() {}

func NewExtensionParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionParamContext {
	var p = new(ExtensionParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_extensionParam

	return p
}

func (s *ExtensionParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionParamContext) NAME() antlr.TerminalNode {
	return s.GetToken(kslParserNAME, 0)
}

func (s *ExtensionParamContext) EXPAND() antlr.TerminalNode {
	return s.GetToken(kslParserEXPAND, 0)
}

func (s *ExtensionParamContext) AllSTRING_DELIM() []antlr.TerminalNode {
	return s.GetTokens(kslParserSTRING_DELIM)
}

func (s *ExtensionParamContext) STRING_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(kslParserSTRING_DELIM, i)
}

func (s *ExtensionParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensionParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterExtensionParam(s)
	}
}

func (s *ExtensionParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitExtensionParam(s)
	}
}

func (s *ExtensionParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitExtensionParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) ExtensionParam() (localctx IExtensionParamContext) {
	localctx = NewExtensionParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, kslParserRULE_extensionParam)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(kslParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(kslParserEXPAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(kslParserSTRING_DELIM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		_la = p.GetTokenStream().LA(1)

		if _la <= 0 || _la == kslParserSTRING_DELIM {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(78)
		p.Match(kslParserSTRING_DELIM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExtensionParamsContext is an interface to support dynamic dispatch.
type IExtensionParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExtensionParam() []IExtensionParamContext
	ExtensionParam(i int) IExtensionParamContext
	AllARG_DELIM() []antlr.TerminalNode
	ARG_DELIM(i int) antlr.TerminalNode

	// IsExtensionParamsContext differentiates from other interfaces.
	IsExtensionParamsContext()
}

type ExtensionParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionParamsContext() *ExtensionParamsContext {
	var p = new(ExtensionParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extensionParams
	return p
}

func InitEmptyExtensionParamsContext(p *ExtensionParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extensionParams
}

func (*ExtensionParamsContext) IsExtensionParamsContext() {}

func NewExtensionParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionParamsContext {
	var p = new(ExtensionParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_extensionParams

	return p
}

func (s *ExtensionParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionParamsContext) AllExtensionParam() []IExtensionParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtensionParamContext); ok {
			len++
		}
	}

	tst := make([]IExtensionParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtensionParamContext); ok {
			tst[i] = t.(IExtensionParamContext)
			i++
		}
	}

	return tst
}

func (s *ExtensionParamsContext) ExtensionParam(i int) IExtensionParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtensionParamContext)
}

func (s *ExtensionParamsContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(kslParserARG_DELIM)
}

func (s *ExtensionParamsContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(kslParserARG_DELIM, i)
}

func (s *ExtensionParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensionParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterExtensionParams(s)
	}
}

func (s *ExtensionParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitExtensionParams(s)
	}
}

func (s *ExtensionParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitExtensionParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) ExtensionParams() (localctx IExtensionParamsContext) {
	localctx = NewExtensionParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, kslParserRULE_extensionParams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.ExtensionParam()
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserARG_DELIM {
		{
			p.SetState(81)
			p.Match(kslParserARG_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.ExtensionParam()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExtensionReferenceContext is an interface to support dynamic dispatch.
type IExtensionReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTENSION_CALL() antlr.TerminalNode
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	ExtensionParams() IExtensionParamsContext
	RPAREN() antlr.TerminalNode
	RESOLVE() antlr.TerminalNode

	// IsExtensionReferenceContext differentiates from other interfaces.
	IsExtensionReferenceContext()
}

type ExtensionReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionReferenceContext() *ExtensionReferenceContext {
	var p = new(ExtensionReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extensionReference
	return p
}

func InitEmptyExtensionReferenceContext(p *ExtensionReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extensionReference
}

func (*ExtensionReferenceContext) IsExtensionReferenceContext() {}

func NewExtensionReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionReferenceContext {
	var p = new(ExtensionReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_extensionReference

	return p
}

func (s *ExtensionReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionReferenceContext) EXTENSION_CALL() antlr.TerminalNode {
	return s.GetToken(kslParserEXTENSION_CALL, 0)
}

func (s *ExtensionReferenceContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(kslParserNAME)
}

func (s *ExtensionReferenceContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(kslParserNAME, i)
}

func (s *ExtensionReferenceContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserLPAREN, 0)
}

func (s *ExtensionReferenceContext) ExtensionParams() IExtensionParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtensionParamsContext)
}

func (s *ExtensionReferenceContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserRPAREN, 0)
}

func (s *ExtensionReferenceContext) RESOLVE() antlr.TerminalNode {
	return s.GetToken(kslParserRESOLVE, 0)
}

func (s *ExtensionReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensionReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterExtensionReference(s)
	}
}

func (s *ExtensionReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitExtensionReference(s)
	}
}

func (s *ExtensionReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitExtensionReference(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) ExtensionReference() (localctx IExtensionReferenceContext) {
	localctx = NewExtensionReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, kslParserRULE_extensionReference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(kslParserEXTENSION_CALL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(89)
			p.Match(kslParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Match(kslParserRESOLVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(93)
		p.Match(kslParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(kslParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.ExtensionParams()
	}
	{
		p.SetState(96)
		p.Match(kslParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationContext is an interface to support dynamic dispatch.
type IRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RELATION() antlr.TerminalNode
	NAME() antlr.TerminalNode
	EXPAND() antlr.TerminalNode
	RelationBody() IRelationBodyContext
	AllExtensionReference() []IExtensionReferenceContext
	ExtensionReference(i int) IExtensionReferenceContext
	ACCESS() antlr.TerminalNode

	// IsRelationContext differentiates from other interfaces.
	IsRelationContext()
}

type RelationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationContext() *RelationContext {
	var p = new(RelationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_relation
	return p
}

func InitEmptyRelationContext(p *RelationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_relation
}

func (*RelationContext) IsRelationContext() {}

func NewRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationContext {
	var p = new(RelationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_relation

	return p
}

func (s *RelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationContext) RELATION() antlr.TerminalNode {
	return s.GetToken(kslParserRELATION, 0)
}

func (s *RelationContext) NAME() antlr.TerminalNode {
	return s.GetToken(kslParserNAME, 0)
}

func (s *RelationContext) EXPAND() antlr.TerminalNode {
	return s.GetToken(kslParserEXPAND, 0)
}

func (s *RelationContext) RelationBody() IRelationBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationBodyContext)
}

func (s *RelationContext) AllExtensionReference() []IExtensionReferenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtensionReferenceContext); ok {
			len++
		}
	}

	tst := make([]IExtensionReferenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtensionReferenceContext); ok {
			tst[i] = t.(IExtensionReferenceContext)
			i++
		}
	}

	return tst
}

func (s *RelationContext) ExtensionReference(i int) IExtensionReferenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionReferenceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtensionReferenceContext)
}

func (s *RelationContext) ACCESS() antlr.TerminalNode {
	return s.GetToken(kslParserACCESS, 0)
}

func (s *RelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterRelation(s)
	}
}

func (s *RelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitRelation(s)
	}
}

func (s *RelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) Relation() (localctx IRelationContext) {
	localctx = NewRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, kslParserRULE_relation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserEXTENSION_CALL {
		{
			p.SetState(98)
			p.ExtensionReference()
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(104)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(107)
		p.Match(kslParserRELATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(kslParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(kslParserEXPAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.relationBody(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationBodyContext is an interface to support dynamic dispatch.
type IRelationBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRelationBodyContext differentiates from other interfaces.
	IsRelationBodyContext()
}

type RelationBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationBodyContext() *RelationBodyContext {
	var p = new(RelationBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_relationBody
	return p
}

func InitEmptyRelationBodyContext(p *RelationBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_relationBody
}

func (*RelationBodyContext) IsRelationBodyContext() {}

func NewRelationBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationBodyContext {
	var p = new(RelationBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_relationBody

	return p
}

func (s *RelationBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationBodyContext) CopyAll(ctx *RelationBodyContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RelationBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ORContext struct {
	RelationBodyContext
}

func NewORContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ORContext {
	var p = new(ORContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *ORContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ORContext) AllRelationBody() []IRelationBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationBodyContext); ok {
			len++
		}
	}

	tst := make([]IRelationBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationBodyContext); ok {
			tst[i] = t.(IRelationBodyContext)
			i++
		}
	}

	return tst
}

func (s *ORContext) RelationBody(i int) IRelationBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationBodyContext)
}

func (s *ORContext) OR() antlr.TerminalNode {
	return s.GetToken(kslParserOR, 0)
}

func (s *ORContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterOR(s)
	}
}

func (s *ORContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitOR(s)
	}
}

func (s *ORContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitOR(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReferenceContext struct {
	RelationBodyContext
}

func NewReferenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReferenceContext {
	var p = new(ReferenceContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) NAME() antlr.TerminalNode {
	return s.GetToken(kslParserNAME, 0)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitReference(s)
	}
}

func (s *ReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitReference(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndContext struct {
	RelationBodyContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllRelationBody() []IRelationBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationBodyContext); ok {
			len++
		}
	}

	tst := make([]IRelationBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationBodyContext); ok {
			tst[i] = t.(IRelationBodyContext)
			i++
		}
	}

	return tst
}

func (s *AndContext) RelationBody(i int) IRelationBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationBodyContext)
}

func (s *AndContext) AND() antlr.TerminalNode {
	return s.GetToken(kslParserAND, 0)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelfContext struct {
	RelationBodyContext
}

func NewSelfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelfContext {
	var p = new(SelfContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *SelfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(kslParserLSQUARE, 0)
}

func (s *SelfContext) NAME() antlr.TerminalNode {
	return s.GetToken(kslParserNAME, 0)
}

func (s *SelfContext) RSQARE() antlr.TerminalNode {
	return s.GetToken(kslParserRSQARE, 0)
}

func (s *SelfContext) CARDINALITY() antlr.TerminalNode {
	return s.GetToken(kslParserCARDINALITY, 0)
}

func (s *SelfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterSelf(s)
	}
}

func (s *SelfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitSelf(s)
	}
}

func (s *SelfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitSelf(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubRelationContext struct {
	RelationBodyContext
}

func NewSubRelationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubRelationContext {
	var p = new(SubRelationContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *SubRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubRelationContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(kslParserNAME)
}

func (s *SubRelationContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(kslParserNAME, i)
}

func (s *SubRelationContext) RESOLVE() antlr.TerminalNode {
	return s.GetToken(kslParserRESOLVE, 0)
}

func (s *SubRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterSubRelation(s)
	}
}

func (s *SubRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitSubRelation(s)
	}
}

func (s *SubRelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitSubRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnlessContext struct {
	RelationBodyContext
}

func NewUnlessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnlessContext {
	var p = new(UnlessContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *UnlessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnlessContext) AllRelationBody() []IRelationBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationBodyContext); ok {
			len++
		}
	}

	tst := make([]IRelationBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationBodyContext); ok {
			tst[i] = t.(IRelationBodyContext)
			i++
		}
	}

	return tst
}

func (s *UnlessContext) RelationBody(i int) IRelationBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationBodyContext)
}

func (s *UnlessContext) UNLESS() antlr.TerminalNode {
	return s.GetToken(kslParserUNLESS, 0)
}

func (s *UnlessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterUnless(s)
	}
}

func (s *UnlessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitUnless(s)
	}
}

func (s *UnlessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitUnless(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenContext struct {
	RelationBodyContext
}

func NewParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenContext {
	var p = new(ParenContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *ParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserLPAREN, 0)
}

func (s *ParenContext) RelationBody() IRelationBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationBodyContext)
}

func (s *ParenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserRPAREN, 0)
}

func (s *ParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterParen(s)
	}
}

func (s *ParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitParen(s)
	}
}

func (s *ParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitParen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) RelationBody() (localctx IRelationBodyContext) {
	return p.relationBody(0)
}

func (p *kslParser) relationBody(_p int) (localctx IRelationBodyContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewRelationBodyContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelationBodyContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, kslParserRULE_relationBody, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(113)
			p.Match(kslParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kslParserCARDINALITY {
			{
				p.SetState(114)
				p.Match(kslParserCARDINALITY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(117)
			p.Match(kslParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Match(kslParserRSQARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewReferenceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(119)
			p.Match(kslParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSubRelationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(120)
			p.Match(kslParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(kslParserRESOLVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Match(kslParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(123)
			p.Match(kslParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.relationBody(0)
		}
		{
			p.SetState(125)
			p.Match(kslParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndContext(p, NewRelationBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_relationBody)
				p.SetState(129)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(130)
					p.Match(kslParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(131)
					p.relationBody(4)
				}

			case 2:
				localctx = NewORContext(p, NewRelationBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_relationBody)
				p.SetState(132)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(133)
					p.Match(kslParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(134)
					p.relationBody(3)
				}

			case 3:
				localctx = NewUnlessContext(p, NewRelationBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_relationBody)
				p.SetState(135)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(136)
					p.Match(kslParserUNLESS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(137)
					p.relationBody(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamNamesContext is an interface to support dynamic dispatch.
type IParamNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode
	ARG_DELIM() antlr.TerminalNode

	// IsParamNamesContext differentiates from other interfaces.
	IsParamNamesContext()
}

type ParamNamesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamNamesContext() *ParamNamesContext {
	var p = new(ParamNamesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_paramNames
	return p
}

func InitEmptyParamNamesContext(p *ParamNamesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_paramNames
}

func (*ParamNamesContext) IsParamNamesContext() {}

func NewParamNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamNamesContext {
	var p = new(ParamNamesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_paramNames

	return p
}

func (s *ParamNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamNamesContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(kslParserNAME)
}

func (s *ParamNamesContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(kslParserNAME, i)
}

func (s *ParamNamesContext) ARG_DELIM() antlr.TerminalNode {
	return s.GetToken(kslParserARG_DELIM, 0)
}

func (s *ParamNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterParamNames(s)
	}
}

func (s *ParamNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitParamNames(s)
	}
}

func (s *ParamNamesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitParamNames(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) ParamNames() (localctx IParamNamesContext) {
	localctx = NewParamNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, kslParserRULE_paramNames)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(kslParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserARG_DELIM {
		{
			p.SetState(144)
			p.Match(kslParserARG_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(kslParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExtensionContext is an interface to support dynamic dispatch.
type IExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTENSION() antlr.TerminalNode
	NAME() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	ParamNames() IParamNamesContext
	RPAREN() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	ACCESS() antlr.TerminalNode
	AllDynamicType() []IDynamicTypeContext
	DynamicType(i int) IDynamicTypeContext

	// IsExtensionContext differentiates from other interfaces.
	IsExtensionContext()
}

type ExtensionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionContext() *ExtensionContext {
	var p = new(ExtensionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extension
	return p
}

func InitEmptyExtensionContext(p *ExtensionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extension
}

func (*ExtensionContext) IsExtensionContext() {}

func NewExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionContext {
	var p = new(ExtensionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_extension

	return p
}

func (s *ExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionContext) EXTENSION() antlr.TerminalNode {
	return s.GetToken(kslParserEXTENSION, 0)
}

func (s *ExtensionContext) NAME() antlr.TerminalNode {
	return s.GetToken(kslParserNAME, 0)
}

func (s *ExtensionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserLPAREN, 0)
}

func (s *ExtensionContext) ParamNames() IParamNamesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamNamesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamNamesContext)
}

func (s *ExtensionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserRPAREN, 0)
}

func (s *ExtensionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserLBRACE, 0)
}

func (s *ExtensionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserRBRACE, 0)
}

func (s *ExtensionContext) ACCESS() antlr.TerminalNode {
	return s.GetToken(kslParserACCESS, 0)
}

func (s *ExtensionContext) AllDynamicType() []IDynamicTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicTypeContext); ok {
			len++
		}
	}

	tst := make([]IDynamicTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicTypeContext); ok {
			tst[i] = t.(IDynamicTypeContext)
			i++
		}
	}

	return tst
}

func (s *ExtensionContext) DynamicType(i int) IDynamicTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicTypeContext)
}

func (s *ExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterExtension(s)
	}
}

func (s *ExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitExtension(s)
	}
}

func (s *ExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) Extension() (localctx IExtensionContext) {
	localctx = NewExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, kslParserRULE_extension)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(148)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(151)
		p.Match(kslParserEXTENSION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(kslParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(kslParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.ParamNames()
	}
	{
		p.SetState(155)
		p.Match(kslParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(kslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == kslParserTYPE {
		{
			p.SetState(157)
			p.DynamicType()
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(162)
		p.Match(kslParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDynamicTypeContext is an interface to support dynamic dispatch.
type IDynamicTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	DynamicName() IDynamicNameContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllDynamicRelation() []IDynamicRelationContext
	DynamicRelation(i int) IDynamicRelationContext

	// IsDynamicTypeContext differentiates from other interfaces.
	IsDynamicTypeContext()
}

type DynamicTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicTypeContext() *DynamicTypeContext {
	var p = new(DynamicTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicType
	return p
}

func InitEmptyDynamicTypeContext(p *DynamicTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicType
}

func (*DynamicTypeContext) IsDynamicTypeContext() {}

func NewDynamicTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicTypeContext {
	var p = new(DynamicTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_dynamicType

	return p
}

func (s *DynamicTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DynamicTypeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(kslParserTYPE, 0)
}

func (s *DynamicTypeContext) DynamicName() IDynamicNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicNameContext)
}

func (s *DynamicTypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserLBRACE, 0)
}

func (s *DynamicTypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserRBRACE, 0)
}

func (s *DynamicTypeContext) AllDynamicRelation() []IDynamicRelationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicRelationContext); ok {
			len++
		}
	}

	tst := make([]IDynamicRelationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicRelationContext); ok {
			tst[i] = t.(IDynamicRelationContext)
			i++
		}
	}

	return tst
}

func (s *DynamicTypeContext) DynamicRelation(i int) IDynamicRelationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicRelationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicRelationContext)
}

func (s *DynamicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DynamicTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicType(s)
	}
}

func (s *DynamicTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicType(s)
	}
}

func (s *DynamicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) DynamicType() (localctx IDynamicTypeContext) {
	localctx = NewDynamicTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, kslParserRULE_dynamicType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(kslParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.DynamicName()
	}
	{
		p.SetState(166)
		p.Match(kslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4195360) != 0 {
		{
			p.SetState(167)
			p.DynamicRelation()
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(173)
		p.Match(kslParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDynamicRelationContext is an interface to support dynamic dispatch.
type IDynamicRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RELATION() antlr.TerminalNode
	DynamicName() IDynamicNameContext
	EXPAND() antlr.TerminalNode
	DynamicBody() IDynamicBodyContext
	ALLOW_DUPLICATES() antlr.TerminalNode
	ACCESS() antlr.TerminalNode

	// IsDynamicRelationContext differentiates from other interfaces.
	IsDynamicRelationContext()
}

type DynamicRelationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicRelationContext() *DynamicRelationContext {
	var p = new(DynamicRelationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicRelation
	return p
}

func InitEmptyDynamicRelationContext(p *DynamicRelationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicRelation
}

func (*DynamicRelationContext) IsDynamicRelationContext() {}

func NewDynamicRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicRelationContext {
	var p = new(DynamicRelationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_dynamicRelation

	return p
}

func (s *DynamicRelationContext) GetParser() antlr.Parser { return s.parser }

func (s *DynamicRelationContext) RELATION() antlr.TerminalNode {
	return s.GetToken(kslParserRELATION, 0)
}

func (s *DynamicRelationContext) DynamicName() IDynamicNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicNameContext)
}

func (s *DynamicRelationContext) EXPAND() antlr.TerminalNode {
	return s.GetToken(kslParserEXPAND, 0)
}

func (s *DynamicRelationContext) DynamicBody() IDynamicBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicBodyContext)
}

func (s *DynamicRelationContext) ALLOW_DUPLICATES() antlr.TerminalNode {
	return s.GetToken(kslParserALLOW_DUPLICATES, 0)
}

func (s *DynamicRelationContext) ACCESS() antlr.TerminalNode {
	return s.GetToken(kslParserACCESS, 0)
}

func (s *DynamicRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicRelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DynamicRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicRelation(s)
	}
}

func (s *DynamicRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicRelation(s)
	}
}

func (s *DynamicRelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) DynamicRelation() (localctx IDynamicRelationContext) {
	localctx = NewDynamicRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, kslParserRULE_dynamicRelation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserALLOW_DUPLICATES {
		{
			p.SetState(175)
			p.Match(kslParserALLOW_DUPLICATES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(178)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(181)
		p.Match(kslParserRELATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.DynamicName()
	}
	{
		p.SetState(183)
		p.Match(kslParserEXPAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.dynamicBody(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDynamicNameContext is an interface to support dynamic dispatch.
type IDynamicNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDynamicNameContext differentiates from other interfaces.
	IsDynamicNameContext()
}

type DynamicNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicNameContext() *DynamicNameContext {
	var p = new(DynamicNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicName
	return p
}

func InitEmptyDynamicNameContext(p *DynamicNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicName
}

func (*DynamicNameContext) IsDynamicNameContext() {}

func NewDynamicNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicNameContext {
	var p = new(DynamicNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_dynamicName

	return p
}

func (s *DynamicNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DynamicNameContext) CopyAll(ctx *DynamicNameContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DynamicNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VariableContext struct {
	DynamicNameContext
}

func NewVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableContext {
	var p = new(VariableContext)

	InitEmptyDynamicNameContext(&p.DynamicNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicNameContext))

	return p
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) VARREF() antlr.TerminalNode {
	return s.GetToken(kslParserVARREF, 0)
}

func (s *VariableContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserLBRACE, 0)
}

func (s *VariableContext) NAME() antlr.TerminalNode {
	return s.GetToken(kslParserNAME, 0)
}

func (s *VariableContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserRBRACE, 0)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralContext struct {
	DynamicNameContext
}

func NewLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralContext {
	var p = new(LiteralContext)

	InitEmptyDynamicNameContext(&p.DynamicNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicNameContext))

	return p
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) NAME() antlr.TerminalNode {
	return s.GetToken(kslParserNAME, 0)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type TemplateContext struct {
	DynamicNameContext
}

func NewTemplateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TemplateContext {
	var p = new(TemplateContext)

	InitEmptyDynamicNameContext(&p.DynamicNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicNameContext))

	return p
}

func (s *TemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateContext) AllTEMPLATE_DELIM() []antlr.TerminalNode {
	return s.GetTokens(kslParserTEMPLATE_DELIM)
}

func (s *TemplateContext) TEMPLATE_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(kslParserTEMPLATE_DELIM, i)
}

func (s *TemplateContext) AllDynamicName() []IDynamicNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicNameContext); ok {
			len++
		}
	}

	tst := make([]IDynamicNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicNameContext); ok {
			tst[i] = t.(IDynamicNameContext)
			i++
		}
	}

	return tst
}

func (s *TemplateContext) DynamicName(i int) IDynamicNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicNameContext)
}

func (s *TemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterTemplate(s)
	}
}

func (s *TemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitTemplate(s)
	}
}

func (s *TemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) DynamicName() (localctx IDynamicNameContext) {
	localctx = NewDynamicNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, kslParserRULE_dynamicName)
	var _alt int

	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kslParserNAME:
		localctx = NewLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Match(kslParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case kslParserVARREF:
		localctx = NewVariableContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			p.Match(kslParserVARREF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Match(kslParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Match(kslParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Match(kslParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case kslParserTEMPLATE_DELIM:
		localctx = NewTemplateContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.Match(kslParserTEMPLATE_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(192)
					p.DynamicName()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(kslParserTEMPLATE_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDynamicBodyContext is an interface to support dynamic dispatch.
type IDynamicBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDynamicBodyContext differentiates from other interfaces.
	IsDynamicBodyContext()
}

type DynamicBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicBodyContext() *DynamicBodyContext {
	var p = new(DynamicBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicBody
	return p
}

func InitEmptyDynamicBodyContext(p *DynamicBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicBody
}

func (*DynamicBodyContext) IsDynamicBodyContext() {}

func NewDynamicBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicBodyContext {
	var p = new(DynamicBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_dynamicBody

	return p
}

func (s *DynamicBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *DynamicBodyContext) CopyAll(ctx *DynamicBodyContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DynamicBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DynamicUnlessContext struct {
	DynamicBodyContext
}

func NewDynamicUnlessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicUnlessContext {
	var p = new(DynamicUnlessContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicUnlessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicUnlessContext) AllDynamicBody() []IDynamicBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			len++
		}
	}

	tst := make([]IDynamicBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicBodyContext); ok {
			tst[i] = t.(IDynamicBodyContext)
			i++
		}
	}

	return tst
}

func (s *DynamicUnlessContext) DynamicBody(i int) IDynamicBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicBodyContext)
}

func (s *DynamicUnlessContext) UNLESS() antlr.TerminalNode {
	return s.GetToken(kslParserUNLESS, 0)
}

func (s *DynamicUnlessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicUnless(s)
	}
}

func (s *DynamicUnlessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicUnless(s)
	}
}

func (s *DynamicUnlessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicUnless(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicORContext struct {
	DynamicBodyContext
}

func NewDynamicORContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicORContext {
	var p = new(DynamicORContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicORContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicORContext) AllDynamicBody() []IDynamicBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			len++
		}
	}

	tst := make([]IDynamicBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicBodyContext); ok {
			tst[i] = t.(IDynamicBodyContext)
			i++
		}
	}

	return tst
}

func (s *DynamicORContext) DynamicBody(i int) IDynamicBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicBodyContext)
}

func (s *DynamicORContext) OR() antlr.TerminalNode {
	return s.GetToken(kslParserOR, 0)
}

func (s *DynamicORContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicOR(s)
	}
}

func (s *DynamicORContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicOR(s)
	}
}

func (s *DynamicORContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicOR(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicParenContext struct {
	DynamicBodyContext
}

func NewDynamicParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicParenContext {
	var p = new(DynamicParenContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicParenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserLPAREN, 0)
}

func (s *DynamicParenContext) DynamicBody() IDynamicBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicBodyContext)
}

func (s *DynamicParenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserRPAREN, 0)
}

func (s *DynamicParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicParen(s)
	}
}

func (s *DynamicParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicParen(s)
	}
}

func (s *DynamicParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicParen(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicSelfContext struct {
	DynamicBodyContext
}

func NewDynamicSelfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicSelfContext {
	var p = new(DynamicSelfContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicSelfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicSelfContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(kslParserLSQUARE, 0)
}

func (s *DynamicSelfContext) DynamicName() IDynamicNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicNameContext)
}

func (s *DynamicSelfContext) RSQARE() antlr.TerminalNode {
	return s.GetToken(kslParserRSQARE, 0)
}

func (s *DynamicSelfContext) CARDINALITY() antlr.TerminalNode {
	return s.GetToken(kslParserCARDINALITY, 0)
}

func (s *DynamicSelfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicSelf(s)
	}
}

func (s *DynamicSelfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicSelf(s)
	}
}

func (s *DynamicSelfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicSelf(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicAndContext struct {
	DynamicBodyContext
}

func NewDynamicAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicAndContext {
	var p = new(DynamicAndContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicAndContext) AllDynamicBody() []IDynamicBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			len++
		}
	}

	tst := make([]IDynamicBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicBodyContext); ok {
			tst[i] = t.(IDynamicBodyContext)
			i++
		}
	}

	return tst
}

func (s *DynamicAndContext) DynamicBody(i int) IDynamicBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicBodyContext)
}

func (s *DynamicAndContext) AND() antlr.TerminalNode {
	return s.GetToken(kslParserAND, 0)
}

func (s *DynamicAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicAnd(s)
	}
}

func (s *DynamicAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicAnd(s)
	}
}

func (s *DynamicAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicReferenceContext struct {
	DynamicBodyContext
}

func NewDynamicReferenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicReferenceContext {
	var p = new(DynamicReferenceContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicReferenceContext) DynamicName() IDynamicNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicNameContext)
}

func (s *DynamicReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicReference(s)
	}
}

func (s *DynamicReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicReference(s)
	}
}

func (s *DynamicReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicReference(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicSubRelationContext struct {
	DynamicBodyContext
}

func NewDynamicSubRelationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicSubRelationContext {
	var p = new(DynamicSubRelationContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicSubRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicSubRelationContext) AllDynamicName() []IDynamicNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicNameContext); ok {
			len++
		}
	}

	tst := make([]IDynamicNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicNameContext); ok {
			tst[i] = t.(IDynamicNameContext)
			i++
		}
	}

	return tst
}

func (s *DynamicSubRelationContext) DynamicName(i int) IDynamicNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicNameContext)
}

func (s *DynamicSubRelationContext) RESOLVE() antlr.TerminalNode {
	return s.GetToken(kslParserRESOLVE, 0)
}

func (s *DynamicSubRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicSubRelation(s)
	}
}

func (s *DynamicSubRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicSubRelation(s)
	}
}

func (s *DynamicSubRelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicSubRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) DynamicBody() (localctx IDynamicBodyContext) {
	return p.dynamicBody(0)
}

func (p *kslParser) dynamicBody(_p int) (localctx IDynamicBodyContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewDynamicBodyContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDynamicBodyContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, kslParserRULE_dynamicBody, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDynamicSelfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(202)
			p.Match(kslParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kslParserCARDINALITY {
			{
				p.SetState(203)
				p.Match(kslParserCARDINALITY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(206)
			p.DynamicName()
		}
		{
			p.SetState(207)
			p.Match(kslParserRSQARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDynamicReferenceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(209)
			p.DynamicName()
		}

	case 3:
		localctx = NewDynamicSubRelationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(210)
			p.DynamicName()
		}
		{
			p.SetState(211)
			p.Match(kslParserRESOLVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.DynamicName()
		}

	case 4:
		localctx = NewDynamicParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(214)
			p.Match(kslParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.dynamicBody(0)
		}
		{
			p.SetState(216)
			p.Match(kslParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(229)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDynamicAndContext(p, NewDynamicBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_dynamicBody)
				p.SetState(220)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(221)
					p.Match(kslParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(222)
					p.dynamicBody(4)
				}

			case 2:
				localctx = NewDynamicORContext(p, NewDynamicBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_dynamicBody)
				p.SetState(223)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(224)
					p.Match(kslParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(225)
					p.dynamicBody(3)
				}

			case 3:
				localctx = NewDynamicUnlessContext(p, NewDynamicBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_dynamicBody)
				p.SetState(226)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(227)
					p.Match(kslParserUNLESS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(228)
					p.dynamicBody(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *kslParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *RelationBodyContext = nil
		if localctx != nil {
			t = localctx.(*RelationBodyContext)
		}
		return p.RelationBody_Sempred(t, predIndex)

	case 16:
		var t *DynamicBodyContext = nil
		if localctx != nil {
			t = localctx.(*DynamicBodyContext)
		}
		return p.DynamicBody_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *kslParser) RelationBody_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *kslParser) DynamicBody_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}