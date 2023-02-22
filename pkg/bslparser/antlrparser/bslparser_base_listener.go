// Code generated from BSLParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // BSLParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBSLParserListener is a complete listener for a parse tree produced by BSLParser.
type BaseBSLParserListener struct{}

var _ BSLParserListener = &BaseBSLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBSLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBSLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBSLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBSLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseBSLParserListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseBSLParserListener) ExitFile(ctx *FileContext) {}

// EnterShebang is called when production shebang is entered.
func (s *BaseBSLParserListener) EnterShebang(ctx *ShebangContext) {}

// ExitShebang is called when production shebang is exited.
func (s *BaseBSLParserListener) ExitShebang(ctx *ShebangContext) {}

// EnterUsedLib is called when production usedLib is entered.
func (s *BaseBSLParserListener) EnterUsedLib(ctx *UsedLibContext) {}

// ExitUsedLib is called when production usedLib is exited.
func (s *BaseBSLParserListener) ExitUsedLib(ctx *UsedLibContext) {}

// EnterUse is called when production use is entered.
func (s *BaseBSLParserListener) EnterUse(ctx *UseContext) {}

// ExitUse is called when production use is exited.
func (s *BaseBSLParserListener) ExitUse(ctx *UseContext) {}

// EnterRegionStart is called when production regionStart is entered.
func (s *BaseBSLParserListener) EnterRegionStart(ctx *RegionStartContext) {}

// ExitRegionStart is called when production regionStart is exited.
func (s *BaseBSLParserListener) ExitRegionStart(ctx *RegionStartContext) {}

// EnterRegionEnd is called when production regionEnd is entered.
func (s *BaseBSLParserListener) EnterRegionEnd(ctx *RegionEndContext) {}

// ExitRegionEnd is called when production regionEnd is exited.
func (s *BaseBSLParserListener) ExitRegionEnd(ctx *RegionEndContext) {}

// EnterRegionName is called when production regionName is entered.
func (s *BaseBSLParserListener) EnterRegionName(ctx *RegionNameContext) {}

// ExitRegionName is called when production regionName is exited.
func (s *BaseBSLParserListener) ExitRegionName(ctx *RegionNameContext) {}

// EnterPreproc_if is called when production preproc_if is entered.
func (s *BaseBSLParserListener) EnterPreproc_if(ctx *Preproc_ifContext) {}

// ExitPreproc_if is called when production preproc_if is exited.
func (s *BaseBSLParserListener) ExitPreproc_if(ctx *Preproc_ifContext) {}

// EnterPreproc_elsif is called when production preproc_elsif is entered.
func (s *BaseBSLParserListener) EnterPreproc_elsif(ctx *Preproc_elsifContext) {}

// ExitPreproc_elsif is called when production preproc_elsif is exited.
func (s *BaseBSLParserListener) ExitPreproc_elsif(ctx *Preproc_elsifContext) {}

// EnterPreproc_else is called when production preproc_else is entered.
func (s *BaseBSLParserListener) EnterPreproc_else(ctx *Preproc_elseContext) {}

// ExitPreproc_else is called when production preproc_else is exited.
func (s *BaseBSLParserListener) ExitPreproc_else(ctx *Preproc_elseContext) {}

// EnterPreproc_endif is called when production preproc_endif is entered.
func (s *BaseBSLParserListener) EnterPreproc_endif(ctx *Preproc_endifContext) {}

// ExitPreproc_endif is called when production preproc_endif is exited.
func (s *BaseBSLParserListener) ExitPreproc_endif(ctx *Preproc_endifContext) {}

// EnterPreproc_expression is called when production preproc_expression is entered.
func (s *BaseBSLParserListener) EnterPreproc_expression(ctx *Preproc_expressionContext) {}

// ExitPreproc_expression is called when production preproc_expression is exited.
func (s *BaseBSLParserListener) ExitPreproc_expression(ctx *Preproc_expressionContext) {}

