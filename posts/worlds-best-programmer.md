+++
title = "10 things holding you back from becoming the worlds best developer"
date = "2024-01-08T17:36:40+01:00"
author = "Theodor René Carlsen (Theo)"
authorTwitter = "theodorc_" #do not include @
cover = ""
tags = ["10xDeveloper", "programming", "tips&tricks"]
keywords = ["best", "10xDeveloper"]
description = "Do these 10 simple things different and your life will change"
showFullContent = false
readingTime = false
+++

As a last minute addition to a "Tech show n tell" at my previous employer I
quickly put together this list of "the top 10 things holding you back from
coming the worlds best developer/programmer". I enjoyed making the list so to
make sure it doesn't get lost in the void I'm posting it here as well. Its been
translated from Norwegian and some minor adjustments has been made so that it
works as a blogpost. 

So follow along and make the necessary adjustments so you can be that 10x
developer you've always dreamed of.

(Also I've moved to Copenhagen🇩🇰. Please get in touch if you are in the city and
want to hang or chat. k thx bye)

## 1. You have a functioning Caps lock key

F*ck caps lock. Why would you ever use caps lock. Please bind it to Ctrl and
your wrists will thank me later

## 2. You're not using brands in Typescript
Don't let your code be stringly typed!! Add some compile time differentiation
between your strings:

```ts
type Creds = {
  username: string;
  password: string;
};

function logIn(username: string, password: string) {/* do stuff */}
logIn(creds.password, creds.username);
// No type error, but obviously wrong

// Typescript/javascript voodoo magic. Don't worry about it
declare const brand: unique symbol;

// Woohoo brands
type Brand<T, TBrand> = T & { [brand]: TBrand };
type Username = Brand<string, "Username">;
type Password = Brand<string, "Password">;

type CredsSafe = {
  username: Username
  password: Password;
};
function logIn(username: Username, password: Password){}

logIn(creds.password, creds.username);
<!-- Argument of type 'Password' is not assignable to parameter of type 'Username'. -->
<!--   Type 'Password' is not assignable to type '{ [brand]: "Username"; }'. -->
<!--     Types of property '[brand]' are incompatible. -->
<!--       Type '"Password"' is not assignable to type '"Username"'.(2345) -->
```

Be aware that you can't differentiate between `Username` and `Password` at
runtime, but it's still convenient.

In the end they work like normal strings, and can even be passed to functions
that take "string" as a parameter. But it will error if you are using brands and
it doesn't match the brand set in the function declaration. 

## 3. Your .gitconfig is still the default!!

```bash
> cat ~/.gitconfig
# just do `git <ur alias>`
[alias]
unstage = restore --staged # why isn't this already in git??
amend = commit --amend --no-edit # amend and don't reword commit message
pr = !gh pr create # The github cli is pretty nice! Creates a PR
prview = !gh pr view --web # View pull request in browser
view = !gh repo view --web # View repo in browser
pf = push --force-with-lease # Always --force-with-lease, never --force

prfix = !git commit -a --amend --no-edit && git push --force-with-lease
# -a : Add all tracked files
# --amend: amend to previous commit
# --no-edit: dont prompt for rewording of commit message
# Use this when your colleagues highlight typos or wants you
# to rename a function in your PR

[pull]
rebase = true # 🌶 take: always rebase
autoStash = true # Basically: git stash && git pull && git stash pop

[push]
autoSetupRemote = true  
# No more "There is no tracking information for the current branch."
# It just does what you expect when pushing commits from a new branch

```

Also add this to your {bash,zsh}.rc
```bash
alias lr="ls -ltrh"
# -t Sort by time, newest first
# -r reverse
# -h human readable size
```

## 4. You do crazy🥴 git stuff without doing a backup

Just copy the folder, do wacky stuff in the original folder and if all hell
breaks loose just start over in the backup folder. There is no shame in this!
Swallow your pride! And I'm saying this as person that has an above average
interest in git.

## 5. You actively try to memorize🧠 cli commands!!

Just use `Ctrl-r` and search through your history. If you use a command often
enough you will remember it. Anything else is not worth memorizing.

