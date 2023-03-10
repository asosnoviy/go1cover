/**
 * This file is a part of BSL Parser.
 *
 * Copyright © 2018-2022
 * Alexey Sosnoviy <labotamy@gmail.com>, Nikita Fedkin <nixel2007@gmail.com>, Sergey Batanov <sergey.batanov@dmpas.ru>
 *
 * SPDX-License-Identifier: LGPL-3.0-or-later
 *
 * BSL Parser is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 3.0 of the License, or (at your option) any later version.
 *
 * BSL Parser is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with BSL Parser.
 */
lexer grammar BSLLexer;
options { caseInsensitive = true; }

channels {
    // для хранения удаленного блока
    PREPROC_DELETE_CHANNEL
}

// commons
fragment DIGIT: [0-9];
LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);
WHITE_SPACE: [ \t\f\r\n]+ -> channel(HIDDEN);

// separators
DOT: '.' -> pushMode(DOT_MODE);
LBRACK: '[';
RBRACK: ']';
LPAREN: '(';
RPAREN: ')';
COLON: ':';
SEMICOLON: ';';
COMMA: ',';
ASSIGN: '=';
PLUS: '+';
MINUS: '-';
LESS_OR_EQUAL: '<=';
NOT_EQUAL: '<>';
LESS: '<';
GREATER_OR_EQUAL: '>=';
GREATER: '>';
MUL: '*';
QUOTIENT: '/';
MODULO: '%';
QUESTION: '?';
AMPERSAND: '&' -> pushMode(ANNOTATION_MODE);
PREPROC_DELETE
    : '#' [ \t]*
    (
          '\u0423\u0414\u0410\u041b\u0415\u041d\u0418\u0415'
        | 'DELETE'
    )
    -> pushMode(PREPROC_DELETE_MODE), channel(PREPROC_DELETE_CHANNEL)
    ;
PREPROC_INSERT
    : '#' [ \t]*
    (
          '\u0412\u0421\u0422\u0410\u0412\u041a\u0410'
        | 'INSERT'
    )
    -> channel(HIDDEN)
    ;
PREPROC_ENDINSERT
    : '#' [ \t]*
    (
          '\u041a\u041e\u041d\u0415\u0426\u0412\u0421\u0422\u0410\u0412\u041a\u0418'
        | 'ENDINSERT'
    )
    -> channel(HIDDEN)
    ;
HASH: '#' -> pushMode(PREPROCESSOR_MODE);

fragment SQUOTE: '\'';
BAR: '|';
TILDA: '~' -> pushMode(LABEL_MODE);

// fragment А: '\u0410' | '\u0430';
// fragment Б: '\u0411' | '\u0431';
// fragment В: '\u0412' | '\u0432';
// fragment Г: '\u0413' | '\u0433';
// fragment Д: '\u0414' | '\u0434';
// fragment Е: '\u0401' | '\u0451';
// fragment Е: '\u0415' | '\u0435';
// fragment Ж: '\u0416' | '\u0436';
// fragment З: '\u0417' | '\u0437';
// fragment И: '\u0418' | '\u0438';
// fragment Й: '\u0419' | '\u0439';
// fragment К: '\u041a' | '\u043a';
// fragment Л: '\u041b' | '\u043b';
// fragment М: '\u041c' | '\u043c';
// fragment Н: '\u041d' | '\u043d';
// fragment О: '\u041e' | '\u043e';
// fragment П: '\u041f' | '\u043f';
// fragment Р: '\u0420' | '\u0440';
// fragment С: '\u0421' | '\u0441';
// fragment Т: '\u0422' | '\u0442';
// fragment У: '\u0423' | '\u0443';
// fragment Ф:'\u0424' | '\u0444';
// fragment Х:'\u0425' | '\u0445';
// fragment Ц: '\u0426' | '\u0446';
// fragment Ч: '\u0427' | '\u0447';
// fragment Ш: '\u0428' | '\u0448';
// fragment Щ: '\u0429' | '\u0449';
// fragment Ъ:'\u042a' | '\u044a';
// fragment Ы: '\u042b' | '\u044b';
// fragment Ь: '\u042c' | '\u044c';
// fragment Э:'\u042d' | '\u044d';
// fragment Ю: '\u042e' | '\u044e';
// fragment Я: '\u042f' | '\u044f';
// fragment A: 'A' | 'a';
// fragment B: 'B' | 'b';
// fragment C: 'C' | 'c';
// fragment D: 'D' | 'd';
// fragment I: 'I' | 'i';
// fragment E: 'E' | 'e';
// fragment F: 'F' | 'f';
// fragment G: 'G' | 'g';
// fragment U: 'U' | 'u';
// fragment K: 'K' | 'k';
// fragment L: 'L' | 'l';
// fragment M: 'M' | 'm';
// fragment N: 'N' | 'n';
// fragment O: 'O' | 'o';
// fragment P: 'P' | 'p';
// fragment R: 'R' | 'r';
// fragment S: 'S' | 's';
// fragment T: 'T' | 't';
// fragment V: 'V' | 'v';
// fragment H: 'H' | 'h';
// fragment W: 'W' | 'w';
// fragment X: 'X' | 'x';
// fragment Y: 'Y' | 'y';

