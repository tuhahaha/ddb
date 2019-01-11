// Code generated by goyacc -l -o parser.go parser.y. DO NOT EDIT.

package dataManager

import __yyfmt__ "fmt"
import (
	"strconv"
)

// 这个函数用来将最终获得的statement赋值给lex
func setStmtResult(l yyLexer, v Statement) {
	l.(*lex).statement = v
}

// 这个函数用来设置FragmentName, FragmentNum，我们在语法分析的最后阶段使用
func setFragment(f *Fragment) {
	for i, frag := range f.FragmentSpec.Fragments {
		frag.FragmentName = f.TableName + "." + strconv.Itoa(i+1)
	}
	f.FragmentNum = len(f.FragmentSpec.Fragments)
}

type yySymType struct {
	yys            int
	stmtType       int
	stmt           Statement
	str            string
	strs           []string //
	num            string   // 用来存储NUMBER的数值
	dbddl          DBDDL
	column         *ColumnDefinition // 单个column的定义 如 id int
	columns        TableSpec         // 存储多个column的定义
	ddl            DDL               // ddl statement
	b              bool              // 用于 notnull 和 primary key
	insert         Insert            // insert statement
	delete         Delete            // delete statement
	fragment       Fragment
	fragment_def   FragmentDefinition
	fragment_def_p *FragmentDefinition
	fragment_spec  FragmentSpec
	allocate       Allocate
	load           Load
}

const LEX_ERROR = 57346
const UNUSED = 57347
const OR = 57348
const AND = 57349
const CREATE = 57350
const ALTER = 57351
const DROP = 57352
const UPDATE = 57353
const DELETE = 57354
const INSERT = 57355
const LOAD = 57356
const INTO = 57357
const DATABASE = 57358
const SCHEMA = 57359
const TABLE = 57360
const PRIMARY = 57361
const VALUES = 57362
const PARTITION = 57363
const BY = 57364
const IN = 57365
const FROM = 57366
const WHERE = 57367
const TO = 57368
const NOT = 57369
const NULL = 57370
const KEY = 57371
const INT = 57372
const FLOAT = 57373
const DOUBLE = 57374
const CHAR = 57375
const VARCHAR = 57376
const NCHAR = 57377
const TEXT = 57378
const NUMBER = 57379
const IDENT = 57380
const STRING = 57381
const NAMEWITHDOT = 57382
const FRAGMENT = 57383
const HORIZONTALLY = 57384
const VERTICALLY = 57385
const ALLOCATE = 57386

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"UNUSED",
	"OR",
	"AND",
	"CREATE",
	"ALTER",
	"DROP",
	"UPDATE",
	"DELETE",
	"INSERT",
	"LOAD",
	"INTO",
	"DATABASE",
	"SCHEMA",
	"TABLE",
	"PRIMARY",
	"VALUES",
	"PARTITION",
	"BY",
	"IN",
	"FROM",
	"WHERE",
	"TO",
	"NOT",
	"NULL",
	"KEY",
	"INT",
	"FLOAT",
	"DOUBLE",
	"CHAR",
	"VARCHAR",
	"NCHAR",
	"TEXT",
	"NUMBER",
	"IDENT",
	"STRING",
	"NAMEWITHDOT",
	"FRAGMENT",
	"HORIZONTALLY",
	"VERTICALLY",
	"ALLOCATE",
	"';'",
	"'('",
	"')'",
	"','",
	"'&'",
	"'>'",
	"'='",
	"'<'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 120