* Bonus tip: Use [fzf](https://github.com/junegunn/fzf) or some other fuzzy
  searcher with "Ctrl-r" 
* Bonus tip numero dos🌶: Use Github Copilot Cli

# 6. Your prompt is on a single line!

Add some space! Take a breather 🫁😫

I'm sorry to put Scott Hanselman as an example here, I really like him, but
please add a newline!! It hurts looking at it. (I also despise the Powerline
theme, but I won't go into that. It's a personal preference)

![Scott Hanselmans prompt](/img/scott_screen.png)

My prompt:
```bash
~ dev/blank/val |main*|                                               [18:42:23]
> echo "hello"
"hello" 
```

- Full path from $HOME indicated as a tilde(~)
- Current branch and if there are any changes
- Removed unnecessary info I already know  (like username and hostname)
- Simple prompt arrow, or "🐊crocodiles mouth" as we say in Norwegian
- Timestamp for when the command was run. Good for long running tasks, but
  thinking about removing it
- I'm also considering adding the exit code of the command, but I don't use that
  info often

## 7. You're not using `Ctrl+k` in Slack or Github 😤

`Ctrl+k` is usually search or a command bar for your current application.

It's sooo nice. Just try it. It's also supported many other places(sorted by how
often I use it):
  - MDN
  - Tailwind docs
  - Spotify(!)
  - New React dos
  - Nextjs docs

++++ many more

I must warn you that it will be very annoying when a page doesn't implement
this. (Like my blog!)

Bonus: `Ctrl+l` to jump to the urlbar. Possibly `Cmd+l` on Macs

Bonus 2: Using [vimium](https://vimium.github.io/) for navigation and to click
links in the browser

## 8. You see emails when you open your email client 📧

And by that I mean, why are you not using inbox zero📥??? I think this was a big
hype a couple of years ago, and I never stopped. 

**Here's how it works:**

You check your emails. If any of the emails needs no action from you just
archive them. If you need to handle them somehow, do it right away. If you can't
do it now, postpone it until later when you can(Gmail has a button for this).
    Any other emails with like tickets or attachments I need later I just
    favorite or label and then archive.

This way you get a peace of mind when you open your email client, knowing that
if there are no emails, there is nothing for you to do. 

In practice its quite different, but you get the idea. Sometimes I keep an email
around until I've handled it. I currently have 4 emails in my mailbox/client.

Bonus: Add keybindings in Gmail. You do it in the settings
[somewhere](https://support.google.com/mail/answer/6594?hl=en&co=GENIE.Platform%3DDesktop)
and then you can archive by just pressing "e" and navigate with vim-keybindings

## 9. You remember the difference between `??` and `||` in Javascript

- You are too good at your job
- Touch grass
- You should google it once a week and then double check later when submitting the PR

## 10. You use Vim

I have probably saved 27 seconds over the last 5 years I've been using
Vim/Neovim. And I've lost 7 months to configuring stuff and debugging. Just use
VScode with Vim keybindings unless you are looking for a new hobby.

(Check out my [dotfiles](https://github.com/TheodorRene/dotfiles/tree/main/nvim)
repo for neovim inspiration or see my previous
[blogpost](/posts/top5-nvim-plugins) on some cool plugins.


# 11. (Bonus tip because why artificially limit yourselves) You have one big playlist with everything!👿

Make a new one every month!! I've done this for the last 3 years and its fun to
look back at. It also keeps my playlist fresh and in tune with my rapidly
changing music taste when shuffling the current month. They are all in a folder,
but Spotify doesn't allow me to share it so here is my
[playlist](https://open.spotify.com/playlist/2u2FRGC4IIkW6y7xXpUwFQ) for January
thus far.

And yes, I should probably start using Last.fm, but I'm in a reverse sunk cost
fallacy where there are so many years I haven't used it that i feel bad for the
lost data 😢 Oh well. Maybe one day

Superbonus: switch up your perfume when your going into a new *era*. I'm currently
wearing my "new job in Copenhagen"-cologne. Quite rare!

# Fin

That's all! Now you will finally be the best. Enjoy your new life. Also share
your personal list with me either on X(`@theodorc_`) or
Mastodon(`theodorc@snabelen.no`). Any typos or grammatical errors can be fixed
with a PR or an issue on
[Github](https://github.com/TheodorRene/blog/blob/main/content/posts/worlds-best-programmer.md).(There
might be some because this blogpost is 🧙*free from AI*🧙: No LLMs was used in the
making of it and thus no graphics card were hurt in the process😄. 

Cowsay moo:

 ______________
< Takk for meg >
 --------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
