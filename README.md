combinatory(1) -- creates a list of possible combinations of chars for a word
=============================================================================

## SYNOPSYS

`combinatory` worder [--entry <string>|-e <string>...] [--howmany] [--save-to-file] [--basedir <path>]

## DESCRIPTION

Taking a word as input, it will create a list of possible variations for that
word using different characters that matches writing the keyboard with a
modifier key.

## EXAMPLE

    $ combinatory worder --entry "ab" --howmany
    60
    $ combinatory worder --entry "ab"
    ab
    Ab
    4b
    $b
    áb
    Áb
    äb
    Äb
    àb
    Àb
    ãb
    Ãb
    ab
    aB
    a8
    a*
    Ab
    AB
    A8
    A*
    4b
    4B
    48
    4*
    $b
    $B
    $8
    $*
    áb
    áB
    á8
    á*
    Áb
    ÁB
    Á8
    Á*
    äb
    äB
    ä8
    ä*
    Äb
    ÄB
    Ä8
    Ä*
    àb
    àB
    à8
    à*
    Àb
    ÀB
    À8
    À*
    ãb
    ãB
    ã8
    ã*
    Ãb
    ÃB
    Ã8
    Ã*
    $ combinatory worder --entry "ab" --save-to-file
    $ ls
    ab.txt

## AUTHOR AND LICENSE

© 2016, Jose-Luis Rivas `<me@ghostbar.co>`. Licensed under the MIT terms, you
can find a copy distributed with the source code under the file `LICENSE`.
