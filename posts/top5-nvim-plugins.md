+++
title= "5 great Neovim plugins for 2023 🤙"
date = "2023-01-02T16:04:00+01:00"
author = "Theodor "
authorTwitter = "theodorc_" #do not include @
cover = ""
tags = ["neovim", "cli", "config", "2023", "plugin", "git", "linux"]
keywords = ["", ""]
description = "Up to date plugins for your Neovim config in 2023. At least some of my favorites"
showFullContent = false
readingTime = true
draft = false
+++

I finally finished my master thesis June 2022 (will write a post on it someday),
and I started in my first "real" job in August. The project that I was put on
was a brand new Next.js project with typescript+tailwind+hasura(postgres).
Initially I started using VScode(which seems to be turning into a default for
javascript projects 🤷), but I was quickly hitting its limits on
configurability. One Friday night I rewrote my neovim config into lua, and now
I'm working on honing the config until I get the comfiest editing setup. With
this, quite laborious and long, process I have collected a few neat and nice
plugins. Here is an unranked list on great neovim plugins (mostly written in lua) out of
my now 40ish installed
[plugins]("https://github.com/TheodorRene/dotfiles/blob/master/nvim/lua/plugins.lua").
(I assume you have LSP set up so I will not highlight the basic plugins to get
LSP working in this post)

(and yes the syntax highlighting looks kinda funky. I had to set it to vimscript
since lua isn't supported(???) by the Hugo theme I'm using)

## 1. Neoscroll

