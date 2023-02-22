@REM java -jar ./antlr-4.11.1-complete.jar -Dlanguage=Go -package parser  -o parser *.g4
@REM java -jar ./antlr-4.11.1-complete.jar -Dlanguage=Go -no-visitor -package parser *.g4

java -Xmx500M -jar ./tools/antlr-4.12.0-complete.jar -Dlanguage=Go -no-visitor -o antlrparser *.g4 