// EnterPreproc_logicalOperand is called when production preproc_logicalOperand is entered.
func (s *BaseBSLParserListener) EnterPreproc_logicalOperand(ctx *Preproc_logicalOperandContext) {}

// ExitPreproc_logicalOperand is called when production preproc_logicalOperand is exited.
func (s *BaseBSLParserListener) ExitPreproc_logicalOperand(ctx *Preproc_logicalOperandContext) {}

// EnterPreproc_logicalExpression is called when production preproc_logicalExpression is entered.
func (s *BaseBSLParserListener) EnterPreproc_logicalExpression(ctx *Preproc_logicalExpressionContext) {
}

// ExitPreproc_logicalExpression is called when production preproc_logicalExpression is exited.
func (s *BaseBSLParserListener) ExitPreproc_logicalExpression(ctx *Preproc_logicalExpressionContext) {
}

// EnterPreproc_symbol is called when production preproc_symbol is entered.
func (s *BaseBSLParserListener) EnterPreproc_symbol(ctx *Preproc_symbolContext) {}

// ExitPreproc_symbol is called when production preproc_symbol is exited.
func (s *BaseBSLParserListener) ExitPreproc_symbol(ctx *Preproc_symbolContext) {}

// EnterPreproc_unknownSymbol is called when production preproc_unknownSymbol is entered.
func (s *BaseBSLParserListener) EnterPreproc_unknownSymbol(ctx *Preproc_unknownSymbolContext) {}

// ExitPreproc_unknownSymbol is called when production preproc_unknownSymbol is exited.
func (s *BaseBSLParserListener) ExitPreproc_unknownSymbol(ctx *Preproc_unknownSymbolContext) {}

// EnterPreproc_boolOperation is called when production preproc_boolOperation is entered.
func (s *BaseBSLParserListener) EnterPreproc_boolOperation(ctx *Preproc_boolOperationContext) {}

// ExitPreproc_boolOperation is called when production preproc_boolOperation is exited.
func (s *BaseBSLParserListener) ExitPreproc_boolOperation(ctx *Preproc_boolOperationContext) {}

// EnterPreprocessor is called when production preprocessor is entered.
func (s *BaseBSLParserListener) EnterPreprocessor(ctx *PreprocessorContext) {}

// ExitPreprocessor is called when production preprocessor is exited.
func (s *BaseBSLParserListener) ExitPreprocessor(ctx *PreprocessorContext) {}

// EnterCompilerDirectiveSymbol is called when production compilerDirectiveSymbol is entered.
func (s *BaseBSLParserListener) EnterCompilerDirectiveSymbol(ctx *CompilerDirectiveSymbolContext) {}

// ExitCompilerDirectiveSymbol is called when production compilerDirectiveSymbol is exited.
func (s *BaseBSLParserListener) ExitCompilerDirectiveSymbol(ctx *CompilerDirectiveSymbolContext) {}

// EnterCompilerDirective is called when production compilerDirective is entered.
func (s *BaseBSLParserListener) EnterCompilerDirective(ctx *CompilerDirectiveContext) {}

// ExitCompilerDirective is called when production compilerDirective is exited.
func (s *BaseBSLParserListener) ExitCompilerDirective(ctx *CompilerDirectiveContext) {}

// EnterAnnotationName is called when production annotationName is entered.
func (s *BaseBSLParserListener) EnterAnnotationName(ctx *AnnotationNameContext) {}

// ExitAnnotationName is called when production annotationName is exited.
func (s *BaseBSLParserListener) ExitAnnotationName(ctx *AnnotationNameContext) {}

// EnterAnnotationParamName is called when production annotationParamName is entered.
func (s *BaseBSLParserListener) EnterAnnotationParamName(ctx *AnnotationParamNameContext) {}