Vertical movement has been and is still my limiting factor when it comes to
speed. I've tried using `ctrl+d` and `ctrl+u` for page up and down, and also
`ctrl+e` and `ctrl+y` for scrolling line-by-line but I always get dizzy and
disoriented. The first step in fixing this is a tip from
[ThePrimagen](https://www.youtube.com/channel/UC8ENHE5xdFSwx71u3fDH5Xw):

```vim
nmap("<C-u>", "<C-u>zz");
nmap("<C-d>", "<C-d>zz");
```

This remapping keeps the cursor in the middle of the screen when going page up
and down.

```vim
vim.o.scrolloff = 8;
```

This options makes it so that the buffer starts scrolling earlier to always show
8 lines above or under the cursor. If you are still unsure what this does
remember vim/nvims help pages are pretty good `:h scrolloff`

[Neoscroll]("https://github.com/karb94/neoscroll.nvim") adds smooth motion when
using scrolling keybindings; mimicking the act of scrolling with a mouse. Much
more pleasant for the eye, and convenient when that sort of file exploration is
needed.

![Neoscroll in action](/img/neoscroll.gif)

Unfortunately the framerate of the gif does not do the improvement justice so
you just have to trust me.

## 2. Leap

Another plugins thats helped with my vertical speed is
[Leap]("https://github.com/ggandor/leap.nvim"). It adds a new motion using `s`
("leap" forwards) and `S` ("leap" backwards). Pressing `s` and two characters it
will jump to the first occurence of that pair of characters. With duplicates it
shows you a third character to press to leap to that destination. It feels like
magic when you use it since you look at where you want to go, do the keystrokes,
and boom you're there.

![Leap](/img/leap.gif)

Here I want to leap to `tabline`. I press `sta` in sequence, noticing that there
are duplicates and have to press `s` as well which is highlighted in pink/purple.

## 3. Neogit/Gitsigns/Diffview

I'm cheating I know but I couldn't just highlight one of them, and I should
probably do a whole post on working with git inside vim. Git is actually
something I enjoy working with. I currently have 26 aliases in my .gitconfig and
I've set up [difftastic]("https://difftastic.wilfred.me.uk/") as my external
difftool and VScode as my mergetool(haven't gotten there with neovim as of yet).
So I'll keep this short since we are cheating.

(The keymappings are mostly my own, and does not necessarily reflect the default
mappings introduced with the plugins)

### [GitSigns]("https://github.com/lewis6991/gitsigns.nvim")

- Introduces a new motion: Jump between hunks using `[c` and `]c`
- Easily stage and unstage hunks. Keybinding: `ctrl-g+h+s`. Mnemonic: git
  hunk stage
- Show new or deleted lines. Keybinding `ctrl-g+t+{d,l}`. Mnemonic: git
  toggle {delete_lines,lines}

### [Diffview]("https://github.com/sindrets/diffview.nvim")

- Nice diffviewer (duh). `ctrl+g+d+o`
- Easily diff against a branch `ctrl+g+d+d` to diff against `develop`
- Supports staging hunks
- Show possible commands by pressing `g?`

![Diffviewer](/img/diffviewer.png)

### [Neogit]("https://github.com/TimUntersberger/neogit")

- An attempt to create [Magit]("https://magit.vc/")(git plugin for emacs) in Neovim. Mapped to `Alt+s`
- Nice keybindings for basic and advanced operations
- Integration with diffview

![Neogit](/img/neogit.png)

Right part of the screen is Neogit after pressing `c` to commit. It prompts with
options to add, e.g `a` to amend or `h` to disable hooks. Then `c` again will
open a buffer to edit the commit message.

## 4. Telescope

[Telescope]("https://github.com/nvim-telescope/telescope.nvim") is a given if you have ever delved into neovim configuration. Its
building itself up as a standard way of picking things from a list when using
neovim. Its made by one on the neovim maintainers and is a solid piece of work
with many integrations. These are the ones I use:

```vim
nmap('gd', ':lua require"telescope.builtin".lsp_definitions()<CR>', "Go to definitions")
nmap('gi', ':lua require"telescope.builtin".lsp_implementations()<CR>', "Go to implementations")
nmap('<C-x>d', ':lua require"telescope.builtin".diagnostics()<CR>', "Show diagnostics")
nmap('<C-x>f', '<CMD> Telescope current_buffer_fuzzy_find <CR>', "Current buffer fuzzy find")
nmap('<C-x>b', '<CMD> Telescope buffers <CR>', "Search open buffers")
nmap('<C-p>', ':lua require"telescope.builtin".git_files{use_git_root=false} <CR>', "Search git files")
nmap('<leader>p', ':lua require"telescope.builtin".commands() <CR>', "Search commands")
nmap('<C-f>', '<CMD> Telescope live_grep <CR>', "Grep through files")
```

(`nmap` is a function I made. Many create their own functions which are less
verbose than whats built in. This is mine:

```vim
local function nmap(comb, cmd, desc)
	vim.api.nvim_set_keymap('n', comb, cmd, {noremap = true, silent = true, desc = desc})
end
```

)

I especially use `ctrl+p` (same keybindings as VScode if I ever have to return
to the dark side) to search for files in git repos:

![Git files](/img/git_files.png)

and `ctr+f` to search through the content of
all files:

![Live Grep](/img/live_grep.png)

Both feature a nice preview of the file, but I've disabled this for `ctrl+p`.

## 5. Which-key

With all these plugin I've hit a limit on the number of keybindings my feeble
brain can remember. What I've started to do is to namespace actions with a
certain `ctrl+<char>` command. Maybe not the most popular choice, looks very
emacsy, but I'm not a big fan of using the leader key since it feels to slow. An
example is all git related actions are grouped under `ctrl+g`:

```vim
nmap('<C-g>s', '<CMD> Git status <CR>', "GIT: status")
nmap('<C-g>w', ':vsplit term://git status <CR>', "GIT: status but with colors")
nmap('<C-g>p', '<CMD> Git pull  <CR>', "GIT: Pull")
nmap('<C-g>do', '<CMD> DiffviewOpen <CR>', "GIT: Show diff")
nmap('<C-g>dq', '<CMD> DiffviewClose <CR>', "GIT: Close diff")
nmap('<C-g>g', ':Git ', "Git: Fugitive")
nmap('<C-g>dd', '<CMD> DiffviewOpen develop...HEAD <CR>', "GIT: Diff develop")
nmap('<A-g>', '<CMD> Neogit <CR>', "Neogit") -- Except for this one 🤷
```

Gitsigns mappings: (see GitSigns documentation on where they should be placed
as `gs` is not in scope in this example)

```vim
nmap('<C-g>hs', ':Gitsigns stage_hunk<CR>', {desc='stage hunk'})
nmap('<C-g>hr', ':Gitsigns reset_hunk<CR>', {desc='reset hunk'})
nmap('<C-g>hu', gs.undo_stage_hunk, {desc='undo stage hunk'})
nmap('<C-g>td', gs.toggle_deleted, {desc='toggle deleted lines'})
nmap('<C-g>tl', gs.toggle_linehl, {desc='toggle line highlight'})
```

When I press `ctrl+g` alone a "HUD" will open, displaying all possible actions.
What you might have noticed is that there is a `desc` options for each keymap.
This information is displayed in the HUD. See image below:

![Which key](/img/which-key.png)

## End

These were some of my favorite plugins. Let me know if you have any tips of
other plugins I should use. The once I use now can be found
[here]("https://github.com/TheodorRene/dotfiles"). In the future I want to have
a smoother Vim-session experience, and maybe a more visually appealing way of
working with lsp-features(renaming etc). Good luck ricing your setup and
hopefully the ecosytem will thrive and continue to be better. Support your local
plugin maintainer ✊. I'm @theodorc\_ on
[Twitter]("https://twitter.com/theodorc_") and @theodorc@snabelen.no on Mastodon

(And yes, I have line numbers disabled 😎)
