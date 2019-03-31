# Walidator

Uproszczony walidator składni dla podzbioru języka OWL 2.0 zapisanego z pomocą wersji funkcjonalnej jego składni.
Celem jest walidacja tekstu zawierającego szereg aksjomatów określających wzajmne relacje pomiędzy klasami lub jednostkami.
Zakres sprawdzanej składni zawiera następujące konstrukcje: `SubClassOf`, `EquivalentClasses`, `DisjointClasses`, `SameIndividual`,
`DifferentIndividuals`, `ObjectIntersectionOf`, `ObjectUnionOf`, `ObjectComplementOf` i `ObjectOneOf`. Nie uwzględniono możliwości
dodawania adnotacji.

## Gramatyka BNF

```
ALPHA ::= [a-zA-Z]+

start ::= <program>
<program> ::= <axiom> <program>

<id> ::= ':'ALPHA
<ids> ::= <id> <ids> | eps

<classExpression> ::= <id> | <objectIntersectionOf> | <objectUnionOf> | <objectComplementOf> | <objectOneOf>
<classExpressions> ::= <classExpression> <classExpressions> | eps

<axiom> ::= <subClassOf> | <equivalentClasses> | <disjointClasses> | <sameIndividual> | <differentIndividuals> | eps

<subClassOf> ::= 'SubClassOf' '(' <classExpression> <classExpression> ')'
<equivalentClasses> ::= 'EquivalentClasses' '(' <classExpression> <classExpression> <classExpressions> ')'
<disjointClasses> ::= 'DisjointClasses' '(' <classExpression> <classExpression> <classExpressions> ')'
<sameIndividual> ::= 'SameIndividual' '(' <id> <id> <ids> ')'
<differentIndividuals> ::= 'DifferentIndividuals' '(' <id> <id> <ids> ')'

<objectIntersectionOf> ::= 'ObjectIntersectionOf' '(' <classExpression> <classExpression> <classExpressions> ')'
<objectUnionOf> ::= 'ObjectUnionOf' '(' <classExpression> <classExpression> <classExpressions> ')'
<objectComplementOf> ::= 'ObjectComplementOf' '(' <classExpression> ')'
<objectOneOf> ::= 'ObjectOneOf' '(' <id> <ids> ')'
```

## Uruchomienie programu

Wymagania:
* zainstalowany [Golang](https://golang.org/dl/)

Instalacja walidatora:
```
go get github.com/b4k3r/walidator
```

Uruchomienie walidatora na przykładowym pliku:
```
walidator test_program.txt
```



