# Torrentd


## Bencoding Grammar

```
<value>     ::= <integer> | <string> | <list> | <dictionary>

<integer>   ::= 'i' <digits> 'e'
<string>    ::= <length>:<data>
<list>      ::= 'l' <value>* 'e'
<dictionary>::= 'd' (<string> <value>)* 'e'
```
