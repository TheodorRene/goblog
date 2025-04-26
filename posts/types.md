+++
title = "Types"
date = "2022-01-17T23:55:39+01:00"
author = ""
authorTwitter = "" #do not include @
cover = ""
tags = ["", ""]
keywords = ["", ""]
description = ""
showFullContent = false
readingTime = false
draft = true
+++

Types is something I've taken for granted in programming languages. They just
exist, and are probably necessary. Gradually when going from dynamically typed
languages (types not known at compile time), like Python to statically typed
(types known at compile time) languages like Java I taken to appreciating types.
And I think few people appreciate and understand the importance. In one way they
have to be there because of technical limitations but in other ways it's a very
important tool for writing software that's easier to work. This blogpost
describes my intuition and how I understand types and their importance. Most of
it will be applicable for any programmer but the ideas are quite common when
working within the functional programming paradigm

This post will be affected by my learnings of Haskell the last few years.

Types are all about setting boundaries for what a function can take as input and
// Its also when defining data, records etc
what it will return on output. Giving limiting types will reduce the amount of
possible input and same with the output type. 

A good programming language should the user reduce the cognitive load. Reducing
whats stored in my temporary memory. This lets the user focus on the
important things.

# Partial functions
Functions should be able to handle any input parameters it gets. In many languages this is
not upheld by the compiler but up to the programmer. In most idioms its a bad
practice. In many languages like Python this is done with exceptions. Bad input?
Throw exception and cross your fingers that it's catched somewhere, else the
program will crash and possibly halt. This is not very safe, but very
convenient.


# String and lists
Strings and lists are something you can't live without in a programming language
but at the same time the cause of many problems. What makes these two difficult
is at compile time you don't know much about them. Lets look at an example:

```haskell
data BeatleMember = George | Ringo | Paul | John
data HairCut = Bowlcut | LongHair 
myFunc :: BeatleMember -> HairCut

myOtherFunc :: String -> Int
```

Indexing a list is the cause of much confusion. It's by its inherent nature a
partial function. Lets just write it like a function so that syntatic sugar does
not cloud the mind

```python
# Indexing is just a function
# python     haskell(infix)   haskell (as a function)
  list[0] => list !! 0     == (!!) list 0
```

```haskell
:t (!!)
(!!) :: [a] -> Int -> a
[0 ,0] !! 3
*** Exception: Prelude.!!: index too large
```

It will not return a something for any input it gets! That why mostly don't
index lists in functional languages. Exception are frowned upon, a function is a
function. Input and output, plain and simple.

Having exceptions ruins the flow.

Some function will be as general as possible. So adapt to the situation.

# Pure function are just like a table lookup
BeatleMember have 4 possibilities, HairCut has 2 possibilites. Just by looking
at the type declaration we know that the final set of possible input/output
pairs are $4 * 2 = 8$. (Want an efficiency tip? Since the final set is very
small why can't it just be table lookup? If the function was very compute heavy
this could be a possibility. As long as the function doesn't have any side
effect like checking a database or doing a network request.)

Lets look at myOtherFunc. By looking at the types we know very little about how
it works. Lets pretend that String can only hold 2048 bytes and that Int is 64
bytes. That give $ 2^2048 * 2^64 $ possible input and output pairs. Try to wrap
your head around that!

Most importantly when you are restricting the type you limiting how many states
the function can exist in. If you take in an argument of type BeatleMember your
function only needs to handle 4 different possible inputs. 

## In the end..

Learn some Haskell or some other statically typed functional programming
language. It really changed my views on functions, input, output and how it all
links together. 
