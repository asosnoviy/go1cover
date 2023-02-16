// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // BSLParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BSLParserListener is a complete listener for a parse tree produced by BSLParser.
type BSLParserListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterShebang is called when entering the shebang production.
	EnterShebang(c *ShebangContext)

	// EnterUsedLib is called when entering the usedLib production.
	EnterUsedLib(c *UsedLibContext)

	// EnterUse is called when entering the use production.
	EnterUse(c *UseContext)

	// EnterRegionStart is called when entering the regionStart production.
	EnterRegionStart(c *RegionStartContext)

	// EnterRegionEnd is called when entering the regionEnd production.
	EnterRegionEnd(c *RegionEndContext)

	// EnterRegionName is called when entering the regionName production.
	EnterRegionName(c *RegionNameContext)

	// EnterPreproc_if is called when entering the preproc_if production.
	EnterPreproc_if(c *Preproc_ifContext)

	// EnterPreproc_elsif is called when entering the preproc_elsif production.
	EnterPreproc_elsif(c *Preproc_elsifContext)

	// EnterPreproc_else is called when entering the preproc_else production.
	EnterPreproc_else(c *Preproc_elseContext)

	// EnterPreproc_endif is called when entering the preproc_endif production.
	EnterPreproc_endif(c *Preproc_endifContext)

	// EnterPreproc_expression is called when entering the preproc_expression production.
	EnterPreproc_expression(c *Preproc_expressionContext)

	// EnterPreproc_logicalOperand is called when entering the preproc_logicalOperand production.
	EnterPreproc_logicalOperand(c *Preproc_logicalOperandContext)

	// EnterPreproc_logicalExpression is called when entering the preproc_logicalExpression production.
	EnterPreproc_logicalExpression(c *Preproc_logicalExpressionContext)

	// EnterPreproc_symbol is called when entering the preproc_symbol production.
	EnterPreproc_symbol(c *Preproc_symbolContext)

	// EnterPreproc_unknownSymbol is called when entering the preproc_unknownSymbol production.
	EnterPreproc_unknownSymbol(c *Preproc_unknownSymbolContext)

	// EnterPreproc_boolOperation is called when entering the preproc_boolOperation production.
	EnterPreproc_boolOperation(c *Preproc_boolOperationContext)

	// EnterPreprocessor is called when entering the preprocessor production.
	EnterPreprocessor(c *PreprocessorContext)

	// EnterCompilerDirectiveSymbol is called when entering the compilerDirectiveSymbol production.
	EnterCompilerDirectiveSymbol(c *CompilerDirectiveSymbolContext)

	// EnterCompilerDirective is called when entering the compilerDirective production.
	EnterCompilerDirective(c *CompilerDirectiveContext)

	// EnterAnnotationName is called when entering the annotationName production.
	EnterAnnotationName(c *AnnotationNameContext)

	// EnterAnnotationParamName is called when entering the annotationParamName production.
	EnterAnnotationParamName(c *AnnotationParamNameContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterAnnotationParams is called when entering the annotationParams production.
	EnterAnnotationParams(c *AnnotationParamsContext)

	// EnterAnnotationParam is called when entering the annotationParam production.
	EnterAnnotationParam(c *AnnotationParamContext)

	// EnterVar_name is called when entering the var_name production.
	EnterVar_name(c *Var_nameContext)

	// EnterModuleVars is called when entering the moduleVars production.
	EnterModuleVars(c *ModuleVarsContext)

	// EnterModuleVar is called when entering the moduleVar production.
	EnterModuleVar(c *ModuleVarContext)

	// EnterModuleVarsList is called when entering the moduleVarsList production.
	EnterModuleVarsList(c *ModuleVarsListContext)

	// EnterModuleVarDeclaration is called when entering the moduleVarDeclaration production.
	EnterModuleVarDeclaration(c *ModuleVarDeclarationContext)

	// EnterSubVars is called when entering the subVars production.
	EnterSubVars(c *SubVarsContext)

	// EnterSubVar is called when entering the subVar production.
	EnterSubVar(c *SubVarContext)

	// EnterSubVarsList is called when entering the subVarsList production.
	EnterSubVarsList(c *SubVarsListContext)

	// EnterSubVarDeclaration is called when entering the subVarDeclaration production.
	EnterSubVarDeclaration(c *SubVarDeclarationContext)

	// EnterSubName is called when entering the subName production.
	EnterSubName(c *SubNameContext)

	// EnterSubs is called when entering the subs production.
	EnterSubs(c *SubsContext)

	// EnterSub is called when entering the sub production.
	EnterSub(c *SubContext)

	// EnterProcedure is called when entering the procedure production.
	EnterProcedure(c *ProcedureContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterProcDeclaration is called when entering the procDeclaration production.
	EnterProcDeclaration(c *ProcDeclarationContext)

	// EnterFuncDeclaration is called when entering the funcDeclaration production.
	EnterFuncDeclaration(c *FuncDeclarationContext)

	// EnterSubCodeBlock is called when entering the subCodeBlock production.
	EnterSubCodeBlock(c *SubCodeBlockContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterRaiseStatement is called when entering the raiseStatement production.
	EnterRaiseStatement(c *RaiseStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterIfBranch is called when entering the ifBranch production.
	EnterIfBranch(c *IfBranchContext)

	// EnterElsifBranch is called when entering the elsifBranch production.
	EnterElsifBranch(c *ElsifBranchContext)

	// EnterElseBranch is called when entering the elseBranch production.
	EnterElseBranch(c *ElseBranchContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForEachStatement is called when entering the forEachStatement production.
	EnterForEachStatement(c *ForEachStatementContext)

	// EnterTryStatement is called when entering the tryStatement production.
	EnterTryStatement(c *TryStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterExecuteStatement is called when entering the executeStatement production.
	EnterExecuteStatement(c *ExecuteStatementContext)

	// EnterCallStatement is called when entering the callStatement production.
	EnterCallStatement(c *CallStatementContext)

	// EnterWaitStatement is called when entering the waitStatement production.
	EnterWaitStatement(c *WaitStatementContext)

	// EnterLabelName is called when entering the labelName production.
	EnterLabelName(c *LabelNameContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterGotoStatement is called when entering the gotoStatement production.
	EnterGotoStatement(c *GotoStatementContext)

	// EnterTryCodeBlock is called when entering the tryCodeBlock production.
	EnterTryCodeBlock(c *TryCodeBlockContext)

	// EnterExceptCodeBlock is called when entering the exceptCodeBlock production.
	EnterExceptCodeBlock(c *ExceptCodeBlockContext)

	// EnterEvent is called when entering the event production.
	EnterEvent(c *EventContext)

	// EnterHandler is called when entering the handler production.
	EnterHandler(c *HandlerContext)

	// EnterAddHandlerStatement is called when entering the addHandlerStatement production.
	EnterAddHandlerStatement(c *AddHandlerStatementContext)

	// EnterRemoveHandlerStatement is called when entering the removeHandlerStatement production.
	EnterRemoveHandlerStatement(c *RemoveHandlerStatementContext)

	// EnterTernaryOperator is called when entering the ternaryOperator production.
	EnterTernaryOperator(c *TernaryOperatorContext)

	// EnterWaitExpression is called when entering the waitExpression production.
	EnterWaitExpression(c *WaitExpressionContext)

	// EnterFileCodeBlockBeforeSub is called when entering the fileCodeBlockBeforeSub production.
	EnterFileCodeBlockBeforeSub(c *FileCodeBlockBeforeSubContext)

	// EnterFileCodeBlock is called when entering the fileCodeBlock production.
	EnterFileCodeBlock(c *FileCodeBlockContext)

	// EnterCodeBlock is called when entering the codeBlock production.
	EnterCodeBlock(c *CodeBlockContext)

	// EnterNumeric is called when entering the numeric production.
	EnterNumeric(c *NumericContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterDefaultValue is called when entering the defaultValue production.
	EnterDefaultValue(c *DefaultValueContext)

	// EnterConstValue is called when entering the constValue production.
	EnterConstValue(c *ConstValueContext)

	// EnterMultilineString is called when entering the multilineString production.
	EnterMultilineString(c *MultilineStringContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterCallParamList is called when entering the callParamList production.
	EnterCallParamList(c *CallParamListContext)

	// EnterCallParam is called when entering the callParam production.
	EnterCallParam(c *CallParamContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterCompareOperation is called when entering the compareOperation production.
	EnterCompareOperation(c *CompareOperationContext)

	// EnterBoolOperation is called when entering the boolOperation production.
	EnterBoolOperation(c *BoolOperationContext)

	// EnterUnaryModifier is called when entering the unaryModifier production.
	EnterUnaryModifier(c *UnaryModifierContext)

	// EnterMember is called when entering the member production.
	EnterMember(c *MemberContext)

	// EnterNewwExpression is called when entering the newwExpression production.
	EnterNewwExpression(c *NewwExpressionContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *MethodCallContext)

	// EnterGlobalMethodCall is called when entering the globalMethodCall production.
	EnterGlobalMethodCall(c *GlobalMethodCallContext)

	// EnterMethodName is called when entering the methodName production.
	EnterMethodName(c *MethodNameContext)

	// EnterComplexIdentifier is called when entering the complexIdentifier production.
	EnterComplexIdentifier(c *ComplexIdentifierContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterAcceptor is called when entering the acceptor production.
	EnterAcceptor(c *AcceptorContext)

	// EnterLValue is called when entering the lValue production.
	EnterLValue(c *LValueContext)

	// EnterAccessCall is called when entering the accessCall production.
	EnterAccessCall(c *AccessCallContext)

	// EnterAccessIndex is called when entering the accessIndex production.
	EnterAccessIndex(c *AccessIndexContext)

	// EnterAccessProperty is called when entering the accessProperty production.
	EnterAccessProperty(c *AccessPropertyContext)

	// EnterDoCall is called when entering the doCall production.
	EnterDoCall(c *DoCallContext)

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitShebang is called when exiting the shebang production.
	ExitShebang(c *ShebangContext)

	// ExitUsedLib is called when exiting the usedLib production.
	ExitUsedLib(c *UsedLibContext)

	// ExitUse is called when exiting the use production.
	ExitUse(c *UseContext)

	// ExitRegionStart is called when exiting the regionStart production.
	ExitRegionStart(c *RegionStartContext)

	// ExitRegionEnd is called when exiting the regionEnd production.
	ExitRegionEnd(c *RegionEndContext)

	// ExitRegionName is called when exiting the regionName production.
	ExitRegionName(c *RegionNameContext)

	// ExitPreproc_if is called when exiting the preproc_if production.
	ExitPreproc_if(c *Preproc_ifContext)

	// ExitPreproc_elsif is called when exiting the preproc_elsif production.
	ExitPreproc_elsif(c *Preproc_elsifContext)

	// ExitPreproc_else is called when exiting the preproc_else production.
	ExitPreproc_else(c *Preproc_elseContext)

	// ExitPreproc_endif is called when exiting the preproc_endif production.
	ExitPreproc_endif(c *Preproc_endifContext)

	// ExitPreproc_expression is called when exiting the preproc_expression production.
	ExitPreproc_expression(c *Preproc_expressionContext)

	// ExitPreproc_logicalOperand is called when exiting the preproc_logicalOperand production.
	ExitPreproc_logicalOperand(c *Preproc_logicalOperandContext)

	// ExitPreproc_logicalExpression is called when exiting the preproc_logicalExpression production.
	ExitPreproc_logicalExpression(c *Preproc_logicalExpressionContext)

	// ExitPreproc_symbol is called when exiting the preproc_symbol production.
	ExitPreproc_symbol(c *Preproc_symbolContext)

	// ExitPreproc_unknownSymbol is called when exiting the preproc_unknownSymbol production.
	ExitPreproc_unknownSymbol(c *Preproc_unknownSymbolContext)

	// ExitPreproc_boolOperation is called when exiting the preproc_boolOperation production.
	ExitPreproc_boolOperation(c *Preproc_boolOperationContext)

	// ExitPreprocessor is called when exiting the preprocessor production.
	ExitPreprocessor(c *PreprocessorContext)

	// ExitCompilerDirectiveSymbol is called when exiting the compilerDirectiveSymbol production.
	ExitCompilerDirectiveSymbol(c *CompilerDirectiveSymbolContext)

	// ExitCompilerDirective is called when exiting the compilerDirective production.
	ExitCompilerDirective(c *CompilerDirectiveContext)

	// ExitAnnotationName is called when exiting the annotationName production.
	ExitAnnotationName(c *AnnotationNameContext)

	// ExitAnnotationParamName is called when exiting the annotationParamName production.
	ExitAnnotationParamName(c *AnnotationParamNameContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitAnnotationParams is called when exiting the annotationParams production.
	ExitAnnotationParams(c *AnnotationParamsContext)

	// ExitAnnotationParam is called when exiting the annotationParam production.
	ExitAnnotationParam(c *AnnotationParamContext)

	// ExitVar_name is called when exiting the var_name production.
	ExitVar_name(c *Var_nameContext)

	// ExitModuleVars is called when exiting the moduleVars production.
	ExitModuleVars(c *ModuleVarsContext)

	// ExitModuleVar is called when exiting the moduleVar production.
	ExitModuleVar(c *ModuleVarContext)

	// ExitModuleVarsList is called when exiting the moduleVarsList production.
	ExitModuleVarsList(c *ModuleVarsListContext)

	// ExitModuleVarDeclaration is called when exiting the moduleVarDeclaration production.
	ExitModuleVarDeclaration(c *ModuleVarDeclarationContext)

	// ExitSubVars is called when exiting the subVars production.
	ExitSubVars(c *SubVarsContext)

	// ExitSubVar is called when exiting the subVar production.
	ExitSubVar(c *SubVarContext)

	// ExitSubVarsList is called when exiting the subVarsList production.
	ExitSubVarsList(c *SubVarsListContext)

	// ExitSubVarDeclaration is called when exiting the subVarDeclaration production.
	ExitSubVarDeclaration(c *SubVarDeclarationContext)

	// ExitSubName is called when exiting the subName production.
	ExitSubName(c *SubNameContext)

	// ExitSubs is called when exiting the subs production.
	ExitSubs(c *SubsContext)

	// ExitSub is called when exiting the sub production.
	ExitSub(c *SubContext)

	// ExitProcedure is called when exiting the procedure production.
	ExitProcedure(c *ProcedureContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitProcDeclaration is called when exiting the procDeclaration production.
	ExitProcDeclaration(c *ProcDeclarationContext)

	// ExitFuncDeclaration is called when exiting the funcDeclaration production.
	ExitFuncDeclaration(c *FuncDeclarationContext)

	// ExitSubCodeBlock is called when exiting the subCodeBlock production.
	ExitSubCodeBlock(c *SubCodeBlockContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitRaiseStatement is called when exiting the raiseStatement production.
	ExitRaiseStatement(c *RaiseStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitIfBranch is called when exiting the ifBranch production.
	ExitIfBranch(c *IfBranchContext)

	// ExitElsifBranch is called when exiting the elsifBranch production.
	ExitElsifBranch(c *ElsifBranchContext)

	// ExitElseBranch is called when exiting the elseBranch production.
	ExitElseBranch(c *ElseBranchContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForEachStatement is called when exiting the forEachStatement production.
	ExitForEachStatement(c *ForEachStatementContext)

	// ExitTryStatement is called when exiting the tryStatement production.
	ExitTryStatement(c *TryStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitExecuteStatement is called when exiting the executeStatement production.
	ExitExecuteStatement(c *ExecuteStatementContext)

	// ExitCallStatement is called when exiting the callStatement production.
	ExitCallStatement(c *CallStatementContext)

	// ExitWaitStatement is called when exiting the waitStatement production.
	ExitWaitStatement(c *WaitStatementContext)

	// ExitLabelName is called when exiting the labelName production.
	ExitLabelName(c *LabelNameContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitGotoStatement is called when exiting the gotoStatement production.
	ExitGotoStatement(c *GotoStatementContext)

	// ExitTryCodeBlock is called when exiting the tryCodeBlock production.
	ExitTryCodeBlock(c *TryCodeBlockContext)

	// ExitExceptCodeBlock is called when exiting the exceptCodeBlock production.
	ExitExceptCodeBlock(c *ExceptCodeBlockContext)

	// ExitEvent is called when exiting the event production.
	ExitEvent(c *EventContext)

	// ExitHandler is called when exiting the handler production.
	ExitHandler(c *HandlerContext)

	// ExitAddHandlerStatement is called when exiting the addHandlerStatement production.
	ExitAddHandlerStatement(c *AddHandlerStatementContext)

	// ExitRemoveHandlerStatement is called when exiting the removeHandlerStatement production.
	ExitRemoveHandlerStatement(c *RemoveHandlerStatementContext)

	// ExitTernaryOperator is called when exiting the ternaryOperator production.
	ExitTernaryOperator(c *TernaryOperatorContext)

	// ExitWaitExpression is called when exiting the waitExpression production.
	ExitWaitExpression(c *WaitExpressionContext)

	// ExitFileCodeBlockBeforeSub is called when exiting the fileCodeBlockBeforeSub production.
	ExitFileCodeBlockBeforeSub(c *FileCodeBlockBeforeSubContext)

	// ExitFileCodeBlock is called when exiting the fileCodeBlock production.
	ExitFileCodeBlock(c *FileCodeBlockContext)

	// ExitCodeBlock is called when exiting the codeBlock production.
	ExitCodeBlock(c *CodeBlockContext)

	// ExitNumeric is called when exiting the numeric production.
	ExitNumeric(c *NumericContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitDefaultValue is called when exiting the defaultValue production.
	ExitDefaultValue(c *DefaultValueContext)

	// ExitConstValue is called when exiting the constValue production.
	ExitConstValue(c *ConstValueContext)

	// ExitMultilineString is called when exiting the multilineString production.
	ExitMultilineString(c *MultilineStringContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitCallParamList is called when exiting the callParamList production.
	ExitCallParamList(c *CallParamListContext)

	// ExitCallParam is called when exiting the callParam production.
	ExitCallParam(c *CallParamContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitCompareOperation is called when exiting the compareOperation production.
	ExitCompareOperation(c *CompareOperationContext)

	// ExitBoolOperation is called when exiting the boolOperation production.
	ExitBoolOperation(c *BoolOperationContext)

	// ExitUnaryModifier is called when exiting the unaryModifier production.
	ExitUnaryModifier(c *UnaryModifierContext)

	// ExitMember is called when exiting the member production.
	ExitMember(c *MemberContext)

	// ExitNewwExpression is called when exiting the newwExpression production.
	ExitNewwExpression(c *NewwExpressionContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *MethodCallContext)

	// ExitGlobalMethodCall is called when exiting the globalMethodCall production.
	ExitGlobalMethodCall(c *GlobalMethodCallContext)

	// ExitMethodName is called when exiting the methodName production.
	ExitMethodName(c *MethodNameContext)

	// ExitComplexIdentifier is called when exiting the complexIdentifier production.
	ExitComplexIdentifier(c *ComplexIdentifierContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitAcceptor is called when exiting the acceptor production.
	ExitAcceptor(c *AcceptorContext)

	// ExitLValue is called when exiting the lValue production.
	ExitLValue(c *LValueContext)

	// ExitAccessCall is called when exiting the accessCall production.
	ExitAccessCall(c *AccessCallContext)

	// ExitAccessIndex is called when exiting the accessIndex production.
	ExitAccessIndex(c *AccessIndexContext)

	// ExitAccessProperty is called when exiting the accessProperty production.
	ExitAccessProperty(c *AccessPropertyContext)

	// ExitDoCall is called when exiting the doCall production.
	ExitDoCall(c *DoCallContext)

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)
}
