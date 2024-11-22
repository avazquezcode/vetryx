# GoVetryx
[![Go](https://github.com/avazquezcode/vetryx/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/avazquezcode/vetryx/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/avazquezcode/vetryx/graph/badge.svg?token=BU2ZL47MNK)](https://codecov.io/gh/avazquezcode/vetryx)
<a href="https://goreportcard.com/report/github.com/avazquezcode/vetryx"><img src="https://goreportcard.com/badge/github.com/avazquezcode/vetryx" alt="Go Report Card" /></a>

> [!NOTE]  
> This is a really simple (toy) language, just developed for fun a few months ago.

## Vetryx
After reading the **AMAZING**: [Crafting Interpreters](https://www.amazon.com/dp/0990582930), I wanted to try out the ideas of the first half of the book (Tree-Walk interpreter).

I mainly followed the **Java** implementation of the _Lox_ language, and ported it to Golang, changing some really tiny things in the process, to experiment a bit on the topic. Hence I named it differently: Vetryx :)

## About the language
- Refer to this [doc](LANGUAGE.md) to understand the syntax of the language, and some of its rules.

Some things I changed in comparison to the _Lox_ language:

- [x] Parentheses to wrap if/while conditions are not necessary (but can be added if wanted)
- [x] Added modulus operator
- [x] Changed some reserved words & chars used for some operators
- [x] Added support for short variable declarator (eg: `a := 1`)
- [x] Added support for sleep, min and max native fns
- [x] Support `break` and `continue` in while loop

## Playground

I developed a playground for it, using Next.js and NextUI.
You can test it online [here](https://govetryx.agustinvazquez.me/).

## Ack
Thanks a lot Robert Nystrom for writing such a pleasant book to read! One of the nicest books about software that I've read in the past years...