var yyAct = [...]int{

	52, 74, 63, 60, 56, 59, 51, 80, 81, 82,
	83, 105, 104, 87, 86, 103, 90, 3, 110, 111,
	84, 11, 98, 99, 120, 13, 12, 17, 65, 66,
	64, 97, 54, 43, 25, 24, 23, 22, 21, 20,
	19, 41, 42, 75, 30, 76, 101, 47, 102, 53,
	61, 31, 57, 46, 18, 32, 88, 16, 35, 116,
	108, 68, 69, 70, 72, 92, 114, 119, 115, 89,
	33, 95, 39, 93, 45, 40, 113, 29, 44, 36,
	37, 38, 100, 26, 49, 27, 48, 28, 1, 107,
	106, 79, 78, 109, 10, 77, 55, 7, 73, 6,
	9, 117, 91, 62, 58, 15, 14, 8, 85, 96,
	71, 50, 118, 112, 94, 67, 5, 34, 4, 2,
}
var yyPact = [...]int{

	13, -1000, -1000, -1000, -5, -6, -7, -8, -9, -10,
	-11, 67, 72, 53, -1000, -1000, 4, 17, 17, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 20, 17, 17, 17,
	46, 51, -1000, -1, -1000, -1000, -13, 58, 49, 15,
	8, 71, 69, 11, -14, 14, -1000, -1000, 12, -16,
	-19, -1000, 31, -1000, 6, 85, -1000, -42, -28, 7,
	-1000, -42, -32, -1000, 11, -1000, 11, 44, -1000, -1000,
	-1000, -1000, -15, -25, -1000, -1000, -1000, 14, -1000, -1000,
	9, -36, -1000, -39, 12, 12, -1000, -1000, -1000, 23,
	-16, -29, -1000, -1000, 47, 40, -1000, 22, -1000, 6,
	-1000, -1000, -1000, -1000, -1000, -1000, 7, -1000, -1000, -1000,
	-1000, 11, -1000, -1000, 38, -1000, -23, -1000, -1000, -1000,
	-1000,
}
var yyPgo = [...]int{

	0, 119, 118, 117, 116, 51, 0, 115, 114, 113,
	6, 111, 110, 109, 108, 7, 107, 106, 105, 104,
	103, 5, 2, 3, 102, 100, 99, 98, 1, 97,
	96, 4, 95, 94, 88,
}
var yyR1 = [...]int{

	0, 34, 34, 1, 1, 1, 1, 1, 1, 1,
	2, 3, 4, 5, 11, 11, 10, 6, 7, 7,
	7, 7, 12, 13, 8, 8, 9, 9, 9, 16,
	16, 17, 18, 19, 19, 21, 21, 23, 14, 14,
	14, 15, 15, 15, 15, 15, 15, 20, 20, 22,
	24, 24, 25, 26, 27, 27, 28, 28, 29, 29,
	30, 30, 32, 32, 31, 31, 33,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	3, 1, 6, 1, 1, 3, 4, 1, 1, 1,
	1, 1, 2, 3, 0, 2, 0, 1, 2, 1,
	1, 5, 5, 1, 3, 1, 3, 3, 1, 1,
	1, 1, 1, 1, 2, 2, 2, 1, 3, 3,
	1, 3, 4, 7, 1, 3, 1, 1, 3, 5,
	1, 3, 1, 1, 3, 3, 4,
}
var yyChk = [...]int{

	-1000, -34, -1, 4, -2, -4, -26, -29, -16, -25,
	-33, 8, 13, 12, -17, -18, 44, 14, 41, 45,
	45, 45, 45, 45, 45, 45, 16, 18, 15, 24,
	40, -5, 38, -5, -3, 38, -5, -5, -5, 26,
	24, 42, 43, 46, 20, 25, 38, 39, 15, 15,
	-11, -10, -6, 38, 46, -30, -31, 38, -19, -21,
	-23, 38, -20, -22, 46, 47, 48, -7, 30, 31,
	32, -12, 33, -27, -28, 37, 39, -32, 7, 6,
	-15, 50, 51, 52, 48, -14, 7, 6, 49, -15,
	48, -24, -6, -10, -8, 27, -13, 46, 47, 48,
	-31, 37, 39, 51, 51, 50, -21, -23, 37, -22,
	47, 48, -9, 29, 19, 28, 37, -28, -6, 29,
	47,
}
var yyDef = [...]int{

	0, -2, 1, 2, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 29, 30, 0, 0, 0, 3,
	4, 5, 6, 7, 8, 9, 0, 0, 0, 0,
	0, 0, 13, 0, 10, 11, 0, 0, 58, 0,
	0, 0, 0, 0, 0, 0, 52, 66, 0, 0,
	0, 14, 0, 17, 0, 59, 60, 0, 31, 33,
	35, 0, 32, 47, 0, 12, 0, 24, 18, 19,
	20, 21, 0, 0, 54, 56, 57, 0, 62, 63,
	0, 41, 42, 43, 0, 0, 38, 39, 40, 0,
	0, 0, 50, 15, 26, 0, 22, 0, 53, 0,
	61, 64, 65, 44, 45, 46, 34, 36, 37, 48,
	49, 0, 16, 27, 0, 25, 0, 55, 51, 28,
	23,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 49, 3,
	46, 47, 3, 3, 48, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 45,
	52, 51, 50,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			setStmtResult(yylex, yyDollar[1].stmt)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].dbddl
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].ddl
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].insert
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].delete
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].fragment
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].allocate
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].load
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.dbddl = DBDDL{Action: "CREATE", DbName: yyDollar[3].str, Part_DB: nil}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = yyDollar[1].str
		}
	case 12:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.ddl = DDL{Action: "CREATE", TableName: yyDollar[3].str, TableSpec: &(yyDollar[5].columns)}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = yyDollar[1].str
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.columns.Columns = append(yyVAL.columns.Columns, yyDollar[1].column)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.columns.Columns = append(yyVAL.columns.Columns, yyDollar[3].column)
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.column = &ColumnDefinition{Name: yyDollar[1].str, DataType: yyDollar[2].str, Not_Null: yyDollar[3].b, Primarykey: yyDollar[4].b}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = yyDollar[1].str
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = "int"
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = "float"
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = "double"
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = yyDollar[1].str
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.str = "char(" + yyDollar[2].str + ")"
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.str = yyDollar[2].num
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.b = false
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.b = true
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.b = false
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.b = true
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.b = true
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.fragment = yyDollar[1].fragment
			setFragment(&(yyVAL.fragment))
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.fragment = yyDollar[1].fragment
			setFragment(&(yyVAL.fragment))
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.fragment = Fragment{TableName: yyDollar[2].str, FragmentMode: "H", FragmentNum: 0, FragmentSpec: &(yyDollar[5].fragment_spec)}
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.fragment = Fragment{TableName: yyDollar[2].str, FragmentMode: "V", FragmentNum: 0, FragmentSpec: &(yyDollar[5].fragment_spec)}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			// FIXME
			temp := yyDollar[1].fragment_def
			yyVAL.fragment_spec.Fragments = append(yyVAL.fragment_spec.Fragments, &(temp))
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			// FIXME
			temp := yyDollar[3].fragment_def
			yyVAL.fragment_spec.Fragments = append(yyVAL.fragment_spec.Fragments, &(temp))
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.fragment_def.ColumnList = append(yyVAL.fragment_def.ColumnList, yyDollar[1].strs[0]) // 水平分片的条件属性赋给columnlist作为分片属性存储
			yyVAL.fragment_def.Limit += yyDollar[1].strs[1]
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.fragment_def.ColumnList = append(yyVAL.fragment_def.ColumnList, yyDollar[3].strs[0])
			yyVAL.fragment_def.Limit += yyDollar[2].str + yyDollar[3].strs[1]
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.strs = []string{"", ""}
			yyVAL.strs[0] = yyDollar[1].str                                     // $$[0]
			yyVAL.strs[1] = yyDollar[1].str + yyDollar[2].str + yyDollar[3].num // $$[1] 我们用用一个字符串切片存储 columnname 和 limit
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = " and "
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = " | "
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = " & "
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = ">"
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = "="
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = "<"
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.str = ">="
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.str = "<="
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.str = "<>"
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.fragment_spec.Fragments = append(yyVAL.fragment_spec.Fragments, yyDollar[1].fragment_def_p)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.fragment_spec.Fragments = append(yyVAL.fragment_spec.Fragments, yyDollar[3].fragment_def_p)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.fragment_def_p = &FragmentDefinition{FragmentName: "", ColumnList: yyDollar[2].strs, Limit: "*"}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.strs = append(yyVAL.strs, yyDollar[1].str)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.strs = append(yyVAL.strs, yyDollar[3].str)
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.allocate = Allocate{FragmentName: yyDollar[2].str, Site: yyDollar[4].str}
		}
	case 53:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.insert = Insert{TableName: yyDollar[3].str, InsertMap: nil, InsertValues: yyDollar[6].strs} // InsertColumns: 我们会在后面根据gdd进行更新map
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.strs = append(yyVAL.strs, yyDollar[1].str)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.strs = append(yyVAL.strs, yyDollar[3].str)
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = yyDollar[1].num
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = yyDollar[1].str
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.delete = Delete{TargetName: yyDollar[3].str, Condition: "*"}
		}
	case 59:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.delete = Delete{TargetName: yyDollar[3].str, Condition: yyDollar[5].strs[1], DeleteCols: yyDollar[5].strs[0]}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.strs = append(yyVAL.strs, yyDollar[1].strs[0])
			yyVAL.strs = append(yyVAL.strs, yyDollar[1].strs[1])
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.strs[1] += yyDollar[2].str + yyDollar[3].strs[1] // 同上， 使用$$ 和 $1 一样
			yyVAL.strs[0] += "," + yyDollar[3].strs[0]             // 只加逗号不加空格 谢谢 :)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = " and "
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.str = " or "
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.strs = append(yyVAL.strs, yyDollar[1].str) // for delete vertically table with condition where
			yyVAL.strs = append(yyVAL.strs, yyDollar[1].str+yyDollar[2].str+yyDollar[3].num)

		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.strs = append(yyVAL.strs, yyDollar[1].str) // for delete vertically table with condition where
			yyVAL.strs = append(yyVAL.strs, yyDollar[1].str+yyDollar[2].str+yyDollar[3].str)
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.load = Load{tableName: yyDollar[2].str, filePath: yyDollar[4].str}
		}
	}
	goto yystack /* stack new state and value */
}
