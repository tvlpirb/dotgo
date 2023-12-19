# Dotgo
A simple dotfiles manager written for myself. This is actually a rewrite of [Plain Dotfiles](https://github.com/tvlpirb/plain-dots) to practice learning Go. I also plan to add more features and make use of the Bubblecharm library.

## How it works
You should have a themes directory with a structure as follows:
'''
$ tree themes
themes
├── hypr-ags
├── hypr-catppuccin
├── hypr-catppuccin (copy)
├── hypr-catreborn
├── hypr-eww-gruv
├── hypr-gruvbox
├── hypr-win11
├── root-hypr
└── sway-catppuccin
'''
where each directory contains all the ~/.config files for your chosen theme.

The tool then symlinks anything in the theme directory to the ~/.config directory.

This is pretty simple and does the job perfectly for my needs.