// ExitAnnotationParamName is called when production annotationParamName is exited.
func (s *BaseBSLParserListener) ExitAnnotationParamName(ctx *AnnotationParamNameContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseBSLParserListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseBSLParserListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterAnnotationParams is called when production annotationParams is entered.
func (s *BaseBSLParserListener) EnterAnnotationParams(ctx *AnnotationParamsContext) {}

// ExitAnnotationParams is called when production annotationParams is exited.
func (s *BaseBSLParserListener) ExitAnnotationParams(ctx *AnnotationParamsContext) {}

// EnterAnnotationParam is called when production annotationParam is entered.
func (s *BaseBSLParserListener) EnterAnnotationParam(ctx *AnnotationParamContext) {}

// ExitAnnotationParam is called when production annotationParam is exited.
func (s *BaseBSLParserListener) ExitAnnotationParam(ctx *AnnotationParamContext) {}

// EnterVar_name is called when production var_name is entered.
func (s *BaseBSLParserListener) EnterVar_name(ctx *Var_nameContext) {}

// ExitVar_name is called when production var_name is exited.
func (s *BaseBSLParserListener) ExitVar_name(ctx *Var_nameContext) {}

// EnterModuleVars is called when production moduleVars is entered.
func (s *BaseBSLParserListener) EnterModuleVars(ctx *ModuleVarsContext) {}

// ExitModuleVars is called when production moduleVars is exited.
func (s *BaseBSLParserListener) ExitModuleVars(ctx *ModuleVarsContext) {}

// EnterModuleVar is called when production moduleVar is entered.
func (s *BaseBSLParserListener) EnterModuleVar(ctx *ModuleVarContext) {}

// ExitModuleVar is called when production moduleVar is exited.
func (s *BaseBSLParserListener) ExitModuleVar(ctx *ModuleVarContext) {}

// EnterModuleVarsList is called when production moduleVarsList is entered.
func (s *BaseBSLParserListener) EnterModuleVarsList(ctx *ModuleVarsListContext) {}

// ExitModuleVarsList is called when production moduleVarsList is exited.
func (s *BaseBSLParserListener) ExitModuleVarsList(ctx *ModuleVarsListContext) {}

// EnterModuleVarDeclaration is called when production moduleVarDeclaration is entered.
func (s *BaseBSLParserListener) EnterModuleVarDeclaration(ctx *ModuleVarDeclarationContext) {}

// ExitModuleVarDeclaration is called when production moduleVarDeclaration is exited.
func (s *BaseBSLParserListener) ExitModuleVarDeclaration(ctx *ModuleVarDeclarationContext) {}

// EnterSubVars is called when production subVars is entered.
func (s *BaseBSLParserListener) EnterSubVars(ctx *SubVarsContext) {}

// ExitSubVars is called when production subVars is exited.
func (s *BaseBSLParserListener) ExitSubVars(ctx *SubVarsContext) {}

// EnterSubVar is called when production subVar is entered.
func (s *BaseBSLParserListener) EnterSubVar(ctx *SubVarContext) {}

// ExitSubVar is called when production subVar is exited.
func (s *BaseBSLParserListener) ExitSubVar(ctx *SubVarContext) {}

// EnterSubVarsList is called when production subVarsList is entered.
func (s *BaseBSLParserListener) EnterSubVarsList(ctx *SubVarsListContext) {}

// ExitSubVarsList is called when production subVarsList is exited.
func (s *BaseBSLParserListener) ExitSubVarsList(ctx *SubVarsListContext) {}

// EnterSubVarDeclaration is called when production subVarDeclaration is entered.
func (s *BaseBSLParserListener) EnterSubVarDeclaration(ctx *SubVarDeclarationContext) {}

// ExitSubVarDeclaration is called when production subVarDeclaration is exited.
func (s *BaseBSLParserListener) ExitSubVarDeclaration(ctx *SubVarDeclarationContext) {}

// EnterSubName is called when production subName is entered.
func (s *BaseBSLParserListener) EnterSubName(ctx *SubNameContext) {}

// ExitSubName is called when production subName is exited.
func (s *BaseBSLParserListener) ExitSubName(ctx *SubNameContext) {}

// EnterSubs is called when production subs is entered.
func (s *BaseBSLParserListener) EnterSubs(ctx *SubsContext) {}

// ExitSubs is called when production subs is exited.
func (s *BaseBSLParserListener) ExitSubs(ctx *SubsContext) {}

// EnterSub is called when production sub is entered.
func (s *BaseBSLParserListener) EnterSub(ctx *SubContext) {}

// ExitSub is called when production sub is exited.
func (s *BaseBSLParserListener) ExitSub(ctx *SubContext) {}

// EnterProcedure is called when production procedure is entered.
func (s *BaseBSLParserListener) EnterProcedure(ctx *ProcedureContext) {}

// ExitProcedure is called when production procedure is exited.
func (s *BaseBSLParserListener) ExitProcedure(ctx *ProcedureContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseBSLParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseBSLParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterProcDeclaration is called when production procDeclaration is entered.
func (s *BaseBSLParserListener) EnterProcDeclaration(ctx *ProcDeclarationContext) {}

// ExitProcDeclaration is called when production procDeclaration is exited.
func (s *BaseBSLParserListener) ExitProcDeclaration(ctx *ProcDeclarationContext) {}

// EnterFuncDeclaration is called when production funcDeclaration is entered.
func (s *BaseBSLParserListener) EnterFuncDeclaration(ctx *FuncDeclarationContext) {}

// ExitFuncDeclaration is called when production funcDeclaration is exited.
func (s *BaseBSLParserListener) ExitFuncDeclaration(ctx *FuncDeclarationContext) {}

// EnterSubCodeBlock is called when production subCodeBlock is entered.
func (s *BaseBSLParserListener) EnterSubCodeBlock(ctx *SubCodeBlockContext) {}

// ExitSubCodeBlock is called when production subCodeBlock is exited.
func (s *BaseBSLParserListener) ExitSubCodeBlock(ctx *SubCodeBlockContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseBSLParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseBSLParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseBSLParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseBSLParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterRaiseStatement is called when production raiseStatement is entered.
func (s *BaseBSLParserListener) EnterRaiseStatement(ctx *RaiseStatementContext) {}

// ExitRaiseStatement is called when production raiseStatement is exited.
func (s *BaseBSLParserListener) ExitRaiseStatement(ctx *RaiseStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseBSLParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseBSLParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterIfBranch is called when production ifBranch is entered.
func (s *BaseBSLParserListener) EnterIfBranch(ctx *IfBranchContext) {}

// ExitIfBranch is called when production ifBranch is exited.
func (s *BaseBSLParserListener) ExitIfBranch(ctx *IfBranchContext) {}

// EnterElsifBranch is called when production elsifBranch is entered.
func (s *BaseBSLParserListener) EnterElsifBranch(ctx *ElsifBranchContext) {}

// ExitElsifBranch is called when production elsifBranch is exited.
func (s *BaseBSLParserListener) ExitElsifBranch(ctx *ElsifBranchContext) {}

// EnterElseBranch is called when production elseBranch is entered.
func (s *BaseBSLParserListener) EnterElseBranch(ctx *ElseBranchContext) {}

// ExitElseBranch is called when production elseBranch is exited.
func (s *BaseBSLParserListener) ExitElseBranch(ctx *ElseBranchContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseBSLParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseBSLParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseBSLParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseBSLParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForEachStatement is called when production forEachStatement is entered.
func (s *BaseBSLParserListener) EnterForEachStatement(ctx *ForEachStatementContext) {}

// ExitForEachStatement is called when production forEachStatement is exited.
func (s *BaseBSLParserListener) ExitForEachStatement(ctx *ForEachStatementContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseBSLParserListener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseBSLParserListener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseBSLParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseBSLParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterExecuteStatement is called when production executeStatement is entered.
func (s *BaseBSLParserListener) EnterExecuteStatement(ctx *ExecuteStatementContext) {}

// ExitExecuteStatement is called when production executeStatement is exited.
func (s *BaseBSLParserListener) ExitExecuteStatement(ctx *ExecuteStatementContext) {}

// EnterCallStatement is called when production callStatement is entered.
func (s *BaseBSLParserListener) EnterCallStatement(ctx *CallStatementContext) {}

// ExitCallStatement is called when production callStatement is exited.
func (s *BaseBSLParserListener) ExitCallStatement(ctx *CallStatementContext) {}

// EnterWaitStatement is called when production waitStatement is entered.
func (s *BaseBSLParserListener) EnterWaitStatement(ctx *WaitStatementContext) {}

// ExitWaitStatement is called when production waitStatement is exited.
func (s *BaseBSLParserListener) ExitWaitStatement(ctx *WaitStatementContext) {}

// EnterLabelName is called when production labelName is entered.
func (s *BaseBSLParserListener) EnterLabelName(ctx *LabelNameContext) {}

// ExitLabelName is called when production labelName is exited.
func (s *BaseBSLParserListener) ExitLabelName(ctx *LabelNameContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseBSLParserListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseBSLParserListener) ExitLabel(ctx *LabelContext) {}

// EnterGotoStatement is called when production gotoStatement is entered.
func (s *BaseBSLParserListener) EnterGotoStatement(ctx *GotoStatementContext) {}

// ExitGotoStatement is called when production gotoStatement is exited.
func (s *BaseBSLParserListener) ExitGotoStatement(ctx *GotoStatementContext) {}

// EnterTryCodeBlock is called when production tryCodeBlock is entered.
func (s *BaseBSLParserListener) EnterTryCodeBlock(ctx *TryCodeBlockContext) {}

// ExitTryCodeBlock is called when production tryCodeBlock is exited.
func (s *BaseBSLParserListener) ExitTryCodeBlock(ctx *TryCodeBlockContext) {}

// EnterExceptCodeBlock is called when production exceptCodeBlock is entered.
func (s *BaseBSLParserListener) EnterExceptCodeBlock(ctx *ExceptCodeBlockContext) {}

// ExitExceptCodeBlock is called when production exceptCodeBlock is exited.
func (s *BaseBSLParserListener) ExitExceptCodeBlock(ctx *ExceptCodeBlockContext) {}

// EnterEvent is called when production event is entered.
func (s *BaseBSLParserListener) EnterEvent(ctx *EventContext) {}

// ExitEvent is called when production event is exited.
func (s *BaseBSLParserListener) ExitEvent(ctx *EventContext) {}

// EnterHandler is called when production handler is entered.
func (s *BaseBSLParserListener) EnterHandler(ctx *HandlerContext) {}

// ExitHandler is called when production handler is exited.
func (s *BaseBSLParserListener) ExitHandler(ctx *HandlerContext) {}

// EnterAddHandlerStatement is called when production addHandlerStatement is entered.
func (s *BaseBSLParserListener) EnterAddHandlerStatement(ctx *AddHandlerStatementContext) {}

// ExitAddHandlerStatement is called when production addHandlerStatement is exited.
func (s *BaseBSLParserListener) ExitAddHandlerStatement(ctx *AddHandlerStatementContext) {}

// EnterRemoveHandlerStatement is called when production removeHandlerStatement is entered.
func (s *BaseBSLParserListener) EnterRemoveHandlerStatement(ctx *RemoveHandlerStatementContext) {}

// ExitRemoveHandlerStatement is called when production removeHandlerStatement is exited.
func (s *BaseBSLParserListener) ExitRemoveHandlerStatement(ctx *RemoveHandlerStatementContext) {}

// EnterTernaryOperator is called when production ternaryOperator is entered.
func (s *BaseBSLParserListener) EnterTernaryOperator(ctx *TernaryOperatorContext) {}

// ExitTernaryOperator is called when production ternaryOperator is exited.
func (s *BaseBSLParserListener) ExitTernaryOperator(ctx *TernaryOperatorContext) {}

// EnterWaitExpression is called when production waitExpression is entered.
func (s *BaseBSLParserListener) EnterWaitExpression(ctx *WaitExpressionContext) {}

// ExitWaitExpression is called when production waitExpression is exited.
func (s *BaseBSLParserListener) ExitWaitExpression(ctx *WaitExpressionContext) {}

// EnterFileCodeBlockBeforeSub is called when production fileCodeBlockBeforeSub is entered.
func (s *BaseBSLParserListener) EnterFileCodeBlockBeforeSub(ctx *FileCodeBlockBeforeSubContext) {}

// ExitFileCodeBlockBeforeSub is called when production fileCodeBlockBeforeSub is exited.
func (s *BaseBSLParserListener) ExitFileCodeBlockBeforeSub(ctx *FileCodeBlockBeforeSubContext) {}

// EnterFileCodeBlock is called when production fileCodeBlock is entered.
func (s *BaseBSLParserListener) EnterFileCodeBlock(ctx *FileCodeBlockContext) {}

// ExitFileCodeBlock is called when production fileCodeBlock is exited.
func (s *BaseBSLParserListener) ExitFileCodeBlock(ctx *FileCodeBlockContext) {}

// EnterCodeBlock is called when production codeBlock is entered.
func (s *BaseBSLParserListener) EnterCodeBlock(ctx *CodeBlockContext) {}

// ExitCodeBlock is called when production codeBlock is exited.
func (s *BaseBSLParserListener) ExitCodeBlock(ctx *CodeBlockContext) {}

// EnterNumeric is called when production numeric is entered.
func (s *BaseBSLParserListener) EnterNumeric(ctx *NumericContext) {}

// ExitNumeric is called when production numeric is exited.
func (s *BaseBSLParserListener) ExitNumeric(ctx *NumericContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseBSLParserListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseBSLParserListener) ExitParamList(ctx *ParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BaseBSLParserListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseBSLParserListener) ExitParam(ctx *ParamContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseBSLParserListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseBSLParserListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterConstValue is called when production constValue is entered.
func (s *BaseBSLParserListener) EnterConstValue(ctx *ConstValueContext) {}

// ExitConstValue is called when production constValue is exited.
func (s *BaseBSLParserListener) ExitConstValue(ctx *ConstValueContext) {}

// EnterMultilineString is called when production multilineString is entered.
func (s *BaseBSLParserListener) EnterMultilineString(ctx *MultilineStringContext) {}

// ExitMultilineString is called when production multilineString is exited.
func (s *BaseBSLParserListener) ExitMultilineString(ctx *MultilineStringContext) {}

// EnterString is called when production string is entered.
func (s *BaseBSLParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseBSLParserListener) ExitString(ctx *StringContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseBSLParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseBSLParserListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseBSLParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseBSLParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterCallParamList is called when production callParamList is entered.
func (s *BaseBSLParserListener) EnterCallParamList(ctx *CallParamListContext) {}

// ExitCallParamList is called when production callParamList is exited.
func (s *BaseBSLParserListener) ExitCallParamList(ctx *CallParamListContext) {}

// EnterCallParam is called when production callParam is entered.
func (s *BaseBSLParserListener) EnterCallParam(ctx *CallParamContext) {}

// ExitCallParam is called when production callParam is exited.
func (s *BaseBSLParserListener) ExitCallParam(ctx *CallParamContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBSLParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBSLParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaseBSLParserListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaseBSLParserListener) ExitOperation(ctx *OperationContext) {}

// EnterCompareOperation is called when production compareOperation is entered.
func (s *BaseBSLParserListener) EnterCompareOperation(ctx *CompareOperationContext) {}

// ExitCompareOperation is called when production compareOperation is exited.
func (s *BaseBSLParserListener) ExitCompareOperation(ctx *CompareOperationContext) {}

// EnterBoolOperation is called when production boolOperation is entered.
func (s *BaseBSLParserListener) EnterBoolOperation(ctx *BoolOperationContext) {}

// ExitBoolOperation is called when production boolOperation is exited.
func (s *BaseBSLParserListener) ExitBoolOperation(ctx *BoolOperationContext) {}

// EnterUnaryModifier is called when production unaryModifier is entered.
func (s *BaseBSLParserListener) EnterUnaryModifier(ctx *UnaryModifierContext) {}

// ExitUnaryModifier is called when production unaryModifier is exited.
func (s *BaseBSLParserListener) ExitUnaryModifier(ctx *UnaryModifierContext) {}

// EnterMember is called when production member is entered.
func (s *BaseBSLParserListener) EnterMember(ctx *MemberContext) {}

// ExitMember is called when production member is exited.
func (s *BaseBSLParserListener) ExitMember(ctx *MemberContext) {}

// EnterNewwExpression is called when production newwExpression is entered.
func (s *BaseBSLParserListener) EnterNewwExpression(ctx *NewwExpressionContext) {}

// ExitNewwExpression is called when production newwExpression is exited.
func (s *BaseBSLParserListener) ExitNewwExpression(ctx *NewwExpressionContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseBSLParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseBSLParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BaseBSLParserListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseBSLParserListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterGlobalMethodCall is called when production globalMethodCall is entered.
func (s *BaseBSLParserListener) EnterGlobalMethodCall(ctx *GlobalMethodCallContext) {}

// ExitGlobalMethodCall is called when production globalMethodCall is exited.
func (s *BaseBSLParserListener) ExitGlobalMethodCall(ctx *GlobalMethodCallContext) {}

// EnterMethodName is called when production methodName is entered.
func (s *BaseBSLParserListener) EnterMethodName(ctx *MethodNameContext) {}

// ExitMethodName is called when production methodName is exited.
func (s *BaseBSLParserListener) ExitMethodName(ctx *MethodNameContext) {}

// EnterComplexIdentifier is called when production complexIdentifier is entered.
func (s *BaseBSLParserListener) EnterComplexIdentifier(ctx *ComplexIdentifierContext) {}

// ExitComplexIdentifier is called when production complexIdentifier is exited.
func (s *BaseBSLParserListener) ExitComplexIdentifier(ctx *ComplexIdentifierContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseBSLParserListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseBSLParserListener) ExitModifier(ctx *ModifierContext) {}

// EnterAcceptor is called when production acceptor is entered.
func (s *BaseBSLParserListener) EnterAcceptor(ctx *AcceptorContext) {}

// ExitAcceptor is called when production acceptor is exited.
func (s *BaseBSLParserListener) ExitAcceptor(ctx *AcceptorContext) {}

// EnterLValue is called when production lValue is entered.
func (s *BaseBSLParserListener) EnterLValue(ctx *LValueContext) {}

// ExitLValue is called when production lValue is exited.
func (s *BaseBSLParserListener) ExitLValue(ctx *LValueContext) {}

// EnterAccessCall is called when production accessCall is entered.
func (s *BaseBSLParserListener) EnterAccessCall(ctx *AccessCallContext) {}

// ExitAccessCall is called when production accessCall is exited.
func (s *BaseBSLParserListener) ExitAccessCall(ctx *AccessCallContext) {}

// EnterAccessIndex is called when production accessIndex is entered.
func (s *BaseBSLParserListener) EnterAccessIndex(ctx *AccessIndexContext) {}

// ExitAccessIndex is called when production accessIndex is exited.
func (s *BaseBSLParserListener) ExitAccessIndex(ctx *AccessIndexContext) {}

// EnterAccessProperty is called when production accessProperty is entered.
func (s *BaseBSLParserListener) EnterAccessProperty(ctx *AccessPropertyContext) {}

// ExitAccessProperty is called when production accessProperty is exited.
func (s *BaseBSLParserListener) ExitAccessProperty(ctx *AccessPropertyContext) {}

// EnterDoCall is called when production doCall is entered.
func (s *BaseBSLParserListener) EnterDoCall(ctx *DoCallContext) {}

// ExitDoCall is called when production doCall is exited.
func (s *BaseBSLParserListener) ExitDoCall(ctx *DoCallContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BaseBSLParserListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BaseBSLParserListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}
