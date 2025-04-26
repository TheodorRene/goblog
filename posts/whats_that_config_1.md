+++
title = "ZSH - What's that config?"
date = "2022-01-13T19:34:17+01:00"
author = "Theodor René Carlsen"
authorTwitter = "teddytroll" #do not include @
cover = ""
tags = ["config", "dotfiles"]
keywords = ["", ""]
description = ""
showFullContent = false
readingTime = true
draft = false
+++

I have been using a lot of CLI tools the last few years and have accumulated a
plethora of configurations, AKA dotfiles. They are all gathered in a git repo on
[Github](https://github.com/theodorene/dotfiles) with a script that adds
[symlinks](https://en.wikipedia.org/wiki/Symbolic_link) so the files are
put in the right place. With this experience, some might be intrigued by my
specific configs and possibly learn some from it. I will try to make a series on
each one of them. The first one is my Zsh config.

## Aliases
Aliases are great for those hard-to-remember commands or commands that are used
often. The aliases are introduced in the order of how often I use them.

```bash
# Number of aliases
> grep -c alias .zshrc 
34
```

### lr

```bash
alias lr="ls -ltrh"
```

Capital letters indicate the flag explained: show as List, sort by Time, in
Reverse, with Human readable size outputs (kilobytes instead of bytes). In
english, show me the most recent files on the bottom. Those files are often more
of interest than the standard stort for "ls -l".

### Pomodoro

```bash
alias work="sleep $((60*20)) && aplay /usr/share/sounds/speech-dispatcher/test.wav &&  notify-send -t 5000 -u critical '5 minutter pause'"
alias break="sleep $((60*5)) && aplay /usr/share/sounds/speech-dispatcher/test.wav &&  notify-send -t 5000 -u critical '5 minutter pause'"
```

I stole the idea of a Bash Pomodoro timer from
[Carl](https://github.com/calidbaba/) but nevertheless, here is my solution. The
audiofile can be whatever you want, but this was the only file I could find on
my computer. Notify-send gives a nice notification window in the top right
corner of my screen.

### Clipboard

```bash
alias clip="xclip -selection c"
```

Using the clipboard in the terminal has always been a pain in the ass for me.
This alias lets me pipe output into my clipboard. Xclip is not installed by
default on distros like Ubuntu. Example use:

```bash
> ls -l | clip
```

### Installing packages

```bash
alias pacupgrade="sudo pacman -Syyu"
alias pacsearch="sudo pacman -Ss"
alias pacinstall="sudo pacman -S"
alias pacremove="sudo pacman -R"
alias pacclean='sudo paccache -r && sudo pacman -Qtdq | sudo pacman -Rns -'
```
Until recently, I used Manjaro on both my laptop and desktop computer. Then
these commands were handy as I thought the ergonomics of Pacman wasn't quite
there. In the future, I want to make a bash function that is package
manager independent. At least for Apt and Pacman. 

### VPN
```bash
alias vpn="sudo openconnect -bq --user=$USER vpn.ntnu.no"
```
At my university, a VPN is necessary to view some resources. Openconnect works
great from the command line. This prompts for my password, which I get from my
*pass*-function described in the next section

## Functions

### Getting password

```bash
pass() {
    bw get password $1 | xclip -selection c && echo "Password copied"
}
```

I'm a big fan of Bitwarden as a password manager. It is open source, the free
solution is comprehensive, and the device support is decent. It has CLI tool and
with it I fetch some of my passwords. Using pass like in the following snippet
brings my NTNU password into my clipboard.

```bash
> pass feide
? Master password: [input is hidden]
Password copied
```

### Getting the wifi password

```bash
# This doesnt work all the time
wifi_pass(){
    sudo grep -oP '^psk=\K\w+' /etc/NetworkManager/system-connections/$(nmcli -t -f name connection show --active | head -n1).nmconnection
}
```

It has to be a better way for this, but it works most of the time, and that's
enough for me. Quite convenient. It searches for a specific line in a file for
the current connected wireless network.

### Using a USB-drive

```bash
mnt(){
    sudo mount $1 /mnt/usb && cd /mnt/usb && ls
}
umnt(){
    sudo umount /mnt/usb
}
```
Some window managers on Linux, like i3, will not do a lot for you. One of those
things is auto mounting USB-drives. I could probably use some other service or
some systemd magic to make this happen, but I like this. Use:

```bash
> mnt /dev/<your device>
```

### Searching
```bash
tgrep(){
    command grep -RIin --color=auto --exclude-dir={.git,venv,node_modules} $1 *
}
```
You might be familiar with "ag" and "rg". They are great alternatives to grep
that is much faster. Sometimes you might be on a system without those programs.
This function makes grep work like "ag" and "rg" except for the speed.

```bash
pdfgrep(){
    command pdfgrep -i $1 *
}
```
I have also done this with pdfgrep. It's a program for searching through PDFs.
Quite convenient when you have a home exam and want to search through old exam
solutions (When it's allowed, of course).

## Other stuff

### Pager
```bash
export LESS="-RFX"
```
This line fixes a pet peeve for me. Some programs like git use a "pager" for
displaying data. With git a "git diff" might have little output, and I don't
like the pager to be opened for so little data. With this environment variable,
Less will only open if the data is of a certain size, other times it will just
display the data in the terminal.

### Plugins

```bash
plugins=(
    git # adds aliases for git (gst == git status etc)
    wd # adds shortcuts to directories 
    catimg # try to display an image in the terminal
    last-working-dir # open last working directory when opening a new terminal
    fzf # fuzzy searching through history and directories
    extract # unzip most archive formats with the command "x"
    sudo # double tap Esc to run the last command with sudo
    docker # autocomplete for docker commands
    autojump # jump to often visited directories
)
```
I will not explain all of these since some are self-explanatory. Autojump is a
bit hard to explain, but it tries its best to jump to the correct directory by
just giving the name of the directory.

```bash
theodorc@abba ~
> j blog
/home/theodorc/dev/blog/content
theodorc@abba ~/dev/blog/content
> j b
/home/theodorc/dev/blog
theodorc@abba ~/dev/blog
>
```
As you can see, it's trying its best

Example use of extract:

```bash
> x my_zip_file.zip
> x my_zip_file.tar
> x my_zip_file.tar.gz
```
It handles all these cases as you would expect.

### Man
I am fond of using man pages for documentation and often expect
program authors to include a man page that explains the use. 

```bash
export MANPAGER='nvim +Man!'
```
This environment variable lets me use Neovim for viewing man-pages. It gives
nicer colors, and I'm used to Vim for movement. Not that useful.


```bash
man(){
    command man $1 || $_ --help
}
```
Some programs do not include a man page. With this command, the program is run
with "--help" which has evolved into a common way to display the manual.

## Finally
Those were some snippets from my zsh config. Please do tweet
[@teddytroll](https://twitter.com/Teddytroll) if you include something into your
config or found anything of interest. See you on the next installation of
"What's that config". It will probably be my Neovim config.