// literals
TRUE
    :
     '\u0418\u0421\u0422\u0418\u041d\u0410'
    | 'TRUE'
    ;
FALSE
    :
    '\u041b\u041e\u0416\u042c'
    | 'FALSE'
    ;
UNDEFINED
    :
     '\u041d\u0415\u041e\u041f\u0420\u0415\u0414\u0415\u041b\u0415\u041d\u041e'
    | 'UNDEFINED'
    ;
NULL
    :
    'NULL'
    ;
DECIMAL: DIGIT+;
DATETIME: SQUOTE(~['\n\r])*SQUOTE?; // TODO: Честная регулярка

FLOAT : DIGIT+ '.' DIGIT*;
STRING: '"' (~[\r\n"] | '""')* '"';
STRINGSTART: '"' (~["\n\r]| '""')*;
STRINGTAIL: BAR (~["\n\r] | '""')* '"';
STRINGPART: BAR (~[\r\n"] | '""')*;

// keywords
PROCEDURE_KEYWORD
    :
     '\u041f\u0440\u043e\u0446\u0435\u0434\u0443\u0440\u0430'
    | 'PROCEDURE'
    ;
FUNCTION_KEYWORD
    :
    '\u0424\u0423\u041d\u041a\u0426\u0418\u042f'
    |'FUNCTION'
    ;
ENDPROCEDURE_KEYWORD
    :'\u041a\u043e\u043d\u0435\u0446\u041f\u0440\u043e\u0446\u0435\u0434\u0443\u0440\u044b'
    | 'ENDPROCEDURE'
    ;
ENDFUNCTION_KEYWORD
    :'\u041a\u041e\u041d\u0415\u0426\u0424\u0423\u041d\u041a\u0426\u0418\u0418'
    |'ENDFUNCTION'
    ;
EXPORT_KEYWORD
    :
    '\u042d\u041a\u0421\u041f\u041e\u0420\u0422'
    |'EXPORT'
    ;
VAL_KEYWORD
    :
    '\u0417\u041d\u0410\u0427'
    |'VAL'
    ;
ENDIF_KEYWORD
    :
    '\u041a\u041e\u041d\u0415\u0426\u0415\u0421\u041b\u0418'
    |'ENDIF'
    ;
ENDDO_KEYWORD
    :
    '\u041a\u041e\u041d\u0415\u0426\u0426\u0418\u041a\u041b\u0410'
    |'ENDDO'
    ;
IF_KEYWORD
    :
    '\u0415\u0421\u041b\u0418'
    |'IF'
    ;
ELSIF_KEYWORD
    :
    '\u0418\u041d\u0410\u0427\u0415\u0415\u0421\u041b\u0418'
    |'ELSIF'
    ;
ELSE_KEYWORD
    :
    '\u0418\u041d\u0410\u0427\u0415'
    |'ELSE'
    ;
THEN_KEYWORD
    :
    '\u0422\u041e\u0413\u0414\u0410'
    |'THEN'
    ;
WHILE_KEYWORD
    :
    '\u041f\u041e\u041a\u0410'
    |'WHILE'
    ;
DO_KEYWORD
    :
    '\u0426\u0418\u041a\u041b'
    |'DO'
    ;
FOR_KEYWORD
    :
    '\u0414\u041b\u042f'
    |'FOR'
    ;
TO_KEYWORD
    :
    '\u041f\u041e'
    |'TO'
    ;
EACH_KEYWORD
    :
    '\u041a\u0410\u0416\u0414\u041e\u0413\u041e'
    |'EACH'
    ;
IN_KEYWORD
    :
    '\u0418\u0417'
    |'IN'
    ;
TRY_KEYWORD
    :
    '\u041f\u041e\u041f\u042b\u0422\u041a\u0410'
    |'TRY'
    ;
EXCEPT_KEYWORD
    :
    '\u0418\u0421\u041a\u041b\u042e\u0427\u0415\u041d\u0418\u0415'
    |'EXCEPT'
    ;
ENDTRY_KEYWORD
    :
    '\u041a\u041e\u041d\u0415\u0426\u041f\u041e\u041f\u042b\u0422\u041a\u0418'
    |'ENDTRY'
    ;
RETURN_KEYWORD
    :
    '\u0412\u041e\u0417\u0412\u0420\u0410\u0422'
    |'RETURN'
    ;
CONTINUE_KEYWORD
    :
    '\u041f\u0420\u041e\u0414\u041e\u041b\u0416\u0418\u0422\u042c'
    |'CONTINUE'
    ;
RAISE_KEYWORD
    :
    '\u0412\u042b\u0417\u0412\u0410\u0422\u042c\u0418\u0421\u041a\u041b\u042e\u0427\u0415\u041d\u0418\u0415'
    |'RAISE'
    ;
VAR_KEYWORD
    :
    '\u041f\u0415\u0420\u0415\u041c'
    |'VAR'
    ;
NOT_KEYWORD
    :
    '\u041d\u0415'
    |'NOT'
    ;
OR_KEYWORD
    :
    '\u0418\u041b\u0418'
    |'OR'
    ;
AND_KEYWORD
    :
    '\u0418'
    |'AND'
    ;
NEW_KEYWORD
    :
    '\u041d\u041e\u0412\u042b\u0419'
    |'NEW'
    ;
GOTO_KEYWORD
    :
    '\u041f\u0415\u0420\u0415\u0419\u0422\u0418'
    |'GOTO'
    ;
BREAK_KEYWORD
    :
    '\u041f\u0420\u0415\u0420\u0412\u0410\u0422\u042c'
    |'BREAK'
    ;
EXECUTE_KEYWORD
    :
    '\u0412\u042b\u041f\u041e\u041b\u041d\u0418\u0422\u042c'
    |'EXECUTE'
    ;
ADDHANDLER_KEYWORD
    :
    '\u0414\u041e\u0411\u0410\u0412\u0418\u0422\u042c\u041e\u0411\u0420\u0410\u0411\u041e\u0422\u0427\u0418\u041a'
    |'ADDHANDLER'
    ;
REMOVEHANDLER_KEYWORD
    :
    '\u0423\u0414\u0410\u041b\u0418\u0422\u042c\u041e\u0411\u0420\u0410\u0411\u041e\u0422\u0427\u0418\u041a'
    |'REMOVEHANDLER'
    ;
ASYNC_KEYWORD
    : ('\u0410\u0421\u0418\u041d\u0425'
    |'ASYNC') -> pushMode(ASYNC_MODE)
    ;

fragment LETTER: [\p{Letter}] | '_';
IDENTIFIER : LETTER ( LETTER | DIGIT )*;

UNKNOWN: . -> channel(HIDDEN);

mode PREPROCESSOR_MODE;

PREPROC_EXCLAMATION_MARK: '!';
PREPROC_LPAREN: '(';
PREPROC_RPAREN: ')';

PREPROC_STRING: '"' (~["\n\r])* '"';

PREPROC_USE_KEYWORD
    :
    ('\u0418\u0421\u041f\u041e\u041b\u042c\u0417\u041e\u0412\u0410\u0422\u042c'
    |'USE') -> pushMode(USE_MODE);

PREPROC_REGION
    :
    ('\u041e\u0411\u041b\u0410\u0421\u0422\u042c'
    |'REGION' ) -> pushMode(REGION_MODE)
    ;
PREPROC_END_REGION
    :
    ('\u041a\u043e\u043d\u0435\u0446\u041e\u0431\u043b\u0430\u0441\u0442\u0438'
    |'ENDREGION' )
    ;

PREPROC_NOT_KEYWORD
    :
     '\u041d\u0415'
    |'NOT'
    ;
PREPROC_OR_KEYWORD
    :
     '\u0418\u041b\u0418'
    |'OR'
    ;
PREPROC_AND_KEYWORD
    :
     '\u0418'
    |'AND'
    ;

PREPROC_IF_KEYWORD
    :
     '\u0415\u0421\u041b\u0418'
    |'IF'
    ;
PREPROC_THEN_KEYWORD
    :
     '\u0422\u041e\u0413\u0414\u0410'
    |'THEN'
    ;
PREPROC_ELSIF_KEYWORD
    :
     '\u0418\u041d\u0410\u0427\u0415\u0415\u0421\u041b\u0418'
    |'ELSIF'
    ;
PREPROC_ENDIF_KEYWORD
    :
     '\u041a\u041e\u041d\u0415\u0426\u0415\u0421\u041b\u0418'
    |'ENDIF'
    ;
PREPROC_ELSE_KEYWORD
    :
     '\u0418\u041d\u0410\u0427\u0415'
    |'ELSE'
    ;

PREPROC_MOBILEAPPCLIENT_SYMBOL
    :
     '\u041c\u041e\u0411\u0418\u041b\u042c\u041d\u041e\u0415'
     '\u041f\u0420\u0418\u041b\u041e\u0416\u0415\u041d\u0418\u0415'
     '\u041a\u041b\u0418\u0415\u041d\u0422'
    |'MOBILE'
     'APP'
     'CLIENT'
    ;
PREPROC_MOBILEAPPSERVER_SYMBOL
    :
     '\u041c\u041e\u0411\u0418\u041b\u042c\u041d\u041e\u0415'
     '\u041f\u0420\u0418\u041b\u041e\u0416\u0415\u041d\u0418\u0415'
     '\u0421\u0415\u0420\u0412\u0415\u0420'
    |'MOBILE'
     'APP'
     'SERVER'
    ;
PREPROC_MOBILECLIENT_SYMBOL
    :
     '\u041c\u041e\u0411\u0418\u041b\u042c\u041d\u042b\u0419'
     '\u041a\u041b\u0418\u0415\u041d\u0422'
    |'MOBILE'
     'CLIENT'
    ;
PREPROC_THICKCLIENTORDINARYAPPLICATION_SYMBOL
    :
    '\u0422\u041e\u041b\u0421\u0422\u042b\u0419'
     '\u041a\u041b\u0418\u0415\u041d\u0422'
     '\u041e\u0411\u042b\u0427\u041d\u041e\u0415'
     '\u041f\u0420\u0418\u041b\u041e\u0416\u0415\u041d\u0418\u0415'
    |'THICK'
     'CLIENT'
     'ORDINARY'
     'APPLICATION'
    ;
PREPROC_THICKCLIENTMANAGEDAPPLICATION_SYMBOL
    :
     '\u0422\u041e\u041b\u0421\u0422\u042b\u0419'
     '\u041a\u041b\u0418\u0415\u041d\u0422'
     '\u0423\u041f\u0420\u0410\u0412\u041b\u042f\u0415\u041c\u041e\u0415'
     '\u041f\u0420\u0418\u041b\u041e\u0416\u0415\u041d\u0418\u0415'
    |'THICK'
     'CLIENT'
     'MANAGED'
     'APPLICATION'
    ;
PREPROC_EXTERNALCONNECTION_SYMBOL
    :
     '\u0412\u041d\u0415\u0428\u041d\u0415\u0415'
     '\u0421\u041e\u0415\u0414\u0418\u041d\u0415\u041d\u0418\u0415'
    |'EXTERNAL'
     'CONNECTION'
    ;
PREPROC_THINCLIENT_SYMBOL
    :
     '\u0422\u041e\u041d\u041a\u0418\u0419'
     '\u041a\u041b\u0418\u0415\u041d\u0422'
    |'THIN'
     'CLIENT'
    ;
PREPROC_WEBCLIENT_SYMBOL
    :
     '\u0412\u0415\u0411'
     '\u041a\u041b\u0418\u0415\u041d\u0422'
    |'WEB'
     'CLIENT'
    ;
PREPROC_ATCLIENT_SYMBOL
    :
     '\u041d\u0410\u041a\u041b\u0418\u0415\u041d\u0422\u0415'
    |'ATCLIENT'
    ;
PREPROC_CLIENT_SYMBOL
    :
     '\u041a\u041b\u0418\u0415\u041d\u0422'
    |'CLIENT'
    ;
PREPROC_ATSERVER_SYMBOL
    :
     '\u041d\u0430\u0421\u0435\u0440\u0432\u0435\u0440\u0435'
    |'ATSERVER'
    ;
PREPROC_SERVER_SYMBOL
    :
     '\u0421\u0415\u0420\u0412\u0415\u0420'
    |'SERVER'
    ;
PREPROC_MOBILE_STANDALONE_SERVER
    :
     '\u041c\u041e\u0411\u0418\u041b\u042c\u041d\u042b\u0419'
     '\u0410\u0412\u0422\u041e\u041d\u041e\u041c\u041d\u042b\u0419'
     '\u0421\u0415\u0420\u0412\u0415\u0420'
    |'MOBILESTANDALONESERVER'
    ;

PREPROC_LINUX   :'LINUX';
PREPROC_WINDOWS :'WINDOWS';
PREPROC_MACOS   :'MACOS';

PREPROC_IDENTIFIER :LETTER (LETTER |DIGIT )*;

PREPROC_WHITE_SPACE: [ \t\f]+ -> channel(HIDDEN), type(WHITE_SPACE);
PREPROC_LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN), type(LINE_COMMENT);
PREPROC_NEWLINE: '\r'?'\n' -> popMode, channel(HIDDEN);

PREPROC_ANY: ~[\r\n];

mode ANNOTATION_MODE;

ANNOTATION_ATSERVERNOCONTEXT_SYMBOL
    : (
     '\u041d\u0410\u0421\u0415\u0420\u0412\u0415\u0420\u0415'
     '\u0411\u0415\u0417\u041a\u041e\u041d\u0422\u0415\u041a\u0421\u0422\u0410'
    |'ATSERVER'
     'NOCONTEXT'
    ) -> popMode
    ;

ANNOTATION_ATCLIENTATSERVERNOCONTEXT_SYMBOL
    : (
     '\u041d\u0410\u041a\u041b\u0418\u0415\u041d\u0422\u0415'
     '\u041d\u0410\u0421\u0415\u0420\u0412\u0415\u0420\u0415'
     '\u0411\u0415\u0417\u041a\u041e\u041d\u0422\u0415\u041a\u0421\u0422\u0410'
    |'ATCLIENT'
     'ATSERVER'
     'NOCONTEXT'
    ) -> popMode
    ;

ANNOTATION_ATCLIENTATSERVER_SYMBOL
    : (
     '\u041d\u0410\u041a\u041b\u0418\u0415\u041d\u0422\u0415'
     '\u041d\u0410\u0421\u0415\u0420\u0412\u0415\u0420\u0415'
    |'ATCLIENT'
     'ATSERVER'
    ) -> popMode
    ;

ANNOTATION_ATCLIENT_SYMBOL
    : (
     '\u041d\u0410\u041a\u041b\u0418\u0415\u041d\u0422\u0415'
    |'ATCLIENT'
    ) -> popMode
    ;

ANNOTATION_ATSERVER_SYMBOL
    : ('\u041d\u0410\u0421\u0415\u0420\u0412\u0415\u0420\u0415'
    |'ATSERVER'
    ) -> popMode
    ;

ANNOTATION_BEFORE_SYMBOL
    : ('\u041f\u0415\u0420\u0415\u0414'
    |'BEFORE'
    ) -> popMode
    ;

ANNOTATION_AFTER_SYMBOL
    : ('\u041f\u041e\u0421\u041b\u0415'
    |'AFTER'
    ) -> popMode
    ;

ANNOTATION_AROUND_SYMBOL
    : ('\u0412\u041c\u0415\u0421\u0422\u041e'
    |'AROUND'
    ) -> popMode
    ;

ANNOTATION_CHANGEANDVALIDATE_SYMBOL
    : ('\u0418\u0417\u041c\u0415\u041d\u0415\u041d\u0418\u0415\u0418\u041a\u041e\u041d\u0422\u0420\u041e\u041b\u042c'
    |'CHANGEANDVALIDATE'
    ) -> popMode
    ;

ANNOTATION_CUSTOM_SYMBOL
    : (
    LETTER ( LETTER | DIGIT )*
    ) -> popMode
    ;

ANNOTATION_WHITE_SPACE
    : [ \n\r\t\f]+
    -> channel(HIDDEN),
       type(WHITE_SPACE)
    ;

ANNOTATION_UNKNOWN
    : .
    -> channel(HIDDEN)
    ;

mode LABEL_MODE;
LABEL_IDENTIFIER : LETTER ( LETTER | DIGIT )* -> type(IDENTIFIER), popMode;

mode REGION_MODE;
REGION_WHITE_SPACE
    : [ \t\f]+
    -> channel(HIDDEN),
       type(WHITE_SPACE)
    ;
REGION_IDENTIFIER : LETTER ( LETTER | DIGIT )* -> type(PREPROC_IDENTIFIER), popMode;

mode USE_MODE;
fragment USE_LETTER: [\p{Letter}] | '_' | '-';
USE_WHITE_SPACE
    : [ \t\f]+
    -> channel(HIDDEN),
       type(WHITE_SPACE)
    ;
USE_STRING : '"' (~["\n\r])* '"' -> type(PREPROC_STRING), popMode;
USE_IDENTIFIER : ( USE_LETTER | DIGIT )+ -> type(PREPROC_IDENTIFIER), popMode;

mode DOT_MODE;
DOT_WHITE_SPACE
    : [ \t\f]+
    -> channel(HIDDEN),
       type(WHITE_SPACE)
    ;
DOT_IDENTIFIER : LETTER ( LETTER | DIGIT )* -> type(IDENTIFIER), popMode;

mode PREPROC_DELETE_MODE;
PREPROC_ENDDELETE
    : '#' [ \t]*
    (
         '\u041a\u041e\u041d\u0415\u0426\u0423\u0414\u0410\u041b\u0415\u041d\u0418\u042f'
        |'ENDDELETE'
    )
    -> popMode, channel(PREPROC_DELETE_CHANNEL);
PREPROC_DELETE_WHITE_SPACE: [ \t\f]+ -> channel(HIDDEN), type(WHITE_SPACE);
PREPROC_DELETE_LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN), type(LINE_COMMENT);
PREPROC_DELETE_NEWLINE: '\r'?'\n' -> channel(HIDDEN), type(PREPROC_NEWLINE);
PREPROC_DELETE_ANY: . -> channel(PREPROC_DELETE_CHANNEL);

mode ASYNC_MODE;
Async_LINE_COMMENT: LINE_COMMENT -> type(LINE_COMMENT), channel(HIDDEN);
Async_WHITE_SPACE: WHITE_SPACE -> type(WHITE_SPACE), channel(HIDDEN);

// separators
Async_DOT: DOT -> type(DOT), pushMode(DOT_MODE);
Async_LBRACK: LBRACK -> type(LBRACK);
Async_RBRACK: RBRACK -> type(RBRACK);
Async_LPAREN: LPAREN -> type(LPAREN);
Async_RPAREN: RPAREN -> type(RPAREN);
Async_COLON: COLON -> type(COLON);
Async_SEMICOLON: SEMICOLON -> type(SEMICOLON);
Async_COMMA: COMMA -> type(COMMA);
Async_ASSIGN: ASSIGN -> type(ASSIGN);
Async_PLUS: PLUS -> type(PLUS);
Async_MINUS: MINUS -> type(MINUS);
Async_LESS_OR_EQUAL: LESS_OR_EQUAL -> type(LESS_OR_EQUAL);
Async_NOT_EQUAL: NOT_EQUAL -> type(NOT_EQUAL);
Async_LESS: LESS -> type(LESS);
Async_GREATER_OR_EQUAL: GREATER_OR_EQUAL -> type(GREATER_OR_EQUAL);
Async_GREATER: GREATER -> type(GREATER);
Async_MUL: MUL -> type(MUL);
Async_QUOTIENT: QUOTIENT -> type(QUOTIENT);
Async_MODULO: MODULO -> type(MODULO);
Async_QUESTION: QUESTION -> type(QUESTION);
Async_AMPERSAND: AMPERSAND -> type(AMPERSAND), pushMode(ANNOTATION_MODE);
Async_PREPROC_DELETE: PREPROC_DELETE ->
    type(PREPROC_DELETE),
    pushMode(PREPROC_DELETE_MODE),
    channel(PREPROC_DELETE_CHANNEL);
Async_PREPROC_INSERT: PREPROC_INSERT -> type(PREPROC_INSERT), channel(HIDDEN);
Async_PREPROC_ENDINSERT: PREPROC_ENDINSERT -> type(PREPROC_ENDINSERT), channel(HIDDEN);

Async_HASH: HASH -> type(HASH), pushMode(PREPROCESSOR_MODE);
Async_BAR: BAR -> type(BAR);
Async_TILDA: TILDA -> type(TILDA), pushMode(LABEL_MODE);

// literals
Async_TRUE: TRUE -> type(TRUE);
Async_FALSE: FALSE -> type(FALSE);
Async_UNDEFINED: UNDEFINED -> type(UNDEFINED);
Async_NULL: NULL -> type(NULL);
Async_DECIMAL: DECIMAL -> type(DECIMAL);
Async_DATETIME: DATETIME -> type(DATETIME);
Async_FLOAT: FLOAT -> type(FLOAT);
Async_STRING: STRING -> type(STRING);
Async_STRINGSTART: STRINGSTART -> type(STRINGSTART);
Async_STRINGTAIL: STRINGTAIL -> type(STRINGTAIL);
Async_STRINGPART: STRINGPART -> type(STRINGPART);

// keywords
Async_PROCEDURE_KEYWORD: PROCEDURE_KEYWORD -> type(PROCEDURE_KEYWORD);
Async_FUNCTION_KEYWORD: FUNCTION_KEYWORD -> type(FUNCTION_KEYWORD);
Async_ENDPROCEDURE_KEYWORD: ENDPROCEDURE_KEYWORD -> type(ENDPROCEDURE_KEYWORD), popMode;
Async_ENDFUNCTION_KEYWORD: ENDFUNCTION_KEYWORD -> type(ENDFUNCTION_KEYWORD), popMode;
Async_EXPORT_KEYWORD: EXPORT_KEYWORD -> type(EXPORT_KEYWORD);
Async_VAL_KEYWORD: VAL_KEYWORD -> type(VAL_KEYWORD);
Async_ENDIF_KEYWORD: ENDIF_KEYWORD -> type(ENDIF_KEYWORD);
Async_ENDDO_KEYWORD: ENDDO_KEYWORD -> type(ENDDO_KEYWORD);
Async_IF_KEYWORD: IF_KEYWORD -> type(IF_KEYWORD);
Async_ELSIF_KEYWORD: ELSIF_KEYWORD -> type(ELSIF_KEYWORD);
Async_ELSE_KEYWORD: ELSE_KEYWORD -> type(ELSE_KEYWORD);
Async_THEN_KEYWORD: THEN_KEYWORD -> type(THEN_KEYWORD);
Async_WHILE_KEYWORD: WHILE_KEYWORD -> type(WHILE_KEYWORD);
Async_DO_KEYWORD: DO_KEYWORD -> type(DO_KEYWORD);
Async_FOR_KEYWORD: FOR_KEYWORD -> type(FOR_KEYWORD);
Async_TO_KEYWORD: TO_KEYWORD -> type(TO_KEYWORD);
Async_EACH_KEYWORD: EACH_KEYWORD -> type(EACH_KEYWORD);
Async_IN_KEYWORD: IN_KEYWORD -> type(IN_KEYWORD);
Async_TRY_KEYWORD: TRY_KEYWORD -> type(TRY_KEYWORD);
Async_EXCEPT_KEYWORD: EXCEPT_KEYWORD -> type(EXCEPT_KEYWORD);
Async_ENDTRY_KEYWORD: ENDTRY_KEYWORD -> type(ENDTRY_KEYWORD);
Async_RETURN_KEYWORD: RETURN_KEYWORD -> type(RETURN_KEYWORD);
Async_CONTINUE_KEYWORD: CONTINUE_KEYWORD -> type(CONTINUE_KEYWORD);
Async_RAISE_KEYWORD: RAISE_KEYWORD -> type(RAISE_KEYWORD);
Async_VAR_KEYWORD: VAR_KEYWORD -> type(VAR_KEYWORD);
Async_NOT_KEYWORD: NOT_KEYWORD -> type(NOT_KEYWORD);
Async_OR_KEYWORD: OR_KEYWORD -> type(OR_KEYWORD);
Async_AND_KEYWORD: AND_KEYWORD -> type(AND_KEYWORD);
Async_NEW_KEYWORD: NEW_KEYWORD -> type(NEW_KEYWORD);
Async_GOTO_KEYWORD: GOTO_KEYWORD -> type(GOTO_KEYWORD);
Async_BREAK_KEYWORD: BREAK_KEYWORD -> type(BREAK_KEYWORD);
Async_EXECUTE_KEYWORD: EXECUTE_KEYWORD -> type(EXECUTE_KEYWORD);
Async_ADDHANDLER_KEYWORD: ADDHANDLER_KEYWORD -> type(ADDHANDLER_KEYWORD);
Async_REMOVEHANDLER_KEYWORD: REMOVEHANDLER_KEYWORD -> type(REMOVEHANDLER_KEYWORD);
AWAIT_KEYWORD
    : ('\u0416\u0414\u0410\u0422\u042c'
    |'AWAIT')
    ;

// всегда в конце мода
Async_IDENTIFIER: IDENTIFIER -> type(IDENTIFIER);
Async_UNKNOWN: UNKNOWN -> type(UNKNOWN);
