package main

const titleAsset = `
x('-.x.-.xxx('-.xxxxxx_x('-.xxxx_x('-.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
(xOOx)xx/xx(xOOx).-.x(x(OOxx)xx(x(OOxx)xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
,--.x,--.xx/x.x--.x/_.'xxxxx\x_.'xxxxx\xx,--.xxx,--.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
|xx|x|xx|xx|x\-.xx\(__...--''(__...--''xxx\xx'.'xx/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
|xxx.|xx|.-'-'xx|xx||xx/xx|x|x|xx/xx|x|x.-')xxxxx/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
|xxxxxxx|x\|x|_.'xx||xx|_.'x|x|xx|_.'x|(OOxx\xxx/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
|xx.-.xx|xx|xx.-.xx||xx.___.'x|xx.___.'x|xxx/xx/\_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
|xx|x|xx|xx|xx|x|xx||xx|xxxxxx|xx|xxxxxx'-./xx/.__)xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
'--'x'--'xx'--'x'--''--'xxxxxx'--'xxxxxxxx'--'xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
x('-.x.-.xxx('-.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx('\x.-')x/'xxx('-.xxxxx('-.xxxxxxx.-')x_xx
(xOOx)xx/xx(xOOx).-.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx'.(xOOx),'x_(xxOO)xx_(xxOO)xxxxx(xOOx)x)x
,--.x,--.xx/x.x--.x/x,--.xxxxxx,--.xxxxxx.-'),-----.x,--./xx.--.xx(,------.(,------.,--./x,--,'xx
|xx|x|xx|xx|x\-.xx\xx|xx|.-')xx|xx|.-')x(xOO'xx.-.xx'|xxxxxx|xx|xxx|xx.---'x|xx.---'|xxx\x|xx|\xx
|xxx.|xx|.-'-'xx|xx|x|xx|xOOx)x|xx|xOOx)/xxx|xx|x|xx||xx|xxx|xx|,xx|xx|xxxxx|xx|xxxx|xxxx\|xx|x)x
|xxxxxxx|x\|x|_.'xx|x|xx|'-'x|x|xx|'-'x|\_)x|xx|\|xx||xx|.'.|xx|_)(|xx'--.x(|xx'--.x|xx.xxxxx|/xx
|xx.-.xx|xx|xx.-.xx|(|xx'---.'(|xx'---.'xx\x|xx|x|xx||xxxxxxxxx|xxx|xx.--'xx|xx.--'x|xx|\xxxx|xxx
|xx|x|xx|xx|xx|x|xx|x|xxxxxx|xx|xxxxxx|xxxx''xx'-'xx'|xxx,'.xxx|xxx|xx'---.x|xx'---.|xx|x\xxx|xxx
'--'x'--'xx'--'x'--'x'------'xx'------'xxxxxx'-----'x'--'xxx'--'xxx'------'x'------''--'xx'--'xxx`

const batAsset = `
xx_xxx,_,xxx_
x/xx'=)x(='xx\
/.-.-.\x/.-.-.\
'xxxxxx"xxxxxx'
`

const tombstoneAsset = `
xxxxxxxxxxxxxxxxxxxxxx.--.x.-,xxxxxxx.-..-.__
xxxxxxxxxxxxxxxxxxxx.'(´.-´x\_.-'-./´xx|\_(x"\__
xxxxxxxxxxxxxxxxx__.>\x';xx_;---,._|xxx/x__/´'--)
xxxxxxxxxxxxxxxx/.--.xx:x|/'x_.--.<|xx/xx|x|
xxxxxxxxxxxx_..-'xxxx´\xxxxx/'x/´xx/_/x_/_/
xxxxxxxxxxxxx>_.-´´-.x´Yxx/'x_;---.´|/))))
xxxxxxxxxxxx'´x.-''.x\|:xx\.'xxx__,x.-'"´
xxxxxxxxxxxxx.'--._x´-:xx\/:xx/'xx'.\xxxxxxxxxxxxx_|_
xxxxxxxxxxxxxxxxx/.'´\x:;xxx/'xxxxxx´-xxxxxxxxx´-|-´
xxxxxxxxxxxxxxxx-´xxxx|xxxxx|xxxxxxxxxxxxxxxxxxxxxx|
xxxxxxxxxxxxxxxxxxxxxxx:.;x:x|xxxxxxxxxxxxxxxxxx.-'~^~´-.
xxxxxxxxxxxxxxxxxxxxxxx|:xxxx|xxxxxxxxxxxxxxxx.'x_xxxxx_x'.
xxxxxxxxxxxxxxxxxxxxxxx|:.xxx|xxxxxxxxxxxxxxxx|x|_)x|x|_)x|
xxxxxxxxxxxxxxxxxxxxxxx:.x:xx|xxxxxxxxxxxxxxxx|x|x\x|x|xxx|
xxxxxxxxxxxxxxxxxxxxx.xxx.x:x;xxxxxxxxxxxxxxxx|xxxxxxxxxxx|
xxxxxxxxxxxx-."-/\\\/:::.xxxx´\."-._'."-"_\\-|xxxxxxxxxxx|///."-
xxxxxxxxxxxx"x-."-.\\"-."//.-".´-."_\\-.".-\\´=.........=´//-".
`

