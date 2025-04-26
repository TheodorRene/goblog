+++
title = "Deleting my backup by accident"
date = "2022-02-02T13:43:42+01:00"
author = "Theodor René Carlsen"
authorTwitter = "teddytroll" #do not include @
cover = ""
tags = ["", ""]
keywords = ["backup", "cli"]
description = ""
showFullContent = false
readingTime = true
+++

In my post about [notetaking](/posts/notetaking) I displayed my simple backup
solution (It's more like mirroring/online syncing since I don't keep old data
because of storage limitations). It's been working great but the problem I had
sometimes is that I download big files to my notetaking directories, which will
fill the drive of the VM I'm renting (It's only about 25 GBs). 

```bash
> cat /etc/cron.daily/copy_notes
#!/bin/bash
rsync -a --delete /home/theodorc/notater api.theodorc.no:/home/theodorc/notater
```

I wanted to exclude any file with the prefix "no_backup" to solve this. This way
I could mark files that I didn't want to backup. Rsync supports this through the
"--exclude" flag, but I was unsure of the syntax of the pattern to match. So I
made a test directory to try it out with two example files.


```bash
> mkdir test && cd test
~/test> touch gurba no_backup # create two empty files
~/test> ls
gurba no_backup
```
With not much thought, I copied my original oneliner for backing up my files,
added the "--exclude" flag and set my current directory as the source and the home
folder on my server as the destination.

```bash
~/test> rsync -a --delete --exclude="^no_backup.*" . api.theodorc.no:~
```

This command ended up deleting almost every file in the home directory of my
server. Do you see why?

```bash
theodorc@disco> ssh api.theodorc.no # yes, the hostname of my laptop is "disco"
theodorc@api> ls
dev gurba no_backup_gurba
```

My little script uses the "--delete" flag to delete files that no longer exist
on my local computer. If I delete something locally, I most likely don't care
about it being in my backup. (I should probably change this assumption, since
accidentally deleting a file locally will save me. Bricking my computer has been
the biggest worry). Rsync checks my home directory on the remote and deletes
everything that doesn't not match the files in the current directory on the
local machine. 

So in the end the syntax was wrong and I lost my backup :--) Till next time,
be careful with which flags I'm using.

Luckily I have all the files locally, so on the next backup run they will
return. Not so lucky is two of my projects. Images for my [homepage](https://theodorc.no)
(which I'm already thinking about retiring) was deleted and the database for my
[ChessPuzzleBot](https://twitter.com/ChessDaily) was deleted. Oh well.


