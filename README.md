# dotmux

> Dotfiles &amp; scripts for tmux

* * *

## Install

    git clone https://github.com/leny/dotmux.git ~/.dotmux
    ln -sfv ~/.dotmux/tmux.conf ~/.tmux.conf
    cd ~/.dotmux
    go get -d ./scripts/...
    go build -o bin/systats scripts/systats/main.go
