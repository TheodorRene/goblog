+++
title = "Notetaking"
date = "2021-12-27T20:21:52+01:00"
author = ""
authorTwitter = "" #do not include @
cover = ""
tags = ["notetaking"]
keywords = ["", ""]
description = ""
showFullContent = false
readingTime = false
+++

# Writing notes using Markdown and Vim - My setup

I'm currently in my ~~fourth~~fifth year of studying computer science and during this
time I've needed to write a lot of notes. After roughly a year using Onenote and
ended up on a combination of tools I really liked. Vim and markdown. In this
post I will go through why I made that choice, what my setup is and some final
thoughts.

My first year I used Onenote, a note taking app part of Microsofts office. It
worked okay, it let me do notetaking using a digital pen and allowed export to
pdf. It also synced nicely to the cloud. After some time I fell in love with Vim
on Linux and that left me without the possibility of using Onenote. So I tried
note taking in Vim. I really like using my preferred editor (software engineers
sometimes have one) that I was fluent with. Using the OSs filesystem to create a
hierarchy for the notes also was just easier. It let me do all my normal
operations I have grown accustomed to in the terminal. I also felt more in
control of the notes, not being scared of Onenote crashing on me and messing up
whatever file type it is using to represent the notepads. With Markdown it was
just me and a text file. 

Why markdown?
* Easy to compile to pdf, html whatever with Pandoc
* Simple syntax
* Popular format
* Pandocs version supports latex for math

My setup now consists of setting the line length to 80 so it autowraps when
hitting a limit. I have enabled Vims built in spellchecker and some handy
keybindings to interact with it. My settings are gathered in a
[Vim-plugin](https://github.com/TheodorRene/skriveleif). Mostly to see how you
make a plugin, it is not a very comprehensive solution.

Currently backup (and in some way online syncing) is done with a simple
cron-job

```bash
> cat /etc/cron.daily/copy_notes
#!/bin/bash
rsync -a --delete /home/theodorc/notater api.theodorc.no:/home/theodorc/notater
```

## Future plans
I have heard good things about Obsidian which can give a better structure and
possibly easier online syncing.




