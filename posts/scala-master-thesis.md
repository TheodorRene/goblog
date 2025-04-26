+++
title = "Scala Master Thesis"
date = "2022-04-19T14:22:47+02:00"
author = "Theodor René Carlsen"
authorTwitter = "teddytroll" #do not include @
cover = ""
tags = ["", ""]
keywords = ["scala", "pl", "master thesis"]
description = ""
showFullContent = false
readingTime = true
draft = true
+++

## Our master thesis
Short description of our master thesis. I worked with Eivind Kopperud on
"Tempura: A Graphx extension for temporal graphs in Apache
Spark"[Link]("https://home.samfundet.no/~theodorc/final_complete_thesis_final.pdf").
TLDR: Developed a temporal extension for Graphx, which is a library for handling
big graphs in Spark. Through this work we had to augment existing datasets to
get the properties we wanted from a graph. Two models were implemented and
benchmarked against each other. This was all done in Scala, a language we had
used somewhat in previous courses but not very much. Again with this a noted
down a few thoughts I had on the language. This post will cover what I liked,
disliked and what I though was missing, including some
examples.

## RDDs
The datatype we had to work with to stay withthin the distributed spark
abastraction are RDDs, Resilient Distributed Dataset. In practice they can be
used like how you use immutable lists. No indexing, just iterating through using
filter, map, reduce etc. 

```scala
// Type signature for dataset of tuples
val listOfTuples: RDD[(Int,Int)]
// RDD<(Int, Int)> with C-like syntax as this RDD is generic
```

## Named tuples
Transforming and working with the data you often had to create some ad hoc
datastructure. Lets say you had Graph datastructure but you wanted to add a
timestamp to it. This can be done using the product type tuple:

```scala
val graphs: RDD[Graph] = getGraphs()
val timestamps:RDD[Instant] = getTimestamps()
val transform = graph
                    .join(timestamps) // Type is now RDD[(Graph,Instant)]
                    .map(tup => {
                        tup._1.timestamp = tup._2
                    }
                    )
```
This example could be make clearer with tuple destructuring or named tuples. Suggested syntax
change:

```scala
                    .map((graph,timestamp) => {
                        graph.timestamp = timestamp
                    }
                    )
                    // or
                    .join(timestamps) // Join now supports named tuples is now RDD[NamedTuple(Graph,Instant)]
                    .map(tup => {
                        tup.graph.timestamp = tup.timestamp
                    }
                    )
```
Right now the destructuring can be done with a clunky case statement:
DONT THINK THIS IS THE RIGHT SYNTAX
```scala
                    .map({case (graph, timestamp) => {
                        graph.timestamp = timestamp
                    }
                    )
```

# Implicits are bonkers
Type classes in Haskell
Implicit parameters
Does it confuse more than it helps?
Easy to abuse

# Verbosity of the language
Verbose ADTs. Scala3 with breaking changes leaves a severed ecosystem (Just saw
Severance, please watch it)

# Alternatives to Scala
What I'm looking for in a language:
* Good ecosystem (Possibly with interop)
* Functional first (immutable datastructures etc)
* Expressive language (Possibly with a macro system)
* A well thought out error handling system
* Simple default when it comes to HTTP server, database connection, test
  framework

F#, Kotlin, Clojure Systems programming language


