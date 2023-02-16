package bslparser

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/asosnoviy/go1cover/pkg/bslparser/parser"
)

func getLines(p parser.IFileContext) map[int]bool {

	result := map[int]bool{}
	desc := antlr.TreesDescendants(p)

	for _, pt := range desc {
		if !mustCovered(pt) {
			continue
		}

		if filterByName(pt) {
			continue
		}

		lineNo := getline(pt)
		if lineNo > 0 {
			_, has := result[lineNo]
			if !has {
				result[lineNo] = false
			}
		}
	}

	return result

}

func getline(pt antlr.ParseTree) int {

	s, ok := pt.(antlr.ParserRuleContext)
	if ok {
		_, ok := s.(*parser.MethodCallContext)
		if !ok {
			return s.GetStart().GetLine()
		}

		methodcall := getRootParent(s, parser.BSLParserRULE_complexIdentifier)
		if methodcall != nil {
			return methodcall.GetStart().GetLine()
		}

	}

	v, ok := pt.(*antlr.TerminalNodeImpl)
	if ok {
		return v.GetSymbol().GetLine()
	}

	return 0
}

func getRootParent(node antlr.ParserRuleContext, ruleindex int) antlr.ParserRuleContext {

	parent := node.GetParent()
	if parent == nil {
		return nil
	}

	s, ok := parent.(antlr.ParserRuleContext)
	if ok {
		if getRuleIndex(s) == ruleindex {
			return s
		}

		return getRootParent(s, ruleindex)
	}

	return nil
}

func getRuleIndex(node antlr.ParserRuleContext) int {

	s, ok := node.(antlr.TerminalNode)
	if ok {
		return s.GetSymbol().GetTokenType()
	}

	return node.GetRuleIndex()

}

func mustCovered(p antlr.ParseTree) bool {

	v, ok := p.(antlr.TerminalNode)
	if ok {
		node := v.GetSymbol().GetTokenType()
		return node == parser.BSLLexerENDDO_KEYWORD ||
			node == parser.BSLLexerENDFUNCTION_KEYWORD ||
			node == parser.BSLLexerENDPROCEDURE_KEYWORD ||
			node == parser.BSLLexerENDTRY_KEYWORD ||
			node == parser.BSLLexerENDIF_KEYWORD ||
			node == parser.BSLLexerDO_KEYWORD
	}

	s, ok := p.(antlr.ParserRuleContext)
	if ok {

		_, ok := s.(*parser.AssignmentContext)
		if ok {
			return true
		}

		_, ok = s.(*parser.CallStatementContext)
		if ok {
			return true
		}

		_, ok = s.(*parser.GotoStatementContext)
		if ok {
			return true
		}

		_, ok = s.(*parser.ReturnStatementContext)
		if ok {
			return true
		}
		_, ok = s.(*parser.BreakStatementContext)
		if ok {
			return true
		}
		_, ok = s.(*parser.ContinueStatementContext)
		if ok {
			return true
		}
		_, ok = s.(*parser.IfStatementContext)
		if ok {
			return true
		}
		_, ok = s.(*parser.ElsifBranchContext)
		if ok {
			return true
		}
		_, ok = s.(*parser.RaiseStatementContext)
		if ok {
			return true
		}
		_, ok = s.(*parser.ForEachStatementContext)
		if ok {
			return true
		}
		_, ok = s.(*parser.ForStatementContext)
		if ok {
			return true
		}
		_, ok = s.(*parser.WhileStatementContext)
		if ok {
			return true
		}
		_, ok = s.(*parser.GlobalMethodCallContext)
		if ok {
			return true
		}
		_, ok = s.(*parser.MethodCallContext)
		if ok {
			return true
		}
		_, ok = s.(*parser.ExecuteStatementContext)
		if ok {
			return true
		}
		_, ok = s.(*parser.AddHandlerStatementContext)
		if ok {
			return true
		}
		_, ok = s.(*parser.RemoveHandlerStatementContext)
		if ok {
			return true
		}

	}

	return false
}

var filterMethodsName map[string]bool = map[string]bool{
	"eval":               true,
	"вычислить":          true,
	"boolean":            true,
	"булево":             true,
	"number":             true,
	"число":              true,
	"string":             true,
	"строка":             true,
	"date":               true,
	"дата":               true,
	"type":               true,
	"тип":                true,
	"typeof":             true,
	"типзнч":             true,
	"strlen":             true,
	"стрдлина":           true,
	"triml":              true,
	"сокрл":              true,
	"trimr":              true,
	"сокрп":              true,
	"trimall":            true,
	"сокрлп":             true,
	"left":               true,
	"лев":                true,
	"right":              true,
	"прав":               true,
	"mid":                true,
	"сред":               true,
	"strpos":             true,
	"upper":              true,
	"врег":               true,
	"lower":              true,
	"нрег":               true,
	"title":              true,
	"трег":               true,
	"char":               true,
	"символ":             true,
	"charcode":           true,
	"кодсимвола":         true,
	"isblankstring":      true,
	"пустаястрока":       true,
	"strreplace":         true,
	"стрзаменить":        true,
	"strgetline":         true,
	"стрполучитьстроку":  true,
	"strlinecount":       true,
	"стрчислострок":      true,
	"stroccurrencecount": true,
	"стрчисловхождений":  true,
	"year":               true,
	"год":                true,
	"month":              true,
	"месяц":              true,
	"day":                true,
	"день":               true,
	"hour":               true,
	"час":                true,
	"minute":             true,
	"минута":             true,
	"second":             true,
	"секунда":            true,
	"begofyear":          true,
	"началогода":         true,
	"begofmonth":         true,
	"началомесяца":       true,
	"begofweek":          true,
	"началонедели":       true,
	"begofday":           true,
	"началодня":          true,
	"begofhour":          true,
	"началочаса":         true,
	"begofminute":        true,
	"началоминуты":       true,
	"begofquarter":       true,
	"началоквартала":     true,
	"endofyear":          true,
	"конецгода":          true,
	"endofmonth":         true,
	"конецмесяца":        true,
	"endofweek":          true,
	"конецнедели":        true,
	"endofday":           true,
	"конецдня":           true,
	"endofhour":          true,
	"конецчаса":          true,
	"endofminute":        true,
	"конецминуты":        true,
	"endofquarter":       true,
	"конецквартала":      true,
	"weekofyear":         true,
	"неделягода":         true,
	"dayofyear":          true,
	"деньгода":           true,
	"weekday":            true,
	"деньнедели":         true,
	"addmonth":           true,
	"добавитьмесяц":      true,
	"currentdate":        true,
	"текущаядата":        true,
	"int":                true,
	"цел":                true,
	"round":              true,
	"окр":                true,
	"log":                true,
	"log10":              true,
	"sin":                true,
	"cos":                true,
	"tan":                true,
	"asin":               true,
	"acos":               true,
	"atan":               true,
	"exp":                true,
	"pow":                true,
	"sqrt":               true,
	"min":                true,
	"мин":                true,
	"max":                true,
	"макс":               true,
	"format":             true,
	"формат":             true,
	"errorinfo":          true,
	"информацияобошибке": true,
	"errordescription":   true,
	"описаниеошибки":     true,
	"найти":              true,
	"find":               true,
}

func filterByName(p antlr.ParseTree) bool {

	v, ok := p.(*parser.GlobalMethodCallContext)
	if ok {
		name := strings.ToLower(v.MethodName().GetText())
		_, ok := filterMethodsName[name]
		return ok
	}

	return false
}
