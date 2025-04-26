+++
title = "Config_rant"
date = "2023-05-15T10:11:12+02:00"
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

Fighting with Traefik with project and started thinking about config

Preface: I'm not familiar with kubernetes, but familier with ops in general. My
thought about this will probably change when I inevitably has to learn the inner
workings.


You have 3 ways to define options for Traefik, 4 IMO since things can be defined
on each service(docker container in my case). It's very easy to create invalid
config and options that crash with each other. The thing that annoys me is that
this has to be validated at run time, when can be known at compile time. At
least some things, but it gets difficult, when the config is distributed. 

Declarative configuration is totally reliant on documentation for the developer
ot understand what is happening. LSP support in some way to get documenation
could be a solution, but often you just have to google and hope their
documentation is up to date and well structured.


api:
    insecure: true


In isolation, what does this mean?


Is the problem yaml? Traefik? or tooling?

Some tooling is often tied to VSCode, which I would argue is unhealthy for the
ecosystem. I don't what a chromeification of the IDE space, where one company
controls how developers code. 

What are names and what are 
