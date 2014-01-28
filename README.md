Shashin
=======

Shashin (Japanese word for picture) is a simple command line program that can resize images.

It supports jpegs and pngs only.

## Install

    go get github.com/nwjlyons/shashin

## Usage

    Usage: shashin [options...] /file/path/to/image.jpg

    Options:
      -w Width to resize image to.
      -h Height to resize image to.
      -g Covert image to grayscale.

Example:

    $ ls
    berries.jpg

    $ shashin -w=1000 -g berries.jpg

    $ ls
    berries-1000x750-grayscale.jpg berries.jpg