const pumpkin_1 = `
xxxxxxxx__J"L__
xxxx,-'´--...--'-.
xxx/xx/\xx___xx/\xx\
xxJxxx""xx\_/xx""xxxL
xx|xxxxxxxxxxxxxxxxx|
xxJxxxx/\/VWV\/\xxxxF
xxx\xx(/\/\_/\/\)xx/
xxxx"-._xxxxxxx_,-"
xxxxxxxx"""""""
`

const houseAsset = `                     _____
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_.-""xxxxx""-._
xxxxxxxxxxxxxxxxxxxxxxxxxxxxx.'xxxxxxxxxxxxxxx'.
xxxxxxxxxxxxxxxx.x.x.x.x.xx.'xxxxxxxxxxxxxxxxxxx'.
xxxxxxxxxxxxxxxx!-!-!-!-!x/xxx.-..xxxxxxxxxxxxxxxx\
xxxxxxxxxxxxxxxx!_!,!_!_!/xxxx|__Hx_xxxx.-\_)´-.xxx\
xxxxxxxxxxxxxx,/´,/'_´\,´\,xx[____|_]xx/.-.x.-,_\xxx;
xxxxxxxxxxxx,/',/'/_|_\´\,´\,|=xxx|=|xxxxxx\(xxxxxxx;
xxxxxxxxxx,/'x|/x||"""||x\|x´\,x=x|x|xxxxxxx´xxxxxxx|
xxxxxxxxxx|xxx#|x||___||x|x#xx|=xx|x|xxxxxxxxxxxxxxx|
xxxxxxxx,/'x#xx|x[_____]x|xxx#´\,x|=|xxxxxxxxxxxxxxx;
xxxxxx,/',-----'xxxxxx=xx'-----,´\--'---,/\,--,xxxx/
xxxxx´""|xxx.;;;,=xxxxxx,;;;,xxx|#xx#x,//xx\\,'\,x/
xxxxxxxx|x=//___\\x=xxx//___\\x=|x#x,//',;;,'\\,#\,
xxxxxxxx|xx||xxx||xxxxx||xxx||xx|#,//'x//||\\x'\\,´\,
xxxxxxxx|xx||xxx||xxxxx||xxx||xx|-|/|x||_||_||x|\|_x_|
xxxxxxxx|xx||xxx||xxx=x||xxx||xx|=xx|x|.----.|=|___]
xxxxxxxx|=x||___||x=xxx||___||x=|xx=|x||xxxx||x|x||
xxxxxxxx|x[_______]xxx[_______]x|--.|x||____||x|x||
xxxxxxxx;_______=______=_____x__;xxx|[________]|x||
xxxxxx,/'#xxxx#xxx#xxxxxx#xxxxxxx#xx'----------''\|
xxxx,/'xxxx#xxxxxxxx#xxxxxxx#xxxxxxxxx#xxxxx#xxx#x'\,
xx,/'___#____#__#_____#___#_______#_______#____#___#'\,
xx´""[____________________________________________]|""´
xxxxx_[_|xxx.-----.xx=-xxxxxxx___________xxxx||_]_||
xxxx|xx_|x.",-"|"-,',xxx()x=x|.--..-..--.|x=x|_xx|||
xxxx|_/x|/x/_\_|_/_\x\x/__\xx||__||_||__||xxx|x\_|\\
xxxx(_)x||x.-------.x|x|xx|xx|.--..-..--.|xx=|x(_)x||
xxxx/x\x||x|xxxxxxx|x|x|()|xx||__||_||__||xxx|x/x\x||
xxxx\_/x||x'-------'x|xx)(xxx|___________|xxx|x\_/x||
xxxx(_)x||.---------.|xx\/xxx|.---------.|=xx|x(_)x||
xxxx/x\x|||xxx___xxx||xxxx=xx||xxxxxxxxx||xxx|x/x\x||
xxxx\_/x|||xx{___}xx||xxxxxxx||xxxxxxxxx||xxx|x\_/x||
xxxx(_)x|||xx((_))xx||x-=xxx=||_________||x=x|_(_)_//
xxxx/x\x|||xxx'-'xxx||xxx_x.-'-----------'-.x|x/x\__\
xxxx\_/_|||xxxxxxx()||xx[_]"""""""""""""""""[_]\_/\\\\
xxx[x__x]||xxxxxxxxx||x=|x|==.==.==.==.==.==|x|__]|||||
xxxj|xx||||xxxxxxxxx||xx|x|xx|xx|xx|xx|xx|xx|x|xx|====|
xxxg|__|||;).xxxxxxx||--|_|=='=='=='=='=='==|_|xx||||||
xx_s____/´---´\x____||_.____._____._____.____.|__|////
x|xxxxx|xx9.9xx|====='x|xxxx|xxxx/xx\xxxx\xxx|xxxx|-'
x'------\xwwwx/---'----'----'---'----'---'---'----'
xxxxxxxxx´---'
`

const cloudC0 = `
xxxxxxxxx_ _ __.
xxxxxxx('       ).
xxxxxx(          ).
xxxxx_(          '''.
x.=('(           .   )
((        (..____.:'-'
'(        )_),
xx' _______.:'   )
xxxxxxxxxx-----'`

const cloudC1 = `
xxxx.--
x.+(   )
x(   .  )
(   (   ))
x'- __.'`

const ghostC0 = `
x.-.
(o o)
| O \
x\   \
xx` + "`" + `~~~'
`

const ghostC1 = `
xxx.-.
xx(o o)
xx/ O |
x/   /
` + "`" + `~~~'
`
