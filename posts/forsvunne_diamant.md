+++
title = "Forsvunne_diamant"
date = "2023-11-04T22:25:51+01:00"
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

A popular board game when I was a kid is "Den forsvunne diamant", or "Jakten på
den forsvunne diamant" as I thought it was
[called](https://en.wikipedia.org/wiki/False_memory#Mandela_effect). Roughly
translated its the "The missing diamond". You throw a dice which decides how
many steps you move and the you look under pieces on the board. Under them you
could possibly find the diamond! This gives me my cheap segway in to todays
blogpost. What happens to the missing commit? Lets put on our adventurer
hat(pith helmet?) and find it

I've always recommended rebasing when I join teams because I argue that it leads
to a cleaner history and a simple mental model. Instead of a complex tree that
diverges and meets again you get a single linear history which you can travel
back and forth on. Unfortunately it just leads being messing up their git
history, breaking Github somehow or just reverting to merge. Which in the ends
is fine. Pragmatism is something I've learned this first year in the industry
and its an awakener for me.

But there is one case with Git and rebase that I don't understand. A simple
workflow that might lose changes! Or so I think. Lets look at my illustration:

![illustration](/img/diamond_git.png)

This is my hypothesis, and through this blogpost we will found out what actually
happens. Will the change X return? Or is Git smart than I think? Stick around!
